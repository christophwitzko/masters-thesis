package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/christophwitzko/master-thesis/pkg/application/benchmark"
	"github.com/christophwitzko/master-thesis/pkg/cli"
	"github.com/christophwitzko/master-thesis/pkg/logger"
	"github.com/christophwitzko/master-thesis/pkg/netutil"
	"github.com/christophwitzko/master-thesis/pkg/setup"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

func main() {
	log := logger.New()
	rootCmd := &cobra.Command{
		Use:   "application-benchmark-runner",
		Short: "application benchmark runner tool",
		Long:  "This tool is used to run the application benchmarks using artillery.",
		Run:   cli.WrapRunE(log, rootRun),
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}
	rootCmd.Flags().String("reference", "", "git reference or source path of the desired application benchmark config")
	rootCmd.Flags().String("git-repository", "", "git repository to use for installing the applications")
	rootCmd.Flags().String("benchmark-directory", "/tmp/.appbench", "directory to use for running the application benchmarks")
	rootCmd.Flags().String("config", "./artillery/config.yaml", "location of the application benchmark config relative to the repository root or provided source path")
	rootCmd.Flags().StringArray("target", []string{"127.0.0.1:3000"}, "target to run the application benchmark on")
	rootCmd.Flags().String("results-output", "", "path where the results should be stored [e.g. gs://ab-results/app]")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func rootRun(log *logger.Logger, cmd *cobra.Command, args []string) error {
	referenceOrPath := cli.MustGetString(cmd, "reference")
	gitRepository := cli.MustGetString(cmd, "git-repository")
	benchmarkDirectory := cli.MustGetString(cmd, "benchmark-directory")
	relConfigFile := cli.MustGetString(cmd, "config")
	targets := cli.MustGetStringArray(cmd, "target")
	resultsOutputPath := cli.MustGetString(cmd, "results-output")

	if referenceOrPath == "" {
		return fmt.Errorf("source path or git reference is required")
	}

	log.Info("setting up environment...")
	applicationBenchmarkPath, err := setup.ApplicationBenchmarkPath(log, benchmarkDirectory, gitRepository, referenceOrPath)
	if err != nil {
		return err
	}

	appBenchConfigFile := cli.GetAbsolutePath(filepath.Join(applicationBenchmarkPath, relConfigFile))
	log.Infof("using application benchmark config file: %s", appBenchConfigFile)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*30)
	defer cancel()
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	appBenchConfig := &benchmark.Config{
		ConfigFile: appBenchConfigFile,
		OutputPath: resultsOutputPath,
	}
	err = appBenchConfig.Validate()
	if err != nil {
		return err
	}

	log.Info("waiting for targets to be ready....")
	errGroup, groupCtx := errgroup.WithContext(ctx)
	for _, target := range targets {
		target := target
		errGroup.Go(func() error {
			return netutil.WaitForPortOpen(groupCtx, target)
		})
	}
	err = errGroup.Wait()
	if err != nil {
		return err
	}

	log.Info("starting artillery...")
	errGroup, groupCtx = errgroup.WithContext(ctx)
	for _, target := range targets {
		target := target
		errGroup.Go(func() error {
			return benchmark.RunArtillery(groupCtx, log, appBenchConfig, target)
		})
	}
	err = errGroup.Wait()
	if err != nil {
		return err
	}
	log.Infof("done.")
	return nil
}

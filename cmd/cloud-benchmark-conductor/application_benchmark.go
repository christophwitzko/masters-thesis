package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/christophwitzko/master-thesis/pkg/assets"
	"github.com/christophwitzko/master-thesis/pkg/cli"
	"github.com/christophwitzko/master-thesis/pkg/config"
	"github.com/christophwitzko/master-thesis/pkg/gcloud"
	"github.com/christophwitzko/master-thesis/pkg/gcloud/actions"
	"github.com/christophwitzko/master-thesis/pkg/logger"
	"github.com/spf13/cobra"
)

func applicationBenchmarkCmd(log *logger.Logger) *cobra.Command {
	return &cobra.Command{
		Use:     "application-benchmark",
		Aliases: []string{"ab", "app"},
		Short:   "Run application benchmarks in the cloud",
		Run:     cli.WrapRunE(log, applicationBenchmarkRun),
	}
}

func applicationBenchmarkRun(log *logger.Logger, cmd *cobra.Command, args []string) error {
	conf, err := config.NewConductorConfig(cmd)
	if err != nil {
		return err
	}
	service, err := gcloud.NewService(conf)
	if err != nil {
		return err
	}
	defer service.Close()

	ctx, cancel := context.WithTimeout(context.Background(), conf.Timeout)
	defer cancel()
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log.Info("setting up firewall rules...")
	err = service.EnsureFirewallRules(ctx)
	if err != nil {
		return err
	}
	log.Info("running application benchmarks...")
	instance, err := service.GetOrCreateInstance(ctx, "test")
	if err != nil {
		return err
	}
	defer instance.Close()

	err = instance.ExecuteActions(ctx,
		actions.NewActionInstallGo(log),
		actions.NewActionInstallBinary(log, "application-runner", assets.ApplicationRunner),
	)
	if err != nil {
		return err
	}
	err = instance.RunWithLogger(ctx, func(stdout, stderr string) {
		log.Infof("|app-runner| %s%s", stdout, stderr)
	}, "application-runner --v1 main --v2 main --git-repository='https://github.com/christophwitzko/go-benchmark-tests.git' --bind 0.0.0.0")
	if err != nil {
		return err
	}
	log.Info("done")
	return nil
}

package main

import (
	"fmt"
	"log"
	"master-thesis/pkg/benchmark"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "microbenchmark-runner",
	Short: "microbenchmark runner tool",
	Run:   wrapRunE(rootRun),
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func main() {
	rootCmd.Flags().StringP("input-file", "i", "", "input file")
	rootCmd.Flags().StringP("source-path", "s", "", "source path")
	rootCmd.MarkFlagsMutuallyExclusive("source-path", "input-file")

	rootCmd.Flags().SortFlags = true

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func rootRun(cmd *cobra.Command, args []string) error {
	inputFile, _ := cmd.Flags().GetString("input-file")
	sourcePath, _ := cmd.Flags().GetString("source-path")
	if inputFile == "" && sourcePath == "" {
		return fmt.Errorf("either --input-file or --source-path must be set")
	}
	if inputFile != "" {
		log.Printf("loading %s", inputFile)
		return nil
	}

	functions, err := benchmark.GetFunctions(sourcePath)
	if err != nil {
		return err
	}

	for _, function := range functions {
		log.Printf("running %s (%s)", function.Name, function.Directory)
		res, err := benchmark.RunFunction(function)
		if err != nil {
			return err
		}
		log.Printf("%#v", res)
	}
	fmt.Println("done")
	return nil
}

func wrapRunE(fn func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if err := fn(cmd, args); err != nil {
			log.Fatalf("ERROR: %v", err)
		}
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/arschles/pare/commands/run"
	"github.com/spf13/cobra"
)

var (
	version = "devel"
)

func main() {
	opts := struct {
		Version bool
	}{}
	cmd := &cobra.Command{
		Use:                "pare",
		Short:              "The build tool for modern software development",
		SilenceUsage:       true,
		DisableSuggestions: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Version {
				logger.Printf("pare %s", version)
				return nil
			}

			fmt.Print(cmd.UsageString())
			return nil
		},
	}
	cmdFlags := cmd.Flags()
	cmdFlags.BoolVarP(&opts.Version, "version", "v", false, "show the version")
	cmdFlags.Parse(os.Args)

	cmd.AddCommand(run.Root())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

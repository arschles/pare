package run

import (
	"github.com/spf13/cobra"
)

// Root returns the command for the root of the 'pare run' command tree
func Root() *cobra.Command {
	return &cobra.Command{
		Use:     "run [target]",
		Aliases: []string{"r"},
		Short:   "run a target",
		Example: `pare run mytarget`,
		RunE:    run,
	}
}

package cli

import (
	"github.com/spf13/cobra"
)

// NewRootCommand ...
func NewRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "cli",
		Short: "Operate resource",
	}
}

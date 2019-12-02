package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RootCommand ...
type RootCommand struct{ *cobra.Command }

// OperationCommand ...
type OperationCommand struct{ *cobra.Command }

// CreateContentCommand ...
type CreateContentCommand struct{ *cobra.Command }

// ListContentCommand ...
type ListContentCommand struct{ *cobra.Command }

// GetContentCommand ...
type GetContentCommand struct{ *cobra.Command }

// ModifyContentCommand ...
type ModifyContentCommand struct{ *cobra.Command }

// RemoveContentCommand ...
type RemoveContentCommand struct{ *cobra.Command }

func unimplemented() func(*cobra.Command, []string) error {
	return func(*cobra.Command, []string) error { return fmt.Errorf("unimplemented") }
}

package cli

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kanopy",
		Short: "Kanopy CLI",
	}

	cmd.AddCommand(NewApplyCommand())

	return cmd
}

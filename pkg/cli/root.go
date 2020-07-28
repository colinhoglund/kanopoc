package cli

import (
	"github.com/spf13/cobra"
)

type RootCommand struct{ *cobra.Command }

func NewRootCommand() *RootCommand {
	cmd := &RootCommand{&cobra.Command{
		Use:   "kanopy",
		Short: "Kanopy CLI",
	}}
	cmd.AddCommand(NewApplyCommand().Command)
	return cmd
}

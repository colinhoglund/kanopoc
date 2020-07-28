package cli

import (
	"kanopoc/pkg/controller/modules"

	"github.com/spf13/cobra"
)

type RootCommand struct{ *cobra.Command }

func NewRootCommand() *RootCommand {
	cmd := &RootCommand{&cobra.Command{}}
	cmd.AddCommand(NewApplyCommand().Command)
	return cmd
}

type ApplyCommand struct{ *cobra.Command }

func NewApplyCommand() *ApplyCommand {
	cmd := &ApplyCommand{&cobra.Command{
		Use: "apply",
	}}
	cmd.Run = cmd.run
	return cmd
}

func (c *ApplyCommand) run(*cobra.Command, []string) {
	m := modules.New()
	m.Dump()
	m.Apply()
	m.Dump()
	m.Apply()
	m.Dump()
}

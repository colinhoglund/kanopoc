package cli

import (
	"kanopoc/pkg/controller/modules"
	"kanopoc/pkg/provider/helm"

	"github.com/spf13/cobra"
)

type ApplyCommand struct{ *cobra.Command }

func NewApplyCommand() *ApplyCommand {
	cmd := &ApplyCommand{&cobra.Command{
		Use:   "apply",
		Short: "apply all modules",
	}}
	cmd.Run = cmd.run
	return cmd
}

func (c *ApplyCommand) run(*cobra.Command, []string) {
	h := helm.New()
	m := modules.New(h)

	m.Apply()
	m.Dump()
}

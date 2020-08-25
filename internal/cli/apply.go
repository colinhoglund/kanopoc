package cli

import (
	"kanopoc/pkg/controller/modules"
	"kanopoc/pkg/provider/helm"

	"github.com/spf13/cobra"
)

func NewApplyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apply [cluster]",
		Short: "apply all modules",
		Args:  cobra.ExactArgs(1),
	}

	configFlags(cmd)
	dryRunFlag(cmd)

	cmd.Run = applyRun

	return cmd
}

func applyRun(cmd *cobra.Command, args []string) {
	m := modules.New(helm.New())
	m.Apply()
	m.Dump()
}

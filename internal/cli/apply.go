package cli

import (
	"path/filepath"

	"github.com/colinhoglund/kanopoc/pkg/config"
	"github.com/colinhoglund/kanopoc/pkg/config/hierarchy"
	"github.com/colinhoglund/kanopoc/pkg/controller/modules"
	"github.com/colinhoglund/kanopoc/pkg/provider/helm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewApplyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apply [cluster]",
		Short: "apply all modules",
		Args:  cobra.ExactArgs(1),
	}

	configFlags(cmd)
	dryRunFlag(cmd)

	cmd.RunE = applyRunE

	return cmd
}

func applyRunE(cmd *cobra.Command, args []string) error {
	configFile := viper.GetString("configFile")
	dataDir := viper.GetString("dataDir")

	loadOrder := []string{
		"global",
		filepath.Join("cluster", args[0]),
	}

	conf, err := config.NewFromHierarchy(hierarchy.New(configFile, dataDir, loadOrder...))
	if err != nil {
		return err
	}

	m := modules.New(helm.New(), conf)

	if err := m.Apply(); err != nil {
		return err
	}

	m.Dump()

	return nil
}

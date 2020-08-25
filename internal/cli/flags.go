package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func viperBindPFlagEnv(key, env string, flag *pflag.Flag) {
	if err := viper.BindPFlag(key, flag); err != nil {
		panic(err)
	}

	if err := viper.BindEnv(key, env); err != nil {
		panic(err)
	}
}

func configFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("config-file", "c", "kanopy.yml", "Kanopy configuration file")
	viperBindPFlagEnv("configFile", "KANOPY_CONFIG_FILE", cmd.Flags().Lookup("config-file"))

	cmd.Flags().StringP("data-dir", "d", "", "Kanopy cluster configuration data directory")
	viperBindPFlagEnv("dataDir", "KANOPY_DATA_DIR", cmd.Flags().Lookup("data-dir"))

	cmd.PreRun = func(cmd *cobra.Command, args []string) {
		viper.SetConfigFile(viper.GetString("configFile"))
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
}

func dryRunFlag(cmd *cobra.Command) {
	cmd.Flags().Bool("dry-run", false, "Show changes without actually applying them")
	viperBindPFlagEnv("dryRun", "KANOPY_DRY_RUN", cmd.Flags().Lookup("dry-run"))
}

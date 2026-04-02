package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alexhokl/helper/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	noColor bool
)

const applicationName = "confluence-cli"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:               applicationName,
	Short:             "A CLI application to talk to Confluence APIs",
	SilenceUsage:      true,
	PersistentPreRunE: validateConfiguration,
}

func Execute() {
	rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	defaultConfigDesc := fmt.Sprintf("$HOME/.config/%s/config.yaml", applicationName)
	if configDir, err := os.UserConfigDir(); err == nil {
		defaultConfigDesc = filepath.Join(configDir, applicationName, "config.yaml")
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default: %s)", defaultConfigDesc))
	rootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "Disable color output")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	cli.ConfigureViper(cfgFile, "confluence-cli", false, "")
}

func validateConfiguration(_ *cobra.Command, _ []string) error {
	if viper.GetString("email") == "" {
		return fmt.Errorf("email is not configured")
	}
	if viper.GetString("api_key") == "" {
		return fmt.Errorf("api_key is not configured")
	}
	if viper.GetString("organization") == "" {
		return fmt.Errorf("organization is not configured")
	}

	return nil
}

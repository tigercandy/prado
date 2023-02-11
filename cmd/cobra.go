package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tigercandy/prado/cmd/api"
	"github.com/tigercandy/prado/cmd/version"
	"os"
)

var rootCmd = &cobra.Command{
	Use:               "prado",
	Short:             "-v",
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	Long:              `prado`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("require at least one arg")
		}
		return nil
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		usageStr := `欢迎使用Prado` + "\n" + `使用-h查看帮助命令`
		fmt.Println(usageStr)
	},
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
	rootCmd.AddCommand(version.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

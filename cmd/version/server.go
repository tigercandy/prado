package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tigercandy/prado/global"
)

var (
	StartCmd = &cobra.Command{
		Use:     "version",
		Short:   "Version info",
		Example: "prado version",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	fmt.Println(global.Version)
	return nil
}

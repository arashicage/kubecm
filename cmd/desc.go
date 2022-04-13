package cmd

import (
	"github.com/spf13/cobra"

	"kubecm/config"
)

// descCmd 打印配置
var descCmd = &cobra.Command{
	Use:   "desc",
	Short: "打印配置",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		name := args[0]
		config.GetConfig().Desc(name)

	},
}

func init() {

	rootCmd.AddCommand(descCmd)
}

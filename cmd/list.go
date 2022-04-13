package cmd

import (
	"github.com/spf13/cobra"

	"kubecm/config"
)

// listCmd 列出配置
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列出配置",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		config.GetConfig().List()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

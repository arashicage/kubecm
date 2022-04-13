package cmd

import (
	"kubecm/config"

	"github.com/spf13/cobra"
)

// delCmd 删除配置
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "删除配置",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		name := args[0]
		config.GetConfig().Del(name).Sync()
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}

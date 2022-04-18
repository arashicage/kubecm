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

var delCmdHelpTmpl = `
NAME
    del
	
DESCRIPTION
    del 命令用于移除指定别名的 kubeconfig。

EXAMPLES
    kubecm del <name>

`

func init() {

	rootCmd.AddCommand(delCmd)
	delCmd.SetHelpTemplate(delCmdHelpTmpl)
}

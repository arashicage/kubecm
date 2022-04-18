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

var listCmdHelpTmpl = `
NAME
    list
	
DESCRIPTION
    list 命令用于列出受管的 kubeconfig 的别名和路径。

EXAMPLES
    kubecm list

`

func init() {

	rootCmd.AddCommand(listCmd)
	listCmd.SetHelpTemplate(listCmdHelpTmpl)
}

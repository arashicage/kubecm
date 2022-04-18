package cmd

import (
	"github.com/spf13/cobra"

	"kubecm/config"
)

// switchCmd 切换配置
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "切换配置",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		name := args[0]
		config.GetConfig().Switch(name).Sync()
	},
}

var switchCmdHelpTmpl = `
NAME
    switch
	
DESCRIPTION
    switch 命令切换到指定别名的 kubeconfig 。

EXAMPLES
    kubecm switch <name>

`

func init() {

	rootCmd.AddCommand(switchCmd)
	switchCmd.SetHelpTemplate(switchCmdHelpTmpl)
}

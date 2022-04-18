package cmd

import (
	"github.com/spf13/cobra"

	"kubecm/config"
)

// descCmd 打印配置
var descCmd = &cobra.Command{
	Use:   "desc",
	Short: "打印配置",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var name string
		if len(args) == 1 {
			name = args[0]
		} else {
			name = config.GetConfig().Current
		}
		config.GetConfig().Desc(name)
	},
}

var descCmdHelpTmpl = `
NAME
    desc
	
DESCRIPTION
    desc 命令用于打印指定别名的 kubeconfig 的内容。

EXAMPLES
    kubecm desc <name>

`

func init() {

	rootCmd.AddCommand(descCmd)
	descCmd.SetHelpTemplate(descCmdHelpTmpl)
}

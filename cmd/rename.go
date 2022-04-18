package cmd

import (
	"github.com/spf13/cobra"

	"kubecm/config"
)

// renameCmd 配置更名
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "配置更名",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		oldName, newName := args[0], args[1]
		config.GetConfig().Rename(oldName, newName).Sync()
	},
}

var renameCmdHelpTmpl = `
NAME
    rename
	
DESCRIPTION
    rename 命令调整受管的 kubeconfig 的别名。

EXAMPLES
    kubecm rename old new

`

func init() {

	rootCmd.AddCommand(renameCmd)
	renameCmd.SetHelpTemplate(renameCmdHelpTmpl)
}

package cmd

import (
	"github.com/spf13/cobra"

	"kubecm/config"
)

// addCmd 新增配置
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "新增配置",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		name := args[0]
		path, _ := cmd.Flags().GetString("file")
		move, _ := cmd.Flags().GetBool("move")
		config.GetConfig().Add(name, path, move).Sync()
	},
}

var addCmdHelpTmpl = `
NAME
    add
	
DESCRIPTION
    add 命令用于将 kubeconfig 加入 kubecm 的管理，kubecm 将通过别名对其进行引用。kubeconfig 将被复制
    或移动(当您使用了 --move)到 ${HOME}/.kube/kubeconfig 目录下，kubecm 在 ${HOME}/.kube/kubecm.yaml 
    中保存了其管理的 kubeconfig 的清单。

EXAMPLES
    kubecm add <name> -f path/to/kubeconfig 
    kubecm add <name> -f path/to/kubeconfig --move

`

func init() {

	rootCmd.AddCommand(addCmd)
	addCmd.SetHelpTemplate(addCmdHelpTmpl)
	addCmd.PersistentFlags().StringP("file", "f", "", "用于指定 kubeconfig 的路径 (required)")
	addCmd.PersistentFlags().BoolP("move", "m", false, "如果使用该选项，将移动 kubeconfig 文件而非拷贝")
}

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

		path, _ := cmd.Flags().GetString("file")
		move, _ := cmd.Flags().GetBool("move")
		config.GetConfig().Add(args[0], path, move).Sync()
	},
}

func init() {

	rootCmd.AddCommand(addCmd)
	addCmd.SetHelpTemplate("kubecm add foo -f path/to/config")
	addCmd.PersistentFlags().StringP("file", "f", "", "Path to kubeconfig")
	addCmd.PersistentFlags().BoolP("move", "m", false, "Move file to KubeConfigVault instead of copy")
}

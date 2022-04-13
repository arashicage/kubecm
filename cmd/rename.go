package cmd

import (
	"github.com/spf13/cobra"

	"kubecm/config"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "配置更名",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		oldName, newName := args[0], args[1]
		config.GetConfig().Rename(oldName, newName).Sync()

	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
}

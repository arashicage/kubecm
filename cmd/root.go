package cmd

import (
	"os"

	"kubecm/utils"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"

	"kubecm/config"
)

var rootCmd = &cobra.Command{
	Use:   "kubecm",
	Short: "Kubecm 是一个用于管理和切换 kubeconfig 的工具",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func initKubecm() {

	conf, err := homedir.Expand(config.AppRC)
	utils.CheckErr(err)

	kubeConfigPath, err := homedir.Expand(config.KubeConfigVault)
	utils.CheckErr(err)

	err = os.MkdirAll(kubeConfigPath, 0755)
	utils.CheckErr(err)

	if !utils.FileExists(conf) {
		f, err := os.Create(conf)
		defer func(f *os.File) {
			_ = f.Close()
		}(f)
		utils.CheckErr(err)
	}

}

func init() {

	cobra.OnInitialize(initKubecm)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
}

func Execute() {

	_ = rootCmd.Execute()
}

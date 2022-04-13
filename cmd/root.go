package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"

	"kubecm/config"
)

var rootCmd = &cobra.Command{
	Use:   "kubecm",
	Short: "Kubecm 是一个用于管理和切换 kubeconfig 的工具",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func FileExists(name string) bool {

	info, err := os.Lstat(name)
	if err == nil {
		return !info.IsDir()
	}
	return !os.IsNotExist(err)
}

func initKubecm() {

	conf, err := homedir.Expand(config.AppRC)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	kubeConfigPath, err := homedir.Expand(config.KubeConfigVault)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := os.MkdirAll(kubeConfigPath, 0755); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if FileExists(conf) {
		return
	}
	f, err := os.Create(conf)
	defer f.Close()
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
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

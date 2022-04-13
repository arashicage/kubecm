package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version          string = "unknown"
	OS_Arch          string = "unknown"
	GitHASH          string = "unknown"
	BuildAt          string = "unknown"
	versionFormatter string = `dinoctl, build at %s
Version: %s
OS/Arch: %s
GitHash: %s, %s for short
`
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "版本信息",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf(versionFormatter, BuildAt, Version, OS_Arch, GitHASH, GitHASH[:7]))
	},
}

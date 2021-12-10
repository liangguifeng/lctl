package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd 声明命令
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "lctl tools is help developers tools",
	Long:  `lctl tools a CLI library for developer.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}

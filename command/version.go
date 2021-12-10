package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

// 声明当前版本
const LCTL_VERSION = "v1.1.0"

// versionCmd 声明命令
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show current lctl tools version",
	Long:  "show current lctl tools version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("lctl tools version %s\n", LCTL_VERSION)
		fmt.Println(` 
 __         ______     ______   __        
/\ \       /\  ___\   /\__  _\ /\ \       
\ \ \____  \ \ \____  \/_/\ \/ \ \ \____  
 \ \_____\  \ \_____\    \ \_\  \ \_____\ 
  \/_____/   \/_____/     \/_/   \/_____/ 
`)
	},
}

// init 默认执行
func init() {
	rootCmd.AddCommand(versionCmd)
}

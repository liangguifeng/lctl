package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

var genpbCmd = &cobra.Command{
	Use:   "genpb",
	Short: "generate proto files.",
	Long: `generate proto files
including file.pb.go、file_grpc.pb.go、file.pb.gw.go、file.swagger.json and swagger_json.go`,
	Run: func(cmd *cobra.Command, args []string) {
		// 拦截错误
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[info] Recover error : %v", err)
			}
		}()

		// 获取当前的proto目录所有.proto文件
		pwd, _ := os.Getwd()

		//获取当前目录下的所有文件或目录信息
		filepath.Walk(pwd+"/proto", func(path string, info os.FileInfo, err error) error {
			fmt.Println(path)        //打印path信息
			fmt.Println(info.Name()) //打印文件或目录名
			return nil
		})
		//
		//var output []byte
		//var err error
		//cmd2 := exec.Command("bash", "-c", "protoc")
		//if output, err = cmd2.CombinedOutput(); err != nil {
		//	log.Fatal(err)
		//}
		//log.Printf("[info] Recover error : %v", string(output))
	},
}

func init() {
	rootCmd.AddCommand(genpbCmd)
}

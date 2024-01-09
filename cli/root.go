package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// 用于接收命令行参数
	version bool
)
var RootCmd = &cobra.Command{
	Use:     "oss-cloudstation-cli",
	Long:    "oss-cloudstation-cli 存储工具",
	Short:   "oss-cloudstation-cli 存储工具",
	Example: "oss-cloudstation-cli cmds",
	RunE: func(cmd *cobra.Command, args []string) error {
		if version {
			fmt.Println("oss-cloudstation-cli version 0.0.1")
		}

		return nil
	},
}

func init() {
	f := RootCmd.PersistentFlags()
	//f.StringVarP(&ossProvider, "provider", "p", "local", "oss storage provider [aliyun/txyun/minio]")
	f.BoolVarP(&version, "version", "v", false, "oss station version 版本信息")
}

package cli

import (
	"fmt"
	"oss-station/store"
	"oss-station/store/aliyun"
	localminio "oss-station/store/local-minio"
	"oss-station/store/txyun"

	"github.com/spf13/cobra"
)

var (
	ossProvider  string
	endpoint     string
	accessKey    string
	accessSecret string
)

var UploadCmd = &cobra.Command{
	Use:     "upload",
	Long:    "upload 文件上传",
	Short:   "upload 文件上传",
	Example: "upload -f filename",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			uploader store.Uploader
			err      error
		)
		switch ossProvider {
		case "local":
			localminio.NewLocalminioStore(&localminio.Options{
				Endpoint:        endpoint,
				AccessKeyID:     accessKey,
				SecretAccessKey: accessSecret,
			})
		case "aliyun":
			aliyun.NewAliOssStore(&aliyun.Options{})
		case "txyun":
			txyun.NewTxyunStore()
		default:
			return fmt.Errorf("oss provider %s not support", ossProvider)
		}
		return nil
	},
}

func init() {
	f := UploadCmd.PersistentFlags()
	f.StringVarP(&ossProvider, "provider", "p", "local", "oss storage provider [aliyun/txyun/minio]")
	f.StringVarP(&endpoint, "endpoint", "e", "http://minio.kkk987happy.com", "oss storage provider endpoint")
	f.StringVarP(&accessKey, "accessKey", "a", "", "oss storage provider accessKey")
	f.StringVarP(&accessSecret, "accessSecret", "s", "", "oss storage provider accessSecret")

	RootCmd.AddCommand(UploadCmd)
}

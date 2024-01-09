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
	bucketName   string
	uploadFile   string
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
			uploader, err = localminio.NewLocalminioStore(&localminio.Options{
				Endpoint:        endpoint,
				AccessKeyID:     accessKey,
				SecretAccessKey: accessSecret,
			})
		case "aliyun":
			uploader, err = aliyun.NewAliOssStore(&aliyun.Options{
				Endpoint:     endpoint,
				AccessKey:    accessKey,
				AccessSecret: accessSecret,
			})
		case "txyun":
			uploader = txyun.NewTxyunStore()
		default:
			return fmt.Errorf("oss provider %s not support", ossProvider)
		}
		if err != nil {
			return err
		}

		//使用uploader上传文件
		uploader.Upload(bucketName, uploadFile, uploadFile)
		return nil
	},
}

func init() {
	f := UploadCmd.PersistentFlags()
	f.StringVarP(&ossProvider, "provider", "p", "local", "oss storage provider [aliyun/txyun/minio]")
	f.StringVarP(&endpoint, "endpoint", "e", "http://minio.kkk987happy.com", "oss storage provider endpoint")
	f.StringVarP(&accessKey, "access_key", "a", "", "oss storage provider accessKey")
	f.StringVarP(&accessSecret, "access_secret", "s", "", "oss storage provider accessSecret")
	f.StringVarP(&bucketName, "bucket_name", "b", "", "oss storage provider bucket")
	f.StringVarP(&uploadFile, "file_name", "f", "", "oss storage provider file")
	RootCmd.AddCommand(UploadCmd)
}

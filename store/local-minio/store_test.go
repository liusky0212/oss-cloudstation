package localminio_test

import (
	"os"
	"oss-station/store"
	localminio "oss-station/store/local-minio"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	uploader store.Uploader
)

var (
	BucketName = os.Getenv("MINIO_BUCKET_NAME")
	UploadFile = os.Getenv("UPLOAD_FLIE")
	ObjectKey  = os.Getenv("OBJECTKEY")
	// Endpoint        = os.Getenv("MINIO_SERVER")
	// AccessKeyID     = os.Getenv("MINIO_AK")
	// SecretAccessKey = os.Getenv("MINIO_SK")
	// UseSSLstr       = os.Getenv("MINIO_USESSL")
)

// localminio upload 测试用例
func TestUpload(t *testing.T) {
	//使用assert库编写测试用例 断言语句
	//获取断言实例
	should := assert.New(t)

	err := uploader.Upload(BucketName, ObjectKey, UploadFile)
	if should.NoError(err) {
		//没有error开始下一个步骤
		t.Log("upload succeess")
	}
}

// 通过init编写uploader实例化
func init() {
	localminio, err := localminio.NewDefaultLocalminioStore()
	if err != nil {
		panic(err)
	}			
	uploader = localminio	
}

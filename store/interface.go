package store

type OSSInfo struct {
	Endpoint     string
	AccessKey    string
	AccessSecret string
}

// 定义如何上传文件到bucket
type Uploader interface {
	Upload(bucketName string, objectKey string, fileName string) error
}

// 定义如何下载文件到本地
type Downloader interface {
	Download(bucketName string, objectKey string, fileName string) error
}

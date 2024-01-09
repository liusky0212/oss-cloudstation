package store


// 定义如何上传文件到bucket
type Uploader interface {
	Upload(bucketName string,objectKey string,fileName string) error
}

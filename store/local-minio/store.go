package localminio

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	localminio "oss-station/store"
	"strconv"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type LocalminioStore struct {
	client *minio.Client
}

type Options struct {
	// Endpoint        string
	// AccessKeyID     string
	// SecretAccessKey string
	UseSSL  bool
	OSSInfo *localminio.OSSInfo // 匿名结构体，继承OSSInfo结构体
}

var (
	//对象是否实现接口的约束
	_ localminio.Uploader = &LocalminioStore{}
)

// 参数校验
func (o *Options) Validate() error {
	if o.OSSInfo.Endpoint == "" {
		return fmt.Errorf("minio地址为空")
	}
	if o.OSSInfo.AccessKey == "" || o.OSSInfo.AccessSecret == "" {
		return fmt.Errorf("accessKeyID 或 secretAccessKey 为空")
	}
	return nil
}

// localminioStore构造函数
// 接受一个 Options 结构体作为参数，这个结构体包含了创建 MinIO 客户端所需的所有配置信息
func NewLocalminioStore(opts *Options) (*LocalminioStore, error) {
	//校验参数
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	//NEW minioClient
	minioClient, err := minio.New(opts.OSSInfo.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(opts.OSSInfo.AccessKey, opts.OSSInfo.AccessSecret, ""),
		Secure: opts.UseSSL,
	})
	if err != nil {
		return nil, err
	}

	log.Printf("%#v\n", minioClient) // minioClient is now setup
	log.Printf("minio实例创建成功")
	return &LocalminioStore{
		client: minioClient,
	}, nil
}

// 使用环境变量创建 LocalminioStore 实例
func NewDefaultLocalminioStore(opts *Options) (*LocalminioStore, error) {
	useSSL, err := strconv.ParseBool(os.Getenv("MINIO_USESSL"))
	if err != nil {
		return nil, err
	}
	return NewLocalminioStore(&Options{
		OSSInfo: &localminio.OSSInfo{
			Endpoint:     os.Getenv("MINIO_SERVER"),
			AccessKey:    os.Getenv("MINIO_AK"),
			AccessSecret: os.Getenv("MINIO_SK"),
		},
		UseSSL: useSSL,
	})
}

// 封装minio上传文件接口
func uploadfile(localminioStore *LocalminioStore, bucketName, objectKey, uploadFile string) error {
	//读取文件
	file, err := os.Open(uploadFile)
	if err != nil {
		fmt.Printf("读取上传文件失败， %v", err)
		return err
	}
	log.Printf("读取上传文件成功")
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.Printf("获取上传文件stat成功")
	// log.Printf("%#v\n", localminioStore.client)
	// log.Println(bucketName, objectKey, file, fileStat.Size())
	log.Println(localminioStore.client.EndpointURL().User.Username())
	uploadInfo, err := localminioStore.client.PutObject(context.Background(), bucketName, objectKey, file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.Printf("%s 文件已上传到minio  ，文件size：%d\n", uploadFile, uploadInfo.Size)
	return nil
}

// 封装minio获取下载链接接口
func downloadFileURL(localminioStore *LocalminioStore, bucketName, objectKey string, expires time.Duration) error {
	// Set request parameters for content-disposition.
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=\"%s\"", objectKey))

	presignedURL, err := localminioStore.client.PresignedGetObject(context.Background(), bucketName, objectKey, expires, reqParams)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("下载链接 %s ", presignedURL)

	return nil
}

// minio上传文件并获取下载链接
func (localminioStore *LocalminioStore) Upload(bucketName string, objectKey string, filename string) error {
	if err := uploadfile(localminioStore, bucketName, objectKey, filename); err != nil {
		fmt.Printf("上传到bucket失败, %s", err)
		return err
	}
	//3.打印下载链接
	fmt.Println("\n注意: 文件下载有效期为1天, 中转站保存时间为3天, 请及时下载")
	if err := downloadFileURL(localminioStore, bucketName, objectKey, time.Second*24*60*60); err != nil {
		fmt.Printf("获取下载链接失败, %s", err)
		return err
	}
	fmt.Println("\n文件上传成功")
	return nil
}

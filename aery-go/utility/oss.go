package utility

import (
	"aery-go/internal/consts"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/dustin/go-humanize"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// Goss 结构体
type Goss struct{}

type StsCredentials struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
	Expiration      string
}

// 共享的 bucket 和 Goss 实例
var (
	bucket *oss.Bucket
	Oss    Goss
)

// InitOssBucket 初始化 OSS Bucket
func InitOssBucket() *oss.Bucket {
	// 创建OSSBucket实例
	if bucket == nil {
		client, err := oss.New("https://oss-cn-fuzhou.aliyuncs.com", consts.AccessKeyId, consts.AccessKeySecret)
		if err != nil {
			panic(err)
		}
		Bucket, err := client.Bucket("aery")
		if err != nil {
			panic(err)
		}
		bucket = Bucket
	}
	return bucket
}

func (g *Goss) Getsts() (*sts.AssumeRoleResponse, error) {
	// 创建 STS 客户端
	client, err := sts.NewClientWithAccessKey("cn-fuzhou", consts.AccessKeystsId, consts.AccessKeystsSecret)
	if err != nil {
	}

	// 创建 AssumeRole 请求
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "HTTPS"
	request.RoleArn = "acs:ram::1273012912686046:role/sts"
	request.RoleSessionName = "s-session"
	request.DurationSeconds = "3600"

	// 发送请求获取 STS 凭证
	response, err := client.AssumeRole(request)
	if err != nil {
		return nil, err
	}
	return response, nil

	//// 使用 STS 凭证创建 OSS 客户端
	//ossClient, err := oss.New("oss-cn-fuzhou.aliyuncs.com", response.Credentials.AccessKeyId, response.Credentials.AccessKeySecret, oss.SecurityToken(response.Credentials.SecurityToken))
	//if err != nil {
	//	return err
	//}
	//
	//// 获取指定的桶
	//buckets, err := ossClient.Bucket("aery")
	//if err != nil {
	//	return err
	//}
	//
	//// 下载文件到指定的本地路径
	//err = buckets.GetObjectToFile(name, lcoal)
	//if err != nil {
	//	return err
	//}
	//return nil
}

// UploadFile 用于将本地文件上传到OSS存储桶。
func (g *Goss) UploadFile(objectName, localFileName string) error {
	// 上传文件。
	err := bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		return err
	}
	return nil
}

// DownloadFile 用于从OSS存储桶下载一个文件到本地路径。
func (g *Goss) DownloadFile(objectName, downloadedFileName string) error { // 改为大写
	err := bucket.GetObjectToFile(objectName, downloadedFileName)
	if err != nil {
		return err
	}
	return nil
}

type WriteCounter struct {
	Total     uint64
	FileSize  uint64
	StartTime time.Time
	LastTime  time.Time
	LastTotal uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	percent := float64(wc.Total) / float64(wc.FileSize) * 100
	elapsed := time.Since(wc.LastTime).Seconds()
	speed := float64(wc.Total-wc.LastTotal) / elapsed

	wc.LastTime = time.Now()
	wc.LastTotal = wc.Total

	fmt.Printf("\r%s", strings.Repeat(" ", 60))
	fmt.Printf("\rDownloading... %s / %s (%.2f%%) at %s/s",
		humanize.Bytes(wc.Total),
		humanize.Bytes(wc.FileSize),
		percent,
		humanize.Bytes(uint64(speed)),
	)
}

func (g *Goss) DownloadFileWithProgress(objectName, downloadedFileName string) error {
	meta, err := bucket.GetObjectDetailedMeta(objectName)
	if err != nil {
		return err
	}

	fileSize, err := humanize.ParseBytes(meta.Get("Content-Length"))
	if err != nil {
		return err
	}

	reader, err := bucket.GetObject(objectName)
	if err != nil {
		return err
	}
	defer func(reader io.ReadCloser) {
		err := reader.Close()
		if err != nil {

		}
	}(reader)

	// 直接创建目标文件
	destFile, err := os.Create(downloadedFileName)
	if err != nil {
		return err
	}
	defer func() {
		if err := destFile.Close(); err != nil {
			fmt.Println("Error closing destination file:", err)
		}
	}()

	counter := &WriteCounter{
		FileSize:  fileSize,
		StartTime: time.Now(),
		LastTime:  time.Now(),
	}

	// 下载文件并通过 TeeReader 记录进度
	if _, err = io.Copy(destFile, io.TeeReader(reader, counter)); err != nil {
		return err
	}

	fmt.Print("\nDownload completed successfully.")
	return nil
}

// ListObjects 用于列举OSS存储空间中的所有对象。
func (g *Goss) ListObjects() error {
	// 列举文件。
	marker := ""
	for {
		lsRes, err := bucket.ListObjects(oss.Marker(marker))
		if err != nil {
			return err
		}

		// 打印列举文件，默认情况下一次返回100条记录。
		for _, object := range lsRes.Objects {
			log.Printf("Object: %s", object.Key)
		}

		if !lsRes.IsTruncated {
			break
		}
		marker = lsRes.NextMarker
	}

	return nil
}

// DeleteObject 用于删除OSS存储空间中的一个对象。
func (g *Goss) DeleteObject(objectName string) error {
	// 删除文件。
	err := bucket.DeleteObject(objectName)
	if err != nil {
		return err
	}
	return nil
}

package minio

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/minio/minio-go"
)

var (
	internalClient *Client
	once           sync.Once
)

// ErrorInvalidName 无效的文件名
var ErrorInvalidName = errors.New("invalid file name")

// Init 初始化minio客户端
func Init(addr, accessKey, secretKey string) *Client {
	once.Do(func() {
		cli, err := minio.New(addr, accessKey, secretKey, false)
		if err != nil {
			panic(err)
		}
		internalClient = &Client{cli}
	})
	return internalClient
}

// GetClient 获取文件存储客户端
func GetClient() *Client {
	return internalClient
}

// Client minio客户端
type Client struct {
	cli *minio.Client
}

// MinioClient 文件存储客户端
func (a *Client) MinioClient() *minio.Client {
	return a.cli
}

// Store 保存文件
// filename 前3段约束(第一段约束为(前缀)，第二段为bucket(业务类型)，第三段以后为文件key)
func (a *Client) Store(ctx context.Context, filename string, data io.Reader, size int64) error {
	if ctx == nil {
		ctx = context.Background()
	}

	bucket, objName, err := a.parseFilename(filename)
	if err != nil {
		return err
	}

	exists, err := a.cli.BucketExists(bucket)
	if err != nil {
		return err
	} else if !exists {
		err = a.cli.MakeBucket(bucket, "local")
		if err != nil {
			return err
		}
	}

	buf, err := ioutil.ReadAll(data)
	if err != nil {
		return err
	}

	rd := bytes.NewBuffer(buf)
	if size == 0 {
		size = int64(rd.Len())
	}

	_, err = a.cli.PutObjectWithContext(ctx, bucket, objName, rd, size, minio.PutObjectOptions{
		ContentType: http.DetectContentType(buf),
		NumThreads:  2,
	})

	return err
}

// 解析文件名
func (a *Client) parseFilename(filename string) (string, string, error) {
	if len(filename) > 0 && filename[0] == '/' {
		filename = filename[1:]
	}
	names := strings.Split(filename, "/")
	if len(names) < 3 {
		return "", "", ErrorInvalidName
	}

	return strings.ToLower(names[1]), strings.Join(names[2:], "/"), nil
}

// Get 获取文件对象
func (a *Client) Get(ctx context.Context, filename string) (*minio.Object, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	bucketName, objectName, err := a.parseFilename(filename)
	if err != nil {
		return nil, err
	}

	return a.cli.GetObjectWithContext(ctx, bucketName, objectName, minio.GetObjectOptions{})
}

// Stat 文件状态信息
func (a *Client) Stat(filename string) (minio.ObjectInfo, error) {
	bucketName, objectName, err := a.parseFilename(filename)
	if err != nil {
		return minio.ObjectInfo{}, err
	}

	return a.cli.StatObject(bucketName, objectName, minio.StatObjectOptions{})
}

// GetBase64 获取文件的base64数据
func (a *Client) GetBase64(ctx context.Context, filename string) (string, error) {
	obj, err := a.Get(ctx, filename)
	if err != nil {
		return "", err
	}
	defer obj.Close()

	buf, err := ioutil.ReadAll(obj)
	if err != nil {
		return "", err
	}

	s := base64.StdEncoding.EncodeToString(buf)
	return s, nil
}


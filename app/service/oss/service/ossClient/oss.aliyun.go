package ossClient

import (
	"bytes"
	"fmt"
	aliOss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"libong/common/context"
	"libong/common/log"
	"libong/oss/app/service/oss/api"
	"os"
	"path"
	"time"
)

type AliYunConf struct {
	EndPoint        string
	AccessKeyID     string
	AccessKeySecret string
	Bucket          string
	AccessHost      string
}

// AliYunOss .
type AliYunOss struct {
	accessHost string
	bucket     *aliOss.Bucket
	client     *aliOss.Client
}

func newAliYunOss(config *AliYunConf) *AliYunOss {
	var aliYun = &AliYunOss{
		accessHost: config.AccessHost,
	}
	client, err := aliOss.New(config.EndPoint, config.AccessKeyID, config.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	bucket, err := client.Bucket(config.Bucket)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	aliYun.bucket = bucket
	return aliYun
}

// UploadImage .
//func (o *AliYunOss) UploadImage(ctx context.Context, req *UploadOssReq) (*UploadOssResp, error) {
//	id := uuid.CreateUUID()
//	year := time.Now().Format("2006")
//	month := time.Now().Format("01")
//	fileNameWithSuffix := path.Base(req.FileName)
//	//获取文件的后缀(文件类型)
//	fileType := path.Ext(fileNameWithSuffix)
//	// out, err := img.ThumbnailB2B(bs, w, h)
//	objectACL := aliOss.ObjectACL(aliOss.ACLPublicRead)
//	imgName := year + "/" + month + "/" + uuid.CreateUUID() + fileType
//	downloadURL := o.accessHost + "/" + imgName
//
//	//TODO 解决方法
//	//	_ "image/gif"
//	//	_ "image/jpeg"
//	//	_ "image/png"
//	//img, _, err := image.Decode(bytes.NewReader(byteArray))
//	//if err != nil {
//	//	return nil, errors.ImgFormatError
//	//}
//	err := o.bucket.PutObject(imgName, bytes.NewReader(req.ByteArray), objectACL)
//	if err != nil {
//		log.Error("UploadImage:", err)
//		return nil, err
//	}
//	return &UploadOssResp{
//		ID:       id,
//		FileType: fileType,
//		Name:     imgName,
//		URL:      downloadURL,
//		//Width:    int64(img.Bounds().Dx()),
//		//Height:   int64(img.Bounds().Dy()),
//	}, nil
//}

func (o *AliYunOss) Upload(ctx context.Context, req *api.UploadReq) (*api.UploadResp, error) {
	id := uuid.New().String()
	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	fileNameWithSuffix := path.Base(req.Name)
	//获取文件的后缀(文件类型)
	fileType := path.Ext(fileNameWithSuffix)
	objectACL := aliOss.ObjectACL(aliOss.ACLPublicRead)
	var uploadFileName string
	if req.IsKeepOriginalName {
		uploadFileName = req.Name
	} else {
		uploadFileName = id + fileType
	}
	imgName := year + "/" + month + "/" + uploadFileName
	downloadURL := o.accessHost + "/" + imgName
	reader := bytes.NewReader(req.Data)
	err := o.bucket.PutObject(imgName, reader, objectACL)
	if err != nil {
		log.Error(ctx, "UploadFile:", err)
		return nil, err
	}
	//判断是否是图片
	var width int64
	var height int64
	img, _, err := image.Decode(reader)
	if err == nil {
		height = int64(img.Bounds().Dy())
		width = int64(img.Bounds().Dx())
	}
	return &api.UploadResp{
		Id:       id,
		FileType: fileType,
		Name:     imgName,
		Url:      downloadURL,
		Width:    width,
		Height:   height,
	}, nil
}

//func (s *Service) UploadFileBase64(ctx context.Context, uid int64, imageBase64 string) (string, error) {
//	b, _ := regexp.MatchString(`^data:\s*image/(\w+);base64,`, imageBase64)
//	if !b {
//		return "", errors.ImgFormatError
//	}
//	bucket, err := s.client.Bucket(s.conf.OSSConf.Bucket)
//	if err != nil {
//		fmt.Println("Error:", err)
//		os.Exit(-1)
//	}
//	year := time.Now().Format("2006")
//	month := time.Now().Format("01")
//
//	re, _ := regexp.Compile(`^data:\s*image/(\w+);base64,`)
//	allData := re.FindAllSubmatch([]byte(imageBase64), 2)
//	fileType := string(allData[0][1]) //png ，jpeg 后缀获取
//	base64Str := re.ReplaceAllString(imageBase64, "")
//	byte, _ := base64.StdEncoding.DecodeString(base64Str)
//	//获取文件的后缀(文件类型)
//	objectACL := oss1.ObjectACL(oss1.ACLPublicRead)
//	imgName := year + "/" + month + "/" + uuid.CreateUUID() + fileType
//	err = bucket.PutObject(imgName, bytes.NewReader(byte), objectACL)
//	if err != nil {
//		fmt.Println("Error:", err)
//		return "", err
//	}
//	return s.conf.OSSConf.AccessHost + "/" + imgName, nil
//}

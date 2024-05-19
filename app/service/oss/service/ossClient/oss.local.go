package ossClient

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/google/uuid"
	"image"
	"io"
	"libong/common/context"
	customHttp "libong/common/server/http"
	"libong/oss/app/service/oss/api"
	"libong/oss/errors"
	"mime/multipart"
	"net/http"
	"path"
	"time"
)

type LocalOssConf struct {
	AccessKey        string
	AccessSecret     string
	BucketName       string
	BucketServiceUrl string
}

type LocalOss struct {
	Config *LocalOssConf
}

func newLocalOss(config *LocalOssConf) *LocalOss {
	return &LocalOss{
		Config: config,
	}
}
func (o *LocalOss) Upload(ctx context.Context, req *api.UploadReq) (*api.UploadResp, error) {
	id := uuid.New().String()
	//签名
	param := "appId" + o.Config.AccessKey + "appSecret" + o.Config.AccessSecret
	sum := md5.Sum([]byte(param))
	//设置请求文件
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	if cff, err := w.CreateFormFile("file", req.Name); err == nil {
		cff.Write(req.Data)
	}
	w.Close()
	request, err := http.NewRequest("POST", o.Config.BucketServiceUrl, buf)
	if err != nil {
		return nil, err
	}
	//设置请求头
	request.Header.Set("Content-Type", w.FormDataContentType())
	request.Header.Set("sign", hex.EncodeToString(sum[:]))
	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	day := time.Now().Format("02")
	imgPath := o.Config.BucketName + "/" + year + "/" + month + "/" + day
	request.Header.Set("prefix", imgPath)
	//请求
	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.UploadError
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var resp customHttp.BaseResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	var url string
	if resp.Data != nil {
		url, _ = resp.Data.(string)
	}
	//获取文件的后缀(文件类型)
	fileNameWithSuffix := path.Base(req.Name)
	fileType := path.Ext(fileNameWithSuffix)
	//如果是图片获取宽高
	var width int64
	var height int64
	if req.Type == uint32(api.UploadType_UploadTypeImage) {
		img, _, err := image.Decode(bytes.NewReader(req.Data))
		if err == nil {
			height = int64(img.Bounds().Dy())
			width = int64(img.Bounds().Dx())
		}
	}
	return &api.UploadResp{
		Url:      url,
		Id:       id,
		Name:     req.Name,
		Width:    width,
		Height:   height,
		FileType: fileType,
	}, nil
}

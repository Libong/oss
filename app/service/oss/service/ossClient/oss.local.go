package ossClient

import (
	"bytes"
	"github.com/google/uuid"
	googleGrpc "google.golang.org/grpc"
	"image"
	"libong/common/context"
	"libong/common/server/grpc"
	"libong/oss/app/service/oss/api"
	bucketServiceApi "libong/oss/rpc/bucket/api"
	"path"
	"time"
)

type LocalOssConf struct {
	//AccessKey        string //用户名
	//AccessSecret     string //密码
	//BucketName       string
	//BucketServiceUrl string
	BucketService *grpc.Config
}

type LocalOss struct {
	Config        *LocalOssConf
	bucketService bucketServiceApi.BucketServiceClient
}

func newLocalOss(config *LocalOssConf) *LocalOss {
	var (
		bucketConn *googleGrpc.ClientConn
		err        error
	)
	o := &LocalOss{
		Config: config,
	}
	if bucketConn, err = grpc.NewConnection(config.BucketService); err != nil {
		panic(-1)
	}
	o.bucketService = bucketServiceApi.NewBucketServiceClient(bucketConn)
	return o
}
func (o *LocalOss) Upload(ctx context.Context, req *api.UploadReq) (*api.UploadResp, error) {
	id := uuid.New().String()
	appID := ctx.AppID()
	if appID == "" {
		appID = "temp"
	}
	////签名
	//param := "accessKey" + o.Config.AccessKey + "accessSecret" + o.Config.AccessSecret
	//md5Param := commonTool.GetMd5String(param)
	////设置请求文件
	//buf := new(bytes.Buffer)
	//w := multipart.NewWriter(buf)
	//if cff, err := w.CreateFormFile("file", req.Name); err == nil {
	//	cff.Write(req.Data)
	//}
	//w.Close()
	//request, err := http.NewRequest("POST", o.Config.BucketServiceUrl, buf)
	//if err != nil {
	//	return nil, err
	//}
	////设置请求头
	//request.Header.Set("Content-Type", w.FormDataContentType())
	//request.Header.Set("sign", md5Param)
	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	day := time.Now().Format("02")
	imgPath := appID + "/" + year + "/" + month + "/" + day + "/" + req.Name
	//request.Header.Set("prefix", imgPath)
	////请求
	//client := &http.Client{
	//	Timeout: 5 * time.Second,
	//}
	//res, err := client.Do(request)
	//if err != nil {
	//	return nil, err
	//}
	//if res.StatusCode != 200 {
	//	return nil, errors.UploadError
	//}
	//body, err := io.ReadAll(res.Body)
	//if err != nil {
	//	return nil, err
	//}
	//var resp customHttp.BaseResponse
	//err = json.Unmarshal(body, &resp)
	//if err != nil {
	//	return nil, err
	//}
	//if resp.Error != "" {
	//	log.Error(ctx, "Upload fail err:%s", resp.Error)
	//	return nil, errors.UploadError
	//}
	//var url string
	//if resp.Data != nil {
	//	url, _ = resp.Data.(string)
	//}
	addBucketObjectResp, err := o.bucketService.AddBucketObject(ctx, &bucketServiceApi.AddBucketObjectReq{
		Path:    imgPath,
		Data:    req.Data,
		NeedZip: false,
		IsDir:   false,
	})
	if err != nil {
		return nil, err
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
		Url:      addBucketObjectResp.Url,
		Id:       id,
		Name:     req.Name,
		Width:    width,
		Height:   height,
		FileType: fileType,
	}, nil
}

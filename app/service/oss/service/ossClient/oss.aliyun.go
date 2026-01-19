package ossClient

import (
	"bytes"
	"fmt"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"libong/common/context"
	"libong/common/log"
	commonRedis "libong/common/redis"
	"libong/common/snowflake"
	"libong/oss/app/service/oss/api"
	"path"
	"time"
)

const (
	fileUrlRedisKey               = "oss:fileUrl:%s:%s" //app_id
	defaultFileUrlExpiredDuration = 3 * time.Hour
)

type AliYunConf struct {
	BucketName string
}

// AliYunOss .
type AliYunOss struct {
	client *oss.Client
	config *AliYunConf
}

func newAliYunOss(config *AliYunConf) *AliYunOss {
	var provider credentials.CredentialsProvider
	//if env.Env == env.DeployEnvDev {
	//	provider = credentials.NewStaticCredentialsProvider("",
	//		"", "")
	//} else {
	provider = credentials.NewEcsRoleCredentialsProvider()
	//}
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(provider).
		WithRegion("cn-hangzhou").
		WithEndpoint("oss-cn-hangzhou-internal.aliyuncs.com")
	return &AliYunOss{
		client: oss.NewClient(cfg),
		config: config,
	}
}

func (o *AliYunOss) MakeFileUrl(ctx context.Context, keys []string) (map[string]string, error) {
	resp := make(map[string]string)
	if len(keys) == 0 {
		return resp, nil
	}
	var redisKeys []string
	for _, key := range keys {
		redisKeys = append(redisKeys, fmt.Sprintf(fileUrlRedisKey, ctx.AppID(), key))
	}
	kvMap, err := commonRedis.RedisClient.MGet(ctx, redisKeys)
	if err != nil {
		log.Error(ctx, "MakeFileUrl MGet err:%v", err)
	}
	for _, key := range keys {
		if kvMap[key] != nil {
			if url, ok := kvMap[key].(string); ok {
				resp[key] = url
				continue
			}
		}
		result, err := o.client.Presign(ctx, &oss.GetObjectRequest{
			Bucket: &o.config.BucketName,
			Key:    &key,
		}, oss.PresignExpires(defaultFileUrlExpiredDuration))
		log.Info(ctx, "MakeFileUrl PresignExpires %s", key)
		if err != nil {
			log.Error(ctx, "Upload MakeFileUrl err:%v", err)
			return nil, err
		}
		resp[key] = result.URL
	}
	return resp, nil
}

func (o *AliYunOss) Upload(ctx context.Context, req *api.UploadReq) (*api.UploadResp, error) {
	id := snowflake.SnowflakeWorker.NextID()
	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	fileNameWithSuffix := path.Base(req.Name)
	//获取文件的后缀(文件类型)
	fileType := path.Ext(fileNameWithSuffix)
	var fileName string
	if req.IsKeepOriginalName {
		fileName = req.Name
	} else {
		fileName = fmt.Sprintf("%d%s", id, fileType)
	}
	objectName := year + "/" + month + "/" + fileName
	reader := bytes.NewReader(req.Data)

	request := &oss.PutObjectRequest{
		Bucket: oss.Ptr(o.config.BucketName),
		Key:    oss.Ptr(objectName),
		Body:   reader,
	}

	_, err := o.client.PutObject(ctx, request)
	if err != nil {
		log.Error(ctx, "Upload PutObject err:%v", err)
		return nil, err
	}
	result, err := o.client.Presign(ctx, &oss.GetObjectRequest{
		Bucket: &o.config.BucketName,
		Key:    &objectName,
	}, oss.PresignExpires(defaultFileUrlExpiredDuration))
	if err != nil {
		log.Error(ctx, "Upload Presign err:%v", err)
		return nil, err
	}
	//缓存
	err = commonRedis.RedisClient.Set(ctx, fmt.Sprintf(fileUrlRedisKey, ctx.AppID(), objectName), result.URL,
		defaultFileUrlExpiredDuration)
	if err != nil {
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
		Id:       objectName,
		FileType: fileType,
		Name:     req.Name,
		Url:      result.URL,
		Width:    width,
		Height:   height,
	}, nil
}

package http

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"libong/common/context"
	"libong/common/server/http"
	ossServiceApi "libong/oss/app/service/oss/api"
	"libong/oss/errors"
	"strconv"
)

func upload(ctx *http.Context) error {
	file, fileHeader, err := ctx.Gin().Request.FormFile("file")
	if err != nil {
		return errors.ParamsIsInvalid
	}
	defer file.Close()
	fileBody, err := io.ReadAll(file)
	if err != nil {
		return errors.ParamsIsInvalid
	}
	strFileType := ctx.Gin().Query("fileType")
	fileType, err := strconv.ParseInt(strFileType, 10, 64)
	if err != nil {
		return errors.ParamsIsInvalid
	}

	var isKeepOriginalName bool
	strIsKeepOriginalName := ctx.Gin().Query("isKeepOriginalName")
	if strIsKeepOriginalName != "" {
		isKeepOriginalName, err = strconv.ParseBool(strIsKeepOriginalName)
		if err != nil {
			return errors.ParamsIsInvalid
		}
	}
	res, err := svc.Upload(context.FromHTTPContext(ctx), &ossServiceApi.UploadReq{
		Name:               fileHeader.Filename,
		IsKeepOriginalName: isKeepOriginalName,
		Type:               uint32(fileType),
		Data:               fileBody,
	})
	if err != nil {
		return err
	}
	ctx.ResponseData(res)
	return nil
}

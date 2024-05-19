package http

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"libong/common/context"
	"libong/common/server/http"
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
	res, err := svc.Upload(context.FromHTTPContext(ctx), fileBody, fileHeader.Filename, uint32(fileType))
	if err != nil {
		return err
	}
	ctx.ResponseData(res)
	return nil
}

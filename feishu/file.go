package feishu

import (
	"context"
	"io"

	"github.com/chyroc/lark"
)

func (bot *FeishuBot) UploadImage(imageType string, image io.Reader) string {
	res, _, err := bot.lark.File.UploadImage(context.Background(), &lark.UploadImageReq{
		ImageType: lark.ImageType(imageType),
		Image:     image,
	})
	if err != nil {
		panic(err)
	}
	return res.ImageKey
}

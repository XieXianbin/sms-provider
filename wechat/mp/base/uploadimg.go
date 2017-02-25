package base

import (
	"io"
	"os"
	"path/filepath"

	"github.com/XieXianbin/msg-provider/wechat/mp/core"
)

// UploadImage 上传图片到微信服务器, 返回的图片url给其他场景使用, 比如图文消息.
func UploadImage(clt *core.Client, imgFilePath string) (url string, err error) {
	file, err := os.Open(imgFilePath)
	if err != nil {
		return
	}
	defer file.Close()

	return UploadImageFromReader(clt, filepath.Base(imgFilePath), file)
}


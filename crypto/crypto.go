package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

//计算文件hash值,经过测试，txt、PNG、PDF、DOCX、MP4、IMG均可
//大小为385MB的视频文件大约3秒计算出结果，效率较高
func SHA256File(path string) (string, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return "", err
	}

	h := sha256.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func SHA256UploadFile(file multipart.File) (string, error) {
	defer file.Close()
	h := sha256.New()
	_, err := io.Copy(h, file)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func SHA256String(str string) (s string) {
	//使用sha256哈希函数
	h := sha256.New()
	h.Write([]byte(str))
	sum := h.Sum(nil)
	//由于是十六进制表示，因此需要转换
	return hex.EncodeToString(sum)
}

func SHA256Byte(bts []byte) (s string) {
	//使用sha256哈希函数
	h := sha256.New()
	h.Write(bts)
	sum := h.Sum(nil)
	//由于是十六进制表示，因此需要转换
	return hex.EncodeToString(sum)
}

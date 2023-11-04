package upload

import (
	"gin_init/global"
	"gin_init/pkg/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

// CheckSavePath：检查保存目录是否存在，通过调用 os.Stat 方法获取文件的描述信息 FileInfo，并调用 os.IsNotExist 方法进行判断，其原理是利用 os.Stat 方法所返回的 error 值与系统中所定义的 oserror.ErrNotExist 进行判断，以此达到校验效果。
// CheckPermission：检查文件权限是否足够，与 CheckSavePath 方法原理一致，是利用 oserror.ErrPermission 进行判断。
// CheckContainExt：检查文件后缀是否包含在约定的后缀配置项中，需要的是所上传的文件的后缀有可能是大写、小写、大小写等，因此我们需要调用 strings.ToUpper 方法统一转为大写（固定的格式）来进行匹配。
// CheckMaxSize：检查文件大小是否超出最大大小限制。
type FileType int

const TypeImage FileType = iota + 1

func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}

	}

	return false
}

func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}

	return false
}

func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}
func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}

	return nil
}

func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

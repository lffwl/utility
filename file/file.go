package file

import (
	uerror "192.168.0.209/wl/utility/error"
	"192.168.0.209/wl/utility/response"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

// IsExist 文件是否存在
// path 文件路径
// dir 判断目录是否存在
func IsExist(path string, dir ...bool) bool {
	dirExist := false
	if len(dir) > 0 {
		dirExist = dir[0]
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	if fileInfo.IsDir() {
		// 判断目录是否存在
		if dirExist {
			return true
		}
		return false
	}

	return true
}

// CreateMoreDir 创建多级目录
func CreateMoreDir(dirPath string) error {
	if !IsExist(dirPath, true) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		return err
	}
	return nil
}

// CreateFilePathDir 创建文件的目录
func CreateFilePathDir(filePath string) error {
	// 获取完整路径
	over, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}

	return CreateMoreDir(filepath.Dir(over))
}

// Copy 复制文件
func Copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// HttpOutput http文件输出
func HttpOutput(writer http.ResponseWriter, request *http.Request, path string) {
	if !IsExist(path) {
		response.Json(writer, uerror.HighErrorFileErrorCode, "file not exist")
		return
	}
	f, _ := os.Open(path)
	info, _ := f.Stat()
	defer f.Close()
	http.ServeContent(writer, request, info.Name(), info.ModTime(), f)
	return
}

// HttpDownload http文件下载
func HttpDownload(writer http.ResponseWriter, request *http.Request, path string) {
	if !IsExist(path) {
		response.Json(writer, uerror.HighErrorFileErrorCode, "file not exist")
		return
	}
	f, _ := os.Open(path)
	info, _ := f.Stat()
	defer f.Close()

	writer.Header().Set("Content-Type", "application/force-download")
	writer.Header().Set("Accept-Ranges", "bytes")
	writer.Header().Set("Content-Disposition", fmt.Sprintf(`attachment;filename=%s`, url.QueryEscape(filepath.Base(path))))

	http.ServeContent(writer, request, info.Name(), info.ModTime(), f)
	return
}

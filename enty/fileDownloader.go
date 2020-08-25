package enty

import (
	"errors"
	"fileDownload/utils"
	"fmt"
	"net/http"
	"strconv"
)

//文件下载器
type FileDownloader struct {
	FileSize       int
	Url            string
	OutputFileName string
	TotalPart      int
	OutPutDir      string
	DoneFilePart   []FilePart
}

//创建一个 request
func (f *FileDownloader) NewRequest(method string) (*http.Request, error) {
	r, err := http.NewRequest(method, f.Url, nil)
	if err != nil {
		return nil, err
	}
	r.Header.Set("User-Agent", "mojocn")
	return r, nil
}

//head 获取想要下载文件的基本信息(header),使用HTTP Head method
func (f *FileDownloader) Head() (int, error) {
	r, err := f.NewRequest("HEAD")
	if err != nil {
		return 0, err
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode > 299 {
		return 0, errors.New(fmt.Sprintf("can't progress,response is %v", resp.StatusCode))
	}

	//检查是否支持断点续传
	if resp.Header.Get("Accept-Range") != "bytes" {
		return 0, errors.New("服务器不支持断点续传")
	}
	f.OutputFileName = utils.ParseFileInfoFrom(resp)
	return strconv.Atoi(resp.Header.Get("Content-Length"))

}

func (f *FileDownloader) Run() error {
	fileTotalSize, err := f.Head()
	if err != nil {
		return err
	}
	f.FileSize = fileTotalSize

	jobs := make([]FilePart, f.TotalPart)
	eachSize := fileTotalSize / f.TotalPart
	for i := range jobs {
		jobs[i].index = i
		if i == 0 {
			jobs[i].From = 0
		} else {
		}
	}
	fmt.Println(eachSize)
	return err
}

package enty

//文件下载器
type FileDownloader struct {
	FileSize       string
	Url            string
	OutputFileName string
	TotalPart      int
	OutPutDir      string
	DoneFilePart   []FilePart
}

//文件分片
type FilePart struct {
	index int
	From  int
	To    int
	Data  []byte
}

package enty

//文件分片
type FilePart struct {
	index int
	From  int
	To    int
	Data  []byte
}

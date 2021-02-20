package util

import (
	"bytes"
	"io"
	"os"
)

// SaveDataToFile 保存数据到文件
func SaveDataToFile(data []byte, fileName string) {
	out, _ := os.Create(fileName)
	io.Copy(out, bytes.NewReader(data))
}

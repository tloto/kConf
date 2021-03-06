package kConf
import (
	"bytes"
	"os"
	"bufio"
	"io"
)

func GetFileContentWithPath(path string) ([]byte,error) {

	f,err := os.Open(path)
	if err!=nil {
		return nil,err
	}
	defer f.Close()

	var config_DATA []byte

	buffioreader := bufio.NewReader(f)
	for{
		b := make([]byte,1024)
		_,er :=buffioreader.Read(b)
		if er == io.EOF {
		   break
		}
		config_DATA=append(config_DATA, b...)
	}

	index := bytes.IndexByte(config_DATA,0)
	if index < 1 {
	   index = 0
	}
	resultData := config_DATA[:index]
	return resultData,nil
}

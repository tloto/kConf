package kConf

import (
	"fmt"
	"github.com/tidwall/gjson"
)

var config_DATA_BYTE []byte
/*
*InitDir
*path ： file path
*/
func SetFiePath(path string)  {

	databyte,err := GetFileContentWithPath(path)
	if err != nil {
		 fmt.Println("config file open error: ",err.Error())
		 return
	}
	config_DATA_BYTE = databyte
}


func GetConfString(jsonpath string)string{
	result := gjson.GetBytes(config_DATA_BYTE,jsonpath)
	return result.String()
}

func GetConfBool(jsonpath string)bool{
	result := gjson.GetBytes(config_DATA_BYTE,jsonpath)
	return result.Bool()
}

func GetConfint(jsonpath string)int64{
	result := gjson.GetBytes(config_DATA_BYTE,jsonpath)
	return result.Int()
}

func GetConfFloat(jsonpath string)float64{
	result := gjson.GetBytes(config_DATA_BYTE,jsonpath)
	return result.Float()
}
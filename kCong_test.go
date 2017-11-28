package kConf

import "testing"

func init() {
	SetFiePath("./test.json")
}

func TestGetConfString(t *testing.T) {

	str := GetConfString("compilerOptions.module")
	if str == "test" {
		t.Log("test success")
	}else {
		t.Error("test false")
	}
}
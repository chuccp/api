package vmess

import (
	"bytes"
	"encoding/base64"
	"io"
	"net/http"
	"testing"
)

func TestName(t *testing.T) {
	resp, err := http.Get("https://call.chenxing.gq/link/zjNcC6g3mOj3oZOo?sub=3&extend=1")
	if err!=nil{
		t.Log("===============",err)
	}else{

		data,_:=io.ReadAll(resp.Body)

		data,_=base64.StdEncoding.DecodeString(string(data))



		buff:=bytes.NewBuffer(data)
		buff.WriteString(moren)

		println(buff.String())

	}
}

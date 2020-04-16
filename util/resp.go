package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type H struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
	Rows interface{}`json:"rows,omitempty"`
	Totak interface{}`json:"totak,omitempty"`
}

func RespFail(w http.ResponseWriter, mag string)  {
	Response(w,-1,nil,mag)
}
func RespOk(w http.ResponseWriter, data interface{}, mag string)  {
	Response(w,0,data,mag)
	
}


func Response(w http.ResponseWriter, code int, data interface{}, mag string) {

	w.Header().Set("Content-Type", "application/json")
	//設置狀態碼
	w.WriteHeader(http.StatusOK)
	//w.Write([]byte(str))
	h := H{
		Code: code,
		Msg:  mag,
		Data: data,
	}
	//struct轉json
	marshal, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	//輸出
	w.Write(marshal)
}


func RespOkList(w http.ResponseWriter, lists interface{}, total interface{})  {
	ResponseList(w,0,lists,total)

}
func ResponseList(w http.ResponseWriter, code int, data interface{}, totak interface{}) {

	w.Header().Set("Content-Type", "application/json")
	//設置狀態碼
	w.WriteHeader(http.StatusOK)
	//w.Write([]byte(str))
	h := H{
		Code: code,
		Rows:  data,
		Totak: totak,
	}
	//struct轉json
	marshal, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	//輸出
	w.Write(marshal)
}
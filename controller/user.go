package controller

import (
	"app/v1/im/model"
	"app/v1/im/service"
	"app/v1/im/util"
	"fmt"
	"math/rand"
	"net/http"
)


func UserLogin(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	mobile := request.Form.Get("mobile")
	pwd := request.Form.Get("passwd")
	//fmt.Println(request.Form.Get("pwd"))

	user, err := userService.Login(mobile, pwd)
	if err!=nil {
		util.RespFail(writer,err.Error())
	}else {
		util.RespOk(writer,user,"")
	}


}

var userService service.UserService
func UserRegister(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	mobile := r.Form.Get("mobile")
	pwd := r.Form.Get("passwd")
	nickname :=fmt.Sprintf("user%06d",rand.Int31n(2000))
	//fmt.Println(pwd)
	avatar:=""
	sex :=model.SEX_UNKNOW
	userinfo, err := userService.Register(mobile, pwd, nickname, avatar, sex)
	if err!=nil {
		util.RespFail(w,err.Error())

	}else {
		util.RespOk(w,userinfo,"")
	}

}

package controller

import (
	"gitapp.com/v1/im/args"
	"gitapp.com/v1/im/model"
	"gitapp.com/v1/im/service"
	"gitapp.com/v1/im/util"
	"fmt"
	"net/http"
)

var contactService service.ContactService

//添加好友
func Addfriend( w http.ResponseWriter,r *http.Request)  {
	var arg args.ContactArg
	util.Bind(r,&arg)
	err := contactService.AddFriend(arg.Userid, arg.Dstid)
	if err!=nil {
		util.RespFail(w,err.Error())
	}else {
		util.RespOk(w,nil,"添加成功")
	}

}
//好友列表
func LoadFriend( w http.ResponseWriter,r *http.Request)  {
	var arg args.ContactArg
	util.Bind(r, &arg)
	//fmt.Println(arg.Userid)
	friend := contactService.SearchFriend(arg.Userid)
	fmt.Println(friend)
	util.RespOkList(w, friend,len(friend))

}


//創建群

func CreateCommunity( w http.ResponseWriter, r *http.Request){
	var com model.Community
	util.Bind(r,&com)
	community, err := contactService.CreateCommunity(com)
	if err!=nil {
		util.RespFail(w,err.Error())
	}else {
		util.RespOk(w,community,"")
	}
}
//群列表

func LoadCommunity(w http.ResponseWriter,r *http.Request) {
	var arg args.ContactArg

	util.Bind(r,&arg)
	comunitys := contactService.SearchComunity(arg.Userid)
	util.RespOkList(w,comunitys,len(comunitys))
	
}

//加群
func JoinCommunity(w http.ResponseWriter, r *http.Request)  {
	var arg args.ContactArg
	util.Bind(r,&arg)
	err := contactService.JoinCommunity(arg.Userid, arg.Dstid)
	AddGroupId(arg.Userid,arg.Dstid)
	if err!=nil{
		util.RespFail(w,err.Error())
	}else {
		util.RespOk(w,nil,"")
	}

}
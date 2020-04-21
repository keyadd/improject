package controller

import (
	"fmt"
	"app/v1/im/util"
	"golang.org/x/exp/rand"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func Upload( w http.ResponseWriter,r *http.Request)  {
	UploadLocal(w,r)
	
}
func init()  {
	os.MkdirAll("./public",os.ModePerm)

}
func UploadLocal( w http.ResponseWriter,r *http.Request)  {


	//url 格式
	sfile, header, err := r.FormFile("file")
	if err!=nil {
		util.RespFail(w,err.Error())
		return
	}

	suffix :=".png"
	filename := header.Filename
	tmp := strings.Split(filename, ".")
	if  len(tmp)>1{
		suffix = "."+tmp[len(tmp)-1]
	}

	filetype := r.FormValue("filetype")
	if  len(filetype)>0{
		suffix = filetype
	}

	filename = fmt.Sprintf("%d%04d%s",time.Now().Unix(),rand.Int31(),suffix)
	dstfile,err :=os.Create("./public/"+filename)
	if err!=nil {

		util.RespFail(w,err.Error())
		return
	}
	_, err = io.Copy(dstfile, sfile)
	if err!=nil {

		util.RespFail(w,err.Error())
		return
	}
	url :="/public/"+filename
	util.RespOk(w,url,"")

}
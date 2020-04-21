package main

import (
	"app/v1/im/controller"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)





func RegisterView()  {
	files, err := template.ParseGlob("view/**/*")
	if err!=nil {
		log.Fatal(err.Error())
	}
	for _,v:=range files.Templates(){
		tplname := v.Name()
		http.HandleFunc(tplname, func(writer http.ResponseWriter, request *http.Request) {
			files.ExecuteTemplate(writer,tplname,nil)
		})

	}
}



func main() {
	http.HandleFunc("/user/login", controller.UserLogin)
	http.HandleFunc("/user/register", controller.UserRegister)
	http.HandleFunc("/contact/addfriend",controller.Addfriend)
	http.HandleFunc("/contact/friend",controller.LoadFriend)
	http.HandleFunc("/contact/createcommunity",controller.CreateCommunity)
	http.HandleFunc("/contact/loadcommunity", controller.LoadCommunity)
	http.HandleFunc("/contact/joincommunity", controller.JoinCommunity)

	http.HandleFunc("/chat",controller.Chat)
	http.HandleFunc("/attach/upload",controller.Upload)


	http.Handle("/public/", http.FileServer(http.Dir(".")))
	http.Handle("/asset/",http.FileServer(http.Dir(".")))
	RegisterView()
	http.ListenAndServe("0.0.0.0:8089", nil)

}

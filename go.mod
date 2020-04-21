module app/v1/im

go 1.14

require (
    //app/v1/im/module v1
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-xorm/xorm v0.7.9
	github.com/gorilla/websocket v1.4.2
	github.com/kr/pretty v0.2.0 // indirect
	golang.org/x/exp v0.0.0-20190121172915-509febef88a4
	gopkg.in/fatih/set.v0 v0.2.1
)

//replace app/v1/im/module=>.module
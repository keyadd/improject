package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"log"
	"net/http"
	"strconv"
	"sync"
)

const (
	CMD_SINGLE_MSG = 10
	CMD_ROOM_MSG   = 11
	CMD_HEART      = 0
)

//读写锁
var rwlocker sync.RWMutex

type Message struct {
	Id      int64  `json:"id,omitempty" form:"id"` //消息ID
	Userid  int64  `json:"userid,omitempty" form:"userid"` //谁发的
	Cmd     int    `json:"cmd,omitempty" form:"cmd"` //群聊还是私聊
	Dstid   int64  `json:"dstid,omitempty" form:"dstid"`//对端用户ID/群ID
	Media   int    `json:"media,omitempty" form:"media"` //消息按照什么样式展示
	Content string `json:"content,omitempty" form:"content"` //消息的内容
	Pic     string `json:"pic,omitempty" form:"pic"` //预览图片
	Url     string `json:"url,omitempty" form:"url"` //服务的URL
	Memo    string `json:"memo,omitempty" form:"memo"` //简单描述
	Amount  int    `json:"amount,omitempty" form:"amount"` //其他和数字相关的
}


type Node struct {
	Conn *websocket.Conn

	DataQueue chan []byte
	GroupSets set.Interface
}

var clientMap map[int64]*Node = make(map[int64]*Node,0)

var rwlockmap sync.RWMutex
func Chat(w http.ResponseWriter, r *http.Request)  {
	query := r.URL.Query()
	id := query.Get("id")
	token := query.Get("token")
	userId, _ := strconv.ParseInt(id, 10, 64)
	isvalida := checkToken(userId, token)

	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(request *http.Request) bool {
			return isvalida
		},
	}).Upgrade(w, r,nil)
	if err!=nil {
		log.Println(err.Error())
		return

	}
	node := &Node{Conn: conn, DataQueue: make(chan []byte,100), GroupSets: set.New(set.ThreadSafe),}


	comIds:=contactService.SearchComunityIds(userId)
	for _,v:=range comIds {
		node.GroupSets.Add(v)

	}

	rwlockmap.Lock()
	clientMap[userId]=node
	rwlockmap.Unlock()
	go sendproc(node)

	go recvproc(node)
	//sendMsg(userId,[]byte("hello word"))


}
//接收協程
func sendproc(node *Node)  {
	for {
		select {
		case data:=<-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err!=nil {
				log.Println(err.Error())
				return
			}
		}
		
	}
	
}



//todo 添加新的群ID到用户的groupset中
func AddGroupId(userId,gid int64){
	//取得node
	rwlocker.Lock()
	node,ok := clientMap[userId]
	if ok{
		node.GroupSets.Add(gid)
	}
	//clientMap[userId] = node
	rwlocker.Unlock()
	//添加gid到set
}



//推送協程
func recvproc(node *Node)  {
	for  {
		_, data, err := node.Conn.ReadMessage()
		if err!=nil {
			log.Println(err.Error())
			return
		}
		dispatch(data)
		fmt.Printf("recv<=%s",data)

	}
	
}
func dispatch(data []byte)  {
	msg:= Message{}
	err := json.Unmarshal(data, &msg)
	if err!=nil {
		log.Println(err.Error())
		return
	}
	switch msg.Cmd {
	case CMD_SINGLE_MSG://用戶私聊
		sendMsg(msg.Dstid,data)
	case CMD_ROOM_MSG: //群聊
		for _,v:=range clientMap{
			if v.GroupSets.Has(msg.Dstid) {
				v.DataQueue<-data
				
			}
			
		}

	case CMD_HEART:

		
	}
}
//發送消息
func sendMsg(userid int64,msg []byte)  {
	rwlockmap.RLock()
	node,ok := clientMap[userid]
	rwlockmap.RUnlock()
	if ok {
		node.DataQueue<-msg

	}

}

func checkToken(userId int64,token string) bool {
	userinfo := userService.Find(userId)
	return userinfo.Token ==token

}

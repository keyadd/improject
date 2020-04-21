package service

import (
	"errors"
	"fmt"
	"app/v1/im/model"
	"app/v1/im/util"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"time"
)

type UserService struct {
}

func (s *UserService) Register(mobile, pwd, nickname, avatar, sex string) (user model.User, err error) {
	tmp := model.User{}
	//var tmp model.User
	//fmt.Println(tmp)
	_, err = DbEngin.Where("Mobile=?", mobile).Get(&tmp)
	//fmt.Println(err)

	if err != nil {
		return tmp, err
	}
	if tmp.Id > 0 {
		return tmp, errors.New("手機號已經註冊")
	}
	tmp.Mobile = mobile
	tmp.Avatar = avatar
	tmp.Sex = sex
	tmp.Nickname = nickname
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	tmp.Passwd = util.MakePasswd(pwd, "123456")
	tmp.Createat = time.Now()
	tmp.Token = fmt.Sprintf("%08d", rand.Int31())
	//if &tmp!=nil {
	_, err = DbEngin.InsertOne(&tmp)
	//}

	return tmp, err

	//return user,err
}

//登陸
func (s *UserService) Login(mobile, pwd string) (user model.User, err error) {
	tmp := model.User{}
	_, err = DbEngin.Where("Mobile=?", mobile).Get(&tmp)
	if tmp.Id==0 {
		return  tmp,errors.New("用戶不存在")

	}
	fmt.Println(tmp)
	if !util.ValidatePasswd(pwd,"123456",tmp.Passwd) {
		return  tmp,errors.New("密碼不正確")
	}
	//刷新token
	str := fmt.Sprintf("%d", time.Now().Unix())
	token := util.MD5Encode(str)
	tmp.Token=token

	DbEngin.ID(tmp.Id).Cols("token").Update(&tmp)
	return tmp, nil

}


//查詢用戶token
func (s *UserService)Find(userId int64)(user model.User) {
	tmp:=model.User{}
	DbEngin.ID(userId).Get(&tmp)
	return tmp
}

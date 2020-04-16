package service

import (
	"errors"
	"fmt"
	"gitapp.com/v1/im/model"
	"time"
)

type ContactService struct {

}

//添加好友

func (service *ContactService)AddFriend(userid,dstid int64)error{

	if userid == dstid {
		return errors.New("不能添加自己為好友")
	}
	tmp:=model.Contact{}

	DbEngin.Where("ownerid =?",userid).And("dstid=?",dstid).And("cate=?",model.CONCAT_CATE_USER).Get(&tmp)
	if tmp.Id>0 {
		return errors.New("該用戶已經被添加過了")
	}
	session := DbEngin.NewSession()
	session.Begin()
	_, err2 := session.InsertOne(model.Contact{Ownerid: userid, Dstobj: dstid, Cate: model.CONCAT_CATE_USER, Createat: time.Now()})
	_, err3 := session.InsertOne(model.Contact{Ownerid: dstid, Dstobj: userid, Cate: model.CONCAT_CATE_USER, Createat: time.Now()})

	if err2==nil && err3 ==nil {
		session.Commit()
		return nil
	}else {
		session.Rollback()
		if err2!=nil {
			return err2
		}else {
			return err3
		}
	}

}
func (service *ContactService)SearchComunityIds(userId int64)  (comIds []int64){
	contacts := make([]model.Contact, 0)

	comIds = make([]int64, 0)
	DbEngin.Where("Ownerid = ? and cate = ?", userId, model.CONCAT_CATE_COMUNITY).Find(&contacts)
	for _,v :=range contacts{
		comIds = append(comIds, v.Dstobj)
	}
	fmt.Println(comIds)
	return comIds
}

//搜索好友
func (service *ContactService)SearchFriend (userId int64)([]model.User){
	contacts := make([]model.Contact, 0)

	dataId := make([]int64, 0)
	 DbEngin.Where("Ownerid = ? and cate = ?", userId, model.CONCAT_CATE_USER).Find(&contacts)
	for _,v :=range contacts{
		dataId = append(dataId, v.Dstobj)
	}
	fmt.Println(dataId)
	users := make([]model.User,0)
	if len(dataId)==0 {
		return users
	}
	DbEngin.In("id",dataId).Find(&users)
	return users

}
//建群

func (service *ContactService)CreateCommunity(com model.Community)(ret model.Community,err error) {

	if len(com.Name)==0 {
		err := errors.New("沒有群名稱")
		return ret,err
	}
	if com.Ownerid==0 {
		err := errors.New("請先登陸")
		return ret,err
	}
	community := model.Community{Ownerid: com.Ownerid}
	num, err := DbEngin.Count(&community)
	if num>5{
		errors.New("一個用戶只能創建五個群")
		return community,err
	}else{
		com.Createat=time.Now()
		session := DbEngin.NewSession()
		session.Begin()
		_, err := session.InsertOne(&com)
		if err != nil{
			session.Rollback();
			return community,err
		}
		_, err = session.InsertOne(model.Contact{Ownerid: com.Ownerid, Dstobj: com.Id, Cate: model.CONCAT_CATE_COMUNITY, Createat: time.Now()})
		if err!=nil {
			session.Rollback();
		}else {
			session.Commit()
		}
		return community,err
	}
	return 
}
//群列表

func (service *ContactService) SearchComunity(userId int64) ([]model.Community){
	conconts := make([]model.Contact,0)
	comIds :=make([]int64,0)

	DbEngin.Where("ownerid = ? and cate = ?",userId,model.CONCAT_CATE_COMUNITY).Find(&conconts)
	for _,v := range conconts{
		comIds = append(comIds,v.Dstobj);
	}
	coms := make([]model.Community,0)
	if len(comIds)== 0{
		return coms
	}
	DbEngin.In("id",comIds).Find(&coms)
	return coms
}


//加群
func (service *ContactService)JoinCommunity(userId,comId int64)error{
	contact := model.Contact{Ownerid: userId, Dstobj: comId, Cate: model.CONCAT_CATE_COMUNITY}
	DbEngin.Get(&contact)
	if(contact.Id==0){
		contact.Createat = time.Now()
		_,err := DbEngin.InsertOne(contact)
		return err
	}else{
		return nil
	}
}



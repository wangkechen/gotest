package main

import (
	"encoding/hex"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"test/inittest"
	"time"
)

type User struct {
	Id bson.ObjectId `bson:"_id"`
	Username string `bson:"name"`
	Pass string `bson:"pass"`
	Regtime int64 `bson:"regtime"`
	Interests []string `bson:"interests"`
}

const URL string = "127.0.0.1:27017"

var c *mgo.Collection
var session *mgo.Session

func (user User) ToString() string {
	return fmt.Sprintf("%#v",user)
}
func init()  {
	fmt.Println("main package")
	session,_ = mgo.Dial(URL)
	//切换到数据库
	db := session.DB("blog")
	//切换到collection
	c = db.C("mgotest")
}

func add()  {
	/*stu1 := User{
		Id: bson.NewObjectId(),
		Username: "stu1_name",
		Pass: "stu1_pass",
		Regtime: time.Now().Unix(),
		Interests: []string{"象棋","游泳","跑步"},
	}*/
	stu1 := new(User)
	stu1.Id = bson.NewObjectId()
	stu1.Username = "stu1_name"
	stu1.Pass = "stu1_pass"
	stu1.Regtime = time.Now().Unix()
	stu1.Interests = []string{"象棋", "游泳", "跑步"}
	err := c.Insert(stu1)
	fmt.Println(stu1.Id)
	fmt.Println(hex.EncodeToString([]byte(stu1.Id)))
	if err == nil {
		fmt.Println("插入成功")
	} else {
		fmt.Println(err.Error())
		return
	}
}

func find()  {
	var users []User
	c.Find(bson.M{"name": "stu1_name"}).All(&users)
	for _,value := range users{
		fmt.Println(value)
		//fmt.Println(value.ToString())
	}
	//根据objectId进行查询
	idStr := "614d3569a602121c9cc8fe01"
	objectId := bson.ObjectIdHex(idStr)
	fmt.Println(objectId)
	user := new(User)
	c.Find(bson.M{"_id":objectId}).One(&user)
	fmt.Println(user.ToString())
}
//根据id进行修改
func update()  {
	interests := []string{"象棋","游泳","铅球"}
	err := c.Update(bson.M{"_id": bson.ObjectIdHex("614d3231a602122e240faf29")},bson.M{"$set":bson.M{
		"name": "再次修改后的name",
		"interests": interests,
	}})
	if err != nil {
		fmt.Println("修改失败")
	} else {
		fmt.Println("修改成功")
	}
}
func del()  {
	err := c.Remove(bson.M{"_id":bson.ObjectIdHex("614d38f0a602123b2c3f6341")})
	if err != nil {
		fmt.Println("删除失败" + err.Error())
	} else {
		fmt.Println("删除成功")
	}
}
func main11() {
	inittest.Wang()
	//add()
	//find()
	//update()
	//del()
}

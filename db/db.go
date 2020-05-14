package db

import (
	"GiveGift/data"
	"encoding/json"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// 为了解决包循环引用问题
type Userer interface {
	GetId() int32
}

type Anchorer interface {
	GetId() int32
}

const URL string = "127.0.0.1:27017"

func mydb() *mgo.Database {
	session, err := mgo.Dial(URL)
	if err != nil {
		fmt.Println("数据库连接失败")
		return nil
	}
	return session.DB("mydb")
}

// 存入日志表
func AddToDailyDB(js []byte) {
	// 连接表dailyOfGifts
	con := mydb().C("dailyOfGifts")

	var record data.DailyOfGifts
	// defer session.Close()
	json.Unmarshal(js, &record)
	fmt.Println(record)

	if err := con.Insert(record); err == nil {
		fmt.Println("插入日志表成功")
	} else {
		fmt.Println("插入日志表失败", err.Error())
		defer panic(err)
	}
}

// 查询计算本次礼物的总价值和id
func CalculateValue(js []byte) (int32, int32, error)  {
	// 连接表dailyOfGifts
	con := mydb().C("giftMessage")

	// 解析json
	var record data.DailyOfGifts
	json.Unmarshal(js, &record)

	// 通过type找value
	giftMessage := &data.GiftMessage{}
	// M：一张无序的map
	err := con.Find(bson.M{"type": record.GiftType}).One(giftMessage)
	if err != nil {
		fmt.Println("查找礼物总价失败", err.Error())
		return -1, -1, err
	}
	fmt.Println("查找礼物总价成功")
	result := giftMessage.Value * record.GiftNumbers
	//fmt.Println("总价值是", giftType, giftMessage.Value, quantity, result)
	return record.AnchorId, result, nil
}

func Findecords(anchorId int32) ([]byte, error) {
	// 连接表dailyOfGifts
	con := mydb().C("dailyOfGifts")

	// 通过anchorid找所有的record
	var records []data.DailyOfGifts
	m := []bson.M{
		{"$match": bson.M{"anchorId": anchorId}},
		{"$sort": bson.M{"time": -1}},		// 倒叙查询
	}
	err := con.Pipe(m).All(&records)
	if err != nil {
		fmt.Println("查找主播流水失败", err.Error())
		return nil, err
	}
	js, err := json.Marshal(records)
	fmt.Printf("主播流水的json是%s", js)
	return js, err
}
//
////查询
//func find() {
//	//    defer session.Close()
//	var users []User
//	//    c.Find(nil).All(&users)
//	c.Find(bson.M{"name": "stu1_name"}).All(&users)
//	for _, value := range users {
//		fmt.Println(value.ToString())
//	}
//	//根据ObjectId进行查询
//	idStr := "5eb54716f6f47e5a44cde60a"
//	objectId := bson.ObjectIdHex(idStr)
//	user := new(User)
//	c.Find(bson.M{"_id": objectId}).One(user)
//	fmt.Println(user)
//}
//
////根据id进行修改
//func update() {
//	interests := []string{"象棋", "游泳", "跑步"}
//	err := c.Update(bson.M{"_id": bson.ObjectIdHex("5eb54716f6f47e5a44cde60a")}, bson.M{"$set": bson.M{
//		"name":      "修改后的name",
//		"pass":      "修改后的pass",
//		"regtime":   time.Now().Unix(),
//		"interests": interests,
//	}})
//	if err != nil {
//		fmt.Println("修改失败")
//	} else {
//		fmt.Println("修改成功")
//	}
//}
//
////删除
//func del() {
//	err := c.Remove(bson.M{"_id": bson.ObjectIdHex("577fb2d1cde67307e819133d")})
//	if err != nil {
//		fmt.Println("删除失败" + err.Error())
//	} else {
//		fmt.Println("删除成功")
//	}
//}
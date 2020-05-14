package db
// 使用了mongodb的另一种包mongo-driver，为了使用连接池实现
import (
	"GiveGift/data"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)
type Database struct {
	Mongo  * mongo.Client
}
var DB *Database

//初始化
func InitDB() {
	fmt.Println("Init...")
	DB = &Database{
		Mongo: SetConnect(),
	}
}
//
func SetConnect() *mongo.Client {
	fmt.Println("setconnect...")
	ctx , cancel :=context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() //养成良好的习惯，在调用WithTimeout之后defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017").SetMaxPoolSize(9))
	if err != nil {
		fmt.Println("clent为nil")
		log.Print(err)
		return nil
	}
	return client
}

// 利用连接池的 查找所有流水
func Findecords2(anchorId int32) ([]byte, error) {
	client := DB.Mongo
	collection := client.Database("mydb").Collection("dailyOfGifts")
	sort := bson.M{"time": -1}
	filter := bson.M{"anchorId": anchorId}
	findOptions := options.Find().SetSort(sort)
	result,err := collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		fmt.Println("查找主播流水失败", err.Error())
		return nil, err
	}
	var records []data.DailyOfGifts
	// Iterating through the cursor allows us to decode documents one at a time
	for result.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem *data.DailyOfGifts
		err := result.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		records = append(records, *elem)
	}
	js, err := json.Marshal(records)
	fmt.Printf("主播流水的json是%s\r\n", js)
	return js, err
}
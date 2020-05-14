package webService

import (
	"GiveGift/data"
	"GiveGift/giveGift"
	"GiveGift/rank"
	"GiveGift/record"
	"encoding/json"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)
type HandleFnc func(http.ResponseWriter, *http.Request)

const giveGiftForm = `
	<html><body>
		<h1>give gift</h1>
		<form action="#" method="post" name="bar">
			my id:<input type="text" name="userId" />
			anchor id:<input type="text" name="anchorId" /></br>
			gift number:<input type="text" name="giftQuantity" /></br>
			gift type:<input type="text" name="giftType" /></br>
			<input type="submit" value="submit"/>
		</form>
	</body></html>
`

const findRecordsForm = `
	<html><body>
		<h1>流水查询</h1>
		<form action="#" method="get" name="bar">
			需要查询的主播id:<input type="text" name="anchorId" />
			<input type="submit" value="submit"/>
		</form>
	</body></html>	
`

func giveGiftServer(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		fmt.Fprintf(response, giveGiftForm)
	case "POST":
		response.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(response, "thanks for your gift")

		userId,_ := strconv.ParseInt(request.FormValue("userId"), 10, 32)
		anchorId,_ := strconv.ParseInt(request.FormValue("anchorId"), 10, 32)
		quantity,_ := strconv.Atoi(request.FormValue("giftQuantity"))
		giftType,_ := strconv.Atoi(request.FormValue("giftType"))

		record := data.DailyOfGifts{
			AnchorId:int32(anchorId),
			UserId:int32(userId),
			GiftType:int32(giftType),
			GiftNumbers:int32(quantity),
			Time:time.Now(),
			Id:bson.NewObjectId(),
		}
		if js, err := json.Marshal(record); err == nil {
			giveGift.GiveGift(js)
			return
		}
		fmt.Println("json生成失败")
	}
}

func showRankServer(response http.ResponseWriter, request *http.Request) {
	rank, _ := rank.ShowRank()
	response.Header().Set("Content-Type", "application/json")
	switch request.Method {
	case "GET":
		fmt.Fprintf(response, "%s", rank)
	}
}

func findRecords(response http.ResponseWriter, request *http.Request) {
	if request.FormValue("anchorId") == "" {
		fmt.Fprintf(response, findRecordsForm)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	anchorId,_ := strconv.Atoi(request.FormValue("anchorId"))
	fmt.Println("需要查询的主播id为", anchorId, "结果为：")
	records,err := record.FindRecords(int32(anchorId))
	if err != nil {
		io.WriteString(response, "查询失败")
		return
	}
	fmt.Fprintf(response, "%s", records)
}

func StartWebService() {
	http.HandleFunc("/giveGift", logPanics(giveGiftServer))
	http.HandleFunc("/showRank", logPanics(showRankServer))
	http.HandleFunc("/findRecords", logPanics(findRecords))
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}

// 使用闭包的错误处理模式
func logPanics(function HandleFnc) HandleFnc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}
		}()
		function(writer, request)
	}
}
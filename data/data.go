package data

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type DailyOfGifts struct {
	Id        	bson.ObjectId `bson:"_id" json:"_id"`
	AnchorId  	int32      	  `bson:"anchorId" json:"anchorId"`
	UserId    	int32     	  `bson:"userId" json:"userId"`
	Time 	  	time.Time	  `bson:"time" json:"time"`
	GiftNumbers int32		  `bson:"giftNumbers" json:"giftNumbers"`
	GiftType	int32		  `bson:"giftType" json:"giftType"`
}

type GiftMessage struct {
	Id        	bson.ObjectId `bson:"_id"`
	Type		int32		  `bson:"type"`
	Value		int32		  `bson:"value"`
}

type AnchorMessage struct {
	Id        	bson.ObjectId `bson:"_id"`
	AnchorId	int32		  `bson:"anchorId"`
	AnchorName  string	      `bson:"anchorName"`
}

type Rank struct {
	AnchorId 	string		  `json:"anchorId"`
	Value		string		  `json:"value"`
}
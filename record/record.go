package record

import "GiveGift/db"

// 提供一个收礼流水查询接口，查询主播的收礼流水记录，按时间从近到远排序；
func FindRecords (anchorId int32) ([]byte, error) {
	//return db.Findecords(anchorId)
	return db.Findecords2(anchorId)
}

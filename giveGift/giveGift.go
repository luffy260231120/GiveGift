package giveGift


import "GiveGift/db"

func GiveGift(js []byte) error {
	// 把quantity和type存入日志表中
	db.AddToDailyDB(js)
	// 把礼物总价值存入缓存
	if id, result, err := db.CalculateValue(js); err == nil {
		err := db.AddValuesToRedis(id, result)
		return err
	}
	return nil
}
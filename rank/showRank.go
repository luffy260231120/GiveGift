package rank
import "GiveGift/db"
// 提供一个排行榜查询接口，根据主播收礼价值数从大到小排序； 主播之间的排序
func ShowRank() ([]byte, error) {
	return db.GetRank()
}

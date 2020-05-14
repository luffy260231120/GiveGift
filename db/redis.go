package db

import (
	"GiveGift/data"
	"encoding/json"

	"fmt"
	"github.com/garyburd/redigo/redis"
	"reflect"
)
var RedisExpire = 3600 //缓存有效期

// 把anchorid和总价值放入redis中。
// 参数需要id和总价值。
func AddValuesToRedis(anchorId, value int32) error {
	fmt.Println("in redis, anchorid = ", anchorId, "value=", value)

	c := RedisClient.Get()
	// 用完后将连接放回连接池
	defer c.Close()

	// sorted set
	num, err := c.Do("zincrby", "rank", value, anchorId)
	fmt.Println(num, reflect.TypeOf(num))
	if err != nil {
		fmt.Println("插入redis失败", err.Error())
		return err
	} else {
		fmt.Println("成功插入redis")
	}

	return nil
}


func GetRank() ([]byte, error)  {
	fmt.Println("开始从redis中拿出排行榜...")
/*	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return nil, err
	}
	defer c.Close()*/

	c := RedisClient.Get()
	// 用完后将连接放回连接池
	defer c.Close()

	res, err := redis.Values(c.Do("ZREVRANGE", "rank", 0, -1, "withscores"))
	var rank []data.Rank
	if err != nil {
		fmt.Println("zrange in rank failed", err.Error())
	} else {
		// redis取出后放入结构体中
		for i := 0; i < len(res); i += 2 {
			var record data.Rank
			if id, ok := res[i].([]byte); ok {
				record.AnchorId = fmt.Sprintf("%s", id)
			}
			if value, ok := res[i+1].([]byte); ok {
				record.Value = fmt.Sprintf("%s", value)
			}

			rank = append(rank, record)
		}
	}
	//fmt.Println("rank", rank)

	// 转化为json格式
	if bytes, err := json.Marshal(rank); err == nil {
		fmt.Printf("JSON format: %s", bytes)
		return bytes, nil
	} else {
		fmt.Println("json生成失败")
		return nil, err
	}
	return nil, nil
}
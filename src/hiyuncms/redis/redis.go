package redis

import (
	"github.com/garyburd/redigo/redis"
	"hiyuncms/config"
	"log"
	"hiyuncms/controllers/frontend"
	"fmt"
	"encoding/json"
)

var redisConn redis.Conn

func init() {
	var err error
	redisConn, err = redis.Dial("tcp", config.GetValue("redis.address") )
	if  err != nil {
		log.Printf("初始化redis connection 出错！\n")
	}
}

func SetToken(token string, sessionInfo * frontend.UserSession)  {
	bytes, _ := json.Marshal( sessionInfo )
	redisConn.Do(fmt.Sprintf("SET %s %s" , token, string(bytes)))
	redisConn.Do(fmt.Sprintf("EXPIRE %s %d", token, config.GetInt("hiyuncms.server.frontend.session.timeout")))
}

func GetToken(token string) * frontend.UserSession {
	value,err := redis.String(redisConn.Do(fmt.Sprintf("GET %s" , token)))
	if err != nil {
		log.Printf("根据token：%s获取数据失败:%s\n", token, err.Error() )
		return nil
	} else {
		sessionInfo := frontend.UserSession{}
		json.Unmarshal([]byte(value), &sessionInfo)
		return &sessionInfo
	}
}
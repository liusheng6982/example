package frontend

import (
	"hiyuncms/config"
	"log"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

var redisConn redis.Conn

func init() {
	var err error
	redisConn, err = redis.Dial("tcp", config.GetValue("redis.address") )
	if  err != nil {
		log.Printf("初始化redis connection 出错！\n")
	}
}

func SetToken(token string, sessionInfo * UserSession)  {
	bytes, _ := json.Marshal( sessionInfo )
	_, err := redisConn.Do("SET" , token, string(bytes))
	if err != nil {
		log.Printf("set token出错！%s\n", err.Error() )
	}
	_, err = redisConn.Do("EXPIRE", token, config.GetInt("hiyuncms.server.frontend.session.timeout"))
	if err != nil {
		log.Printf("set token出错！%s\n", err.Error() )
	}
}

func GetToken(token string) * UserSession {
	value,err := redis.String(redisConn.Do("GET" , token))
	if err != nil {
		log.Printf("根据token：%s获取数据失败:%s\n", token, err.Error() )
		return nil
	} else {
		if value != "" {
			sessionInfo := UserSession{}
			json.Unmarshal([]byte(value), &sessionInfo)
			return &sessionInfo
		}else {
			return nil
		}

	}
}
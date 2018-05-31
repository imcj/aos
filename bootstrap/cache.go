package bootstrap

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/go-redis/redis"
	"github.com/apex/log"
)

var RedisClients map[string]*redis.Client

type RedisCommand struct {

}

// TODO: 允许redis存在未明明的主机
func (c RedisCommand)Execute() {
	log.Info("RedisCommand")

	RedisClients = make(map[string]*redis.Client)
	var client *redis.Client
	var host string
	var port int
	var password string
	var name int

	for key, value := range viper.Get("redis").(map[string]interface{}) {


		host = value.(map[string]interface{})["host"].(string)
		port = value.(map[string]interface{})["port"].(int)
		password = value.(map[string]interface{})["password"].(string)
		name = value.(map[string]interface{})["name"].(int)
		
		client = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%d", host, port),
			Password: password,
			DB: name,
		})
		pong, _ := client.Ping().Result()
		if "PONG" != pong {
			log.Error("error")
		} else {
			RedisClients[key] = client
		}
	}
}
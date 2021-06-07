package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"mail_log/dao"
	"mail_log/entity"
	"mail_log/util"
	"time"
)

var ctx = context.Background()

type LogSaverServer struct {
	dao   *dao.LogDao
	Redis *redis.Client
}

func (server *LogSaverServer) Init(logDao *dao.LogDao, client *redis.Client) {
	server.dao = logDao
	server.Redis = client
}

func (server LogSaverServer) ListenAndSave(level string) {
	for {
		message, err := server.Redis.RPop(level).Result()
		if err != nil {
			time.Sleep(time.Millisecond)
			continue
		}
		if level == "picture" {
			m := entity.PictureEntity{}
			err = json.Unmarshal([]byte(message), &m)
			if err != nil {
				fmt.Println(err)
				continue
			}
			server.dao.WritePitcure(&m)
		} else {
			server.dao.Write(&entity.LogEntity{Level: level, Message: message})
		}
	}
}

func (server LogSaverServer) Run() {
	for i := range dao.LogItem {
		go server.ListenAndSave(dao.LogItem[i])
	}
	util.PauseForRun()
}

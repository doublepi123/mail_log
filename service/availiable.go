package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mail_log/util"
	"net/http"
)

type AvailiableServer struct {
	db *util.Database
}

func (server *AvailiableServer) Init(db *util.Database) {
	server.db = db
}

func (server *AvailiableServer) Check() interface{} {
	redisErr := server.db.Redis.Set("aha", "testing", 0).Err()
	MysqlErr := server.db.DB.Raw("SELECT VERSION()").Error
	errs := make([]error, 0)
	if redisErr != nil {
		fmt.Println("RedisErr " + fmt.Sprint(redisErr))
		errs = append(errs, redisErr)
	}
	if MysqlErr != nil {
		fmt.Println("MysqlErr " + fmt.Sprint(MysqlErr))
		errs = append(errs, MysqlErr)
	}
	return errs
}

func (server *AvailiableServer) checkAcheck(c *gin.Context) {
	ans := server.Check().([]error)
	if len(ans) == 0 {
		c.JSON(http.StatusOK, gin.H{"Status": "OK"})
		return
	}
	var str string
	for i := range ans {
		str += fmt.Sprint(ans[i])
	}
	fmt.Println(str)
	c.JSON(http.StatusInternalServerError, gin.H{"ERROR": str})
}

func (server *AvailiableServer) ListenAndServer() {
	r := gin.Default()
	r.GET("/check", server.checkAcheck)
	r.Run("127.0.0.1:9999")
	util.PauseForRun()
}

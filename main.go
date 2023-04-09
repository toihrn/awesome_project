package main

import (
	"github.com/toxb11/awesome_project/api"
	"github.com/toxb11/awesome_project/infra/ai"
	"github.com/toxb11/awesome_project/infra/repository/db"
	"github.com/toxb11/awesome_project/infra/repository/redis"
)

func main() {
	ai.InitOpenAIClient()
	r := api.SetupBaseRouter()
	r = api.SetupCaseFileRouter(r)
	db.InitMysql()
	redis.InitRedisClient()
	db.InitCaseRepo()
	db.InitChatSentenceRepo()
	// Listen and Server in 0.0.0.0:8080
	_ = r.Run(":8080")
}

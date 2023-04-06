package main

import (
	"github.com/toxb11/awesome_project/api"
	"github.com/toxb11/awesome_project/infra/ai"
)

func main() {
	ai.InitOpenAIClient()
	r := api.SetupBaseRouter()
	// Listen and Server in 0.0.0.0:8080
	_ = r.Run(":8080")
}

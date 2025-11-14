package main

//https://www.youtube.com/watch?v=wyEYpX5U4Vg&t=2756s
import (
	"api/golang/gin-gonic/config"
	"api/golang/gin-gonic/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	_, err := config.InitializePostges()
	if err != nil {
		logger.Errorf("Failed to connect to database: %v", err)
	}
	//Initialize configs
	err = config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v ", err)
		return
	}

	//Initialize Router
	router.InitRouter()
}

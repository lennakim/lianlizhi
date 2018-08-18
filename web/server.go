package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	if err := InitConfig("config.yml"); err != nil {
		panic(err)
	}

	runmode := viper.GetString("runmode")
	fmt.Printf("%s\n", runmode)

	gin.SetMode(runmode)
	g := gin.New()
	middlewares := []gin.HandlerFunc{}

	db := InitDB()
	defer db.Close()

	Load(
		g,
		middlewares...,
	)

	http.ListenAndServe("0.0.0.0:2930", g)
}

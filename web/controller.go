package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "ok",
		Data: "index",
	})
}

func Goods(c *gin.Context) { // 用户商品
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "ok",
		Data: "goods",
	})
}

func User(c *gin.Context) { // 用户的信息
	id := c.Param("id")

	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "ok",
		Data: id,
	})
}

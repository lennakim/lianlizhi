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

func GoodCreate(c *gin.Context) { // 用户商品
	// 权限验证
	// 图片上传
	// json 验证
	// 设计state
	brand := c.PostForm("brand")
	img := c.PostForm("img")
	directPrice := c.PostForm("direct_price")
	state := c.PostForm("state")
	intr := c.PostForm("intr")

	db := InitDB()
	defer db.Close()

	res := db.MustExec("INSERT INTO goods (brand, img, direct_price, state, intr, created_at, updated_at ) VALUES ($1, $2, $3, $4, $5, now(), now() )",
		brand, img, directPrice, state, intr)

	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "ok",
		Data: res,
	})
}

func UserCreate(c *gin.Context) { // 用户的信息
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	addr := c.PostForm("addr")
	checkoutType := c.PostForm("checkout_type")
	wechat := c.PostForm("wechat")
	alipay := c.PostForm("alipay")
	bankNum := c.PostForm("bank_num")
	bankAddr := c.PostForm("bank_addr")
	intr := c.PostForm("intr")

	c.Header("Content-Type", "application/json; charset=utf-8")

	db := InitDB()
	defer db.Close()

	res := db.MustExec("INSERT INTO users (name, phone, addr, checkout_type, wechat, alipay, bank_num, bank_addr, intr, created_at, updated_at ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, now(), now() )",
		name, phone, addr, checkoutType, wechat, alipay, bankNum, bankAddr, intr)

	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "ok",
		Data: res,
	})
}

func UserShow(c *gin.Context) { // 用户的信息
	id := c.Param("id")

	db := InitDB()
	db = db.Unsafe()
	defer db.Close()

	users := []User{}
	err := db.Select(&users, "Select * from users where id = $1 limit 1", id)

	if err != nil {
		panic(err)
		return
	}

	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "ok",
		Data: users,
	})
}

func UserUpdate(c *gin.Context) { // 用户的信息
	id := c.Param("id")

	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "ok",
		Data: "UserUpdate" + id,
	})
}

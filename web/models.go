package main

import (
	"time"
)

type User struct {
	ID           uint64
	Name         string
	Phone        string
	Addr         string
	Avatar       string
	CheckoutType string
	Wechat       string
	Alipay       string
	BankId       string
	BankAddr     string
	Desc         string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "users"
}

type Good struct {
	UserId      uint64
	Category    string
	Brand       string    // 品牌
	Size        string    // 尺寸
	DirectPrice float64   // 寄卖价格
	UpItemAt    time.Time // 寄卖时间
	SoldPrice   float64   // 卖出价格
	SoldAt      time.Time // 卖出时间
	DownItemAt  time.Time // 下架时间
	State       string    // 状态
	Desc        string    // 描述
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Good) TableName() string {
	return "goods"
}

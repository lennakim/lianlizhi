package main

import (
	"database/sql"
	"time"
)

type User struct {
	ID           int64
	Name         string
	Phone        sql.NullString
	Addr         sql.NullString
	Avatar       sql.NullString
	CheckoutType sql.NullString
	Wechat       sql.NullString
	Alipay       sql.NullString
	BankNum      sql.NullString
	BankAddr     sql.NullString
	Intr         sql.NullString
	CreatedAt    time.Time
	UpdatedAt    time.Time
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

type SmsCode struct {
	Id        uint64
	Phone     string
	code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (SmsCode) TableName() string {
	return "sms_codess"
}

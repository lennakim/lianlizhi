package main

import (
	"database/sql"
	"time"
)

type User struct {
	ID           int64
	Name         string
	Phone        sql.NullString
	Verified     bool
	Addr         sql.NullString
	Avatar       sql.NullString
	CheckoutType sql.NullString `db:"checkout_type"` // 数据库字段 checkout_type <=> CheckoutType 映射
	Wechat       sql.NullString
	Alipay       sql.NullString
	BankNum      sql.NullString `db:bank_num`
	BankAddr     sql.NullString `db:bank_addr`
	Intr         sql.NullString
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

type Good struct {
	UserId      uint64
	Category    string
	Brand       string    // 品牌
	Size        string    // 尺寸
	DirectPrice float64   `db:direct_price` // 寄卖价格
	UpItemAt    time.Time `db:up_item_at`   // 寄卖时间
	SoldPrice   float64   `db:sold_price`   // 卖出价格
	SoldAt      time.Time `db:sold_at`      // 卖出时间
	DownItemAt  time.Time `db:down_item_at` // 下架时间
	State       string    // 状态
	Intr        string    // 描述
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (Good) TableName() string {
	return "goods"
}

type SmsCode struct {
	Id        uint64
	Phone     string
	Code      string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (SmsCode) TableName() string {
	return "sms_codess"
}

-- postgresql
CREATE DATABASE lianlizhi;
\c lianlizhi

CREATE TABLE "users" (
	"id" SERIAL PRIMARY KEY,
	"name" TEXT,
	"phone" TEXT,
	"addr" TEXT,
	"avatar" TEXT,
	"checkout_type" TEXT,
	"wechat" TEXT,
	"alipay" TEXT,
	"bank_num" TEXT,
	"bank_addr" TEXT,
	"intr" TEXT,
	"created_at" TIMESTAMP(0),
	"updated_at" TIMESTAMP(0)
	);

-- CREATE INDEX blocks_hash_inx ON blocks ("hash");

CREATE TABLE "goods" (
	"id" SERIAL PRIMARY KEY,
	"user_id" INT,
	"brand" TEXT,
	"img" TEXT,
	"size" TEXT,
	"direct_price" Float,
	"up_item_at" TIMESTAMP(0),
	"sold_price" Float,
	"sold_at" TIMESTAMP(0), -- 卖出价格
	"down_item_at" TIMESTAMP(0),
	"state" TEXT,
	"intr" TEXT,
	"created_at" TIMESTAMP(0),
	"updated_at" TIMESTAMP(0)
	);

CREATE TABLE "sms_codes" (
	"id" SERIAL PRIMARY KEY,
	"phone" TEXT NOT NULL,
	"code" TEXT NOT NULL,
	"created_at" TIMESTAMP(0),
	"updated_at" TIMESTAMP(0)
)
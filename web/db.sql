-- postgresql
CREATE DATABASE lianlizhi;
\c lianlizhi

CREATE TABLE "users" (
	"id" SERIAL PRIMARY KEY,
	"name" TEXT,
	"phone" TEXT,
	"addr" TEXT,
	"Avatar" TEXT,
	"checkout_type" TEXT,
	"wechat" TEXT,
	"Alipay" TEXT,
	"bank_id" TEXT,
	"bank_addr" TEXT,
	"desc" TEXT,
	"created_at" TIMESTAMP(0),
	"updated_at" TIMESTAMP(0)
	);

-- CREATE INDEX blocks_hash_inx ON blocks ("hash");

CREATE TABLE "goods" (
	"id" SERIAL PRIMARY KEY,
	"user_id" INT,
	"brand" TEXT,
	"size" TEXT,
	"direct_price" Float,
	"up_item_at" TIMESTAMP(0),
	"sold_price" Float,
	"sold_at" TIMESTAMP(0),
	"down_item_at" TIMESTAMP(0),
	"state" TEXT,
	"desc" TEXT,
	"created_at" TIMESTAMP(0),
	"updated_at" TIMESTAMP(0)
	);



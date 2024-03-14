package main

import (
	"fmt"

	"github.com/nats-io/stan.go"
)

func main() {
	dat_1 := []byte(`{
		"order_uid": "b563feb7b2b84b6test",
		"track_number": "WBILMTESTTRACK",
		"entry": "WBIL",
		"delivery": {
		  "name": "Test Testov",
		  "phone": "+9720000000",
		  "zip": "2639809",
		  "city": "Kiryat Mozkin",
		  "address": "Ploshad Mira 15",
		  "region": "Kraiot",
		  "email": "test@gmail.com"
		},
		"payment": {
		  "transaction": "b563feb7b2b84b6test",
		  "request_id": "",
		  "currency": "USD",
		  "provider": "wbpay",
		  "amount": 1817,
		  "payment_dt": 1637907727,
		  "bank": "alpha",
		  "delivery_cost": 1500,
		  "goods_total": 317,
		  "custom_fee": 0
		},
		"items": [
		  {
			"chrt_id": 9934930,
			"track_number": "WBILMTESTTRACK",
			"price": 453,
			"rid": "ab4219087a764ae0btest",
			"name": "Mascaras",
			"sale": 30,
			"size": "0",
			"total_price": 317,
			"nm_id": 2389212,
			"brand": "Vivienne Sabo",
			"status": 202
		  }
		],
		"locale": "en",
		"internal_signature": "",
		"customer_id": "test",
		"delivery_service": "meest",
		"shardkey": "9",
		"sm_id": 99,
		"date_created": "2021-11-26T06:22:19Z",
		"oof_shard": "1"
	  }`)

	// dat_2 := []byte(`{
	// 	"order_uid": "AfKKbEbFfyDUrcLtest",
	// 	"track_number": "TEST2_TRACK_WB",
	// 	"entry": "WBIL",
	// 	"delivery": {
	// 	  "name": "Vitaly Yagodkin",
	// 	  "phone": "+8001007505",
	// 	  "zip": "WIoPtCq",
	// 	  "city": "Cheboksary",
	// 	  "address": "Moskovsky Prospekt 5",
	// 	  "region": "Chuvashia",
	// 	  "email": "yagodkin@mail.com"
	// 	},
	// 	"payment": {
	// 	  "transaction": "AfKKbEbFfyDUrcLtest",
	// 	  "request_id": "",
	// 	  "currency": "RUB",
	// 	  "provider": "wbpay",
	// 	  "amount": 5000,
	// 	  "payment_dt": 1637907727,
	// 	  "bank": "cheb_bank",
	// 	  "delivery_cost": 500,
	// 	  "goods_total": 245,
	// 	  "custom_fee": 0
	// 	},
	// 	"items": [
	// 	  {
	// 		"chrt_id": 1034931,
	// 		"track_number": "TEST2_TRACK_WB",
	// 		"price": 400,
	// 		"rid": "PWMLIcdPfjIbOrBjbtest",
	// 		"name": "sneakers",
	// 		"sale": 30,
	// 		"size": "40",
	// 		"total_price": 134,
	// 		"nm_id": 2389211,
	// 		"brand": "ecco",
	// 		"status": 202
	// 	  }
	// 	],
	// 	"locale": "en",
	// 	"internal_signature": "",
	// 	"customer_id": "test2",
	// 	"delivery_service": "meest",
	// 	"shardkey": "2",
	// 	"sm_id": 21,
	// 	"date_created": "2024-03-13T06:55:19Z",
	// 	"oof_shard": "2"
	//   }`)

	// dat_3 := []byte(`{
	// 	"order_uid": "blfVmTkanwgpQGYtest",
	// 	"track_number": "TEST3_TRACK_WB",
	// 	"entry": "WBIL",
	// 	"delivery": {
	// 	  "name": "Olga Primerchikova",
	// 	  "phone": "+8301007505",
	// 	  "zip": "gTxnUlkF",
	// 	  "city": "Svetlogorsk",
	// 	  "address": "Raspberry lane 1",
	// 	  "region": "Kaliningrad region",
	// 	  "email": "primerchikova@mail.com"
	// 	},
	// 	"payment": {
	// 	  "transaction": "blfVmTkanwgpQGYtest",
	// 	  "request_id": "",
	// 	  "currency": "RUB",
	// 	  "provider": "wbpay",
	// 	  "amount": 3000,
	// 	  "payment_dt": 1637107727,
	// 	  "bank": "zenit",
	// 	  "delivery_cost": 100,
	// 	  "goods_total": 307,
	// 	  "custom_fee": 0
	// 	},
	// 	"items": [
	// 	  {
	// 		"chrt_id": 1034931,
	// 		"track_number": "TEST3_TRACK_WB",
	// 		"price": 4000,
	// 		"rid": "mIqfTAFUaobzCkItest",
	// 		"name": "photo camera",
	// 		"sale": 1,
	// 		"size": "",
	// 		"total_price": 4000,
	// 		"nm_id": 2389218,
	// 		"brand": "zenit",
	// 		"status": 200
	// 	  }
	// 	],
	// 	"locale": "en",
	// 	"internal_signature": "",
	// 	"customer_id": "test3",
	// 	"delivery_service": "cdek",
	// 	"shardkey": "4",
	// 	"sm_id": 40,
	// 	"date_created": "2024-03-10T06:55:19Z",
	// 	"oof_shard": "3"
	//   }`)
	sc, err := stan.Connect("test-cluster", "2")
	if err != nil {
		fmt.Print(err)
	}
	defer sc.Close()
	err = sc.Publish("foo", dat_1)
	//err = sc.Publish("foo", dat_2)
	// err = sc.Publish("foo", dat_3)

}

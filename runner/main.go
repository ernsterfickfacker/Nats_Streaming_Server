package main

import (
	"L0/db"
	cache "L0/src/cache"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	db.DBConnection()
	cache.New(10*time.Minute, 20*time.Minute)
	ords, err := db.FindAllOrders(context.Background())
	if err != nil {
		fmt.Print("Failed to dump db")
	}
	for _, ord := range ords {
		data, _ := json.Marshal(ord)
		cache.LocalCache.Set(ord.OrderUid, string(data), 1*time.Minute)
	}
}

func controllers(router *gin.Engine) {
	db.OrdersController(router)
}

func main() {
	var hnd db.Data_Base
	hnd = db.NewDB()

	router := gin.Default()
	router.Use(cors.Default())
	controllers(router)

	sc, err := stan.Connect("test-cluster", "1")
	if err != nil {
		fmt.Print("Connection failed:", err)
	}
	_, err = sc.Subscribe("foo", func(m *stan.Msg) {
		var data db.Order
		err := json.Unmarshal(m.Data, &data)
		if err != nil {
			fmt.Print("Unmarshal fail")
		}
		hnd.AddOrder(&data, context.Background()) //добавление заказа
	})

	if err != nil {
		fmt.Print("Subscriber fail")
	}
	err = router.Run(":8080")
	if err != nil {
		fmt.Print("Unable to run gin")
	}
}

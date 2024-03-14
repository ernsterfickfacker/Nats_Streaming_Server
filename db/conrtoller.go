package db

import (
	"L0/src/cache"
	"L0/src/pattern"
	"context"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	//new

	"encoding/json"
)

var index = template.Must(template.New("index").Parse(pattern.IndexTmpl))

func OrdersController(router *gin.Engine) {
	v1 := router.Group("/")
	{
		// v1.GET("", func(ctx *gin.Context) {
		// 	orders, err := FindAllOrders(context.Background())
		// 	//result := gin.H{"order_uid": orders, "count": 1}
		// 	if err != nil {
		// 		fmt.Print("Get all; error")

		// 	}

		// 	var uid = make([]string, 0, len(orders))
		// 	for i := 0; i < len(orders); i++ {
		// 		uid = append(uid, orders[i].OrderUid)
		// 	}

		// 	ctx.JSON(http.StatusOK, "Orders_UID")
		// 	ctx.JSON(http.StatusOK, uid)

		// })
		v1.GET("all", func(ctx *gin.Context) {
			orders, err := FindAllOrders(context.Background())
			if err != nil {
				fmt.Print("Get all; error")

			}
			ctx.JSON(http.StatusOK, orders)

		})
		v1.GET(":id", func(c *gin.Context) {
			dat, res := cache.LocalCache.Get(c.Param("id"))
			var ord Order
			var err error
			if !res {
				ord, err = FindOrderById(c.Param("id"), context.Background())
				if err != nil {
					c.JSON(http.StatusNotFound, "id not found")
					return
				} else {
					data, _ := json.Marshal(ord)
					cache.LocalCache.Set(ord.OrderUid, string(data), 10*time.Minute)
				}
			} else {
				err = json.Unmarshal([]byte(dat), &ord)
			}
			if err := index.Execute(c.Writer, ord); err != nil {
				fmt.Printf("index.Execute error: %s\n", err)
			}
		})
	}
}

package db

import (
	"context"
	"fmt"
	"strconv"
)

type Data_Base struct {
}

func NewDB() Data_Base {
	return Data_Base{}
}

func FindOrderById(id string, c context.Context) (Order, error) {
	var ord Order
	var tmp Order_return
	query := `select orderuid, tracknumber, entry,
		delivery_id, payment_id, locale, internalsignature, customerid,
		deliveryservice, shardkey, smid, datecreated, oofshard
	from order_table where orderuid = $1`
	err := Conn.QueryRow(c, query, id).Scan(
		&ord.OrderUid,
		&ord.TrackNumber,
		&ord.Entry,
		&tmp.Delivery_id,
		&tmp.Payment_id,
		&ord.Locale,
		&ord.InternalSignature,
		&ord.CustomerId,
		&ord.DeliveryService,
		&ord.Shardkey,
		&ord.SmId,
		&ord.DateCreated,
		&ord.OofShard,
	)
	if err != nil {
		fmt.Println("Order not found:", err)
		return ord, err
	}
	query = `select name, phone, zip, city, adress, region, email 
		from delivery where delivery_id = $1`
	err = Conn.QueryRow(c, query, tmp.Delivery_id).Scan(
		&ord.Delivery.Name,
		&ord.Delivery.Phone,
		&ord.Delivery.Zip,
		&ord.Delivery.City,
		&ord.Delivery.Address,
		&ord.Delivery.Region,
		&ord.Delivery.Email,
	)
	if err != nil {
		fmt.Println("Delivery not found:", err)
		return ord, err
	}
	query = `select transaction, requestid, currency, provider, amount, 
		paymentdt, bank, deliverycost, goodstotal, customfee 
		from payment where payment_id = $1`
	err = Conn.QueryRow(c, query, tmp.Payment_id).Scan(
		&ord.Payment.Transaction,
		&ord.Payment.RequestId,
		&ord.Payment.Currency,
		&ord.Payment.Provider,
		&ord.Payment.Amount,
		&ord.Payment.PaymentDt,
		&ord.Payment.Bank,
		&ord.Payment.DeliveryCost,
		&ord.Payment.GoodsTotal,
		&ord.Payment.CustomFee,
	)
	if err != nil {
		fmt.Println("Payment not found:", err)
		return ord, err
	}
	var itm Items
	query = `select chrtid, tracknumber, price, rid, name, sale, 
		size, totalprice, mnid, brand, status from items where order_id = $1`
	rows, err := Conn.Query(c, query, ord.OrderUid)
	if err != nil {
		fmt.Println("Item not found:", err)
		return ord, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&itm.ChrtId,
			&itm.TrackNumber,
			&itm.Price,
			&itm.Rid,
			&itm.Name,
			&itm.Sale,
			&itm.Size,
			&itm.TotalPrice,
			&itm.NmId,
			&itm.Brand,
			&itm.Status,
		)
		ord.Items = append(ord.Items, itm)
		if err != nil {
			fmt.Println("Scan fail:", err)
			return ord, err
		}
	}

	return ord, err
}

func FindAllOrders(c context.Context) ([]Order, error) {
	var orders []Order
	var orderuid []string
	var tmp_uid string
	var tmp Order
	query := "select orderuid from order_table"
	rows, err := Conn.Query(c, query)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&tmp_uid,
		)
		orderuid = append(orderuid, tmp_uid)
	}
	if err != nil {
		fmt.Print("Error during reading: ")
	}
	var order_id string
	for _, order_id = range orderuid {
		tmp, _ = FindOrderById(order_id, c)
		orders = append(orders, tmp)
	}
	if err != nil {
		fmt.Print("Error during reading: ")
	}
	return orders, err
}

func (d *Data_Base) AddOrder(ord *Order, c context.Context) {
	//fmt.Print("OK")
	de_id := ord.OrderUid + "_d"
	query := "insert into delivery values($1, $2, $3, $4, $5, $6, $7, $8)"
	var err error
	_, err = Conn.Query(c, query,
		de_id,
		ord.Delivery.Name,
		ord.Delivery.Phone,
		ord.Delivery.Zip,
		ord.Delivery.City,
		ord.Delivery.Address,
		ord.Delivery.Region,
		ord.Delivery.Email,
	)
	if err != nil {
		fmt.Print("Error during insert: ")
	}
	pr_id := ord.OrderUid + "_p"
	query = "insert into payment values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"
	_, err = Conn.Query(c, query,
		pr_id,
		ord.Payment.Transaction,
		ord.Payment.RequestId,
		ord.Payment.Currency,
		ord.Payment.Provider,
		ord.Payment.Amount,
		ord.Payment.PaymentDt,
		ord.Payment.Bank,
		ord.Payment.DeliveryCost,
		ord.Payment.GoodsTotal,
		ord.Payment.CustomFee,
	)
	if err != nil {
		fmt.Print("Error during insert: ")
	}
	var i int = 1
	var itms []Items = ord.Items
	for _, itm := range itms {
		query = "insert into items values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)"
		_, err = Conn.Query(c, query,
			ord.OrderUid,
			ord.OrderUid+"_"+strconv.Itoa(i),
			itm.ChrtId,
			itm.TrackNumber,
			itm.Price,
			itm.Rid,
			itm.Name,
			itm.Sale,
			itm.Size,
			itm.TotalPrice,
			itm.NmId,
			itm.Brand,
			itm.Status,
		)
		if err != nil {
			fmt.Print("Error during insert: ")
		}
		i += 1
	}
	query = "insert into order_table values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, " +
		"$13)"
	_, err = Conn.Exec(c, query,
		ord.OrderUid,
		ord.TrackNumber,
		ord.Entry,
		de_id,
		pr_id,
		//itm_ids,
		ord.Locale,
		ord.InternalSignature,
		ord.CustomerId,
		ord.DeliveryService,
		ord.Shardkey,
		ord.SmId,
		ord.DateCreated,
		ord.OofShard)
	if err != nil {
		fmt.Print("Error during insert: ", err)
	}
}

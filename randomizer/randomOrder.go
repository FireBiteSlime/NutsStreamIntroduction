package randomizer

import (
	"fmt"
	"math/rand"
	"nstream/models"
)

func RandomOrder() models.Order {

	var o models.Order
	o.Order_uid = RandomString(19)
	o.Track_number = RandomString(14)
	o.Entry = RandomString(4)
	o.Delivery.Name = RandomString(rand.Intn(6)) + " " + RandomString(rand.Intn(6))
	o.Delivery.Phone = "+" + RandomString(10)
	o.Delivery.Zip = RandomString(7)
	o.Delivery.City = RandomString(rand.Intn(8))
	o.Delivery.Address = RandomString(rand.Intn(8)) + string(rand.Intn(20))
	o.Delivery.Region = RandomString(rand.Intn(7))
	o.Delivery.Email = RandomString(rand.Intn(8)) + "@gmail.com"
	o.Payment.Transaction = o.Order_uid
	o.Payment.Request_id = string(rand.Intn(200))
	o.Payment.Currency = RandomString(3)
	o.Payment.Provider = RandomString(rand.Intn(5))
	o.Payment.Amount = rand.Intn(5000)
	o.Payment.Payment_dt = rand.Intn(5000000)
	o.Payment.Bank = RandomString(rand.Intn(8))
	o.Payment.Delivery_cost = rand.Intn(50000)
	o.Payment.Goods_total = rand.Intn(1000)
	o.Payment.Custom_fee = rand.Intn(100)
	items := make([]models.Item, 0)
	k := rand.Intn(4) + 1
	fmt.Println(k)
	for i := 0; i < k; i++ {
		var item models.Item
		item.Chrt_id = rand.Intn(100000000)

		item.Track_number = o.Track_number
		item.Price = rand.Intn(100000)
		item.Rid = RandomString(21)
		item.Name = RandomString(10)
		item.Sale = rand.Intn(100)
		item.Size = RandomString(4)
		item.Total_price = rand.Intn(1000)
		item.Nm_id = rand.Intn(100000000)
		item.Brand = RandomString(10)
		item.Status = rand.Intn(1000)
		items = append(items, item)
	}
	o.Items = items
	o.Locale = RandomString(2)
	o.Internal_signature = RandomString(5)
	o.Customer_id = RandomString(5)
	o.Delivery_service = RandomString(rand.Intn(5))
	o.Shardkey = RandomString(rand.Intn(5))
	o.Sm_id = rand.Intn(100)
	o.Date_created = RandomString(rand.Intn(15))
	o.Oof_shard = RandomString(10)
	return o
}

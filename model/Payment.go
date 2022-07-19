package model

type Payment struct {
	Transaction   string
	Request_id    string
	Currency      string
	Provider      string
	Amount        int64
	Payment_dt    int64
	Bank          string
	Delivery_cost int64
	Goods_total   int64
	Custom_fee    int64
}

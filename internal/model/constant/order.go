package constant

type OrderStatus string

const (
	Processed OrderStatus = "processed"
	Finished  OrderStatus = "finished"
	Failed    OrderStatus = "failed"
)

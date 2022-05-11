package entity

type Payment_coin struct {
	Id       int64   `json:"id"`
	Uid      int64   `json:"uid"`
	Account  float64 `json:"account"`
	Cost     float64 `json:"cost"`
}

type Payment_point struct {
	Id       int64   `json:"id"`
	Uid      int64   `json:"uid"`
	Account  float64 `json:"account"`
	Cost     float64 `json:"cost"`
}

func (Payment_coin) TableName() string {
	return "platform_coin"
}

func (Payment_point) TableName() string {
	return "platform_point"
}
package entity

type Percentage_point struct {
	Percentage   float64  `json:"percentage"`
}

type Discount_vip_coin struct{
	Vip         int64     `json:"vip"`
	Disacount   float64   `json:"disacount"`
}

type Discount_vip_point struct {
	Disacount   float64  `json:"disacount"`
}

func (Percentage_point) TableName() string {
	return "platform_point_percentage"
}

func (Discount_vip_coin) TableName() string {
	return "vip_coin_discount"
}

func (Discount_vip_point) TableName() string {
	return "vip_point_discount"
}
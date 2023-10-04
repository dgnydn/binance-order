package domain

import (
	"encoding/json"
	"github.com/adshao/go-binance/v2"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Symbol          string                  `json:"symbol,omitempty"`
	Side            binance.SideType        `json:"side,omitempty"`
	TimeInForceType binance.TimeInForceType `json:"time_in_force_type,omitempty"`
	Quantity        int                     `json:"quantity,omitempty"`
	Price           float64                 `json:"price,omitempty"`
}

func (o *Order) TableName() string {
	return "orders"
}

// IsValid checks if the order is valid
func (o *Order) IsValid() bool {
	return o.Symbol != "" && o.Side != "" && o.TimeInForceType != "" && o.Quantity > 0 && o.Price > 0
}

// ToJSON returns the order as JSON to use in kafka etc...
func (o *Order) ToJSON() string {
	jsonOrder, err := json.Marshal(o)
	if err != nil {
		return err.Error()
	}
	return string(jsonOrder)
}

// FromJSON returns the order from JSON to use in kafka etc...
func (o *Order) FromJSON(jsonOrder string) error {
	return json.Unmarshal([]byte(jsonOrder), &o)
}

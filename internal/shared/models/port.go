// Package models contain shared models/data transfer objects (DTO) accross the application
package models

import "github.com/shopspring/decimal"

// Port represents a harbor and contain correlated data
type Port struct {
	Name        string            `json:"name"`
	City        string            `json:"city"`
	Country     string            `json:"country"`
	Alias       []string          `json:"alias"`
	Regions     []string          `json:"regions"`
	Coordinates []decimal.Decimal `json:"coordinates"`
	Province    string            `json:"province"`
	Timezone    string            `json:"timezone"`
	Unlocs      []string          `json:"unlocs"`
	Code        string            `json:"code"`
}

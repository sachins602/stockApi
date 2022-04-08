package models

import (
	"gorm.io/gorm"
)

type NepseInfo struct {
	Mt      string `json:"mt"`
	Overall struct {
		D  string `json:"d"`
		T  string `json:"t"`
		Q  string `json:"q"`
		Tn string `json:"tn"`
		St string `json:"st"`
		Mc string `json:"mc"`
		Fc string `json:"fc"`
	} `json:"overall"`
	Turnover struct {
		Date   string `json:"date"`
		Detail []struct {
			S  string  `json:"s"`
			N  string  `json:"n"`
			Lp float64 `json:"lp"`
			T  float64 `json:"t"`
			Pc float64 `json:"pc"`
			H  float64 `json:"h"`
			L  float64 `json:"l"`
			Op float64 `json:"op"`
			Q  float64 `json:"q"`
		} `json:"detail"`
	} `json:"turnover"`
	Sector struct {
		Date   string `json:"date"`
		Detail []struct {
			S string  `json:"s"`
			T float64 `json:"t"`
			Q float64 `json:"q"`
		} `json:"detail"`
	} `json:"sector"`
	Broker struct {
		Date   string `json:"date"`
		Detail []struct {
			B string  `json:"b"`
			N string  `json:"n"`
			P float64 `json:"p"`
			S float64 `json:"s"`
			M float64 `json:"m"`
			T float64 `json:"t"`
		} `json:"detail"`
	} `json:"broker"`
	Stock struct {
		Date   string `json:"date"`
		Detail []struct {
			S  string  `json:"s"`
			Lp float64 `json:"lp"`
			C  int     `json:"c"`
			Q  float64 `json:"q"`
		} `json:"detail"`
	} `json:"stock"`
}
type Stock_details struct {
	// gorm.Model
	StockName string `json:"StockName"`
	// N  string  `json:"n"`
	LastPrice   float64 `json:"LastPrice"`
	TurnOver    float64 `json:"TurnOver"`
	Change      float64 `json:"Change"`
	High        float64 `json:"High"`
	Low         float64 `json:"Low"`
	Open        float64 `json:"Open"`
	ShareTraded float64 `json:"ShareTraded"`
}

func GetStocks(db *gorm.DB, Stock_details *[]Stock_details) (err error) {
	err = db.Find(Stock_details).Error
	//Raw("SELECT * FROM `stock_details`(stockName, lastPrice, `turnOver`, `change`, `high`, `low`, `open`, `shareTraded`) VALUES (?, ?, ?, ?, ?, ?, ?, ?);").First(ScripDetail).Error
	if err != nil {
		return err
	}
	return nil
}

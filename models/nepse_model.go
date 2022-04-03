package models

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

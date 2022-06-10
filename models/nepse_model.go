package models

type Historic struct {
	Scrip  string  `json:"Scrip" gorm:"column:Scrip"`
	Time   float64 `json:"Time" gorm:"column:Time"`
	Close  float64 `json:"Close" gorm:"column:Close"`
	Open   float64 `json:"Open" gorm:"column:Open"`
	High   float64 `json:"High" gorm:"column:High"`
	Low    float64 `json:"Low" gorm:"column:Low"`
	Volume float64 `json:"Volume" gorm:"column:Volume"`
}

type Stock struct {
	StockName   string  `json:"StockName" gorm:"column:StockName"`
	LastPrice   float64 `json:"LastPrice" gorm:"column:LastPrice"`
	TurnOver    float64 `json:"TurnOver" gorm:"column:TurnOver"`
	Change      float64 `json:"Change" gorm:"column:Change"`
	High        float64 `json:"High" gorm:"column:High"`
	Low         float64 `json:"Low" gorm:"column:Low"`
	Open        float64 `json:"Open" gorm:"column:Open"`
	ShareTraded float64 `json:"ShareTraded" gorm:"column:ShareTraded"`
}

type Broker struct {
	BrokerNumber float64 `json:"BrokerNumber" gorm:"column:BrokerNumber"`
	BrokerName   string  `json:"BrokerName" gorm:"column:BrokerName"`
	Purchase     float64 `json:"Purchase" gorm:"column:Purchase"`
	Sales        float64 `json:"Sales" gorm:"column:Sales"`
	Matching     float64 `json:"Matching" gorm:"column:Matching"`
	Total        float64 `json:"Total" gorm:"column:Total"`
}

type Sector struct {
	SectorName string  `json:"SectorName" gorm:"column:SectorName"`
	Turnover   float64 `json:"Turnover" gorm:"column:Turnover"`
	Quantity   float64 `json:"Quantity" gorm:"column:Quantity"`
}
type Index struct {
	Alpha               string `json:"Alpha" gorm:"column:Alpha"`
	Beta                string `json:"Beta" gorm:"column:Beta"`
	DailyGain           string `json:"DailyGain" gorm:"column:DailyGain"`
	FullName            string `json:"FullName" gorm:"column:FullName"`
	IndexName           string `json:"IndexName" gorm:"column:IndexName"`
	Ltp                 string `json:"Ltp" gorm:"column:Ltp"`
	Macd                string `json:"Macd" gorm:"column:Macd"`
	Macdsignal          string `json:"Macdsignal" gorm:"column:Macdsignal"`
	Pb                  string `json:"Pb" gorm:"column:Pb"`
	Pe                  string `json:"Pe" gorm:"column:Pe"`
	Peg                 string `json:"Peg" gorm:"column:Peg"`
	Roa                 string `json:"Roa" gorm:"column:Roa"`
	Roe                 string `json:"Roe" gorm:"column:Roe"`
	Rsi                 string `json:"Rsi" gorm:"column:Rsi"`
	SharpeRatio         string `json:"SharpeRatio" gorm:"column:SharpeRatio"`
	SmaTwo              string `json:"SmaTwo" gorm:"column:SmaTwo"`
	TotalDividendYield  string `json:"TotalDividendYield" gorm:"column:TotalDividendYield"`
	TotalNegativeGainer string `json:"TotalNegativeGainer" gorm:"column:TotalNegativeGainer"`
	TotalPositiveGainer string `json:"TotalPositiveGainer" gorm:"column:TotalPositiveGainer"`
	Turnover            string `json:"Turnover" gorm:"column:Turnover"`
	YearlyPercentChange string `json:"YearlyPercentChange" gorm:"column:YearlyPercentChange"`
}

package models

type Historic struct {
	Time   float64 `json:"Time" gorm:"column:Time"`
	Close  float64 `json:"Close" gorm:"column:Close"`
	Open   float64 `json:"Open" gorm:"column:Open"`
	High   float64 `json:"High" gorm:"column:High"`
	Low    float64 `json:"Low" gorm:"column:Low"`
	Volume float64 `json:"Volume" gorm:"column:Volume"`
}
type HistoricPrediction struct {
	Time   float64 `json:"Time" gorm:"column:Time"`
	Close  float64 `json:"Close" gorm:"column:Close"`
	Prediction   float64 `json:"Prediction" gorm:"column:Prediction"`
}

type NepseHistoricPrediction struct {
	Scrip  string `json:"Scrip" gorm:"column:Scrip"`
	Time   float64 `json:"Time" gorm:"column:Time"`
	Close  float64 `json:"Close" gorm:"column:Close"`
	LstmPrediction   float64 `json:"LstmPrediction" gorm:"column:lstm_prediction"`
	GruPrediction   float64 `json:"GruPrediction" gorm:"column:gru_prediction"`
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

type Gainer struct {
	Id            int    `json:"Id" gorm:"column:id"`
	Symbol        string `json:"Symbol" gorm:"column:symbol"`
	CompanyName   string `json:"CompanyName" gorm:"column:company"`
	Ltp           string `json:"Ltp" gorm:"column:ltp"`
	Change        string `json:"Change" gorm:"column:point_change"`
	PercentChange string `json:"PercentChange" gorm:"column:percent_change"`
}

type SubIndex struct {
	Sector        string `json:"Sector" gorm:"column:sector"`
	Turnover      string `json:"Turnover" gorm:"column:turnover"`
	Close         string `json:"Close" gorm:"column:close"`
	Point         string `json:"Point" gorm:"column:point"`
	PercentChange string `json:"PercentChange" gorm:"column:percent_change"`
}

type CompanyDetails struct {
	Scrip                  string `json:"Scrip" gorm:"column:scrip"`
	Name                   string `json:"Name" gorm:"column:name"`
	Sector                 string `json:"Sector" gorm:"column:sector"`
	ShareOutstanding       string `json:"ShareOutstanding" gorm:"column:share_outstanding"`
	MarketPrice            string `json:"MarketPrice" gorm:"column:market_price"`
	PercentChange          string `json:"PercentChange" gorm:"column:percent_change"`
	LastTradedOn           string `json:"LastTradedOn" gorm:"column:last_traded_on"`
	FifitytwoWeekHighLow   string `json:"FifitytwoWeekHighLow" gorm:"column:fifitytwo_week_high_low"`
	OnehundredeightyDayAvg string `json:"OnehundredeightyDayAvg" gorm:"column:onehundredeighty_day_avg"`
	OnehundredtwentyDayAvg string `json:"OnehundredtwentyDayAvg" gorm:"column:onehundredtwenty_day_avg"`
	OneYearYield           string `json:"OneYearYield" gorm:"column:one_year_yield"`
	Eps                    string `json:"Eps" gorm:"column:eps"`
	PeRatio                string `json:"PeRatio" gorm:"column:pe_ratio"`
	BookValue              string `json:"BookValue" gorm:"column:book_value"`
	Pbv                    string `json:"Pbv" gorm:"column:pbv"`
}

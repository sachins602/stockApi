package models

import (
	"gorm.io/gorm"
)

type UserInfo struct {
	FullName    string `json:"full_name"`
	Email       string `json:"email"  gorm:"primary_key"`
	Password    string `json:"password"`
	PhoneNumber int64  `json:"phone_number"`
}

type Portfolio struct {
	Email    string  `json:"email" gorm:"foreign_key"`
	Scrip    string  `json:"scrip" gorm:"foreign_key"`
	Type     string  `json:"type"`
	Quantity int64   `json:"quantity"`
	BuyPrice float64 `json:"buy_price"`
}

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
type Historic struct {
	Scrip  string  `json:"Scrip"`
	Time   float64 `json:"Time"`
	Close  float64 `json:"Close"`
	Open   float64 `json:"Open"`
	High   float64 `json:"High"`
	Low    float64 `json:"Low"`
	Volume float64 `json:"Volume"`
}

type IndexInfo struct {
	Date      string `json:"date"`
	HomeTable []struct {
		Five0_200Sma interface{} `json:"_50_200_sma"`
		Alpha        string      `json:"alpha"`
		Beta         string      `json:"beta"`
		CreatedAt    string      `json:"created_at"`
		DailyGain    string      `json:"daily_gain"`
		FullName     string      `json:"full_name"`
		ID           int64       `json:"id"`
		IndexName    string      `json:"index_name"`
		Indexvalue   struct {
			Current        string `json:"current"`
			FormattedIndex string `json:"formatted_index"`
			PercentChange  string `json:"percent_change"`
			TurnOverValue  string `json:"turn_over_value"`
		} `json:"indexvalue"`
		Investors         interface{} `json:"investors"`
		Ltp               string      `json:"ltp"`
		Macd              string      `json:"macd"`
		Macdsignal        string      `json:"macdsignal"`
		Orders            int64       `json:"orders"`
		Pb                string      `json:"pb"`
		Pe                string      `json:"pe"`
		PeFwd             interface{} `json:"pe_fwd"`
		Peg               string      `json:"peg"`
		Roa               string      `json:"roa"`
		Roe               string      `json:"roe"`
		Rsi               string      `json:"rsi"`
		SeasonalInvestors interface{} `json:"seasonal_investors"`
		SharpeRatio       string      `json:"sharpe_ratio"`
		Sma200            string      `json:"sma_200"`
		Sparkline         struct {
			Current []string `json:"current"`
			Max     string   `json:"max"`
			Min     string   `json:"min"`
			Time    []string `json:"time"`
		} `json:"sparkline"`
		TotalDividendYield  string      `json:"total_dividend_yield"`
		TotalNegativeGainer string      `json:"total_negative_gainer"`
		TotalPositiveGainer string      `json:"total_positive_gainer"`
		Traders             interface{} `json:"traders"`
		TurnoverValues      string      `json:"turnover_values"`
		UpdatedAt           interface{} `json:"updated_at"`
		YearlyPercentChange string      `json:"yearly_percent_change"`
	} `json:"home_table"`
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

type Broker_details struct {
	BrokerNumber float64 `json:"BrokerNumber"`
	BrokerName   string  `json:"BrokerName"`
	Purchase     float64 `json:"Purchase"`
	Sales        float64 `json:"Sales"`
	Matching     float64 `json:"Matching"`
	Total        float64 `json:"Total"`
}

type Sector_details struct {
	SectorName string  `json:"SectorName"`
	Turnover   float64 `json:"Turnover"`
	Quantity   float64 `json:"Quantity"`
}
type Index_details struct {
	Alpha               string `json:"alpha"`
	Beta                string `json:"beta"`
	DailyGain           string `json:"daily_gain"`
	FullName            string `json:"full_name"`
	IndexName           string `json:"index_name"`
	Ltp                 string `json:"ltp"`
	Macd                string `json:"macd"`
	Macdsignal          string `json:"macdsignal"`
	Pb                  string `json:"pb"`
	Pe                  string `json:"pe"`
	Peg                 string `json:"peg"`
	Roa                 string `json:"roa"`
	Roe                 string `json:"roe"`
	Rsi                 string `json:"rsi"`
	SharpeRatio         string `json:"sharpe_ratio"`
	SmaTwo              string `json:"sma_200"`
	TotalDividendYield  string `json:"total_dividend_yield"`
	TotalNegativeGainer string `json:"total_negative_gainer"`
	TotalPositiveGainer string `json:"total_positive_gainer"`
	Turnover            string `json:"turnover_values"`
	YearlyPercentChange string `json:"yearly_percent_change"`
}

func GetStocks(db *gorm.DB, Stock_details *[]Stock_details) (err error) {
	err = db.Find(Stock_details).Error
	//Raw("SELECT * FROM `stock_details`(stockName, lastPrice, `turnOver`, `change`, `high`, `low`, `open`, `shareTraded`) VALUES (?, ?, ?, ?, ?, ?, ?, ?);").First(ScripDetail).Error
	if err != nil {
		return err
	}
	return nil
}

func GetNepseHistory(db *gorm.DB, Historic *[]Historic) (err error) {
	err = db.Find(Historic, "scrip = ?", "ADBL").Error
	if err != nil {
		return err
	}
	return nil
}

func GetBroker(db *gorm.DB, Broker_details *[]Broker_details) (err error) {
	err = db.Find(Broker_details).Error
	if err != nil {
		return err
	}
	return nil
}

func GetSector(db *gorm.DB, Sector_details *[]Sector_details) (err error) {
	err = db.Find(Sector_details).Error
	if err != nil {
		return err
	}
	return nil
}

func GetIndex(db *gorm.DB, Index_details *[]Index_details) (err error) {
	err = db.Find(Index_details).Error
	if err != nil {
		return err
	}
	return nil
}

func GetLoginInfo(db *gorm.DB, LoginInfo *[]UserInfo) (err error) {
	err = db.Find(LoginInfo).Error
	if err != nil {
		return err
	}
	return nil
}

func PostLoginInfo(db *gorm.DB, LoginInfo *UserInfo) (err error) {
	err = db.Create(LoginInfo).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPortfolio(db *gorm.DB, Portfolio *[]Portfolio) (err error) {
	err = db.Find(Portfolio).Error
	if err != nil {
		return err
	}
	return nil
}

func PostPortfolio(db *gorm.DB, Portfolio *Portfolio) (err error) {
	err = db.Create(Portfolio).Error
	if err != nil {
		return err
	}
	return nil
}

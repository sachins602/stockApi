package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

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

func SaveIndexDetailsToDb(index IndexInfo) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/stock")

	if err != nil {
		fmt.Println("error validating sql.Open arguments")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping")
		panic(err.Error())
	}

	if err != nil {
		fmt.Println("error validating db.Query arguments")
	}

	db.Exec("CREATE TABLE IF NOT EXISTS indices (`IndexName` VARCHAR(100) PRIMARY KEY, `FullName` VARCHAR(100), `Turnover` DOUBLE, `DailyGain` DOUBLE, `TotalPositiveGainer` DOUBLE, `TotalNegativeGainer` DOUBLE, `Pe` DOUBLE, `Pb` DOUBLE, `Peg` DOUBLE, `Roe` DOUBLE, `Alpha` DOUBLE, `Beta` DOUBLE,`SharpeRatio` DOUBLE, `Macd` DOUBLE,`Rsi` DOUBLE, `YearlyPercentChange` DOUBLE, `MacdSignal` DOUBLE, `SmaTwo` DOUBLE,`Ltp` DOUBLE, `TotalDividendYield` DOUBLE,`Roa` DOUBLE);")

	for _, v := range index.HomeTable {

		ins, err := db.Prepare("INSERT INTO indices(`IndexName`, `FullName`, `Turnover`, `DailyGain`, `TotalPositiveGainer`, `TotalNegativeGainer`, `Pe`, `Pb`, `Peg`, `Roe`, `Alpha`, `Beta`, `SharpeRatio`, `Macd`, `Rsi`, `YearlyPercentChange`, `MacdSignal`, `SmaTwo`, `Ltp`, `TotalDividendYield`, `Roa`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")

		if err != nil {
			fmt.Println("error validating db.Exec arguments")
		}
		ins.Exec(v.IndexName, v.FullName, v.TurnoverValues, v.DailyGain, v.TotalPositiveGainer, v.TotalNegativeGainer, v.Pe, v.Pb, v.Peg, v.Roe, v.Alpha, v.Beta, v.SharpeRatio, v.Macd, v.Rsi, v.YearlyPercentChange, v.Macdsignal, v.Sma200, v.Ltp, v.TotalDividendYield, v.Roa)
	}

}

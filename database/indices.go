package database

import (
	"database/sql"
	"fmt"
	model "goapi/models"

	_ "github.com/go-sql-driver/mysql"
)

func SaveIndexDetailsToDb(indexInfo model.IndexInfo) {
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

	db.Exec("CREATE TABLE IF NOT EXISTS index_details (`IndexName` VARCHAR(100) PRIMARY KEY, `FullName` VARCHAR(100), `Turnover` DOUBLE, `DailyGain` DOUBLE, `TotalPositiveGainer` DOUBLE, `TotalNegativeGainer` DOUBLE, `Pe` DOUBLE, `Pb` DOUBLE, `Peg` DOUBLE, `Roe` DOUBLE, `Alpha` DOUBLE, `Beta` DOUBLE,`SharpeRatio` DOUBLE, `Macd` DOUBLE,`Rsi` DOUBLE, `YearlyPercentChange` DOUBLE, `MacdSignal` DOUBLE, `SmaTwo` DOUBLE,`Ltp` DOUBLE, `TotalDividendYield` DOUBLE,`Roa` DOUBLE);")

	for _, v := range indexInfo.HomeTable {

		ins, err := db.Prepare("INSERT INTO index_details(`IndexName`, `FullName`, `Turnover`, `DailyGain`, `TotalPositiveGainer`, `TotalNegativeGainer`, `Pe`, `Pb`, `Peg`, `Roe`, `Alpha`, `Beta`, `SharpeRatio`, `Macd`, `Rsi`, `YearlyPercentChange`, `MacdSignal`, `SmaTwo`, `Ltp`, `TotalDividendYield`, `Roa`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")

		if err != nil {
			fmt.Println("error validating db.Exec arguments")
		}
		ins.Exec(v.IndexName, v.FullName, v.TurnoverValues, v.DailyGain, v.TotalPositiveGainer, v.TotalNegativeGainer, v.Pe, v.Pb, v.Peg, v.Roe, v.Alpha, v.Beta, v.SharpeRatio, v.Macd, v.Rsi, v.YearlyPercentChange, v.Macdsignal, v.Sma200, v.Ltp, v.TotalDividendYield, v.Roa)
	}

}

package database

import (
	"database/sql"
	"fmt"
	model "goapi/models"

	_ "github.com/go-sql-driver/mysql"
)

func SaveDetailsToDb(nepseInfo model.NepseInfo) {

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
	db.Exec("CREATE TABLE IF NOT EXISTS stock_details (`StockName` VARCHAR(100) PRIMARY KEY, `LastPrice` DOUBLE, `TurnOver` DOUBLE, `Change` DOUBLE, `High` DOUBLE, `Low` DOUBLE, `Open` DOUBLE, `ShareTraded` DOUBLE);")
	db.Exec("CREATE TABLE IF NOT EXISTS sector_details (`SectorName` VARCHAR(100) PRIMARY KEY, `Turnover` DOUBLE, `Quantity` DOUBLE);")
	db.Exec("CREATE TABLE IF NOT EXISTS broker_details (`BrokerNumber` Double PRIMARY KEY, `BrokerName` VARCHAR(100), `Purchase` DOUBLE, `Sales` DOUBLE, `Matching` DOUBLE, `Total` DOUBLE);")

	for _, v := range nepseInfo.Turnover.Detail {
		// stockName = append(stockName, v.S)
		// lastPrice = append(lastPrice, v.Lp)
		// turnOver = append(turnOver, v.T)
		// change = append(change, v.Pc)
		// high = append(high, v.H)
		// low = append(low, v.L)
		// open = append(open, v.Op)
		// shareTraded = append(shareTraded, v.Q)

		ins, err := db.Prepare("INSERT INTO stock_details(`StockName`, `LastPrice`, `TurnOver`, `Change`, `High`, `Low`, `Open`, `ShareTraded`) VALUES (?, ?, ?, ?, ?, ?, ?, ?);")

		if err != nil {
			fmt.Println("error validating db.Exec arguments")
		}
		ins.Exec(v.S, v.Lp, v.T, v.Pc, v.H, v.L, v.Op, v.Q)
	}

	for _, s := range nepseInfo.Sector.Detail {
		sec, err := db.Prepare("INSERT INTO sector_details(`SectorName`, `Turnover`, `Quantity`) VALUES (?, ?, ?);")

		if err != nil {
			fmt.Println("error validating db.Exec arguments")
		}
		sec.Exec(s.S, s.T, s.Q)
	}

	for _, b := range nepseInfo.Broker.Detail {
		bro, err := db.Prepare("INSERT INTO broker_details(`BrokerNumber`, `BrokerName`, `Purchase`, `Sales`, `Matching`, `Total`) VALUES (?, ?, ?, ?, ?, ?);")

		if err != nil {
			fmt.Println("error validating db.Exec arguments")
		}
		bro.Exec(b.B, b.N, b.P, b.S, b.M, b.T)
	}
}

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

	// var stockName []string
	// var lastPrice []float64
	// var turnOver []float64
	// var change []float64
	// var high []float64
	// var low []float64
	// var open []float64
	// var shareTraded []float64

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

}

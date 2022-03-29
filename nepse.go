package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func nepseDetails() {

	resp, err := http.Get("https://merolagani.com/handlers/webrequesthandler.ashx?type=market_summary")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var nepseInfo NepseInfo

	err = json.Unmarshal([]byte(body), &nepseInfo)

	if err == nil {
		fmt.Println("error")
	}

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

		ins, err := db.Prepare("INSERT INTO stock_details(`stockName`, `lastPrice`, `turnOver`, `change`, `high`, `low`, `open`, `shareTraded`) VALUES (?, ?, ?, ?, ?, ?, ?, ?);")

		if err != nil {
			fmt.Println("error validating db.Exec arguments")
		}
		ins.Exec(v.S, v.Lp, v.T, v.Pc, v.H, v.L, v.Op, v.Q)
	}

}

// CREATE TABLE stock_details(
// `stockName` VARCHAR(100),
// `lastPrice` DOUBLE,
// `turnOver` DOUBLE,
// `change` DOUBLE,
// `high` DOUBLE,
// `low` DOUBLE,
// `open` DOUBLE,
// `shareTraded` DOUBLE
// );

// INSERT INTO stock_details(`stockName`, `lastPrice`, `turnOver`, `change`, `high`, `low`, `open`, `shareTraded`)
// VALUES ('SHIVM', 1047, 79199051.5, -0.19, 1083, 1025, 1065, 75614);

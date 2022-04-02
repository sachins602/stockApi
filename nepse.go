package main

import (
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
	saveDetailsToDb(nepseInfo)
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

package spiders

import (
	"encoding/json"
	"fmt"
	database "goapi/spiderdb"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func NepseDetails() {

	resp, err := http.Get("https://merolagani.com/handlers/webrequesthandler.ashx?type=market_summary")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var nepseInfo database.NepseInfo

	err = json.Unmarshal([]byte(body), &nepseInfo)

	if err == nil {
		fmt.Println("error", err)
	}
	database.SaveDetailsToDb(nepseInfo)
}

// CREATE TABLE stocks(
// `StockName` VARCHAR(100),
// `LastPrice` DOUBLE,
// `TurnOver` DOUBLE,
// `Change` DOUBLE,
// `High` DOUBLE,
// `Low` DOUBLE,
// `Open` DOUBLE,
// `ShareTraded` DOUBLE
// );

// INSERT INTO stocks(`StockName`, `LastPrice`, `TurnOver`, `Change`, `High`, `Low`, `Open`, `ShareTraded`)
// VALUES ('SHIVM', 1047, 79199051.5, -0.19, 1083, 1025, 1065, 75614);

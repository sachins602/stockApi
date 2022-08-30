package spiders

import (
	"encoding/json"
	"fmt"
	database "goapi/spiderdb"
	"io/ioutil"
	"log"
	"net/http"

	//"os"

	_ "github.com/go-sql-driver/mysql"
)

func NepseIndexHistory() {

	// loading from env not working needs work
	resp, err := http.Get("https://nepsealpha.com/trading/1/history?symbol=NEPSE&resolution=1D&from=1325355300&to=1661299200&pass=ok&force=23757&currencyCode=NRS")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var historyData database.HistoryData

	err = json.Unmarshal([]byte(body), &historyData)

	if err != nil {
		fmt.Println("error", err)
	}
	database.SaveIndexHistoryToDb(historyData)
}

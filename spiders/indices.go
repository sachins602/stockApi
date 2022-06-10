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

func IndexDetails() {

	resp, err := http.Get("https://nepsealpha.com/api/smx9841/dashboard_board")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var indices database.IndexInfo

	err = json.Unmarshal([]byte(body), &indices)

	if err != nil {
		fmt.Println("error", err)
	}
	database.SaveIndexDetailsToDb(indices)
}

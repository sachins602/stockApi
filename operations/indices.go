package operations

import (
	"encoding/json"
	"fmt"
	database "goapi/database"
	models "goapi/models"
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

	var indexInfo models.IndexInfo

	err = json.Unmarshal([]byte(body), &indexInfo)

	if err == nil {
		fmt.Println("error")
	}
	database.SaveIndexDetailsToDb(indexInfo)
}

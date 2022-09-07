package spiders

import (
	"encoding/json"
	"fmt"
	database "goapi/spiderdb"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	//"os"

	_ "github.com/go-sql-driver/mysql"
)

func NepseIndexHistory() {

	// loading from env not working needs work
	resp, err := http.Get(os.Getenv("HISTORY_LINK"))
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
	fmt.Println(historyData)
	database.SaveIndexHistoryToDb(historyData)
}

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

func IndexDetails() {

	resp, err := http.Get(os.Getenv("INDEX_LINK"))
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

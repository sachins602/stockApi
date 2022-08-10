package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type NepseLiveInfo struct {
	Mt      string `json:"mt"`
	Overall struct {
		D  string `json:"d"`
		T  string `json:"t"`
		Q  string `json:"q"`
		Tn string `json:"tn"`
		St string `json:"st"`
		Mc string `json:"mc"`
		Fc string `json:"fc"`
	} `json:"overall"`
	Turnover struct {
		Date   string `json:"date"`
		Detail []struct {
			S  string  `json:"s"`
			N  string  `json:"n"`
			Lp float64 `json:"lp"`
			T  float64 `json:"t"`
			Pc float64 `json:"pc"`
			H  float64 `json:"h"`
			L  float64 `json:"l"`
			Op float64 `json:"op"`
			Q  float64 `json:"q"`
		} `json:"detail"`
	} `json:"turnover"`
	Sector struct {
		Date   string `json:"date"`
		Detail []struct {
			S string  `json:"s"`
			T float64 `json:"t"`
			Q float64 `json:"q"`
		} `json:"detail"`
	} `json:"sector"`
	Broker struct {
		Date   string `json:"date"`
		Detail []struct {
			B string  `json:"b"`
			N string  `json:"n"`
			P float64 `json:"p"`
			S float64 `json:"s"`
			M float64 `json:"m"`
			T float64 `json:"t"`
		} `json:"detail"`
	} `json:"broker"`
	Stock struct {
		Date   string `json:"date"`
		Detail []struct {
			S  string  `json:"s"`
			Lp float64 `json:"lp"`
			C  int     `json:"c"`
			Q  float64 `json:"q"`
		} `json:"detail"`
	} `json:"stock"`
}

func SaveLiveDetailsToDb(nepseInfo NepseLiveInfo) {

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

	for _, v := range nepseInfo.Turnover.Detail {
		// stockName = append(stockName, v.S)
		// lastPrice = append(lastPrice, v.Lp)
		// turnOver = append(turnOver, v.T)
		// change = append(change, v.Pc)
		// high = append(high, v.H)
		// low = append(low, v.L)
		// open = append(open, v.Op)
		// shareTraded = append(shareTraded, v.Q)

		//UPDATE `stocks` SET `LastPrice` = 10, `TurnOver` = 10, `Change` = 10, `High` = 10, `Low` = 10, `Open` = 10, `ShareTraded` = 10 WHERE StockName="adbl";
		ins, err := db.Prepare("UPDATE `stocks` SET `LastPrice` = ?, `TurnOver` = ?, `Change` = ?, `High` = ?, `Low` = ?, `Open` = ?, `ShareTraded` = ? WHERE StockName = ?")
		if err != nil {
			fmt.Println("error validating db.Prepare arguments")
			panic(err.Error())
		}
		ins.Exec(v.Lp, v.T, v.Pc, v.H, v.L, v.Op, v.Q, v.S)
	}

}

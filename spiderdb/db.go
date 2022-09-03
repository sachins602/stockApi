package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type NepseInfo struct {
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

func SaveDetailsToDb(nepseInfo NepseInfo) {

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
	db.Exec("CREATE TABLE IF NOT EXISTS stocks (`StockName` VARCHAR(100) PRIMARY KEY, `LastPrice` DOUBLE, `TurnOver` DOUBLE, `Change` DOUBLE, `High` DOUBLE, `Low` DOUBLE, `Open` DOUBLE, `ShareTraded` DOUBLE);")
	db.Exec("CREATE TABLE IF NOT EXISTS sectors (`SectorName` VARCHAR(100) PRIMARY KEY, `Turnover` DOUBLE, `Quantity` DOUBLE);")
	db.Exec("CREATE TABLE IF NOT EXISTS brokers (`BrokerNumber` Double PRIMARY KEY, `BrokerName` VARCHAR(100), `Purchase` DOUBLE, `Sales` DOUBLE, `Matching` DOUBLE, `Total` DOUBLE);")
	db.Exec("DELETE FROM stocks;")
	db.Exec("DELETE FROM sectors;")
	db.Exec("DELETE FROM brokers;")
	for _, v := range nepseInfo.Turnover.Detail {
		// stockName = append(stockName, v.S)
		// lastPrice = append(lastPrice, v.Lp)
		// turnOver = append(turnOver, v.T)
		// change = append(change, v.Pc)
		// high = append(high, v.H)
		// low = append(low, v.L)
		// open = append(open, v.Op)
		// shareTraded = append(shareTraded, v.Q)

		ins, err := db.Prepare("INSERT INTO stocks(`StockName`, `LastPrice`, `TurnOver`, `Change`, `High`, `Low`, `Open`, `ShareTraded`) VALUES (?, ?, ?, ?, ?, ?, ?, ?);")

		if err != nil {
			fmt.Println("error validating db.Exec arguments")
		}
		ins.Exec(v.S, v.Lp, v.T, v.Pc, v.H, v.L, v.Op, v.Q)
	}

	for _, s := range nepseInfo.Sector.Detail {
		sec, err := db.Prepare("INSERT INTO sectors(`SectorName`, `Turnover`, `Quantity`) VALUES (?, ?, ?);")

		if err != nil {
			fmt.Println("error validating db.Exec arguments")
		}
		sec.Exec(s.S, s.T, s.Q)
	}

	for _, b := range nepseInfo.Broker.Detail {
		bro, err := db.Prepare("INSERT INTO brokers(`BrokerNumber`, `BrokerName`, `Purchase`, `Sales`, `Matching`, `Total`) VALUES (?, ?, ?, ?, ?, ?);")

		if err != nil {
			fmt.Println("error validating db.Exec arguments")
		}
		bro.Exec(b.B, b.N, b.P, b.S, b.M, b.T)
	}
}

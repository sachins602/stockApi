package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type HistoryData struct {
	T []int    `json:"t"`
	C []string `json:"c"`
	O []string `json:"o"`
	H []string `json:"h"`
	L []string `json:"l"`
	V []string `json:"v"`
	S string   `json:"s"`
}

func SaveIndexHistoryToDb(nepseHistory HistoryData) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/stock")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.Exec("CREATE TABLE IF NOT EXISTS historic (`Scrip` VARCHAR(100), `Time` DOUBLE  PRIMARY KEY, `Close` DOUBLE)")
	db.Exec("CREATE TABLE IF NOT EXISTS historicChart (`Time` DOUBLE  PRIMARY KEY, `Close` DOUBLE, `Open` DOUBLE, `High` DOUBLE, `Low` DOUBLE, `Volume` DOUBLE)")
	db.Exec("DELETE FROM historic;")
	db.Exec("DELETE FROM historicChart;")
	fmt.Println("here")

	for j := 0; j < len(nepseHistory.T); j++ {
		o := nepseHistory.O[j]
		h := nepseHistory.H[j]
		c := nepseHistory.C[j]
		l := nepseHistory.L[j]
		v := nepseHistory.V[j]
		t := nepseHistory.T[j]

		queryChart := fmt.Sprintf("INSERT INTO historicChart (`Time`, `Close`, `Open`, `High`, `Low`, `Volume`) VALUES (%d, %s, %s, %s, %s, %s)", t, c, o, h, l, v)
		query := fmt.Sprintf("INSERT INTO historic (`Scrip`, `Time`, `Close`) VALUES (%s, %d, %s)", `"Nepse"`, t, c)

		fmt.Println(query)
		db.Exec(queryChart)
		db.Exec(query)

	}
}

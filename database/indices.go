package database

import (
	"database/sql"
	"fmt"
	model "goapi/models"

	_ "github.com/go-sql-driver/mysql"
)

func SaveIndexDetailsToDb(indexInfo model.IndexInfo) {
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

	db.Exec("CREATE TABLE IF NOT EXISTS index_details (`IndexName` VARCHAR(100) PRIMARY KEY, `LastPrice` DOUBLE, `TurnOver` DOUBLE, `Change` DOUBLE, `High` DOUBLE, `Low` DOUBLE, `Open` DOUBLE, `ShareTraded` DOUBLE);")

	for _, v := range indexInfo.Date.Turnover.Detail {

		ins, err := db.Prepare("INSERT INTO index_details(`IndexName`, `LastPrice`, `TurnOver`, `Change`, `High`, `Low`, `Open`, `ShareTraded`) VALUES (?, ?, ?, ?, ?, ?, ?, ?);")

		if err != nil {
			fmt.Println("error validating db.Exec arguments")
		}
		ins.Exec(v.S, v.Lp, v.T, v.Pc, v.H, v.L, v.Op, v.Q)
	}

}

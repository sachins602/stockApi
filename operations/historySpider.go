package operations

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"goapi/models"
	"io/ioutil"
	"log"
	"net/http"
)

func ScrapeAllHistory() {
	scripSymbol := []string{"NEPSE", "ACLBSL", "ADBL", "AHPC", "AIL", "AKJCL", "AKPL", "ALBSL", "ALICL", "API", "BARUN", "BBC", "BFC", "BNHC", "BNT", "BOKL", "BPCL", "CBBL", "CBL", "CCBL", "CFCL", "CGH", "CHCL", "CHDC", "CHL", "CIT", "CLBSL", "CMF1", "CMF2", "CORBL", "CZBIL", "CZBILP", "DDBL", "DHPL", "EBL", "EDBL", "EIC", "EICPO", "ENL", "FMDBL", "FOWAD", "GBBL", "GBIME", "GBLBS", "GFCL", "GHL", "GIC", "GILB", "GIMES1", "GLBSL", "GLH", "GLICL", "GMFBS", "GMFIL", "GRDBL", "GUFL", "HDHPC", "HDL", "HGI", "HIDCL", "HPPL", "HURJA", "ICFC", "ICFCD83", "IGI", "ILBS", "JALPA", "JBBL", "JBLB", "JFL", "JFLPO", "JLI", "JOSHI", "JSLBB", "KBL", "KEF", "KKHC", "KLBSL", "KMCDB", "KPCL", "KRBL", "KSBBL", "LBBL", "LBL", "LEC", "LEMF", "LGIL", "LICN", "LLBS", "LUK", "MBJC", "MBL", "MDB", "MEGA", "MEGAPO", "MEN", "MERO", "MFIL", "MHNL", "MKJC", "MKLB", "MLBBL", "MLBL", "MLBLPO", "MLBS", "MLBSL", "MMF1", "MMFDB", "MNBBL", "MPFL", "MSLB", "NABBC", "NABIL", "NBF2", "NBF3", "NBL", "NBLD85", "NCCB", "NEF", "NESDO", "NFS", "NGPL", "NHDL", "NHPC", "NIBLPF", "NIBSF2", "NICA", "NICAD8283", "NICBF", "NICGF", "NICL", "NICLBSL", "NICSF", "NIFRA", "NIL", "NLBBL", "NLG", "NLIC", "NLICL", "NMB", "NMB50", "NMBHF1", "NMBMF", "NMFBS", "NRIC", "NRN", "NSLB", "NTC", "NUBL", "NYADI", "OHL", "PCBL", "PFL", "PIC", "PICL", "PLI", "PLIC", "PMHPL", "PPCL", "PRIN", "PROFL", "PRVU", "PSF", "RADHI", "RBCL", "RBCLPO", "RHPL", "RLFL", "RLI", "RMDC", "RMF1", "RSDC", "RULB", "RURU", "SABSL", "SADBL", "SAEF", "SAHAS", "SANIMA", "SAPDBL", "SBCF", "SBD87", "SBI", "SBL", "SCB", "SDLBSL", "SEF", "SFCL", "SFMF", "SGI", "SHEL", "SHINE", "SHIVM", "SHL", "SHPC", "SIC", "SICL", "SIFC", "SIGS2", "SIL", "SINDU", "SJCL", "SKBBL", "SLBBL", "SLBSL", "SLCF", "SLI", "SLICL", "SMATA", "SMB", "SMBPO", "SMFBS", "SMFDB", "SPC", "SPDL", "SRBL", "SSHL", "STC", "SWBBL", "TPC", "TRH", "UIC", "ULI", "UMHL", "UMRH", "UNHPL", "UNL", "UPCL", "UPPER", "USLB", "VLBS", "WNLB"}

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

	for i := 0; i < len(scripSymbol); i++ {
		query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (`Time` DOUBLE, `Close` DOUBLE, `Open` DOUBLE, `High` DOUBLE, `Low` DOUBLE, `Volume` DOUBLE)", scripSymbol[i])
		db.Exec(query)
	}

	for _, s := range scripSymbol {
		url := fmt.Sprintf("https://nepsealpha.com/trading/1/history?symbol=%s&resolution=1D&from=1611705600&to=1650240000&pass=ok&force=161259&currencyCode=NRS", s)
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var scripHistory models.ScripHistory

		err = json.Unmarshal([]byte(body), &scripHistory)

		if err == nil {
			fmt.Println("error")
		}

		// for _, v := range scripHistory.T {
		// 	query := fmt.Sprintf("INSERT INTO %s (`Time`) VALUES (%d, %s, %s, %s)", s, t)

		// 	ins, err := db.Prepare("INSERT INTO %s(`IndexName`, `FullName`, `Turnover`, `DailyGain`, `TotalPositiveGainer`, `TotalNegativeGainer`, `Pe`, `Pb`, `Peg`, `Roe`, `Alpha`, `Beta`, `SharpeRatio`, `Macd`, `Rsi`, `YearlyPercentChange`, `MacdSignal`, `SmaTwo`, `Ltp`, `TotalDividendYield`, `Roa`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")

		// 	if err != nil {
		// 		fmt.Println("error validating db.Exec arguments")
		// 	}
		// 	ins.Exec(v)
		// }

		for j := 0; j < len(scripHistory.T); j++ {
			for _, t := range scripHistory.T {
				query := fmt.Sprintf("INSERT INTO %s (`Time`) VALUES (?)", s)
				ins, err := db.Prepare(query)
				if err != nil {
					fmt.Println("error validating db.Exec arguments")
				}
				ins.Exec(t)
			}
			for _, c := range scripHistory.C {
				query := fmt.Sprintf("INSERT INTO %s (`Close`) VALUES (?)", s)
				ins, err := db.Prepare(query)
				if err != nil {
					fmt.Println("error validating db.Exec arguments")
				}
				ins.Exec(c)
			}
			for _, o := range scripHistory.O {
				query := fmt.Sprintf("INSERT INTO %s (`Open`) VALUES (?)", s)
				ins, err := db.Prepare(query)
				if err != nil {
					fmt.Println("error validating db.Exec arguments")
				}
				ins.Exec(o)
			}
			for _, h := range scripHistory.H {
				query := fmt.Sprintf("INSERT INTO %s (`High`) VALUES (?)", s)
				ins, err := db.Prepare(query)
				if err != nil {
					fmt.Println("error validating db.Exec arguments")
				}
				ins.Exec(h)
			}
			for _, l := range scripHistory.L {
				query := fmt.Sprintf("INSERT INTO %s (`Low`) VALUES (?)", s)
				ins, err := db.Prepare(query)
				if err != nil {
					fmt.Println("error validating db.Exec arguments")
				}
				ins.Exec(l)
			}
			for _, v := range scripHistory.V {
				query := fmt.Sprintf("INSERT INTO %s (`Volume`) VALUES (?)", s)
				ins, err := db.Prepare(query)
				if err != nil {
					fmt.Println("error validating db.Exec arguments")
				}
				ins.Exec(v)
			}

		}

		// for _, t := range scripHistory.T {
		// 	query := fmt.Sprintf("INSERT INTO %s (`Time`) VALUES (%d, %s, %s, %s)", s, t)
		// 	db.Exec(query)
		// }

		// for _, c := range scripHistory.C {
		// 	query := fmt.Sprintf("INSERT INTO %s (`Close`) VALUES (%s)", s, c)
		// 	db.Exec(query)
		// }

		// for _, o := range scripHistory.O {
		// 	query := fmt.Sprintf("INSERT INTO %s (`Open`) VALUES (%s)", s, o)
		// 	db.Exec(query)
		// }

		// for _, h := range scripHistory.H {
		// 	query := fmt.Sprintf("INSERT INTO %s (`High`) VALUES (%s)", s, h)
		// 	db.Exec(query)
		// }

		// for _, l := range scripHistory.L {
		// 	query := fmt.Sprintf("INSERT INTO %s (`Low`) VALUES (%s)", s, l)
		// 	db.Exec(query)
		// }

		// for _, v := range scripHistory.V {
		// 	query := fmt.Sprintf("INSERT INTO %s (`Volume`) VALUES (%s)", s, v)
		// 	db.Exec(query)
		// }

		resp.Body.Close()

	}

}

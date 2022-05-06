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
	scripSymbol := []string{"NEPSE", "ACLBSL", "ADBL", "AHPC", "AIL", "AKJCL", "AKPL", "ALBSL", "ALICL", "API", "BARUN", "BBC", "BFC", "BNHC", "BNT", "BOKL", "BPCL", "CBBL", "CBL", "CCBL", "CFCL", "CGH", "CHCL", "CHDC", "CHL", "CIT", "CLBSL", "CORBL", "CZBIL", "DDBL", "DHPL", "EBL", "EDBL", "EIC", "ENL", "FMDBL", "FOWAD", "GBBL", "GBIME", "GBLBS", "GFCL", "GHL", "GIC", "GILB", "GLBSL", "GLH", "GLICL", "GMFBS", "GMFIL", "GRDBL", "GUFL", "HDHPC", "HDL", "HGI", "HIDCL", "HPPL", "HURJA", "ICFC", "IGI", "ILBS", "JALPA", "JBBL", "JBLB", "JFL", "JLI", "JOSHI", "JSLBB", "KBL", "KEF", "KKHC", "KLBSL", "KMCDB", "KPCL", "KRBL", "KSBBL", "LBBL", "LBL", "LEC", "LGIL", "LICN", "LLBS", "LUK", "MBJC", "MBL", "MDB", "MEGA", "MEN", "MERO", "MFIL", "MHNL", "MKJC", "MKLB", "MLBBL", "MLBL", "MLBS", "MLBSL", "MMFDB", "MNBBL", "MPFL", "MSLB", "NABBC", "NABIL", "NBF2", "NBF3", "NBL", "NBLD85", "NCCB", "NEF", "NESDO", "NFS", "NGPL", "NHDL", "NHPC", "NIBLPF", "NIBSF2", "NICA", "NICAD8283", "NICBF", "NICGF", "NICL", "NICLBSL", "NICSF", "NIFRA", "NIL", "NLBBL", "NLG", "NLIC", "NLICL", "NMB", "NRIC", "NRN", "NSLB", "NTC", "NUBL", "NYADI", "OHL", "PCBL", "PFL", "PIC", "PICL", "PLI", "PLIC", "PMHPL", "PPCL", "PRIN", "PROFL", "PRVU", "PSF", "RADHI", "RBCL", "RHPL", "RLFL", "RLI", "RMDC", "RSDC", "RULB", "RURU", "SABSL", "SADBL", "SAEF", "SAHAS", "SANIMA", "SAPDBL", "SBCF", "SBD87", "SBI", "SBL", "SCB", "SDLBSL", "SEF", "SFCL", "SGI", "SHEL", "SHINE", "SHIVM", "SHL", "SHPC", "SIC", "SICL", "SIFC", "SIGS2", "SIL", "SINDU", "SJCL", "SKBBL", "SLBBL", "SLBSL", "SLCF", "SLI", "SLICL", "SMATA", "SMB", "SMFBS", "SMFDB", "SPC", "SPDL", "SRBL", "SSHL", "STC", "SWBBL", "TPC", "TRH", "UIC", "ULI", "UMHL", "UMRH", "UNHPL", "UNL", "UPCL", "UPPER", "USLB", "VLBS", "WNLB"}

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

	tbQuery := "CREATE TABLE IF NOT EXISTS historic (`Scrip` VARCHAR(10),`Time` DOUBLE, `Close` VARCHAR(20), `Open` VARCHAR(20), `High` VARCHAR(20), `Low` VARCHAR(20), `Volume` VARCHAR(20))"
	db.Exec(tbQuery)

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
		// fmt.Println(scripHistory)
		if err != nil {
			fmt.Println("error", err)
		}
		fs := fmt.Sprintf(`"%s"`, s)
		for j := 0; j < len(scripHistory.T); j++ {
			o := scripHistory.O[j]
			h := scripHistory.H[j]
			c := scripHistory.C[j]
			l := scripHistory.L[j]
			v := scripHistory.V[j]
			t := scripHistory.T[j]

			query := fmt.Sprintf("INSERT INTO historic (`Scrip`, `Time`, `Close`, `Open`, `High`, `Low`, `Volume`) VALUES (%s, %d, %s, %s, %s, %s, %s)", fs, t, c, o, h, l, v)
			db.Exec(query)
		}

		resp.Body.Close()

	}

}

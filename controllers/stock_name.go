package controllers

import (
	"goapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// get Stocks
func GetStocks(c *gin.Context) {

	var stocks []models.Stock

	models.DB.Find(&stocks)

	c.JSON(http.StatusOK, stocks)

}

//get Broker
func GetBroker(c *gin.Context) {
	var brokers []models.Broker

	models.DB.Find(&brokers)

	c.JSON(http.StatusOK, brokers)

}

// get sector
func GetSector(c *gin.Context) {
	var sectors []models.Sector

	models.DB.Find(&sectors)

	c.JSON(http.StatusOK, sectors)

}

//get index
func GetIndex(c *gin.Context) {
	var indices []models.Index

	models.DB.Find(&indices)

	c.JSON(http.StatusOK, indices)
}

//get gainer
func GetGainer(c *gin.Context) {
	var gainers []models.Gainer

	models.DB.Table("top_gainers").Find(&gainers)

	c.JSON(http.StatusOK, gainers)
}

//get small gainer
func GetSmallGainer(c *gin.Context) {
	var smallGainers []models.Gainer

	models.DB.Table("top_gainers").Find(&smallGainers, "id < ?", 5)

	c.JSON(http.StatusOK, smallGainers)
}

//get loser
func GetLoser(c *gin.Context) {
	var losers []models.Gainer

	models.DB.Table("top_losers").Find(&losers)

	c.JSON(http.StatusOK, losers)
}

//get small loser
func GetSmallLoser(c *gin.Context) {
	var smallLosers []models.Gainer

	models.DB.Table("top_losers").Find(&smallLosers, "id < ?", 5)

	c.JSON(http.StatusOK, smallLosers)
}

//get subIndex
func GetSubIndex(c *gin.Context) {
	var subIndices []models.SubIndex

	models.DB.Table("sub_indices").Find(&subIndices)

	c.JSON(http.StatusOK, subIndices)
}

//get NEPSE data
func GetNepse(c *gin.Context) {
	var nepse models.Index

	models.DB.First(&nepse, "IndexName = ?", "NEPSE")

	c.JSON(http.StatusOK, nepse)
}

// get NEPSE index history data
func GetNepseIndexHistory(c *gin.Context) {
	var nepseHistoric []models.Historic

	if err := models.DB.Where("time > 1635865190").Table("historic").Find(&nepseHistoric).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, nepseHistoric)
}

//get NEPSE history data
func GetNepseHistory(c *gin.Context) {
	var historics []models.Historic

	var scripIn string = c.Param("scrip")

	sector := map[string][]string{
		"corporate_debentures":         {"NICAD8283", "NBLD85"},
		"microfinance":                 {"ACLBSL", "ALBSL", "CBBL", "CLBSL", "DDBL", "FMDBL", "FOWAD", "GMFBS", "GILB", "GBLBS", "GLBSL", "ILBS", "JALPA", "JSLBB", "JBLB", "KMCDB", "KLBSL", "LLBS", "MLBSL", "MSLB", "MKLB", "MLBS", "MERO", "MMFDB", "MLBBL", "NSLB", "NLBBL", "NESDO", "NICLBSL", "NUBL", "RULB", "RMDC", "RSDC", "SABSL", "SDLBSL", "SMATA", "SLBSL", "SKBBL", "SMFDB", "SMB", "SWBBL", "SMFBS", "SLBBL", "USLB", "VLBS", "WNLB"},
		"commercial_banks":             {"ADBL", "BOKL", "CCBL", "CZBIL", "CBL", "EBL", "GBIME", "KBL", "LBL", "MBL", "MEGA", "NABIL", "NBL", "NCCB", "SBI", "NICA", "NMB", "PRVU", "PCBL", "SANIMA", "SBL", "SCB", "SRBL"},
		"non_life_insurance":           {"AIL", "EIC", "GIC", "HGI", "IGI", "LGIL", "NIL", "NICL", "NLG", "PRIN", "PIC", "PICL", "RBCL", "SIC", "SGI", "SICL", "SIL", "UIC"},
		"hydro_powers":                 {"AKJCL", "API", "AKPL", "AHPC", "BARUN", "BNHC", "BPCL", "CHL", "CHCL", "DHPL", "GHL", "GLH", "HDHPC", "HURJA", "HPPL", "JOSHI", "KPCL", "KKHC", "LEC", "MBJC", "MKJC", "MEN", "MHNL", "NHPC", "NHDL", "NGPL", "NYADI", "PMHPL", "PPCL", "RADHI", "RHPL", "RURU", "SAHAS", "SPC", "SHPC", "SJCL", "SSHL", "SHEL", "SPDL", "TPC", "UNHPL", "UMRH", "UMHL", "UPCL", "UPPER"},
		"life_insurance":               {"ALICL", "GLICL", "JLI", "LICN", "NLICL", "NLIC", "PLI", "PLIC", "RLI", "SLI", "SLICL", "ULI"},
		"finance":                      {"BFC", "CFCL", "GFCL", "GMFIL", "GUFL", "ICFC", "JFL", "MFIL", "MPFL", "NFS", "PFL", "PROFL", "RLFL", "SFCL", "SIFC"},
		"tradings":                     {"BBC", "STC"},
		"manufacturing_and_processing": {"BNT", "HDL", "SHIVM", "UNL"},
		"investment":                   {"CHDC", "CIT", "ENL", "HIDCL", "NIFRA", "NRN"},
		"hotels":                       {"CGH", "OHL", "SHL", "TRH"},
		"development_banks":            {"CORBL", "EDBL", "GBBL", "GRDBL", "JBBL", "KSBBL", "KRBL", "LBBL", "MLBL", "MDB", "MNBBL", "NABBC", "SAPDBL", "SADBL", "SHINE", "SINDU"},
		"mutual_fund":                  {"KEF", "LUK", "NEF", "NIBLPF"},
		"other":                        {"NTC", "NRIC"},
	}

	var realSector string

	for sectorName, sectorScrip := range sector {
		for _, scrip := range sectorScrip {
			if scrip == scripIn {
				realSector = sectorName
			}
		}
	}

	//models.DB.Find(&historics)

	if err := models.DB.Where("scrip = ? AND time > 1622732399", c.Param("scrip")).Table(realSector).Find(&historics).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// var total_data int64
	// models.DB.Where("scrip = ?", c.Param("scrip")).Table(c.Param("sector")).Count(&total_data)

	c.JSON(http.StatusOK, historics)
}

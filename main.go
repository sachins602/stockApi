package main

import (
	"fmt"

	"reflect"
	"sync"

	"time"

	"goapi/controllers"
	"goapi/middlewares"
	"goapi/models"

	sp "goapi/spiders"
	"goapi/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {

	binding.Validator = new(defaultValidator)

	//register custom validations
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bullishmarket", utils.BullishMarket)
	}

	sp.NepseDetails()
	sp.IndexDetails()
	sp.NepseIndexHistory()
	go doEvery(10 * time.Second)
	r := setupRouter()
	_ = r.Run(":8080")

}
func setupRouter() *gin.Engine {

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"*"},
	}))

	models.ConnectDataBase()

	public := r.Group("/api")
	admin := r.Group("/api/admin")
	admin.Use(middlewares.SetMiddlewareAuthentication())

	public.POST("/login", controllers.Login)
	public.POST("/register", controllers.Register)
	//get individual user details
	admin.GET("/user/:username", controllers.GetUserByUsername)

	//news scrape on request
	public.GET("/news", controllers.GetNews)

	public.GET("/stocks", controllers.GetStocks)
	public.GET("/broker", controllers.GetBroker)
	public.GET("/sector", controllers.GetSector)
	public.GET("/index", controllers.GetIndex)
	public.GET("/gainer", controllers.GetGainer)
	public.GET("/smallgainer", controllers.GetSmallGainer)
	public.GET("/loser", controllers.GetLoser)
	public.GET("/smallloser", controllers.GetSmallLoser)
	public.GET("/subindex", controllers.GetSubIndex)

	//individaul stock details
	public.GET("/stock/:scrip", controllers.GetStockByScrip)

	//historic data
	public.GET("/nepse", controllers.GetNepse)
	public.GET("/nepseHistoric", controllers.GetNepseIndexHistory)
	public.GET("/nepseHistory/:scrip", controllers.GetNepseHistory)

	//prediction data
	public.GET("/nepseHistoryPrediction/:scrip", controllers.GetNepseHistoryPrediction)
	public.GET("/nepsePrediction", controllers.GetNepsePrediction)
	public.GET("/nepseOneDayPrediction", controllers.GetNepseOneDayPrediction)

	//porfolio CRUD
	admin.POST("/portfolios", controllers.CreatePortfolio)
	admin.GET("/portfolios/:username", controllers.GetPortfolioByID)
	admin.PATCH("/portfolios/:username", controllers.UpdatePortfolio)
	admin.GET("/portfolios/:username/:scrip", controllers.DeletePortfolio)

	return r
}

func doEvery(d time.Duration) {
	for range time.Tick(d) {
		fmt.Println("Live data Scrapped")
		sp.LiveDetails()
	}
}

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &defaultValidator{}

func (v *defaultValidator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {

		v.lazyinit()

		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}

	return nil
}

func (v *defaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")

		// add any custom validations etc. here
	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

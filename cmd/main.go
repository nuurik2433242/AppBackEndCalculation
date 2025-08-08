package main

import (
	"go/adv-demo/internal/calcuationService"
	"go/adv-demo/internal/db"
	"go/adv-demo/internal/handlers"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)



func main() {
	datebase, err := db.InitDB()

	if err != nil {
		log.Fatalf("Could not connect to database %v", err)
	}

    e := echo.New()

	calcRepo := calcuationService.NewCalculationRepository(datebase)
 	calcService := calcuationService.NewCalculationService(calcRepo)
	calcHandler := handlers.NewCalculationHandler(calcService)

	

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", calcHandler.GetCalculation)
	e.POST("/calculations", calcHandler.PostCalculation)
	e.PATCH("/calculations/:id", calcHandler.PatchCalculation)
	e.DELETE("/calculations/:id", calcHandler.DeleteCalculation)

	e.Start("localhost:8080")
}

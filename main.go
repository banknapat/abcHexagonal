package main

import (
	"abc/database"
	"abc/handler"
	"abc/repository"
	"abc/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Mariadb()
	defer db.Close()

	r := repository.NewRepositoryAdapter(db)
	s := service.NewServiceAdapter(r)
	h := handler.NewHandlerAdaper(s)

	router := gin.Default()

	router.GET("/api/getBrees", h.GetHand)
	router.GET("/api/getBrees/:id", h.GetHandById)
	//router.POST("/add", h.Add)
	//router.PUT("/update/:id", h.Update)
	//router.DELETE("/delete/:id", h.Delete)

	err := router.Run(":9000")
	if err != nil {
		panic(err.Error())
	}

}

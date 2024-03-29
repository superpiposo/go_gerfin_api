package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	s := &http.Server{
		Addr:           ":5050",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	InitializeRoutes(router)
	s.ListenAndServe()

}

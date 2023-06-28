package main

import (
	"fmt"
	"go-log-generator/config"
	"go-log-generator/router"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()

	router.Router(app)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.Common.Port),
		Handler: app,
	}

	srv.ListenAndServe()

}

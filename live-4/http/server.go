package http

import (
	"net/http"
)

type WebServer struct {
	Products *model.Products
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	e := echo.New()
	e.GET("/product", w.getAll)
	e.logger.Fatal(e.Start(":8585"))
}

func (w WebServer) getAll(echo.Context) error {
	return w.JSON(http.StatusOK, w.Products)
}

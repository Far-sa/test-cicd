package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	router := echo.New()

	router.GET("/health-check", checkHealth)

	port := getEnv("HTTP_PORT", "8080")

	if err := router.Start(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf(fmt.Sprintf("could not start server : %v", err))
	}
}

func checkHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "everything is good",
	})
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

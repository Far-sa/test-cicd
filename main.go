package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	router := echo.New()

	router.GET("/health-check", checkHealth)
	router.GET("/serve-config-file", ServeConfigFile)
	router.GET("/show-password", ShowPassword)

	staticFilePath := getEnv("STATIC_URL", "./static")
	router.Static("/static", staticFilePath)

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

func ServeConfigFile(c echo.Context) error {

	filePath := getEnv("CONFIG_FILE_PATH", "config.yml")
	f, err := os.Open(filePath)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	content, rErr := io.ReadAll(f)
	if rErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, rErr.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": string(content),
	})
}

func ShowPassword(c echo.Context) error {

	key := getEnv("SECRET_KEY", "no-password")

	return c.JSON(http.StatusOK, echo.Map{
		"secret-key": key,
	})
}
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

/*
package main

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	fmt.Println("Go Program")
	server := echo.New()
	server.GET(path.Join("/"), Version)

	godotenv.Load()
	port := os.Getenv("PORT")

	address := fmt.Sprintf("%s:%s", "0.0.0.0", port)
	fmt.Println(address)
	server.Start(address)

}
func Version(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]interface{}{"version": 2})
}
*/

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Jsagudelo1704/Go/services"
)

const rutabase = "/api"

func main() {

	services.SetupRoutes(rutabase)
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
	}
	return ":" + port
}

package main

import (
	"CGV/docs"
	"CGV/routes"
)

func main() {

	docs.SwaggerInfo.Title = "CGV API Documentation"
	docs.SwaggerInfo.Description = "Ticketing Application by: \n- 1121009 Kevin N\n- 1121035 Anthony Kevin\n- 1122011 Darren\n- 1122015 Gregorius Reza\n- 1122027 Martinus"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := routes.SetupRouter()
	r.Run(":8080")

}

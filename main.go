package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    // Crea una instancia del enrutador Gin
    r := gin.Default()

    // Define una ruta
    r.GET("/saludo", func(c *gin.Context) {
        c.String(200, "¡Hola desde Gin!")
    })

    // Inicia el servidor en el puerto 8080
    r.Run(":8080")
}

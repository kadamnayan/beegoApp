// main.go
package main

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

// HomeController handles the home route
type HomeController struct {
	web.Controller
}

func (c *HomeController) Get() {
	c.Ctx.WriteString("Welcome to the Home Page!")
}

// FileController handles file routes
type FileController struct {
	web.Controller
}

func (c *FileController) Get() {
	// Get the file parameter from the request
	fileName := c.Ctx.Input.Param(":splat")
	fmt.Println("File requested:", fileName)
	if fileName == "" {

		c.Ctx.WriteString("Requested file not found")
		return
	}

	c.Ctx.WriteString("Requested file: " + fileName)
}

// Main function to set up the Beego application
func main() {
	// Define different routes
	web.Router("/", &HomeController{})
	// Change the route pattern to explicitly match file parameter after '/files/'
	web.Router("/files/*", &FileController{})

	// Start the Beego application
	web.Run()
}

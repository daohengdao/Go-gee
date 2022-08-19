package main

import (
	"gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(200, "<h1>Hello</h1>")
	})
	err := r.Run(":9999")
	if err != nil {
		return
	}

}

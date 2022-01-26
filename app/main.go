package main

import (
	"github.com/gkdgbilal/blog-rest-api/controller"
	"github.com/gkdgbilal/blog-rest-api/model"
)

func main() {
	model.Init()
	controller.Start()
}

package main

import (
	"myBlog/model"
	"myBlog/routes"
)

func main() {
	model.InitDb()
	
	routes.InitRouter()
}

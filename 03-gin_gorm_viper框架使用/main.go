package main

import (
	routers "gin-demo/router"
)

func main() {
	r := routers.InitRouters()
	r.Run(":8080")
}

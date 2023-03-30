package main

import (
	"github.com/byte-cats/autho/db"
	"github.com/bytecats/auth/routing"
	"github.com/bytecats/auth/serve"
)

func main() {
	defer db.DB.Close()

	serve.Serve(routing.InitRoutes())
}

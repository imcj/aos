package bootstrap

import (
	"log"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

type HTTPServerCommand struct {}

func (c HTTPServerCommand)Execute() {
	router := gin.New()

	server := endless.NewServer(":3000", router)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
	log.Printf("123")
}
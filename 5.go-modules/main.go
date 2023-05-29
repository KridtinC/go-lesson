package main

import (
	"go-mod-demo/src/model"

	"github.com/gin-gonic/gin"
)

func main() {
	var _ = model.Cat{}
	gin.New()
}

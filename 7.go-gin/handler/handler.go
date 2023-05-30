package handler

import "github.com/gin-gonic/gin"

type Post struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var posts = map[string]Post{
	"1": {
		ID:   "1",
		Name: "Kan",
	},
}

func GetPostHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(200, posts[id])
}

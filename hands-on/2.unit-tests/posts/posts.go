package posts

import "errors"

type Post struct {
	ID   string
	Name string
}

var posts = map[string]*Post{
	"1": {
		ID:   "1",
		Name: "myPost",
	},
	"2": {
		ID:   "2",
		Name: "mySecondPost",
	},
}

func GetPost(id string) (*Post, error) {
	if _, exist := posts[id]; !exist {
		return nil, errors.New("post not found")
	}
	return posts[id], nil
}

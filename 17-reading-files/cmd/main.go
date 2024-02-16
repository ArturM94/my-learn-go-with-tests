package main

import (
	blogposts "learn-go-with-tests/17-reading-files"
	"log"
	"os"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}

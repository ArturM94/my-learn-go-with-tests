package blogposts

import (
	"errors"
	"fmt"
	"io/fs"
	"strings"
)

const fileExtension = ".md"

var ErrWrongFileExtension = errors.New(fmt.Sprintf("file extension must be %s", fileExtension))

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		if !strings.HasSuffix(f.Name(), fileExtension) {
			return nil, ErrWrongFileExtension
		}
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

package blogposts_test

import (
	blogposts "learn-go-with-tests/17-reading-files"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
U
A
F`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	})
}

func TestFileExtension(t *testing.T) {
	cases := []struct {
		Name  string
		MapFS fstest.MapFS
	}{
		{
			Name: ".txt",
			MapFS: fstest.MapFS{
				"hello world.md":   {Data: []byte("hello")},
				"hello-world2.txt": {Data: []byte("world")},
			},
		},
		{
			Name: ".doc",
			MapFS: fstest.MapFS{
				"hello world.md":   {Data: []byte("hello")},
				"hello-world2.doc": {Data: []byte("world")},
			},
		},
		{
			Name: ".wrong",
			MapFS: fstest.MapFS{
				"hello world.md":     {Data: []byte("hello")},
				"hello-world2.wrong": {Data: []byte("world")},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			_, err := blogposts.NewPostsFromFS(test.MapFS)

			if err != blogposts.ErrWrongFileExtension {
				t.Errorf("got %v want %v", err, blogposts.ErrWrongFileExtension)
			}
		})
	}

}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v want %+v", got, want)
	}
}

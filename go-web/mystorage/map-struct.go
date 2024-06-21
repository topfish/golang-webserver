package mystorage

import (
)

type Post struct{
	Id 	int
	Content string
	Author	string
}

var PostById map[int]*Post
var PostByAuth map[string][]*Post

func Mystorage(post Post) {
	PostById[post.Id] = &post
	PostByAuth[post.Author] = append(PostByAuth[post.Author], &post)
}

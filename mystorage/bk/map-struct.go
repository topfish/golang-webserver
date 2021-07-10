package main

import (
	"fmt"
)

type Post struct{
	Id 	int
	Content	string
	Author	string
}

var PostById map[int]*Post
var PostByAuth map[string][]*Post

func Mystorage(post Post) {
	PostById[post.Id] = &post
	PostByAuth[post.Author] = append(PostByAuth[post.Author], &post)
	fmt.Printf("->%T--%v\n", PostByAuth[post.Author], PostByAuth[post.Author])
}

func main() {
	post1 := Post{1, "hello", "yuyang"}
	post2 := Post{2, "world", "mingze"}
	post3 := Post{3, "I'am", "wangjing"}
	post4 := Post{4, "kitty", "yuyang"}
//	fmt.Println(post1,"\n", post2,"\n",post3,"\n",post4)
	PostById = make(map[int]*Post)
	PostByAuth = make(map[string][]*Post)
	Mystorage(post1)
	Mystorage(post2)
	Mystorage(post3)
	Mystorage(post4)
	for _, post := range PostByAuth["yuyang"]{
		fmt.Println(post)
	}
	for _, post := range PostByAuth["mingze"]{
		fmt.Println(post)
	}
}

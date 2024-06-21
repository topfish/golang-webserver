package mystorage

import (
	"fmt"
//	"webapp/
)

func MystorageDaemon() {
	PostById = make(map[int]*Post)
	PostByAuth = make(map[string][]*Post)
	post1 := Post{1, "hello", "yy"}
	post2 := Post{2, "world", "mz"}
	post3 := Post{3, "I'am", "wj"}
	post4 := Post{4, "kitty", "yy"}

	fmt.Println("I am Main")
	Mystorage(post1)
	Mystorage(post2)
	Mystorage(post3)
	Mystorage(post4)
	fmt.Println(PostById[1])
	fmt.Println(PostById[2])
	for _, post := range PostByAuth{
		size := len(post)
		for{
			fmt.Printf("%v\t",post[size-1])
			size--
			if size == 0{
				break
			}
		}
		fmt.Println()
	}
}

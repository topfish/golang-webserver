package mytemplate

import(
	"fmt"
	"os"
)

func CreateTemp() {
	newtemp, err := os.Create("/home/tem.html")
	if err != nil{
		fmt.Println(err)
	}
	defer newtemp.Close()
	str := `<body>
	{{.}}
	</body>`
	_, err = newtemp.Write([]byte(str))
	if err != nil{
		fmt.Println(err)
	}
}

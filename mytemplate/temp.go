package mytemplate

import(
	"os"
)

func CreateTemp() {
	newtemp, err := os.Create("./mytemplate/tem.html")
	myerr(err)
	defer newtemp.Close()
	
	str := `<body>
	{{.}}
	</body>`

	_, err = newtemp.Write([]byte(str))
	myerr(err)
}

func DeleteTemp(filename string) {
	err := os.Remove(filename)
	myerr(err)
}

func myerr(err error) {
	if err != nil {
		panic(err)
	}
}

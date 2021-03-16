package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	//sample.CreateFolder()



	//sample.TestObj()


	//sample.TestTime()

	//sample.TestReg()


	//testMap()


	//sample.CreateFile()

	//testSlice()
	testStr()
}


func testStr(){
	//fmt.Println(strings.Replace("$sdfg sdf as $sdfg","$sdfg","@@@@",-1))
	a:="namespace:<$project>"
	fmt.Println(strings.Replace(a,"<$project>","123",1))
	fmt.Println(a)

	b:="asdfasdf| 12332"
	fmt.Println(strings.Split(b,"|"))

	fmt.Println(strings.HasPrefix("my string", "my  s"))
	fmt.Println(strings.HasPrefix("my string", "xxx"))

	fmt.Println(strconv.Quote("Hello, 世界"))
	fmt.Println("\"Hello, \"世界\"")
	fmt.Println(strconv.Unquote("\"Hello, \"世界\""))
	fmt.Println(strconv.Quote("Hello, 世界"))
	fmt.Println(strconv.QuoteToASCII("Hello, 世界"))
}

func testMap(){
	var a map[string]string
	a=make(map[string]string)
	a["a"]="a"
	fmt.Println(a)
}

func testSlice(){
	var s []string
	s=append(s,"test")
	fmt.Println(s)


	fmt.Println(strings.Join([]string{"sdfsdf","sdfsdf","apple"}, ","))
}
package sample

import (
	"fmt"
	"os"
)

func CreateFolder(){
	path:="C:\\tmp\\tt\\123\\456/7889"
	if _,err:=os.Stat(path);os.IsNotExist(err){
		err := os.Mkdir(path, 0755)
		fmt.Println("err:",err)
	}
}


func CreateFile(){
	_, err := os.Create("C:\\tmp\\tt\\a.txt")
	if err!=nil{
		fmt.Println("err:",err)
	}
}
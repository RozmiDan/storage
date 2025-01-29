package main

import (
	"fmt"
	"log"

	"github.com/RozmiDan/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("test.txt", []byte("hello"))
	if err != nil{
		log.Fatal("error")
	}
	fmt.Println(file)

	fl, err := st.GetById(file.Id)
	if err != nil{
		panic("err")
	}
	fmt.Println(fl.Name)
}
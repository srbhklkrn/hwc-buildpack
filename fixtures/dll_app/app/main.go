package main

import (
	"fmt"
	"syscall"
	"time"
)

func main() {
	_, err := syscall.LoadLibrary("MySql.Data.dll")
	if err != nil {
		panic(err)
	}
	fmt.Println("dllHandle loaded")
	for {
		fmt.Println("Sleeping for good")
		time.Sleep(time.Hour)
	}
}

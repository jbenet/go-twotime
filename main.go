package main

import (
	"fmt"
	"runtime"
	"time"
)

func boom() int {
	i := 0
	for {
		i++

		if i%10E6 == 0 {
			if i%10E8 == 0 {
				fmt.Println(".")
			} else {
				fmt.Printf(".")
			}
		}

		t1 := time.Now()
		t2 := time.Now()
		t3 := time.Now()
		t4 := time.Now()
		t5 := time.Now()
		t6 := time.Now()
		if t1 == t2 || t2 == t3 || t3 == t4 || t4 == t5 || t5 == t6 {
			return i
		}
	}
}

func main() {
	i := boom()
	fmt.Println("BOOM!")
	fmt.Println(i, "iterations on", runtime.GOOS, runtime.GOARCH)
}

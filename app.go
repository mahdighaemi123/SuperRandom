package main

import (
	"./super_random"
)

func main() {
	sample()
}

func sample(){
	// min -> 0 , max ->100
	println(super_random.SuperRandom(0,100))
}

func test(){
	// min -> 0 , max ->1000
	for i:=0;i<10000 ;i++  {
		println(super_random.SuperRandom(0,1000))
	}
}
package main

import "os"

func main() {

}

func GenerateData() (chan int, error) {
	f, err := os.Open("./file.csv")
	if err != nil {
		return nil, err
	}

	f.



}

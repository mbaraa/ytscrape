package main

import "fmt"

func main() {
	res, err := search("volkswagen das auto")
	if err != nil {
		return
	}
	for _, r := range res {
		// _ = r
		fmt.Printf("%+v\n", r)
	}
}

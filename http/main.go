package main

import "net/http"
import "fmt"
import "os"

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	resp.Body.Read(bs)

	fmt.Println(string(bs))
}

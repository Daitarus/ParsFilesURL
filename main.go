package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var url string
	var file_address string
	// change url_base
	var url_base string
	fmt.Println("Enter your url:")
	fmt.Fscan(os.Stdin, &url_base)
	// change n page
	var n int
	fmt.Println("Enter number of pages:")
	fmt.Fscan(os.Stdin, &n)
	for i := 1; i <= n; i++ {
		//form url and file address
		if i < 10 {
			url = url_base + "0"
			file_address = "tmp/0"
		} else {
			url = url_base
			file_address = "tmp/"
		}
		url += strconv.Itoa(i)
		url += ".jpg"
		file_address += strconv.Itoa(i)
		file_address += ".jpg"

		//get response from url
		response, e := http.Get(url)
		if e != nil {
			log.Fatal(e)
		}
		defer response.Body.Close()

		//create file
		file, err := os.Create(file_address)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// Use io.Copy to just dump the response body to the file. This supports huge files
		_, err = io.Copy(file, response.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("All files uploaded!")
}

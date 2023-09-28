package gowebserver

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func MPrint(){
	fmt.Println("Welcome to Get web video - lco")
}

// create separate Methods

func PerformGetRequest(){
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil{
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code: ", response.Status)
	fmt.Println("Content length: ", response.ContentLength)

	content, _ := io.ReadAll(response.Body)

	var responseString strings.Builder
	
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))
}
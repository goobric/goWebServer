package gowebserver

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func MPrint(){
	fmt.Println("Welcome to Get web video - lco")
	//PerformGetRequest()
	//PerformPostJsonRequest()
	PerformpostFormRequest()

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
	// fmt.Println(responseString.String())
}

func PerformPostJsonRequest(){
	const myurl = "http://localhost:8000/post"

	// fake json payload

	requestBody := strings.NewReader(`{
		"coursename":"Let's go with golang",
		"price": 0,
		"platform":"learnCodeOnline.in"
	}`)
	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil{
		panic(err)
	}
	defer response.Body.Close()
	
	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformpostFormRequest(){
	const myurl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname", "Hitesh")
	data.Add("lastname", "Choudhary")
	data.Add("email", "hitesh@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
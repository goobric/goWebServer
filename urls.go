package gowebserver

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456abc"

func urls(){
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	//parsing the url - extract information
	// url library
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port()) //Port() is a method
	fmt.Println(result.RawQuery)	
}
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var linkRoute string = "https://dummyjson.com/posts"

type PostsRequest struct {
	Posts []PostTypes `json:"posts"`
	Total int32       `json:"total"`
	Skip  byte        `json:"skip"`
	Limit byte        `json:"limit"`
}

type PostTypes struct {
	Id        uint32   `json:"id"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	UserId    uint32   `json:"userId"`
	Tags      []string `json:"tags"`
	Reactions uint16   `json:"reactions"`
}

func main() {
	res, err := http.Get(linkRoute)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	data := PostsRequest{}
	json.Unmarshal(resData, &data)

	fmt.Print(data)

}

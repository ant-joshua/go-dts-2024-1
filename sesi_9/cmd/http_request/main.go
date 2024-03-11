package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type JsonPlaceholderService struct {
	client http.Client
}

func NewJsonPlaceholderService() *JsonPlaceholderService {
	return &JsonPlaceholderService{
		client: http.Client{},
	}
}

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (j JsonPlaceholderService) getPosts() []Post {
	res, err := j.client.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	var posts []Post

	err = json.Unmarshal(body, &posts)

	if err != nil {
		panic(err)
	}

	return posts
}

func getPost(id string) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + id)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	var post Post

	err = json.Unmarshal(body, &post)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", post)
	fmt.Println(post.Id)
	fmt.Println(post.Title)
	fmt.Println(post.Body)
	fmt.Println(post.UserId)
}

func postRequest() {
	data := map[string]interface{}{
		"title":  "hello",
		"body":   "world",
		"userId": 1,
	}

	requestJson, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	var post Post

	err = json.Unmarshal(body, &post)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", post)
	fmt.Println(post.Id)
	fmt.Println(post.Title)
	fmt.Println(post.Body)
	fmt.Println(post.UserId)

	// log.Println(string(body))
}

func updatePost() {
	data := map[string]interface{}{
		"title":  "hello",
		"body":   "world updated",
		"userId": 1,
	}

	requestJson, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("PUT", "https://jsonplaceholder.typicode.com/posts/1", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	var post Post

	err = json.Unmarshal(body, &post)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", post)
	fmt.Println(post.Id)
	fmt.Println(post.Title)
	fmt.Println(post.Body)
	fmt.Println(post.UserId)

	// log.Println(string(body))
}

func main() {
	// posts := getPosts()

	// for _, post := range posts {
	// 	fmt.Println("post id = ", post.Id)
	// 	fmt.Println(post.Title)
	// 	fmt.Println(post.Body)
	// 	fmt.Println(post.UserId)
	// }

	updatePost()

}

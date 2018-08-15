package main

import (
	"fmt"
	"selfLearning/retiever/mock"
	"selfLearning/retiever/real"
	"time"
)

const url = "http://www.weibo.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "grandFather",
			"course": "golang",
		})
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "anthoer weibo.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{
		"http://www.weibo.com"}
	r = &retriever
	inspect(r)
	r = &real.Retreiver{
		"Mozilla/5.0",
		time.Minute,
	}
	inspect(r)

	//type assertion
	if mockRetreiver, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetreiver.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))

}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("> %T %v\n", r, r)

	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("> Contents :", v.Contents)
	case *real.Retreiver:
		fmt.Println("> User Agent:", v.UserAgent)
	}
	fmt.Println()
}

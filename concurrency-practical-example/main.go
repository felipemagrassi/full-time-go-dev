package main

import (
	"fmt"
	"sync"
	"time"
)

type UserProfile struct {
	ID       int
	Comments []string
	Likes    int
	Friends  []int
}

type Response struct {
	data any
	err  error
}

func handleGetUserProfile(id int) (*UserProfile, error) {
	responseChan := make(chan Response, 3)
	wg := sync.WaitGroup{}

	go getComments(id, responseChan, &wg)
	go getLikes(id, responseChan, &wg)
	go getFriends(id, responseChan, &wg)

	wg.Add(3)
	wg.Wait()
	close(responseChan)

	userProfile := &UserProfile{ID: id}

	for response := range responseChan {
		if response.err != nil {
			return nil, response.err
		}

		switch msg := response.data.(type) {
		case []string:
			userProfile.Comments = msg
		case int:
			userProfile.Likes = msg
		case []int:
			userProfile.Friends = msg
		}

	}

	return userProfile, nil
}

func getComments(id int, responseChan chan Response, wg *sync.WaitGroup) {
	time.Sleep(200 * time.Millisecond)
	comments := []string{"Hello", "World", "!"}
	responseChan <- Response{data: comments, err: nil}
	wg.Done()
}

func getLikes(id int, responseChan chan Response, wg *sync.WaitGroup) {
	time.Sleep(300 * time.Millisecond)
	likes := 100
	responseChan <- Response{data: likes, err: nil}
	wg.Done()
}

func getFriends(id int, responseChan chan Response, wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	friends := []int{11, 32, 55, 49}
	responseChan <- Response{data: friends, err: nil}
	wg.Done()
}

func main() {
	start := time.Now()
	userProfile, err := handleGetUserProfile(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(userProfile)
	fmt.Println(time.Since(start))
}

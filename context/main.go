package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	accesses := map[string]string{"library1": "user1", "library2": "user2"}
	ctx := context.WithValue(context.Background(), "usernames", accesses)
	start := time.Now()

	res, err := fetchUser(ctx)

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println("user: ", res)
	fmt.Println("took: ", time.Since(start))
}

func fetchUser(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 1100*time.Millisecond)
	defer cancel()

	val := ctx.Value("usernames").(map[string]string)
	fmt.Println("username: ", val)

	type result struct {
		userId  string
		err     error
		library string
	}

	resultch := make(chan result)

	go func() {
		library := "library1"
		id, err := firstThirdPartyLibraryCall(val[library])
		resultch <- result{userId: id, library: library, err: err}
	}()

	go func() {
		library := "library2"
		id, err := secondThirdPartyLibraryCall(val[library])
		resultch <- result{userId: id, library: library, err: err}
	}()

	select {
	case <-ctx.Done():
		return "", fmt.Errorf("timeout")
	case result := <-resultch:
		return fmt.Sprintf("%s from %s", result.userId, result.library), result.err
	}
}

func firstThirdPartyLibraryCall(username string) (string, error) {
	randomTimeDuration := time.Duration(1000+rand.Intn(1000)) * time.Millisecond
	fmt.Println("firstThirdPartyLibraryCall: ", randomTimeDuration)
	fmt.Println("username: ", username)
	time.Sleep(randomTimeDuration)

	return "1", nil
}

func secondThirdPartyLibraryCall(username string) (string, error) {
	randomTimeDuration := time.Duration(1000+rand.Intn(1000)) * time.Millisecond
	fmt.Println("secondThirdPartyLibraryCall: ", randomTimeDuration)
	fmt.Println("username: ", username)
	time.Sleep(randomTimeDuration)

	return "2", nil
}

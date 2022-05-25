package limit

import (
	"fmt"
	"testing"
	"time"
)

func TestBucket(t *testing.T) {
	generator := NewTokenGenerator(1)
	generator.StartGenerate()
	bucket := NewBucket(10, generator)
	bucket.StartLimit()
	times := 0
	for {
		if bucket.FetchToken() {
			fmt.Println("get token successful")
			times++
		} else {
			fmt.Println("get token fail")
		}

		if times < 10 {
			time.Sleep(2 * time.Second)
		} else {
			time.Sleep(50 * time.Millisecond)
		}
	}
}

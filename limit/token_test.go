package limit

import (
	"fmt"
	"testing"
)

func TestGenerator(t *testing.T) {
	generator := NewTokenGenerator(1)
	generator.StartGenerate()
	for {
		generator.GenerateToken()
		fmt.Println("get token")
	}
}

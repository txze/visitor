package hash

import (
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	var token string

	{
		token = GenerateToken()
		if len(token) != 32 {
			t.Error(token, len(token))
			return
		}
		fmt.Println("case 1:", token)
	}
}

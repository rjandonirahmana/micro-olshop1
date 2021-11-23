package auth

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateValidateToken(t *testing.T) {
	// Testcases list
	testCases := []struct {
		testName string
		id       uint
		exp      int64
	}{
		{
			testName: "success",
			id:       5,
			exp:      time.Now().Add(time.Hour * 8).Unix(),
		},
		{
			testName: "success",
			id:       10,
			exp:      time.Now().Add(time.Hour * 8).Unix(),
		}, {
			testName: "success",
			id:       30,
			exp:      time.Now().Add(time.Hour * 8).Unix(),
		},
	}

	for _, testCase := range testCases {
		newToken, err := NewService("coba", "cobalagi").GenerateToken(testCase.id)
		fmt.Println(newToken)
		assert.Nil(t, err)

		newToken1, err := NewService("coba", "cobalagi").GenerateTokenSeller(testCase.id)
		fmt.Println(newToken1)
		assert.Nil(t, err)

		token, exp, er := NewService("coba", "cobalagi").ValidateToken(newToken)
		assert.Nil(t, er)

		token1, exp1, err := NewService("coba", "cobalagi").ValidateTokenSeller(newToken1)
		assert.NoError(t, err)

		fmt.Println(*exp)

		assert.Equal(t, testCase.id, token)
		assert.Equal(t, testCase.exp, *exp)
		assert.Equal(t, testCase.exp, *exp1)
		assert.Equal(t, testCase.id, token1)

	}
}

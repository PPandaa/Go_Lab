package testLab

import (
	"GoLab/pkg"
	"testing"
)

func TestIsEmptyString(t *testing.T) {

	testCases := []string{"     ", "sss", ""}
	testAnswers := []bool{true, false, true}

	for testIndex, testCase := range testCases {
		if pkg.IsEmptyString(testCase) != testAnswers[testIndex] {
			t.Error("wrong result")
		}
	}

}

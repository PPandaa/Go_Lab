package testLab

import (
	"GoLab/tool"
	"testing"
)

func TestIsEmptyString(t *testing.T) {

	testCases := []string{"     ", "sss", ""}
	testAnswers := []bool{true, false, true}

	for testIndex, testCase := range testCases {
		if tool.IsEmptyString(testCase) != testAnswers[testIndex] {
			t.Error("wrong result")
		}
	}

}

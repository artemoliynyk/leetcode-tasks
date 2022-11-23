package taskN

import (
	"testing"
)

type TestValue struct {
	input    int
	expected int
}

func TestInorderTraversal(t *testing.T) {
	var testData []TestValue
	testData = append(testData, TestValue{1, 2})
	testData = append(testData, TestValue{2, 3})
	testData = append(testData, TestValue{3, 4})
	testData = append(testData, TestValue{4, 0}) // wrong one

	t.Logf("Total test data sets: %d, timeout %dms", len(testData), 0)

	for _, testSet := range testData {
		actual := TaskFunction(testSet.input)

		if testSet.expected != actual {
			t.Errorf("Got value '%v', expected '%v'", actual, testSet.expected)
		}
	}
}

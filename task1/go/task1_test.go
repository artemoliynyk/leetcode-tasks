package task1

import (
	"testing"
	"time"
)

type TestValue struct {
	nums     []int
	target   int
	expected []int
}

func TestTwoSum(t *testing.T) {
	var testData []TestValue
	testData = append(testData, TestValue{[]int{2, 7, 11, 15}, 9, []int{0, 1}})
	testData = append(testData, TestValue{[]int{3, 2, 4}, 6, []int{1, 2}})
	testData = append(testData, TestValue{[]int{3, 3}, 6, []int{0, 1}})

	t.Logf("Total test data sets: %d, timeout %dms", len(testData), 0)

	for _, testSet := range testData {
		start := time.Now()
		actual := twoSum(testSet.nums, testSet.target)
		execTime := time.Since(start)

		t.Logf("Execution time %dms\nNums '%v', target: '%v'", execTime.Microseconds(), testSet.nums, testSet.target)

		if testSet.expected[0] != actual[0] || testSet.expected[1] != actual[1] {
			t.Errorf("Got value '%v', expected '%v'", actual, testSet.expected)
		}
	}
}

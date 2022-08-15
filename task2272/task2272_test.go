package task2272

import (
	"testing"
	"time"
)

type TestValue struct {
	input    string
	expected int
}

const TIMEOUT_CORRECTION = 0.5

func TestLargestVariance(t *testing.T) {
	// local machine performance is greated that leetcode site, dedicated for 138 test, allowing = total timeout / total tasks
	TIMEOUT := (5990 * TIMEOUT_CORRECTION * time.Millisecond) / 138

	var testData []TestValue
	testData = append(testData, TestValue{"abc", 0})
	testData = append(testData, TestValue{"aabbc", 1})
	testData = append(testData, TestValue{"aabbc", 1})
	testData = append(testData, TestValue{"hdivbekdpvjfczschxwylgmfntolkvapgwszvilwdurfcvmmyjxqlwdawcjhgsjbxwtwitkqlhsmefcfhzfjinssxmrwtcsshetadjvactftrffpehzbeaqinqrfbevhxdyroasrlbdnonchcvfiwznpyimqtiqiwyetrikecrqdusytmvuzqnmdlosxficmqctidfldapympuianbsqfrbooukppwfopcujikagcdkznkdhfjqzdqlevcjwrucwtbrksddvhisvmytztqfuknvhhgalueojjzeiiqdspqkmuuzamywnkjjbtqgzkkhjihfrzmpqqtrrruveexvsoychipadoifkezvkapodkobqlgctzaqcoqwtewfblsbdyyicnbtnqupytomttxtyohvsvlznabzbpzpnwdblaecoeemdzcfwraoujqcwbkkhknpdjd", 18})

	t.Logf("Total test data sets: %d, timeout %dms", len(testData), TIMEOUT.Microseconds())

	for _, testSet := range testData {
		start := time.Now()
		actual := largestVariance(testSet.input)
		execTime := time.Since(start)

		t.Logf("Execution time %dms\n\tInput '%s'", execTime.Microseconds(), testSet.input)

		if execTime > TIMEOUT {
			t.Errorf("\tTime limit of %dms exceeded (actual %dms), for input '%s'", TIMEOUT.Microseconds(), execTime.Microseconds(), testSet.input)
		} else if testSet.expected != actual {
			t.Errorf("\tGot value '%d', expected '%d'", actual, testSet.expected)
		}
	}
}

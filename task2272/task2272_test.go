package task2272

import (
	"testing"
	"time"
)

type TestValue struct {
	input    string
	expected int
}

func TestLargestVariance(t *testing.T) {
	var testData []TestValue
	testData = append(testData, TestValue{"abc", 0})
	testData = append(testData, TestValue{"aabbc", 1})
	testData = append(testData, TestValue{"aabbc", 1})
	testData = append(testData, TestValue{"hdivbekdpvjfczschxwylgmfntolkvapgwszvilwdurfcvmmyjxqlwdawcjhgsjbxwtwitkqlhsmefcfhzfjinssxmrwtcsshetadjvactftrffpehzbeaqinqrfbevhxdyroasrlbdnonchcvfiwznpyimqtiqiwyetrikecrqdusytmvuzqnmdlosxficmqctidfldapympuianbsqfrbooukppwfopcujikagcdkznkdhfjqzdqlevcjwrucwtbrksddvhisvmytztqfuknvhhgalueojjzeiiqdspqkmuuzamywnkjjbtqgzkkhjihfrzmpqqtrrruveexvsoychipadoifkezvkapodkobqlgctzaqcoqwtewfblsbdyyicnbtnqupytomttxtyohvsvlznabzbpzpnwdblaecoeemdzcfwraoujqcwbkkhknpdjd", 18})

	t.Logf("Total test data sets: %d", len(testData))
	
	for _, testSet := range testData {
		start := time.Now()
		actual := largestVariance(testSet.input)
		t.Logf("Execution time %s\n\tInput '%s'", testSet.input, time.Since(start))

		if testSet.expected != actual {
			t.Errorf("\tGot value '%d', expected '%d'", actual, testSet.expected)
		}
	}
}

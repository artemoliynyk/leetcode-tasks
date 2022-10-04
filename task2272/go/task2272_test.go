package task2272

import (
	"testing"
)

type TestValue struct {
	input    string
	expected int
}

func TestLargestVariance(t *testing.T) {

	var testData []TestValue
	// testData = append(testData, TestValue{"aaaaa", 0})
	// testData = append(testData, TestValue{"bbbbbbb", 0})
	// testData = append(testData, TestValue{"ab", 0})
	// testData = append(testData, TestValue{"abc", 0})
	testData = append(testData, TestValue{"aabbb", 2})
	// testData = append(testData, TestValue{"aabbc", 1})
	// testData = append(testData, TestValue{"ababab", 1})
	// testData = append(testData, TestValue{"ababc", 1})
	// testData = append(testData, TestValue{"aababbb", 3})
	// testData = append(testData, TestValue{"dsyemilsuwhciclqwprizywgkwkbgcqhtcwfvlw", 5})
	// testData = append(testData, TestValue{"hdivbekdpvjfczschxwylgmfntolkvapgwszvilwdurfcvmmyjxqlwdawcjhgsjbxwtwitkqlhsmefcfhzfjinssxmrwtcsshetadjvactftrffpehzbeaqinqrfbevhxdyroasrlbdnonchcvfiwznpyimqtiqiwyetrikecrqdusytmvuzqnmdlosxficmqctidfldapympuianbsqfrbooukppwfopcujikagcdkznkdhfjqzdqlevcjwrucwtbrksddvhisvmytztqfuknvhhgalueojjzeiiqdspqkmuuzamywnkjjbtqgzkkhjihfrzmpqqtrrruveexvsoychipadoifkezvkapodkobqlgctzaqcoqwtewfblsbdyyicnbtnqupytomttxtyohvsvlznabzbpzpnwdblaecoeemdzcfwraoujqcwbkkhknpdjd", 18})

	for _, testSet := range testData {
		actual := largestVariance(testSet.input)

		if testSet.expected != actual {
			t.Errorf("Got value '%d', expected '%d'", actual, testSet.expected)
		}
	}
}

package main

import (
	"testing"
)

type In struct {
	num1 []int
	num2 []int
}
type Op struct {
	op string
}

var StEd []In
var ExpectedArr []Op

func TestHandle(t *testing.T) {
	StEd = []In{
		In{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		In{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		In{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1, 2, 3, 4, 5}},
		In{[]int{1, 2, 3, 4, 6, 9, 9, 9, 9, 9}, []int{2, 1, 2, 9, 4, 5, 8}},
		In{[]int{2, 7, 6, 5, 4, 3}, []int{9, 2, 9, 1, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}},
		In{[]int{9, 8, 7, 6, 5, 5, 7, 5}, []int{5, 7, 5, 5, 6, 7, 8, 9}},
		In{[]int{7, 1, 8, 7, 5, 7, 2, 1}, []int{7, 7, 7, 7, 7, 7, 7, 7}},
	}
	ExpectedArr = []Op{
		Op{"987666666"},
		Op{"66666"},
		Op{"123511110"},
		Op{"1243249211"},
		Op{"100000000268472"},
		Op{"197531150"},
		Op{"149653498"},
	}
	for i, v := range StEd {
		result := Add(v.num1, v.num2)
		//fmt.Println(result, "\t", ExpectedArr[i].op)
		if result != ExpectedArr[i].op {
			t.Errorf("Expected result did not match the Output")
		}
	}
}

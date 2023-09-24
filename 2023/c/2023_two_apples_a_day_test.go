package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCaseTest2(t *testing.T) {
	assert.Equal(t, 7, solution(2, []int{7, 7, 7}))
}

func TestCaseTest4(t *testing.T) {
	assert.Equal(t, -1, solution(3, []int{1, 9, 1, 1, 4}))
}

func TestCaseTest7(t *testing.T) {
	assert.Equal(t, 102, solution(3, []int{100, 2, 10, 4, 94}))
}

/*
3
6 3 1 2 5
6 5 3 2 1
*/
func TestCaseVal1(t *testing.T) {
	assert.Equal(t, 4, solution(3, []int{6, 3, 1, 2, 5}))
}

/*
4
1 9 1 1 4 9 9
9 9 9 4 1 1 1
*/
func TestCaseVal5(t *testing.T) {
	assert.Equal(t, 6, solution(4, []int{1, 9, 1, 1, 4, 9, 9}))
}

/*
4
4
1 9 10 1 4 6 9
10 9 9 6 4 1 1
*/
func TestCaseVal6(t *testing.T) {
	assert.Equal(t, 6, solution(4, []int{1, 9, 1, 1, 4, 9, 9}))
}

/*
9
5 7 4 9 4 10 4 8 3 7 8 3 7 3 10 5 6
10 10 9 8 8 7 7 7 6 5 5 4 4 4 3 3 3
7 7 4 4 3
13 -> -1
*/
func TestCaseVal10(t *testing.T) {
	assert.Equal(t, -1, solution(9, []int{5, 7, 4, 9, 4, 10, 4, 8, 3, 7, 8, 3, 7, 3, 10, 5, 6}))
}

/*
11
2 2 4 20 22 22 22 26 27 28 30 34 36 37 38 42 42 44 60 62 62
22 -> 64-22 = 42
*/
func TestCaseVal8(t *testing.T) {
	assert.Equal(t, 42, solution(11, []int{4, 28, 36, 62, 22, 2, 30, 37, 20, 42, 42, 22, 38, 27, 22, 44, 62, 2, 34, 60, 26}))
}

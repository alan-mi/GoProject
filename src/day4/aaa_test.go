package main

import "testing"

func add(x, y int) int {
	x += y
	return x
}

func TestAdd(t *testing.T){
	var tests = []struct{
		x int
		y int
		expect int
	}{
		{1,2,3,},
		{9,5,14,},
		{3,6,9,},
	}
	for _,tt:=range tests{
		actual := add(tt.x,tt.y)
		println(actual)
		if actual != tt.expect{
			t.Errorf("add(%d,%d):expect %d,actual %d",tt.x,tt.y,tt.expect,actual)
		}
	}
}

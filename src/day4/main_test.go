package day4

import "testing"

func add(x, y int) int {
	return x + y
}
func TestAdd(t *testing.T) {
	if add(1, 2) != 3 {
		//直接退出测试任务
		t.FailNow()
		//继续执行当前函数
		t.Fail()
		//跳过当前函数
		t.SkipNow()
	}
}

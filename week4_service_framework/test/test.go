package main

import (
	"fmt"
	"reflect"
)

func main(){
	funcNames := []string{"Test1", "Test2", "Test3"}
	mt := &MyTest{
		t1: "test1",
		t2: "test2",
		t3: "test3",
	}
	value := reflect.ValueOf(mt)
	for _, funcName := range funcNames {
		f := value.MethodByName(funcName)
		res := f.Call([]reflect.Value{})
		fmt.Println(res)
	}

}

type MyTest struct {
	t1 string
	t2 string
	t3 string
}

func (t *MyTest) Test1() string {
	return t.t1
}

func (t *MyTest) Test2() string {
	return t.t2
}

func (t *MyTest) Test3() string {
	return t.t3
}
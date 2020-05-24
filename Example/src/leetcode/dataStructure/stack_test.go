package dataStructure

import "testing"

var obj MyStack

func init(){
    obj = Constructor()
}

func TestMyStack_Push(t *testing.T) {
    input := []int{1,2,3,4,5}
    for _, a := range input {
        obj.Push(a)
    }
    t.Log("test push [1,2,3,4,5]")
}

func TestMyStack_Pop(t *testing.T) {

}

func TestMyStack_Top(t *testing.T) {

}

func TestMyStack_Empty(t *testing.T) {

}
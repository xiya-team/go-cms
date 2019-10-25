package test

import (
	"go-cms/pkg/arr"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	var data = []string{"1", "3", "2", "5", "2"}
	result := arr.RemoveRepeatedElement(data)
	//t.Log("result1:", result)

	var reverse = []string{"1", "3", "2", "5", "2"}
	result=arr.Reverse(reverse)
	t.Log("result2:", result)
}
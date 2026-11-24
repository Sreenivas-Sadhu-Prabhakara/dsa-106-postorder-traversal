package main

import (
	"reflect"
	"testing"
)

func TestPostorder(t *testing.T) {
	root := tn(4, tn(2, tn(1, nil, nil), tn(3, nil, nil)), tn(6, tn(5, nil, nil), tn(7, nil, nil)))
	want := []int{1, 3, 2, 5, 7, 6, 4}
	if got := Postorder(root); !reflect.DeepEqual(got, want) {
		t.Errorf("Postorder = %v, want %v", got, want)
	}
	if got := Postorder(nil); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Postorder(nil) = %v", got)
	}
	chain := tn(1, nil, tn(2, tn(3, nil, nil), nil))
	if got := Postorder(chain); !reflect.DeepEqual(got, []int{3, 2, 1}) {
		t.Errorf("Postorder(chain) = %v, want [3 2 1]", got)
	}
}

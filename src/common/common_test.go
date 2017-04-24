package common

import "testing"
import "log"

func TestCamelToSnake(t *testing.T) {
	type pair struct {
		in, out string
	}

	tests := []pair{
		{"HelloWorld", "hello_world"},
		{"TestID", "test_id"},
		{"ID", "id"},
	}

	for _, v := range tests {
		out := CamelToSnake(v.in)
		if out != v.out {
			t.Error(out, " != ", v.out)
		}
	}

	log.Println("TestCamelToSnake succesfull..")
}

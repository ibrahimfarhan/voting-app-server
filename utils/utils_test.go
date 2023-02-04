package utils

import (
	"fmt"
	"testing"
)

func TestStringInSlice(t *testing.T) {
	slice := []string{"s1", "s2", "s3"}

	s1_in_slice := StringInSlice("s1", slice)
	s_in_slice := StringInSlice("s", slice)

	if !s1_in_slice || s_in_slice {
		t.Error("StringInSlice misbehaves")
	}
}

func TestToString(t *testing.T) {
	var i, b interface{}

	i = "string"
	b = nil

	iToString := ToString(i)
	bToString := ToString(b)

	if iToString != "string" || bToString != "" {
		fmt.Println(iToString, bToString)
		t.Fatal("ToString failes")
	}
}

func TestToBool(t *testing.T) {
	var i, b interface{}

	i = true
	b = nil

	iToBool, _ := ToBool(i)
	bToBool, err := ToBool(b)

	if iToBool != true || bToBool != false || err == nil {
		t.Fatal("ToBool failes", iToBool, bToBool)
	}
}

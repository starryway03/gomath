package gomath

import "testing"

func TestSum(t *testing.T){
	want:=28
	s:=Sum(7,1,2,3,4,5,6)
	t.Logf("Sum : %d",s)
	if s!=want{
		t.Errorf("Sum was incorrect, got : %d ,want :%d",s,want)
	}
}

func TestSubstract(t *testing.T){
	want:=-10
	s:=Substract(-7,1,2)
	t.Logf("Sum : %d",s)
	if s!=want{
		t.Errorf("Sum was incorrect, got : %d ,want :%d",s,want)
	}
}

func TestMultiply(t *testing.T){
	want:=42
	s:=Multiply(-7,1,-2,3)
	t.Logf("Sum : %d",s)
	if s!=want{
		t.Errorf("Sum was incorrect, got : %d ,want :%d",s,want)
	}
}
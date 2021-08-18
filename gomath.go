package gomath

func Substract(base int,i ...int) int {
	r:=base
	for _,b:=range i{
		r=r-b
	}
	return r
}
func Sum(i ...int)int  {
	r:=0
	for _,b:=range i{
		r=r+b
	}
	return r
}
func Multiply(i ...int)int  {
	r:=1
	for _,b:=range i{
		r=r*b
	}
	return r
}

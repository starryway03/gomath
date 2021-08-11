package gomath

func Substract(i ...int) int {
	r:=i[0]
	for _,b:=range i{
		r=r-b
	}
	return r
}
func Sum(i ...int)int  {
	r:=i[0]
	for _,b:=range i{
		r=r-b
	}
	return r
}
func Multiply(i ...int)int  {
	r:=i[0]
	for _,b:=range i{
		r=r*b
	}
	return r
}

package main

import "fmt"

func multiplication (){
  nums := []interface{}{1,2,3,4,5,6,7,8,9}
  fmt.Println("Multiplication : ",mulTi(nums))
}
func mulTi (m []interface{})interface{}{
	mul := 1
	for _,n := range m{
		mul = mul * n.(int)
	}
	return mul
}
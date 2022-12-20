package main

import "fmt" 

func numberAddition (){
	nums := []interface{}{1,2,3,4,5,6,7,8,9}
	fmt.Println("sum : ",sum(nums))

}

func sum (s []interface{})interface{}{
	sum := 0
	for _,n := range s{
		sum = sum + n.(int)
		
	}
	return sum
}
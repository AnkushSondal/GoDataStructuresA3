package main

import (
	"fmt"
)

func countOccur(arr []int,num int) int{
	count :=0
	for i := 0; i < len(arr); i++ {
		if arr[i]==num{
			count++
		}
	}
	return count
}

func sortArray(arr []int) []int{
	temp := 0
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i]>=arr[j]{
				temp=arr[i]
				arr[i]=arr[j]
				arr[j]=temp
			}
		}	
	}
	return arr
}

func main(){
	number :=0
	arr := []int{1,2,3,4,1,2,5,2,6,9}
	fmt.Print("Enter number to find the count of : ")
	fmt.Scan(&number)
	count :=countOccur(arr,number)
	fmt.Println("count of",number,"in",arr,":",count)
	arr = sortArray(arr)
	fmt.Println("Sorted array : ",arr)

}


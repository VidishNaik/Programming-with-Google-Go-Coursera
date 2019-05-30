package main

import (
"fmt"
"bufio"
"os"
"strings"
"strconv"
)

func main() {
	nums := []int{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the numbers separated by a space. Maximum 10 allowed")
	s, _ := reader.ReadString('\n')
	temp := strings.Split(strings.TrimSpace(s), " ")
	if len(temp) > 10 {
		fmt.Println("Length of input greater than 10.")
		os.Exit(1)
	}
	for i := range temp{
		num, err := strconv.Atoi(strings.TrimSpace(temp[i]))
		if err != nil {
			break	
		}	
		nums = append(nums, num)
	}
	if len(nums) == len(temp){
		BubbleSort(nums...)
		fmt.Println("Sorted:",nums)
	} else {
		fmt.Println("Error in input")
	}
}

func BubbleSort(nums ...int){
	nums_length := len(nums)
	for i := 0; i < nums_length - 1; i++ {
		for j := 0; j < nums_length - i - 1; j++ {
			if (nums[j] > nums[j + 1]){
				Swap(j, nums...)		
			}
		}
	}
}

func Swap(i int, nums ...int){
	temp := nums[i]
	nums[i] = nums[i + 1]
	nums[i + 1] = temp
}

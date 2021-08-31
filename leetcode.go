package main

import "sort"
func Factorial(a int)int{
	if(a>0){
		return Factorial(a-1)
	}else{
		return 1
	}
}
func numRescueBoats(people []int, limit int) (ans int) {
	sort.Ints(people)
	light, heavy := 0, len(people)-1
	for light <= heavy {
		if people[light]+people[heavy] > limit {
			heavy--
		} else {
			light++
			heavy--
		}
		ans++
	}
	return
}

func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)
	for _, booking := range bookings {
		nums[booking[0]-1] += booking[2]
		if booking[1] < n {
			nums[booking[1]] -= booking[2]
		}
	}
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	return nums
}


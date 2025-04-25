package fourSum18

import (
	"fmt"
	"os"
	"sort"
)

func Run() {
	var array = fourSum([]int{2, 2, 2, 2, 2}, 8)

	solv := fmt.Sprint(array)
	res := "[[2 2 2 2]]"

	if solv != res {
		fmt.Println("Wrong")
		fmt.Printf("Expected: %s\n", res)
		fmt.Printf("Got: %s\n", solv)
		os.Exit(1)
	}

	fmt.Println("Nice")
	os.Exit(0)
}

func fourSum(nums []int, target int) [][]int {
	var res = make([][]int, 0)
	var hash = make(map[string]int)

	for i := range nums {
		for j := range nums {
			if i != j {
				for k := range nums {
					if k != j && k != i {
						for l := range nums {
							if l != k && l != j && l != i {
								if (nums[i] + nums[j] + nums[k] + nums[l]) == target {
									_, ok := hash[getKey(nums[i], nums[j], nums[k], nums[l])]
									if !ok {
										res = append(res, []int{nums[i], nums[j], nums[k], nums[l]})
										hash[getKey(nums[i], nums[j], nums[k], nums[l])] = 1
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return res
}

func getKey(i int, j int, k int, l int) string {
	values := []int{i, j, k, l}
	sort.Ints(values)
	return fmt.Sprint(values)
}

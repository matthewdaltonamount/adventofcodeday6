package main

import (
	"fmt"
	"strconv"
	"strings"
)

const fish_array_string = "1,3,1,5,5,1,1,1,5,1,1,1,3,1,1,4,3,1,1,2,2,4,2,1,3,3,2,4,4,4,1,3,1,1,4,3,1,5,5,1,1,3,4,2,1,5,3,4,5,5,2,5,5,1,5,5,2,1,5,1,1,2,1,1,1,4,4,1,3,3,1,5,4,4,3,4,3,3,1,1,3,4,1,5,5,2,5,2,2,4,1,2,5,2,1,2,5,4,1,1,1,1,1,4,1,1,3,1,5,2,5,1,3,1,5,3,3,2,2,1,5,1,1,1,2,1,1,2,1,1,2,1,5,3,5,2,5,2,2,2,1,1,1,5,5,2,2,1,1,3,4,1,1,3,1,3,5,1,4,1,4,1,3,1,4,1,1,1,1,2,1,4,5,4,5,5,2,1,3,1,4,2,5,1,1,3,5,2,1,2,2,5,1,2,2,4,5,2,1,1,1,1,2,2,3,1,5,5,5,3,2,4,2,4,1,5,3,1,4,4,2,4,2,2,4,4,4,4,1,3,4,3,2,1,3,5,3,1,5,5,4,1,5,1,2,4,2,5,4,1,3,3,1,4,1,3,3,3,1,3,1,1,1,1,4,1,2,3,1,3,3,5,2,3,1,1,1,5,5,4,1,2,3,1,3,1,1,4,1,3,2,2,1,1,1,3,4,3,1,3"
const test_fish_array = "1,3,1,5,5,1,1,1"
const days = 256

func main() {
	//convert fish array to map
	fishMap := convert_fish_input(fish_array_string)

	fmt.Println(fishMap)

	totalFish := 0
	stashMap := make(map[int]int)

	for _, v := range fishMap {
		_, exists := stashMap[v]
		if  exists {
			totalFish += stashMap[v]
		} else {
			iterativeFish := recursive_count(v,days,1)
			totalFish += iterativeFish
			stashMap[v] = iterativeFish
		}
	}
	//output
	fmt.Println(totalFish)
}

func convert_fish_input(fish string) map[int]int {
	m := make(map[int]int)
	split := strings.Split(fish,",")
	for i, k := range split {
		value, err := strconv.Atoi(k)
		if err != nil {
			return nil
		}
		m[i] = value
	}
	return m
}

func recursive_count (fish int, days int, count int) int {
	if days == 0 {
		return count
	}
	switch fish {
	case 0:
		return recursive_count(6,days-1, count) + recursive_count(8, days-1, count)
	default:
		if days - fish >= 0 {
			return recursive_count(fish-fish, days-fish, count)
		} else {
			return count
		}
	}
}




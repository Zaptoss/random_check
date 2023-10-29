package main

import (
	"fmt"
	"math/rand"
)

func random_gen2() []uint32 {
	result := []uint32{}
	for i := 0; i < 625; i++ {
		result = append(result, uint32(rand.Int31()))
	}
	return result
}

func monobit(number_array []uint32) bool {
	var ones uint32
	for i := 0; i < len(number_array); i++ {
		for j := uint32(1 << 31); j >= 1; j >>= 1 {
			if number_array[i]&j != 0 {
				ones += 1
			}
		}
	}
	if ones < 9654 || ones > 10346 {
		return false
	}
	return true
}

func long_series_check(number_array []uint32) bool {
	seria := 0
	digit := 0
	max_seria := 0
	for i := 0; i < len(number_array); i++ {
		for j := uint32(1); j < 1<<31; j <<= 1 {
			if number_array[i]&j != 0 {
				if digit != 1 {
					digit = 1
					if seria > max_seria {
						max_seria = seria
					}
					seria = 1
				} else {
					seria += 1
				}
			} else {
				if digit != 0 {
					digit = 0
					if seria > max_seria {
						max_seria = seria
					}
					seria = 1
				} else {
					seria += 1
				}
			}
		}
	}

	if max_seria <= 36 {
		return true
	}
	return false
}

func poker_test(number_array []uint32) bool {
	var frequency [16]int
	for i := 0; i < len(number_array); i++ {
		for j := 0; j < 32; j += 4 {
			frequency[(number_array[i]>>j)&0b1111] += 1
		}
	}

	res := float64(0)
	count := 0
	for i := 0; i < 16; i++ {
		count += frequency[i]
		res += float64(frequency[i] * frequency[i])
	}
	x := (16/float64(count))*res - float64(count)

	if x < 57.4 && x > 1.03 {
		return true
	}
	return false
}

func series_check(number_array []uint32) bool {
	zero_series := make(map[int]int)
	one_series := make(map[int]int)
	seria := 0
	digit := 0
	for i := 0; i < len(number_array); i++ {
		for j := uint32(1); j < 1<<31; j <<= 1 {
			if number_array[i]&j != 0 {
				if digit != 1 {
					digit = 1
					if seria > 6 {
						seria = 6
					}
					if seria != 0 {
						zero_series[seria] += 1
					}
					seria = 1
				} else {
					seria += 1
				}
			} else {
				if digit != 0 {
					digit = 0
					if seria > 6 {
						seria = 6
					}
					one_series[seria] += 1
					seria = 1
				} else {
					seria += 1
				}
			}
		}
	}

	compare_table := map[int][2]int{1: {2267, 2733}, 2: {1079, 1421}, 3: {502, 748}, 4: {223, 402}, 5: {90, 223}, 6: {90, 223}}
	for i := 1; i <= 6; i++ {
		if zero_series[i] < compare_table[i][0] || zero_series[i] > compare_table[i][1] || one_series[i] < compare_table[i][0] || one_series[i] > compare_table[i][1] {
			return false
		}
	}
	return true
}

func main() {
	var random_number []uint32
	for {
		random_number = random_gen2()
		if monobit(random_number) && series_check(random_number) && poker_test(random_number) && long_series_check(random_number) {
			break
		}
	}
	fmt.Println("The sequence is quite random")
}

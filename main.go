package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

// randomGen2 - функція що генерує випадкову послідовність 20000 біт
// на основі функції rand.Uint32, яка генерує 32 бітні випадкові значення;
// повертає масив uint32
func randomGen2() []uint32 {
	var bitArray []uint32
	for i := 0; i < 625; i++ {
		bitArray = append(bitArray, rand.Uint32())
	}
	return bitArray
}

// getHex виводить послідовність біт у 16-річному форматі
func getHex(bitArray []uint32) string {
	var hexString string
	for i := 0; i < len(bitArray); i++ {
		hexString += strconv.FormatUint(uint64(bitArray[i]), 16)
	}
	return hexString
}

// Тест на відповідність кількості 0/1 стандарту FIPS-140
func monobit(bitArray []uint32) bool {
	var onesCount uint32
	for i := 0; i < len(bitArray); i++ {
		for j := 31; j >= 0; j-- {
			onesCount += (bitArray[i] >> j) & 1
		}
	}

	if onesCount < 9654 || onesCount > 10346 {
		return false
	}
	return true
}

// Тест Покера
func pokerTest(bitArray []uint32) bool {
	// entryFrequency - масив який зберігає кількість входжень кожного з блоків; блоки розміром 4, тому можливих різних блоків - 16
	var entryFrequency [16]float32
	// blockCount - кількість блоків розміром 4
	blockCount := float64(len(bitArray) * 8)
	// quadsSum - сума квадратів кількості входження кожного з блоків
	var quadsSum float64
	for i := 0; i < len(bitArray); i++ {
		for j := 0; j < 32; j += 4 {
			entryFrequency[(bitArray[i]>>j)&0b1111] += 1
		}
	}

	for i := 0; i < 16; i++ {
		quadsSum += float64(entryFrequency[i] * entryFrequency[i])
	}
	X := (16/blockCount)*quadsSum - blockCount

	if X < 1.03 || X > 57.4 {
		return false
	}
	return true
}

// seriesCheck2 приймає на вхід послідовність 20000 біт(big endian);
// true якщо максимальна серія та кількість входжень серій відповідають стандарту FIPS-140
func seriesCheck2(bitArray []uint32) bool {
	// series масив який зберігає серії різної довжини 0 і 1 відповідно
	var series [2][6]int

	// maxSeria - найбільша серія 1 або 0
	maxSeria := 0

	// seria - поточна серія 1 або 0
	seria := 0

	// digit - вказує на серію одиниць(=1) чи нулів(=0)
	digit := (bitArray[0] >> 31) & 1

	// Проходимось в циклі по всім бітам
	for i := 0; i < len(bitArray); i++ {
		for j := 31; j >= 0; j-- {
			// Визначаємо поточний біт
			bit := (bitArray[i] >> j) & 1

			// Якщо поточний біт = поточній серії -> seria += 1
			if bit == digit {
				seria += 1
			} else { // В іншому випадку перевіряємо чи є поточна серія максимальною, якщо так - то замінюємо максимальну на поточну
				if seria > maxSeria {
					maxSeria = seria
				}
				// Якщо серія більша за 6 - прирівнюємо серію до 6 (в таблиці останнє значення 6+)
				if seria > 6 {
					seria = 6
				}
				// Додаємо 1 до комірки що зберігає кількістьвходжень поточної серії, змінюємо поточну серію на протилежну, поточна серія = 1
				series[digit][seria-1] += 1
				digit = bit
				seria = 1
			}
		}
	}

	// Для врахування останньої серії
	if seria > maxSeria {
		maxSeria = seria
	}
	if seria > 6 {
		seria = 6
	}
	series[digit][seria-1] += 1

	// Перевірка чи відповідає стандарту максимальна довжина серії
	if maxSeria > 36 {
		return false
	}

	// Діапазон значень входжень кожної з серій відповідно до стандарту
	compareTable := [6][2]int{{2267, 2733}, {1079, 1421}, {502, 748}, {223, 402}, {90, 223}, {90, 223}}

	// Перевірка належності кількості входжень проміжку зазначеному в стандарті
	for i := 0; i < 6; i++ {
		for j := 0; j < 2; j++ {
			if series[j][i] < compareTable[i][0] || series[j][i] > compareTable[i][1] {
				return false
			}
		}
	}

	return true
}

func main() {
	var randomNumber []uint32
	
	for {
		randomNumber = randomGen2()
		if monobit(randomNumber) && seriesCheck2(randomNumber) && pokerTest(randomNumber) {break}
	}
	
	fmt.Println(getHex(randomNumber))
	fmt.Println("The sequence is quite random")
}

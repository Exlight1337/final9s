package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(1_000_000) + 1
	}
	return data
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxVal := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > maxVal {
			maxVal = data[i]
		}
	}
	return maxVal
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	if chunkSize == 0 {
		return maximum(data)
	}

	maxes := make([]int, CHUNKS)

	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data) // последняя часть забирает остаток
		}

		go func(idx, from, to int) {
			defer wg.Done()
			maxes[idx] = maximum(data[from:to])
		}(i, start, end)
	}

	wg.Wait()
	return maximum(maxes)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)
	fmt.Println()

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	fmt.Println()
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}

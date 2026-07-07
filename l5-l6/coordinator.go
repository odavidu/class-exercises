package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

type KVPair struct {
	key   string
	value string
}

func main() {
	cities := []string{"Tokyo", "Delhi", "Shanghai", "Sao_Paulo", "Mexico_City", "Cairo", "Mumbai", "Beijing", "Dhaka", "Osaka", "New_York", "Karachi", "Buenos_Aires", "Istanbul", "Kolkata", "Lagos", "Moscow", "London", "Paris", "Los_Angeles"}

	//Reads in input file for each city and calls Map function
	ch := make(chan KVPair)
	var wg sync.WaitGroup
	for _, city := range cities {
		wg.Add(1)
		go func(city string) {
			defer wg.Done()
			path := "data/" + city + ".txt"
			input, err := os.ReadFile(path)
			if err != nil {
				panic(err)
			}
			map_out := Map(path, string(input))
			for _, item := range map_out {
				ch <- item
			}
		}(city)
	}

	// Goroutine waits to close the channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	// range over channel repeatedly reads from channel until it is closed
	kv_pairs := make(map[string][]string)
	for item := range ch {
		kv_pairs[item.key] = append(kv_pairs[item.key], item.value)
	}

	// Reduce tasks
	ch = make(chan KVPair)
	for year, pairs := range kv_pairs {
		wg.Add(1)
		go func(year string, pairs []string) {
			defer wg.Done()
			reduceOut := Reduce(year, pairs)
			ch <- KVPair{year, strconv.FormatFloat(reduceOut, 'f', -1, 64)}
		}(year, pairs)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for item := range ch {
		fmt.Println("Highest Temp in", item.key, "was", item.value)
	}
}

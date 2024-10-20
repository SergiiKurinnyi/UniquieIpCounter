package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	TotalIPNumber = 1 << 32
)

type BitSet struct {
	data []uint32
}

func NewBitSet() *BitSet {
	return &BitSet{data: make([]uint32, TotalIPNumber/32)}
}

func (bs *BitSet) Set(index int) {
	bs.data[index/32] |= 1 << (index % 32)
}

func (bs *BitSet) Get(index int) bool {
	return (bs.data[index/32] & (1 << (index % 32))) != 0
}

func ipToLong(ip string) uint32 {
	parts := strings.Split(ip, ".")

	p0, _ := strconv.Atoi(parts[0])
	p1, _ := strconv.Atoi(parts[1])
	p2, _ := strconv.Atoi(parts[2])
	p3, _ := strconv.Atoi(parts[3])

	return (uint32(p0) << 24) | (uint32(p1) << 16) | (uint32(p2) << 8) | uint32(p3)
}

func main() {

	startTime := time.Now()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	startMemory := m.Alloc

	filename := "ip_addresses"

	bitSet := NewBitSet()

	uniqueCount := 0

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ip := scanner.Text()
		ipLong := ipToLong(strings.TrimSpace(ip))

		if !bitSet.Get(int(ipLong)) {
			bitSet.Set(int(ipLong))
			uniqueCount++
		}
	}

	endTime := time.Now()
	runtime.ReadMemStats(&m)
	endMemory := m.Alloc - startMemory

	fmt.Printf("Processing time: %v sec\n", (endTime.Sub(startTime).Seconds()))
	fmt.Printf("Memory allocated:  %v MB\n", (endMemory / 1024 / 1024))
	fmt.Printf("Number of unique IPs: %d\n", uniqueCount)
}

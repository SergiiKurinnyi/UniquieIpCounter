package main

import (
	"bufio"
	"io"
	"strings"
	"testing"
)

func Test_ipToLong_whenParsingIpString_thenShouldReturnLong(t *testing.T) {
	tests := []struct {
		ip       string
		expected uint32
	}{
		{"10.10.10.1", 168430081},
		{"127.0.0.1", 2130706433},
		{"192.168.0.1", 3232235521},
		{"255.255.255.255", 4294967295},
	}

	for _, tt := range tests {
		result := ipToLong(tt.ip)
		if result != tt.expected {
			t.Errorf("Test Failed: ipToLong(%s): expected %d, got %d", tt.ip, tt.expected, result)
		}
	}
}

func Test_ipToLong_whenIpDigitsWithoutDotsMatch_thenParsedValuesNotEqual(t *testing.T) {
	ip1 := "10.10.11.1"
	ip2 := "10.10.1.11"

	long1 := ipToLong(ip1)
	long2 := ipToLong(ip2)

	if long1 == long2 {
		t.Errorf("Test Failed: Expected different values for IPs %s and %s, but got the same: %d", ip1, ip2, long1)
	}
}

func Test_ipToLong_whenSameIp_thenCounterShouldSkipDuplicates(t *testing.T) {
	IPs := `192.168.0.1
			192.168.0.2
			192.168.0.1
			127.0.0.1
			192.168.0.1
			192.168.0.2`

	expectedUniqueCount := 3

	reader := io.NopCloser(strings.NewReader(IPs))
	bitSet := NewBitSet()
	uniqueCount := countUniqueIPsHelper(reader, bitSet)

	if uniqueCount != expectedUniqueCount {
		t.Errorf("Test Failed: Expected %d unique IPs, got %d", expectedUniqueCount, uniqueCount)
	}
}

func Test_BitSet_Set_whenBitSet_thenGetBitShouldReturnTrue(t *testing.T) {
	bitSet := NewBitSet()

	// Set a bit and check if it is marked correctly.
	ipIndex := 168430081
	bitSet.Set(ipIndex)

	if !bitSet.Get(ipIndex) {
		t.Errorf("Test Failed: Expected bit %d to be set", ipIndex)
	}
}

func Test_BitSet_Get_whenBitNotSet_thenGetBitShouldReturnFalse(t *testing.T) {
	bitSet := NewBitSet()

	ipIndex := 3232235521
	bitSet.Set(ipIndex)

	if bitSet.Get(ipIndex + 1) {
		t.Errorf("Test Failed: Expected bit %d NOT to be set", ipIndex+1)
	}
}

func countUniqueIPsHelper(reader io.Reader, bitSet *BitSet) int {
	scanner := bufio.NewScanner(reader)
	uniqueCount := 0

	for scanner.Scan() {
		ip := scanner.Text()
		ipLong := ipToLong(strings.TrimSpace(ip))

		if !bitSet.Get(int(ipLong)) {
			bitSet.Set(int(ipLong))
			uniqueCount++
		}
	}

	return uniqueCount
}

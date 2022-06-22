package utils

import (
	"math/rand"
	"strconv"
	"time"
)

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n uint8) string {
	runes := make([]rune, n)
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	for i := range runes {
		runes[i] = letters[r.Intn(len(runes))]
	}
	return string(runes)
}

func RandNum(n uint8) (uint64, error) {
	runes := make([]rune, n)
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	for i := range runes {
		runes[i] = letters[r.Intn(len(runes))]
	}
	parseUint, err := strconv.ParseUint(string(runes), 10, 64)
	if err != nil {
		return 0, err
	}
	return parseUint, nil
}

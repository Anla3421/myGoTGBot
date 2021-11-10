package random

import (
	"math/rand"
	"time"
)

var (
	randSeed *rand.Rand
	randStr  = []rune("1234567890" + "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func init() {
	randSeed = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// 取得隨機rand
func GetSeed() *rand.Rand {
	return randSeed
}

// 變換randseed
func ChangeRandSeed() {
	randSeed = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// 取得隨機字串
func RandStr(length int) (str string) {
	if length <= 0 {
		length = 32
	}
	slice := make([]rune, length)
	for i := range slice {
		slice[i] = randStr[randSeed.Intn(len(randStr))]
	}
	str = string(slice)
	return
}

// 從客製化的string中 做 隨機字串
func RandStrFromSample(length int, sample []rune) (str string) {
	if length <= 0 {
		length = 32
	}

	if len(sample) <= 0 {
		sample = randStr
	}
	slice := make([]rune, length)
	for i := range slice {
		slice[i] = sample[randSeed.Intn(len(sample))]
	}
	str = string(slice)
	return
}

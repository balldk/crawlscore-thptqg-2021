package src

import (
	"fmt"
	"os"
	"strconv"
)

func toFloat(s string) float32 {
	val, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return -1
	}
	return float32(val)
}

func AppendToFile(filename string, data interface{}) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(fmt.Sprint(data) + "\n"); err != nil {
		panic(err)
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func formatScore(score float32) string {
	if score == -1 {
		return " "
	} else {
		return fmt.Sprint(score)
	}
}

func SBDFormat(areaCode int, i int) string {
	return fmt.Sprintf("%02d%06d", areaCode, i)
}

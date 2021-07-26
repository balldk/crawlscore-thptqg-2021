package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/schollz/progressbar/v3"
)

func upperBoundAreaCode(areaCode int, left int, right int, chBound chan BoundChannel) {
	ch := make(chan *StudentChannel, 1)
	chFail := make(chan string, 1)
	chFinish := make(chan bool, 1)

	defer close(ch)
	defer close(chFail)
	defer close(chFinish)

	if right > left {
		mid := (right + left) / 2
		sbd := SBDFormat(areaCode, mid)

		FetchScore(sbd, ch, chFail, chFinish)

		std := <-ch

		if std.data != nil {
			upperBoundAreaCode(areaCode, mid+1, right, chBound)
		} else {
			upperBoundAreaCode(areaCode, left, mid-1, chBound)
		}
	}

	sbd := SBDFormat(areaCode, right)
	FetchScore(sbd, ch, chFail, chFinish)

	std := <-ch
	if std.data != nil {
		chBound <- BoundChannel{
			areaCode: areaCode,
			bound:    right,
		}
	} else {
		chBound <- BoundChannel{
			areaCode: areaCode,
			bound:    right - 1,
		}
	}
}

func SearchAreaRange() map[int]int {

	fmt.Println("\nSearching SBD range of each area...")
	fmt.Println()
	bar := progressbar.Default(64)

	res := make(map[int]int)
	chBound := make(chan BoundChannel)
	defer close(chBound)

	for areaCode := 1; areaCode <= 64; areaCode++ {
		go upperBoundAreaCode(areaCode, 1, 110000, chBound)
	}

	for i := 1; i <= 64; {
		bound := <-chBound
		bar.Add(1)
		i++
		res[bound.areaCode] = bound.bound
	}

	return res
}

func LoadAreaRangeFile() (data map[int]int) {
	filename := "area_range.json"
	file, err := os.Open(filename)

	if err != nil {
		return nil
	}

	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &data)

	file.Close()
	return data
}

func SaveAreaRangeFile(data map[int]int) {
	filename := "area_range.json"
	file, _ := os.Create(filename)

	byteValue, _ := json.MarshalIndent(data, "", "\t")
	file.Write(byteValue)
	file.Close()
}

func NumberOfStudent(data map[int]int) int {
	sum := 0
	for _, areaCodeRange := range data {
		if areaCodeRange != -1 {
			sum += areaCodeRange
		}
	}
	return sum
}

package src

import (
	"fmt"
	"os"
	"time"

	"github.com/schollz/progressbar/v3"
)

func runInPatch(areaCode int, start int, end int) {
	ch := make(chan *StudentChannel)
	chFail := make(chan string)
	chFinish := make(chan bool)

	for i := start; i <= end; i++ {
		sbd := SBDFormat(areaCode, i)
		go FetchScore(sbd, ch, chFail, chFinish)
	}
	for i := start; i < end; {
		select {
		case each := <-ch:
			if each.data != nil {
				filename := fmt.Sprintf("%s/%02d.csv", os.Getenv("OUTPUT_FOLDER"), areaCode)
				AppendToFile(filename, each.data)
				AppendToFile(fmt.Sprintf("%s/total.csv", os.Getenv("OUTPUT_FOLDER")), each.data)
			}
		case <-chFinish:
			i++
		}
	}
}

func Run(areaCodeRangeMap map[int]int, patchSize int, patchDelayDuration time.Duration) {

	fmt.Println("\nFetching scores...")
	fmt.Println()
	bar := progressbar.Default(int64(NumberOfStudent(areaCodeRangeMap)))

	fileHeader := "SBD\tTên\tNgày Sinh\tGiới tính\tToán\tVăn\tLý\tHoá\tSinh\tKHTN\tLịch Sử\tĐịa Lý\tGDCD\tKHXH\tNgoại Ngữ"

	if _, err := os.Stat("data"); os.IsNotExist(err) {
		os.Mkdir(os.Getenv("OUTPUT_FOLDER"), 0755)
	}
	// Add header to total file
	totalFileName := os.Getenv("TOTAL_FILENAME")
	AppendToFile(fmt.Sprintf("%s/%s", os.Getenv("OUTPUT_FOLDER"), totalFileName), fileHeader)

	// Loop through all areas
	for areaCode, areaCodeRange := range areaCodeRangeMap {
		filename := fmt.Sprintf("%s/%02d.csv", os.Getenv("OUTPUT_FOLDER"), areaCode)

		file, _ := os.Create(filename)
		file.Close()
		AppendToFile(filename, fileHeader)

		for start := 1; start <= areaCodeRange; start += patchSize {
			end := min(start+patchSize, areaCodeRange)

			runInPatch(areaCode, start, end)
			bar.Add(end - start)

			time.Sleep(patchDelayDuration)
		}
	}
	bar.Finish()
}

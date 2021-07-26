package main

import (
	"crawlscore/src"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// ASCII ART
	fmt.Println(`

   ████████╗██╗  ██╗██████╗ ████████╗ ██████╗  ██████╗               ██████╗  ██████╗ ██████╗  ██╗
   ╚══██╔══╝██║  ██║██╔══██╗╚══██╔══╝██╔═══██╗██╔════╝               ╚════██╗██╔═████╗╚════██╗███║
      ██║   ███████║██████╔╝   ██║   ██║   ██║██║  ███╗    █████╗     █████╔╝██║██╔██║ █████╔╝╚██║
      ██║   ██╔══██║██╔═══╝    ██║   ██║▄▄ ██║██║   ██║    ╚════╝    ██╔═══╝ ████╔╝██║██╔═══╝  ██║
      ██║   ██║  ██║██║        ██║   ╚██████╔╝╚██████╔╝              ███████╗╚██████╔╝███████╗ ██║
      ╚═╝   ╚═╝  ╚═╝╚═╝        ╚═╝    ╚══▀▀═╝  ╚═════╝               ╚══════╝ ╚═════╝ ╚══════╝ ╚═╝
	`)

	// Load .env data
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	patchSize, _ := strconv.Atoi(os.Getenv("PATCH_SIZE"))
	patchDelay, _ := strconv.ParseFloat(os.Getenv("PATCH_DELAY"), 64)
	patchDelayDuration := time.Duration(float64(time.Second) * patchDelay)
	fmt.Print("Configurations:\n\n")
	fmt.Printf(" - Data Source: https://thanhnien.vn/giao-duc/tuyen-sinh/2021/tra-cuu-diem-thi-thpt-quoc-gia.html\n")
	fmt.Printf(" - Patch Size: %d\n", patchSize)
	fmt.Printf(" - Patch Delay: %.2f seconds\n", patchDelay)
	fmt.Printf(" - Output Folder: %s/\n", os.Getenv("OUTPUT_FOLDER"))

	// Load area range
	areaCodeRangeMap := src.LoadAreaRangeFile()
	if areaCodeRangeMap == nil {
		areaCodeRangeMap = src.SearchAreaRange()
		src.SaveAreaRangeFile(areaCodeRangeMap)
	}

	// Fetch scores
	src.Run(areaCodeRangeMap, patchSize, patchDelayDuration)

	fmt.Print("\nFinished!\n\n")
}

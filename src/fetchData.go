package src

import (
	"fmt"
	"io"
	"net/http"
)

func FetchRaw(sbd string) string {

	url := fmt.Sprintf("https://thanhnien.vn/ajax/diemthi.aspx?kythi=THPT&nam=2021&city=DDT&text=%s&top=no", sbd)
	res, err := http.Get(url)
	if res.StatusCode == 403 {
		errMsg := "https://thanhnien.vn/ might blocked you. Try to change IP address, reduce patch size/delay and run again."
		errMsg += "\nLast SBD: " + string(sbd)
		panic(errMsg)
	}
	if err != nil {
		fmt.Println(sbd, err)
		return ""
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(sbd, err)
		return ""
	}

	return string(body)
}

func FetchScore(sbd string, ch chan *StudentChannel, chFail chan string, chFinish chan bool) {

	htmlBody := FetchRaw(sbd)
	var std *Student

	if len(htmlBody) == 0 {
		std = nil
		chFail <- sbd
	} else {
		std = ParseStudent(&htmlBody)
	}

	ch <- &StudentChannel{
		id:   sbd,
		data: std,
	}
	chFinish <- true
}

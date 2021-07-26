package src

import (
	"regexp"
)

func ParseScore(data *[]string) Score {

	score := Score{
		math:           toFloat((*data)[6]),
		literature:     toFloat((*data)[7]),
		physics:        toFloat((*data)[8]),
		chemistry:      toFloat((*data)[9]),
		biology:        toFloat((*data)[10]),
		naturalScience: toFloat((*data)[11]),
		history:        toFloat((*data)[12]),
		geography:      toFloat((*data)[13]),
		civic:          toFloat((*data)[14]),
		socialScience:  toFloat((*data)[15]),
		language:       toFloat((*data)[16]),
	}
	return score
}

func ParseStudent(htmlBody *string) *Student {

	r := regexp.MustCompile(`(?m)\s*<td[^>]*">([^<]*)<\/td>`)
	matches := r.FindAllStringSubmatch(*htmlBody, -1)
	if len(matches) == 0 {
		return nil
	}

	var data []string
	for _, match := range matches {
		data = append(data, string(match[1]))
	}

	// stt, _ := strconv.Atoi(data[0])
	std := Student{
		// stt:    stt,
		name:   data[2],
		sbd:    data[3],
		dob:    data[4],
		gender: data[5],
		score:  ParseScore(&data),
	}

	return &std
}

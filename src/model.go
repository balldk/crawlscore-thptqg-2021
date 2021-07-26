package src

import "fmt"

type Score struct {
	math           float32
	literature     float32
	physics        float32
	chemistry      float32
	biology        float32
	naturalScience float32
	history        float32
	geography      float32
	civic          float32
	socialScience  float32
	language       float32
}

type Student struct {
	// stt    int
	name   string
	sbd    string
	dob    string
	gender string
	score  Score
}

type StudentChannel struct {
	id   string
	data *Student
}

type BoundChannel struct {
	areaCode int
	bound    int
}

func (score *Score) String() string {
	res := ""
	res += formatScore(score.math) + "\t"
	res += formatScore(score.literature) + "\t"
	res += formatScore(score.physics) + "\t"
	res += formatScore(score.chemistry) + "\t"
	res += formatScore(score.biology) + "\t"
	res += formatScore(score.naturalScience) + "\t"
	res += formatScore(score.history) + "\t"
	res += formatScore(score.geography) + "\t"
	res += formatScore(score.civic) + "\t"
	res += formatScore(score.socialScience) + "\t"
	res += formatScore(score.language)
	return res
	// return fmt.Sprintf("%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t",
	// 	score.math, score.literature, score.physics, score.chemistry, score.biology,
	// 	score.naturalScience, score.history, score.geography, score.civic, score.socialScience, score.language)
}

func (std *Student) String() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s", std.sbd, std.name, std.dob, std.gender, std.score.String())
}

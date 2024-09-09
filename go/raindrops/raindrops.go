package raindrops

import "strconv"

const Pling string = "Pling"
const Plang string = "Plang"
const Plong string = "Plong"


func Convert(number int) string {
	var answer string
	if number%3 == 0 {
		answer += Pling
	} 
	if number%5 == 0 {
		answer += Plang
	}
	if number%7 == 0 {
		answer += Plong
	}

	if len(answer) == 0 {
		answer = strconv.Itoa(number)
	}
	return answer
}
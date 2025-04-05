package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HumanReadableTime(seconds int) string {
	minutes, hours := 0, 0
	var mer string
	for seconds >= 60 {
		if seconds >= 60 {
			minutes = seconds / 60
			seconds = seconds % 60
		}
		if minutes >= 60 {
			hours = minutes / 60
			minutes %= 60
		}
	}
	mer = strconv.Itoa(hours) + ":" + strconv.Itoa(minutes) + ":" + strconv.Itoa(seconds)
	result := strings.Split(mer, ":")
	for i, r := range result {
		okey, _ := strconv.Atoi(r)
		if okey < 10 {
			result[i] = "0" + result[i]
		}
	}
	return strings.Join(result, ":")
}

func main() {
	fmt.Println(HumanReadableTime(0))
	fmt.Println(HumanReadableTime(59))
	fmt.Println(HumanReadableTime(60))
	fmt.Println(HumanReadableTime(90))
	fmt.Println(HumanReadableTime(3599))
	fmt.Println(HumanReadableTime(3600))
	fmt.Println(HumanReadableTime(45296))
	fmt.Println(HumanReadableTime(86399))
	fmt.Println(HumanReadableTime(86400))
	fmt.Println(HumanReadableTime(359999))
}

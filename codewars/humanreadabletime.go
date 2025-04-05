package codewars

import (
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
		zeroFinder, _ := strconv.Atoi(r)
		if zeroFinder < 10 {
			result[i] = "0" + result[i]
		}
	}
	return strings.Join(result, ":")
}

/*
func HumanReadableTime(s int) string {
    m,s := s / 60, s % 60
    h,m := m / 60, m % 60
    return fmt.Sprintf("%02d:%02d:%02d",h,m,s)
}
*/

/*
func HumanReadableTime(seconds int) string {
  var sec, min, hour string

  hours := seconds / 3600
  seconds = seconds % 3600
  minutes := seconds / 60
  seconds = seconds % 60

  sec = strconv.Itoa(seconds)
  hour = strconv.Itoa(hours)
  min = strconv.Itoa(minutes)

  if(seconds<10){ sec ="0" + strconv.Itoa(seconds) }
  if(hours<10){ hour ="0" + strconv.Itoa(hours) }
  if(minutes<10){ min ="0" + strconv.Itoa(minutes) }

  return hour + ":" + min + ":" + sec
}
*/

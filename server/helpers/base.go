package helpers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

// Date ...
type Date struct {
	day   int
	month time.Month
	year  int
}

// CheckErr ...
func CheckErr(errMsg string, err error) {
	if err != nil {
		beego.Warning(errMsg, err)
	}
}

// BytesToString ...
func BytesToString(data []byte) string {
	return string(data[:])
}

// ArrayToString ...
func ArrayToString(arr []string, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(arr), " ", delim, -1), "[]")
}

// Multiply ...
func Multiply(x, y, z float64) float64 {
	return (x - (y * z))
}

func stringToInt(value string) int {
	result, err := strconv.Atoi(value)
	CheckErr("err", err)
	return result
}

func stringToMonth(value string) time.Month {
	tm, err := time.Parse("01", value)
	CheckErr("err", err)
	return tm.Month()
}

func date2unix(d Date, loc *time.Location) int64 {
	return time.Date(d.year, d.month, d.day, 0, 0, 0, 0, loc).Unix()
}

func primitive(d Date, loc *time.Location) int64 {
	base := Date{2000, time.January, 3}
	seconds := date2unix(d, loc) - date2unix(base, loc)
	weeks := seconds / (7 * 24 * 60 * 60)
	secondIntoWeek := seconds % (7 * 24 * 60 * 60)
	workdays := secondIntoWeek / (24 * 60 * 60)
	if workdays > 5 {
		workdays = 5
	}
	return 5*weeks + workdays
}

func dayCountExcludingWeekends(from, to Date, loc *time.Location) int {
	return int(primitive(to, loc) - primitive(from, loc))
}

// GetTotalDay ...
func GetTotalDay(from string, to string) int {
	loc, err := time.LoadLocation("Asia/Jakarta")
	CheckErr("err", err)

	f := strings.Split(from, "-")
	t := strings.Split(to, "-")

	dateFrom := Date{stringToInt(f[0]), stringToMonth(f[1]), stringToInt(f[2])}
	dateTo := Date{stringToInt(t[0]), stringToMonth(t[1]), stringToInt(t[2])}
	result := dayCountExcludingWeekends(dateFrom, dateTo, loc)
	fmt.Println(from)
	fmt.Println("arr", f[0])
	return result + 1
}

// GetDay ...
func GetDay(date string) int {
	dateDay, _ := time.Parse("2006-01-02", date)
	m := dateDay.Month()
	var i = int(m)
	return i
}

var addedTimeNow = 0

// Now replace helper.Now()
func Now() time.Time {
	return time.Now().AddDate(0, 0, addedTimeNow)
}

// NowLoc ...
func NowLoc(timeLoc string) (
	dateTime time.Time,
	err error,
) {
	// "Asia/Bangkok"
	loc, err := time.LoadLocation(timeLoc)
	if err != nil {
		beego.Warning("failed load location", err)
	}
	dateTime = Now().In(loc)
	return
}

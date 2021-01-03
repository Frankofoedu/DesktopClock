package main

import (
	"strconv"
	"time"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  502,
		Height: 436,
		Title:  "goclock",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(GetDateAndTime)
	app.Run()
}


type MyDate struct {
	Date string
	Time string
}
func GetDateAndTime() MyDate{
	week := [7]string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}

	var cd = time.Now()
	var calcdate = zeroPadding(cd.Year(), 4) +
		"-" +
		zeroPadding(int(cd.Month())+1, 2) +
		"-" +
		zeroPadding(cd.Day(), 2) +
		" " +
		week[cd.Weekday()]
	var calctime = zeroPadding(cd.Hour(), 2) +
		":" +
		zeroPadding(cd.Minute(), 2) +
		":" +
		zeroPadding(cd.Second(), 2)

	mydate := MyDate{
		Date : calcdate,
		Time : calctime,		 
	}

	return mydate
}

func zeroPadding(num int, digit int) string {
	var zero = ""

	for i := 0; i < digit; i++ {
		zero += "0"
	}

	stringNum := strconv.Itoa(num)
	stringNum = zero + stringNum
	length := len(stringNum)
	lastNDigit := stringNum[length-digit : length]

	return lastNDigit
}

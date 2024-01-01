package main

import (
	"fmt"
	"time"
	"github.com/common-nighthawk/go-figure"
)


func main() {


	now :=  time.Now()


	if now.Day() == 31 && now.Month() == 12 {
		fg := figure.NewFigure("Happy Old Year", "", true)
		fg.Print()		
	}

	if now.Day() == 1 && now.Month() == 1 {
		phrase :=  fmt.Sprintf("Happy %d", now.Year())
		fg := figure.NewFigure(phrase, "", true)
		fg.Print()		
	}
}
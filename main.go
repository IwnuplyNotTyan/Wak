package main

import (
	"fmt"
	"os"
	"time"

	"github.com/muesli/termenv"
)

type Calendar struct {
	output *termenv.Output
	year   int
	month  time.Month
}

func NewCalendar() *Calendar {
	output := termenv.NewOutput(os.Stdout)
	now := time.Now()
	return &Calendar{
		output: output,
		year:   now.Year(),
		month:  now.Month(),
	}
}

func (c *Calendar) Render() {
	header := fmt.Sprintf("%s %d", c.month.String(), c.year)
	headerStyle := c.output.String(header).Bold().Foreground(termenv.ANSIBrightCyan)
	fmt.Printf("%s\n\n", headerStyle.String())
	
	days := []string{"  Su ", " Mo ", " Tu ", " We ", " Th ", " Fr ", " Sa"}
	for _, day := range days {
		dayStyle := c.output.String(day).Bold().Foreground(termenv.ANSIYellow)
		fmt.Printf("%4s", dayStyle.String())
	}
	fmt.Println()
	
	firstDay := time.Date(c.year, c.month, 1, 0, 0, 0, 0, time.Local)
	lastDay := firstDay.AddDate(0, 1, -1)
	numDays := lastDay.Day()
	startWeekday := int(firstDay.Weekday())
	
	today := time.Now()
	isCurrentMonth := today.Year() == c.year && today.Month() == c.month
	currentDay := today.Day()
	
	day := 1
	for week := 0; week < 6; week++ {
		for weekday := 0; weekday < 7; weekday++ {
			if week == 0 && weekday < startWeekday {
				// Empty space before month starts
				fmt.Print("    ")
			} else if day <= numDays {
				// Style the day number
				dayStr := fmt.Sprintf("%2d", day)
				var styledDay termenv.Style
				
				if isCurrentMonth && day == currentDay {
					styledDay = c.output.String(dayStr).Bold().Background(termenv.ANSIBrightRed).Foreground(termenv.ANSIBrightWhite)
				} else if weekday == 0 || weekday == 6 {
					styledDay = c.output.String(dayStr).Foreground(termenv.ANSIBrightMagenta)
				} else {
					styledDay = c.output.String(dayStr).Foreground(termenv.ANSIWhite)
				}
				
				fmt.Printf("  %s", styledDay.String())
				day++
			} else {
				fmt.Print("    ")
			}
		}
		fmt.Println()
		if day > numDays {
			break
		}
	}
	
	fmt.Println()
}

func (c *Calendar) NextMonth() {
	if c.month == 12 {
		c.month = 1
		c.year++
	} else {
		c.month++
	}
}

func (c *Calendar) PrevMonth() {
	if c.month == 1 {
		c.month = 12
		c.year--
	} else {
		c.month--
	}
}

func (c *Calendar) NextYear() {
	c.year++
}

func (c *Calendar) PrevYear() {
	c.year--
}

func main() {
	calendar := NewCalendar()
	calendar.Render()
}

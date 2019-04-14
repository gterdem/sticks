package solver

import (
	"fmt"
	"github.com/gterdem/sticks/stickhelper"
	"sort"
	"strconv"
	"strings"
)

//Stick object
type Stick struct {
	startH, startM, endH, endM int8
	status                     int
}

var lightSticks = []*Stick{}
var darkSticks = []*Stick{}
var resultSticks = []Stick{}

//Solve method solves the stick problem
func Solve(input string) {
	createSticksFromInput(input)
}

func createSticksFromInput(input string) {
	before := stickhelper.Before(input, " - ")
	after := stickhelper.After(input, " - ")
	fmt.Print(before, " - ", after)
	createTheSticks(before, true)
	createTheSticks(after, false)

	overlapSticks()
	resultSticks = getResults()
	printResults(resultSticks)
}

// func overlapSticks(lightSticks []*Stick, darkSticks []*Stick) {
func overlapSticks() {
	for i := 0; i < len(lightSticks); i++ {
		for j := 0; j < len(darkSticks); j++ {
			hasOverlap := isOverlapping(lightSticks[i], darkSticks[j])
			if hasOverlap {
				merge(lightSticks[i], darkSticks[j], i)
			}
		}
	}
}

func merge(light *Stick, dark *Stick, index int) {
	if dark.startH == light.startH {
		if dark.startM > light.startM {
			if dark.endH <= light.endH {
				// Stick divided to two
				extraStick := &Stick{startH: dark.endH, startM: dark.endM, endH: light.endH, endM: light.endM, status: stickhelper.Enum.Lit}
				lightSticks = append(lightSticks, extraStick)
			} else {
				//diminish the original light stick to the left
				lightSticks[index].endH = dark.startH
				lightSticks[index].endM = dark.startM
				lightSticks[index].status = stickhelper.Enum.Touched
			}

		}
	} else if dark.startH > light.startH {
		if dark.endH == light.endH {
			if dark.endM < light.endM {
				// Stick divided to two
				extraStick := &Stick{startH: dark.endH, startM: dark.endM, endH: light.endH, endM: light.endM}
				lightSticks = append(lightSticks, extraStick)
				//diminish the original light stick to the left
				lightSticks[index].endH = dark.startH
				lightSticks[index].endM = dark.startM
				lightSticks[index].status = stickhelper.Enum.Touched
			} else {
				//light stick diminished completely
				lightSticks[index].status = stickhelper.Enum.Diminished

			}
		} else if dark.endH > light.endH {
			//diminish the original light stick to the left
			lightSticks[index].endH = dark.startH
			lightSticks[index].endM = dark.startM
			lightSticks[index].status = stickhelper.Enum.Touched
		}
	} else {
		// Dark stick is longer on both left and right side that diminishes light stick completely
		lightSticks[index].status = stickhelper.Enum.Diminished
	}
	if dark.endH <= light.endH {
		if dark.startH >= light.startH {
			// Stick divided to two
			extraStick := &Stick{startH: dark.endH, startM: dark.endM, endH: light.endH, endM: light.endM, status: stickhelper.Enum.Lit}
			lightSticks = append(lightSticks, extraStick)
			//diminish the original light stick to the left
			lightSticks[index].endH = dark.startH
			lightSticks[index].endM = dark.startM
			lightSticks[index].status = stickhelper.Enum.Touched
		} else {
			lightSticks[index].startH = dark.endH
			lightSticks[index].startM = dark.endM
			lightSticks[index].status = stickhelper.Enum.Touched
		}
		// if dark.endM < light.endM {
		//diminish the original light stick to the right

		// }
	}
}
func isOverlapping(light *Stick, dark *Stick) bool {
	// If dark stick is longer or equal to light stick, light stick diminishes
	if dark.startH <= light.startH && dark.endH >= light.endH && light.status != stickhelper.Enum.Touched {
		light.status = stickhelper.Enum.Diminished
	}
	isStartOverlaped := false
	isEndOverlaped := false
	if dark.startH > light.startH {
		isStartOverlaped = true
	} else if dark.startH == light.startH {
		if dark.startM >= light.startM {
			isStartOverlaped = true
		}
	}
	if light.startH <= dark.endH && light.endH <= dark.endH {
		if light.startM <= dark.endM && light.endM <= light.endM {
			isStartOverlaped = true
		}
	}
	if dark.endH < light.endH {
		isEndOverlaped = true
	} else if dark.endH == light.endH {
		if dark.endM <= light.endM {
			isEndOverlaped = true
		}
	}
	if dark.startH == light.startH && dark.startM >= light.startM {
		isEndOverlaped = true
	}
	return isStartOverlaped && isEndOverlaped
}

func createTheSticks(input string, addToLightSticks bool) {
	removedParentheses := strings.Replace(input, "(", "", -1)
	removedParentheses = strings.Replace(removedParentheses, ")", "", -1)
	var splitted = make([]string, 0)
	// There are more then one stick (A time ranges)
	if strings.Contains(removedParentheses, ",") {
		splitted = strings.Split(removedParentheses, ",")
		if addToLightSticks {
			for _, item := range splitted {
				lightSticks = append(lightSticks, createStick(item))
			}
		} else {
			for _, item := range splitted {
				darkSticks = append(darkSticks, createStick(item))
			}
		}
	} else { // There is only one stick
		if addToLightSticks {
			lightSticks = append(lightSticks, createStick(removedParentheses))
		} else {
			darkSticks = append(darkSticks, createStick(removedParentheses))
		}

	}
}
func createStick(inputStr string) *Stick {
	splittedTime := strings.Split(inputStr, "-")
	begin := strings.TrimSpace(splittedTime[0])
	end := strings.TrimSpace(splittedTime[1])
	startHour, e := strconv.Atoi(strings.Split(begin, ":")[0])
	startMinute, e := strconv.Atoi(strings.Split(begin, ":")[1])
	endHour, e := strconv.Atoi(strings.Split(end, ":")[0])
	endMinute, e := strconv.Atoi(strings.Split(end, ":")[1])
	if e != nil {
		panic(e)
	}
	return &Stick{startH: int8(startHour), startM: int8(startMinute), endH: int8(endHour), endM: int8(endMinute), status: stickhelper.Enum.Lit}
}

func createValidInput(input string) string {
	newDinput := strings.TrimSpace(input)
	if strings.Contains(newDinput, "minus") {
		splitted := strings.Split(newDinput, "minus")
		first := splitted[0]
		return first
	}
	return ""
}
func printResults(sticks []Stick) {
	if len(sticks) == 0 {
		fmt.Print("( )")
	}
	for index, item := range sticks {
		startHourStr := "00"
		if item.startH > 0 {
			startHourStr = strconv.Itoa(int(item.startH))
		}
		startMinuteStr := "00"
		if item.startM > 0 {
			startMinuteStr = strconv.Itoa(int(item.startM))
		}
		endHourStr := "00"
		if item.endH > 0 {
			endHourStr = strconv.Itoa(int(item.endH))
		}
		endMinuteStr := "00"
		if item.endM > 0 {
			endMinuteStr = strconv.Itoa(int(item.endM))
		}
		if index == 0 {
			fmt.Print("(")
		}
		fmt.Print(startHourStr, ":", startMinuteStr, "-", endHourStr, ":", endMinuteStr)
		if index != len(sticks)-1 {
			fmt.Print(", ")
		}

		if index == len(sticks)-1 {
			fmt.Print(")")
		}
	}

}

func getResults() []Stick {
	for _, item := range lightSticks {
		if item.status != stickhelper.Enum.Diminished {
			if item.startH != item.endH || item.startM != item.endM { //Left over stick
				resultSticks = append(resultSticks, *item)
			}

		}
	}
	sort.Slice(resultSticks, func(i, j int) bool {
		return resultSticks[i].startH < resultSticks[j].startH
	})
	return resultSticks
}

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
func Solve(input string) []Stick {
	createSticksFromInput(input)
	//this needs to check between stick bounderies for inputs such as:
	// (01:00-03:00, 05:00:07) - (02:00:06) etc
	overlapSticksOuterBoundries()
	overlapSticks()
	return getResults()
}

// StringifyResult is used to return sampled result as string ==> I need to move Stick type to but that required quite a lot refactoring for public accessors etc
func StringifyResult(array []Stick) string {
	var str strings.Builder
	if len(array) == 0 {
		str.WriteString("()")
		return str.String()
	}
	for index, item := range array {
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
			str.WriteString("(")
		}
		str.WriteString(startHourStr)
		str.WriteString(":")
		str.WriteString(startMinuteStr)
		str.WriteString("-")
		str.WriteString(endHourStr)
		str.WriteString(":")
		str.WriteString(endMinuteStr)
		if index != len(array)-1 {
			str.WriteString(", ")
		}

		if index == len(array)-1 {
			str.WriteString(")")
		}
	}
	return str.String()
}

func createSticksFromInput(input string) {
	validInput := createValidInput(input)
	before := stickhelper.Before(validInput, " - ")
	after := stickhelper.After(validInput, " - ")
	if before == "" || after == "" {
		panic("Input format is invalid!")
	}

	// fmt.Print(before, " - ", after, " = ")
	createTheSticks(before, true)
	createTheSticks(after, false)
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
func overlapSticksOuterBoundries() {
	// first sort the light array
	for i := 0; i < len(darkSticks); i++ {
		dark := darkSticks[i]
		for j := 0; j < len(lightSticks); j++ {
			light := lightSticks[j]
			if j == 0 {
				if isDarkStartBiggerThanLightStart(light, dark) && !isDarkEndSmallerThanLightEnd(light, dark) {
					dark.status = stickhelper.Enum.Growing
				}
			}
			if dark.status == stickhelper.Enum.Growing {
				if isDarkEndPointBetweenLightLimits(light, dark) {
					//merge light sticks to ease further merging
					lightSticks[j].startH = lightSticks[j-1].startH
					lightSticks[j].startM = lightSticks[j-1].startM
					diminishLightStickAtIndex(j - 1)
				}
			}
		}
	}
}

func merge(light *Stick, dark *Stick, index int) {
	// Handle overnight times like 23:00-02:00
	if dark.startH > dark.endH {
		dark.endH = dark.endH + 24
	}
	if light.startH > light.endH {
		light.endH = light.endH + 24
	}
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
			if !isAlreadyExist(extraStick) {
				lightSticks = append(lightSticks, extraStick)
			}
			//diminish the original light stick to the left
			lightSticks[index].endH = dark.startH
			lightSticks[index].endM = dark.startM
			lightSticks[index].status = stickhelper.Enum.Touched
		} else {
			lightSticks[index].startH = dark.endH
			lightSticks[index].startM = dark.endM
			lightSticks[index].status = stickhelper.Enum.Touched
		}
	}
}
func isOverlapping(light *Stick, dark *Stick) bool {
	// If dark stick is longer or equal to light stick, light stick diminishes
	if dark.startH <= light.startH && dark.endH >= light.endH && light.status != stickhelper.Enum.Touched {
		light.status = stickhelper.Enum.Diminished
	}
	isStartOverlaped := false
	isEndOverlaped := false
	if dark.startH > light.startH && dark.startH <= dark.endH {
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
func isAlreadyExist(stick *Stick) bool {
	for _, item := range lightSticks {
		if stick.startH == item.startH && stick.startM == item.startM && stick.endH == item.endH && stick.endM == item.endM && stick.status == item.status {
			return true
		}
	}
	return false
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
	var str strings.Builder
	newDinput := strings.TrimSpace(input)
	if strings.Contains(newDinput, "minus") {
		splitted := strings.Split(newDinput, "minus")
		leftSide := strings.TrimSpace(splitted[0])
		rightSide := strings.TrimSpace(splitted[1])
		str.WriteString(leftSide)
		str.WriteString(" - ")
		str.WriteString(rightSide)
		return str.String()
	}
	return input
}

//PrintResults prints the result array in sampled string format - Obsolute
func PrintResults(sticks []Stick) {
	fmt.Print("Result: ", StringifyResult(sticks))
}

func getResults() []Stick {
	hasOvernightValue := false
	for _, item := range lightSticks {
		if item.status != stickhelper.Enum.Diminished {
			if item.startH != item.endH || item.startM != item.endM { //Left over stick
				if item.startH > 24 || item.endH > 24 {
					hasOvernightValue = true
				}
				item.startH = item.startH % 24
				item.endH = item.endH % 24
				resultSticks = append(resultSticks, *item)
			}

		}
	}
	// No need to sort overnight times like (23:00-23:30, 02:00-03:00)
	if !hasOvernightValue {
		sort.Slice(resultSticks, func(i, j int) bool {
			return resultSticks[i].startH < resultSticks[j].startH
		})
	}

	return resultSticks
}

func isDarkStartBiggerThanLightStart(light *Stick, dark *Stick) bool {
	if dark.startH > light.startH || (dark.startH == light.startH && dark.startM > light.startM) {
		return true
	}
	return false
}
func isDarkEndBiggerThanLightStart(light *Stick, dark *Stick) bool {
	if dark.endH > light.startH || (dark.startH == light.startM && dark.startM > light.startM) {
		return true
	}
	return false
}
func isDarkEndSmallerThanLightEnd(light *Stick, dark *Stick) bool {
	if dark.endH < light.endH || (dark.endH == light.endH && dark.endM < light.endM) {
		return true
	}
	return false
}
func isDarkBetweenLightLimits(light *Stick, dark *Stick) bool {
	return isDarkStartBiggerThanLightStart(light, dark) && isDarkEndSmallerThanLightEnd(light, dark)
}
func isDarkBeyondOrEqualToLightLimits(light *Stick, dark *Stick) bool {
	isBiggerOrEqualLeft := false
	isBiggerOrEqualRight := false
	if dark.startH > light.startH || (dark.startH == light.startH && dark.startM < light.startM) {
		isBiggerOrEqualLeft = true
	}
	if dark.endH > light.endH || (dark.endH == light.endH && dark.endM > light.endM) {
		isBiggerOrEqualRight = true
	}
	return isBiggerOrEqualLeft && isBiggerOrEqualRight
}
func isDarkEndPointBetweenLightLimits(light *Stick, dark *Stick) bool {
	isBiggerOrEqualLeft := false
	isLesserOrEqualRight := false
	if dark.endH > light.startH || (dark.endH == light.startH && dark.endM > light.startM) {
		isBiggerOrEqualLeft = true
	}
	if dark.endH < light.endH || (dark.endH == light.endH && dark.endM < light.endM) {
		isLesserOrEqualRight = true
	}
	return isBiggerOrEqualLeft && isLesserOrEqualRight
}

func diminishLightStickAtIndex(index int) {
	lightSticks[index].status = stickhelper.Enum.Diminished
	lightSticks[index].startH = 0
	lightSticks[index].startM = 0
	lightSticks[index].endH = 0
	lightSticks[index].endM = 0
}

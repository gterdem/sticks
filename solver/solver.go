//Package solver contains the solution of the problem.
//I simulated given problem as light and dark photon sort of stick game.
//Timelines of the A time range as light sticks and timelines of the B time range as dark sticks.
//When you overlap the dark sticks into light sticks; light sticks either shrink, diminish or sometimes cause a new light stick emerge.
//After overlapping, remaining light sticks are basically the result
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
	clearLightStickFromDiminishedLights()
	matchSticks()
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

// Matches the light sticks with dark sticks which results with modified light sticks
func matchSticks() {
	for i := 0; i < len(lightSticks); i++ {
		// Overlapping midnight solution
		light := lightSticks[i]
		if light.endH < light.startH {
			light.endH += 24
		}
		if light.status != stickhelper.Enum.Diminished {
			for j := 0; j < len(darkSticks); j++ {
				dark := darkSticks[j]
				// Overlapping midnight solution
				if dark.endH < dark.startH {
					dark.endH += 24
				}
				// Some of the bool controls are unnecessary but my wife says this way she can just read and understand easily
				if isDarkInBetweenLight(light, dark) { // Case 1 where dark stick is between the boundries of light stick hence causing a new light stick creation
					executeOrder12(light, dark, i)
				} else if isDarkStartSmallerOrEqualToLightStart(light, dark) && isDarkEndPointBetweenLightLimits(light, dark) { // Case-2 where dark stick overlaps to the beginning of the light stick, causing the original stick to shrink
					executeOrder24(light, dark, i)
				} else if isDarkStartPointBetweenLightLimits(light, dark) && isDarkEndBiggerOrEqualToLightEnd(light, dark) { // Case-3 where dark stick overlaps to the end of the light stick, causing the original stick to shrink
					executeOrder42(light, dark, i)
				} else if isDarkBeyondOrEqualToLight(light, dark) { // Case-4 where dark sticks compeletly overlaps the light stick which causes light stick's diminishing
					executeOrder66(light, dark, i)
				}
			}
		}
	}
}

// Creates the static sticks from input string where input string format is crucial!
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

// This method is used to create big light sticks where dark sticks tangle more than one light stick
func overlapSticksOuterBoundries() {
	// first sort the light array
	for i := 0; i < len(darkSticks); i++ {
		dark := darkSticks[i]
		for j := 0; j < len(lightSticks); j++ {
			light := lightSticks[j]
			if j == 0 {
				// if isDarkStartBiggerThanLightStart(light, dark) && !isDarkEndSmallerThanLightEnd(light, dark) {
				if isDarkStartPointBetweenLightLimits(light, dark) {
					dark.status = stickhelper.Enum.Growing
				}
			} else if dark.status == stickhelper.Enum.Growing {
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

// My wife was out of idea when i asked for function names
func executeOrder12(light *Stick, dark *Stick, index int) {
	// Choice 1) Create two new sticks and diminish the current one -> Didn't work because of dynamic adding to array and loop
	// Choice 2) Shrink the original light stick to left and add new one stick
	newStick := &Stick{startH: dark.endH, startM: dark.endM, endH: light.endH, endM: light.endM, status: stickhelper.Enum.Lit}

	lightSticks[index].endH = dark.startH
	lightSticks[index].endM = dark.startM
	lightSticks[index].status = stickhelper.Enum.Touched
	lightSticks = append(lightSticks, newStick)
}
func executeOrder24(light *Stick, dark *Stick, index int) {
	// Choice 1) Create new stick and diminish current stick => Doesn't worth the dynamic array hassle
	// newStick := &Stick{startH: dark.endH, startM: dark.endM, endH: light.endH, endM: light.endM, status: stickhelper.Enum.Lit}
	// lightSticks = append(lightSticks, newStick)
	// // lightSticks = append([]*Stick{newStick}, lightSticks...) // Need to prepend to prevent same objects adding to dynamic array and infinite loop
	// diminishLightStickAtIndex(index - 1)
	// Choice 2) Shrink current stick
	lightSticks[index].startH = dark.endH
	lightSticks[index].startM = dark.endM
	lightSticks[index].status = stickhelper.Enum.Touched
}

func executeOrder42(light *Stick, dark *Stick, index int) {
	// Choice 1) Create new stick and diminish the current stick => Doesn't worth the dynamic array hassle
	// newStick := &Stick{startH: light.startH, startM: light.startM, endH: dark.startH, endM: dark.startM, status: stickhelper.Enum.Lit}
	// // lightSticks = append(lightSticks, newStick)
	// lightSticks = append([]*Stick{newStick}, lightSticks...) // Need to prepend to prevent same objects adding to dynamic array and infinite loop
	// diminishLightStickAtIndex(index + 1)
	// Choice 2) Shrink current stick
	lightSticks[index].endH = dark.startH
	lightSticks[index].endM = dark.startM
	lightSticks[index].status = stickhelper.Enum.Touched
}

// Original function idea of my wife at midnight
func executeOrder66(light *Stick, dark *Stick, index int) {
	// Only 1 choice -> diminish the current light stick
	diminishLightStickAtIndex(index)
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

// Converts minus string to - operation. Can be improved
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

//PrintResults prints the result array in sampled string format
func PrintResults(sticks []Stick) {
	fmt.Print("Result: ", StringifyResult(sticks))
}

//Returns the final stick result array
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

// For Case 1
func isDarkInBetweenLight(light *Stick, dark *Stick) bool {
	return isDarkStartBiggerThanLightStart(light, dark) && isDarkEndSmallerThanLightEnd(light, dark)
}
func isDarkStartBiggerThanLightStart(light *Stick, dark *Stick) bool {
	if dark.startH > light.startH || (dark.startH == light.startH && dark.startM > light.startM) {
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

// For Case 2
func isDarkStartSmallerOrEqualToLightStart(light *Stick, dark *Stick) bool {
	if dark.startH <= light.startH || (dark.startH == light.startH && dark.startM <= light.startM) {
		return true
	}
	return false
}

// For Case 3
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

func isDarkEndBiggerOrEqualToLightEnd(light *Stick, dark *Stick) bool {
	if dark.endH >= light.endH || (dark.endH == light.endH && dark.endM >= light.endM) {
		return true
	}
	return false
}

// Case 4
func isDarkBeyondOrEqualToLight(light *Stick, dark *Stick) bool {
	isBiggerOrEqualLeft := false
	isBiggerOrEqualRight := false
	if dark.startH < light.startH || (dark.startH == light.startH && dark.startM <= light.startM) {
		isBiggerOrEqualLeft = true
	}
	if dark.endH > light.endH || (dark.endH == light.endH && dark.endM >= light.endM) {
		isBiggerOrEqualRight = true
	}
	return isBiggerOrEqualLeft && isBiggerOrEqualRight
}

func diminishLightStickAtIndex(index int) {
	lightSticks[index].status = stickhelper.Enum.Diminished
	lightSticks[index].startH = 0
	lightSticks[index].startM = 0
	lightSticks[index].endH = 0
	lightSticks[index].endM = 0
}

func clearLightStickFromDiminishedLights() {
	var cleanArr = []*Stick{}
	for _, stick := range lightSticks {
		if stick.status != stickhelper.Enum.Diminished {
			cleanArr = append(cleanArr, stick)
		}
	}
	lightSticks = cleanArr
}

func isDarkStartPointBetweenLightLimits(light *Stick, dark *Stick) bool {
	isBiggerOrEqualLeft := false
	isLesserOrEqualRight := false
	if dark.startH > light.startH || (dark.startH == light.startH && dark.startM > light.startM) {
		isBiggerOrEqualLeft = true
	}
	if dark.startH < light.endH || (dark.startH == light.endH && dark.startM < light.endM) {
		isLesserOrEqualRight = true
	}
	return isBiggerOrEqualLeft && isLesserOrEqualRight
}

package solver

import (
	"fmt"
	"github.com/gterdem/sticks/stringhelper"
	"strconv"
	"strings"
)

//Stick object
type Stick struct {
	startH, startM, endH, endM int8
}

var lightSticks = []Stick{}
var darkSticks = []Stick{}

//Solve method solves the stick problem
func Solve(input string) {
	createSticksFromInput(input)
}

func createSticksFromInput(input string) {
	before := stringhelper.Before(input, " - ")
	after := stringhelper.After(input, " - ")
	fmt.Println(before)
	fmt.Println(after)
	createTheSticks(before, true)
	createTheSticks(after, false)
	fmt.Println(lightSticks)
	fmt.Println(darkSticks)
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
			darkSticks = append(darkSticks, createStick(removedParentheses))
		}
	} else { // There is only one stick
		if addToLightSticks {
			lightSticks = append(lightSticks, createStick(removedParentheses))
		} else {
			darkSticks = append(darkSticks, createStick(removedParentheses))
		}

	}
}
func createStick(inputStr string) Stick {
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
	return Stick{startH: int8(startHour), startM: int8(startMinute), endH: int8(endHour), endM: int8(endMinute)}
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

// func after(value string, a string) string {
// 	// Get substring after a string.
// 	pos := strings.LastIndex(value, a)
// 	if pos == -1 {
// 		return ""
// 	}
// 	adjustedPos := pos + len(a)
// 	if adjustedPos >= len(value) {
// 		return ""
// 	}
// 	return value[adjustedPos:len(value)]
// }
// fmt.Println("Hello, playground", cap(lightSticks), len(lightSticks), lightSticks)

// 	fmt.Println("fuck", cap(blackSticks), len(blackSticks), blackSticks)
// 	// stick2 := createLightStick(input2)
// 	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
// 	z := strings.SplitN("a,b,c", ",", 0)
// 	fmt.Printf("%q (nil = %v)\n", z, z == nil)

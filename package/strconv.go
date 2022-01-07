package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Cast STR to BOOLEAN
	castStrToBool, err := strconv.ParseBool("True")
	if err == nil {
		fmt.Println(castStrToBool)
	} else {
		fmt.Println(err)
	}

	// Cast STR to INT
	castStrToInt, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(castStrToInt)
	} else {
		fmt.Println(err)
	}
	// Cast STR to INT with ATOI(no need bit or base number)
	castInt, _ := strconv.Atoi("2020")
	fmt.Println(castInt)

	// Cast STR to FLOAT
	castStrToFloat, err := strconv.ParseFloat("2010210", 64)
	if err == nil {
		fmt.Println(castStrToFloat)
	} else {
		fmt.Println(err)
	}

	// Cast BOOLEAN to STR
	castBoolToStr := strconv.FormatBool(true)
	fmt.Println(castBoolToStr)

	// Cast FLOAT to STR
	castFloatToStr := strconv.FormatFloat(2.512, 'g', 8, 64)
	fmt.Println(castFloatToStr)

	// Cast INT to STR
	castIntToStr := strconv.FormatInt(215, 10)
	fmt.Println(castIntToStr)
}

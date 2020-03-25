package util

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	fmt.Println(Now13())

	// test for start and end of a day
	fmt.Println(DayStart(0))
	fmt.Println(DayStart(time.Hour * 24))
	fmt.Println(DayStart(time.Hour * 24 * -1))
	fmt.Println(DayStartTime(0))
	fmt.Println(DayStartTime(24 * time.Hour))

}

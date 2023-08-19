package clockdiff

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleParseTime() {
	var t float64
	var err error

	t, err = ParseTime("8")
	if err != nil {
		log.Printf("Error parsing time: %s", err)
	}
	fmt.Println(t)

	t, err = ParseTime("8:00")
	if err != nil {
		log.Printf("Error parsing time: %s", err)
	}
	fmt.Println(t)

	t, err = ParseTime("08:00")
	if err != nil {
		log.Printf("Error parsing time: %s", err)
	}
	fmt.Println(t)

	t, err = ParseTime("08.00")
	if err != nil {
		log.Printf("Error parsing time: %s", err)
	}
	fmt.Println(t)

	t, err = ParseTime("8.00")
	if err != nil {
		log.Printf("Error parsing time: %s", err)
	}
	fmt.Println(t)

	t, err = ParseTime("8.30")
	if err != nil {
		log.Printf("Error parsing time: %s", err)
	}
	fmt.Println(t)

	t, err = ParseTime("8.45")
	if err != nil {
		log.Printf("Error parsing time: %s", err)
	}
	fmt.Println(t)

	// Output:
	// 8
	// 8
	// 8
	// 8
	// 8
	// 8.5
	// 8.75
}

func Test_ParseTime_errors_on_incorrect_input_format(t *testing.T) {
	var err error

	_, err = ParseTime("24")
	assert.NotNil(t, err, "There are only 24 hours in a day")

	_, err = ParseTime("24.00")
	assert.NotNil(t, err, "There are only 24 hours in a day")

	_, err = ParseTime("08.60")
	assert.NotNil(t, err, "There are only 60 minutes in an hour")

	_, err = ParseTime("08.60")
	assert.NotNil(t, err, "There are only 60 minutes in an hour")

	_, err = ParseTime("08.00.00")
	assert.NotNil(t, err, "Seconds are not supported")

	_, err = ParseTime("08:00.00")
	assert.NotNil(t, err, "Seconds are not supported")

	_, err = ParseTime("08:00:00")
	assert.NotNil(t, err, "Seconds are not supported")

	val := "08:.00"
	_, err = ParseTime(val)
	assert.NotNil(t, err, fmt.Sprintf("Garbage input should not be accepted: %q", val))

	val = "08:000"
	_, err = ParseTime(val)
	assert.NotNil(t, err, fmt.Sprintf("Garbage input should not be accepted: %q", val))

}

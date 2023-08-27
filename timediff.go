// Copyright 2022 timediff Authors
// SPDX-License-Identifier: Apache-2.0

package timediff

import (
	"fmt"
	"strconv"
	"strings"
)

func malformedTimeExpressionError(expr string) error {
	return fmt.Errorf("malformed time expression: %q", expr)
}

func calculateHours(split []string) (float64, error) {

	if len(split) > 2 {
		return -1, malformedTimeExpressionError(
			"too many fields in time format",
		)
	}

	if len(split[0]) > 2 {
		return -1, malformedTimeExpressionError("too many digits in hour field")
	}

	if len(split[1]) != 2 {
		return -1, malformedTimeExpressionError(
			"the minutes field must be two digits",
		)
	}

	hours, err := strconv.ParseInt(split[0], 10, 32)
	if err != nil {
		return -1, fmt.Errorf(
			"error in hours split: %w",
			malformedTimeExpressionError(split[0]),
		)
	}

	if hours > 23 || hours < 0 {
		return -1, fmt.Errorf("hours out of range: %d", hours)
	}

	minutes, err := strconv.ParseInt(split[1], 10, 32)
	if err != nil {
		return -1, fmt.Errorf(
			"error in minutes split: %w",
			malformedTimeExpressionError(split[1]),
		)
	}

	if minutes > 59 || minutes < 0 {
		return -1, fmt.Errorf("minutes out of range: %d", minutes)
	}

	minutesInHours := float64(minutes) / 60.0
	return float64(hours) + minutesInHours, nil

}

// ParseTime converts a time of day represented as a string from hours or hours
// and minutes into a scalar float representing the hours into the day.
//
// It takes a string containing a point in time in the format of 24-hour clock
// hours with minutes being optional. The separator between hours and minutes
// can be either a colon or a period, and seconds are not supported. The
// returned value is a `float64` representing the added hours and minutes as
// hours.
//
// As an example, the time 8 AM can be expressed in the following supported
// ways:
//
//	8
//	8:00
//	08:00
//	8.00
//	08:00
//
// The function wraps and returns any transitive errors, and also generates an
// error if the time expression is invalid.
func ParseTime(s string) (float64, error) {
	var split []string

	if strings.Contains(s, ":") {
		if strings.Contains(s, ".") {
			return -1, fmt.Errorf(
				"string contains both ':' and '.': %w",
				malformedTimeExpressionError(s),
			)
		}
		split = strings.Split(s, ":")
		return calculateHours(split)
	}

	if strings.Contains(s, ".") {
		if strings.Contains(s, ":") {
			return -1, fmt.Errorf(
				"string contains both '.' and ':': %w",
				malformedTimeExpressionError(s),
			)
		}
		split = strings.Split(s, ".")
		return calculateHours(split)
	}

	return calculateHours([]string{s, "00"})
}

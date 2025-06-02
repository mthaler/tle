package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

const LINE_LENGTH = 69

const CHECKSUM_INDEX = 68

func generateChecksum(line string) (int, error) {
	line = strings.TrimSpace(line)

	length := utf8.RuneCountInString(line)

	if length != CHECKSUM_INDEX {
		return -1, fmt.Errorf("Invalid line length: must be %d characters, received %d\n", CHECKSUM_INDEX, length)
	}

	return calculateChecksum(line)
}

/**
* Gets the checksum for the line as a <code>char</code>.
*
* @param line the line with a checksum digit
* @return the checksum
* @throws IllegalArgumentException if <code>line</code> is <code>null</code> or if
* <code>line</code> is not 69 characters in length
 */
func getChecksum(line string) (rune, error) {
	length := utf8.RuneCountInString(line)

	if length != LINE_LENGTH {
		return -1, fmt.Errorf("Line must be 69 characters long: %d\n", length)
	}

	return []rune(line)[CHECKSUM_INDEX], nil
}

func parseChecksum(line string) (int, error) {
	length := utf8.RuneCountInString(line)

	if length != LINE_LENGTH {
		return -1, fmt.Errorf("Line must be 69 characters long: %d\n", length)
	}

	c := []rune(line)[CHECKSUM_INDEX]

	return strconv.Atoi(string(c))
}

/**
 * Verifies the checksum for a line in a TLE.
 *
 * @param line the entire line of the TLE (including the checksum digit)
 * @return <code>true</code> if the checksum if valid, <code>false</code> otherwise
 * @throws IllegalArgumentException if <code>line</code> is <code>null</code> or if the line is
 * not 69 characters in length
 */
func isChecksumValid(line string) (bool, error) {
	line = strings.TrimSpace(line)

	length := utf8.RuneCountInString(line)

	if length != CHECKSUM_INDEX {
		return false, fmt.Errorf("Invalid line length: must be %d characters, received %d\n", CHECKSUM_INDEX, length)
	}

	c := []rune(line)[CHECKSUM_INDEX]

	checksum, err := strconv.Atoi(string(c))

	if err != nil {
		return false, err
	}

	calculatedChecksum := calculateChecksum(line.substring(0, CHECKSUM_INDEX))

	return checksum == calculatedChecksum, nil
}

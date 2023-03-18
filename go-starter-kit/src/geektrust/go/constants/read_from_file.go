package constants

import (
	"bufio"
	"os"
	"strconv"
)

func readLines(path string) []string {
	file, _ := os.Open(path)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func Init_Constants() {
	lines := readLines("magic_numbers.txt")
	ten := 10
	sixty_four := 64
	a, _ := strconv.ParseInt(lines[12], ten, sixty_four)
	Deal_min_threshold = int(a)
	b, _ := strconv.ParseInt(lines[13], ten, sixty_four)
	B4G1_threshold = int(b)
	set_floats(lines)
}

func set_floats(lines []string) {
	sixty_four := 64
	Cert_cost, _ = strconv.ParseFloat(lines[0], sixty_four)
	Degree_cost, _ = strconv.ParseFloat(lines[1], sixty_four)
	Diploma_cost, _ = strconv.ParseFloat(lines[2], sixty_four)
	Degree_discount, _ = strconv.ParseFloat(lines[3], sixty_four)
	Cert_discount, _ = strconv.ParseFloat(lines[4], sixty_four)
	Diploma_discount, _ = strconv.ParseFloat(lines[5], sixty_four)
	Enrollment_fee, _ = strconv.ParseFloat(lines[6], sixty_four)
	Enrollment_threshold, _ = strconv.ParseFloat(lines[7], sixty_four)
	Deal_G5, _ = strconv.ParseFloat(lines[8], sixty_four)
	Deal_G20, _ = strconv.ParseFloat(lines[9], sixty_four)
	Membership_fee, _ = strconv.ParseFloat(lines[10], sixty_four)
	Deal_threshold, _ = strconv.ParseFloat(lines[11], sixty_four)
	set_strings(lines)
}

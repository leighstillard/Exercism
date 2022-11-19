package partyrobot

import (
	"fmt"
	"strconv"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return "Happy birthday " + name + "! You are now " + strconv.Itoa(age) + " years old!"
}

func tableNumber(table int) string {
	var msg string
	switch {
	case table < 10:
		msg = "00" + strconv.Itoa(table)
	case table < 100:
		msg = "0" + strconv.Itoa(table)
	default:
		msg = strconv.Itoa(table)
	}
	return msg
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	message := ""

	message += Welcome(name) + "\n"
	message += "You have been assigned to table " + tableNumber(table) + ". "
	message += "Your table is " + direction + ", exactly " + fmt.Sprintf("%.1f", distance) + " meters from here.\n"
	message += "You will be sitting next to " + neighbor + "."

	return message
}

package partyrobot

import "fmt"

func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

func HappyBirthday(name string, age int) string {
	return "Happy birthday " + name + "! You are now " + fmt.Sprint(age) + " years old!"
}

func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var s1 string = "Welcome to my party, " + name + "!"
	var s2 string = "You have been assigned to table " + fmt.Sprintf("%03d", table) + "."
	var s3 string = "Your table is " + direction + ", exactly " + fmt.Sprintf("%.1f", distance) + " meters from here."
	var s4 string = "You will be sitting next to " + neighbor + "."
	return s1 + "\n" + s2 + " " + s3 + "\n" + s4
}

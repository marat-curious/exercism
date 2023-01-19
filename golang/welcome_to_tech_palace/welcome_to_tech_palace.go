package techpalace

import "strings"

func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var s string = strings.Repeat("*", numStarsPerLine)
	return s + "\n" + welcomeMsg + "\n" + s
}

func CleanupMessage(oldMsg string) string {
	return strings.Trim(oldMsg, "* \n")
}

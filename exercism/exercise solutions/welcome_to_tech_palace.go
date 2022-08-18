package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return ("Welcome to the Tech Palace, " + strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	if numStarsPerLine == 2 {
		return (strings.Repeat("*", numStarsPerLine) + "\nHi\n" + strings.Repeat("*", numStarsPerLine))
	} else {
		return (strings.Repeat("*", numStarsPerLine) + "\nWelcome!\n" + strings.Repeat("*", numStarsPerLine))
	}
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	removedStar := strings.ReplaceAll(oldMsg, "*", "")
	return strings.TrimSpace(removedStar)
}

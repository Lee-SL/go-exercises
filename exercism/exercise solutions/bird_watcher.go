package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	birdCount := 0
	for _, count := range birdsPerDay {
		birdCount += count
	}

	return birdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	startingIndex := (week * 7) - 7
	endIndex := (week * 7)

	birdCount := 0
	for _, count := range birdsPerDay[startingIndex:endIndex] {
		birdCount += count
	}
	return birdCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i <= len(birdsPerDay); i += 2 {
		if i >= len(birdsPerDay) {
			return birdsPerDay
		}
		birdsPerDay[i] += 1
	}
	return birdsPerDay
}

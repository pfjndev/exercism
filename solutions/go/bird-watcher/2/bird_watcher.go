package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
    for _, birds := range birdsPerDay {
        total += birds
    }
    return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	/*
    weekDays := 7
    switch week {
    	case 1:		return TotalBirdCount(birdsPerDay[:weekDays])
    	case 2:		return TotalBirdCount(birdsPerDay[weekDays:week*weekDays])
    	default:	return TotalBirdCount(birdsPerDay[(week-1)*weekDays:week*weekDays])
    }
    */
    // Calculate the start and end offsets for the slice
	start := (week - 1) * 7
	end := start + 7
	
	// Pass the specific "slice" of the week to TotalBirdCount
	return TotalBirdCount(birdsPerDay[start:end])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	
    for k := 0; k < len(birdsPerDay); k += 2 {
        birdsPerDay[k]++
    }
    return birdsPerDay
}
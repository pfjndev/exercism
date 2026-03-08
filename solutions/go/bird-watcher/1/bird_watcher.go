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
	weekDays := 7
    switch week {
    	case 1:		return TotalBirdCount(birdsPerDay[:weekDays])
    	case 2:		return TotalBirdCount(birdsPerDay[weekDays:week*weekDays])
    	default:	return TotalBirdCount(birdsPerDay[(week-1)*weekDays:week*weekDays])
    }
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	birdsPerDay[0] += 1
	
    for k := 2; k < len(birdsPerDay); k += 2 {
        birdsPerDay[k] += 1
    }
    return birdsPerDay
}

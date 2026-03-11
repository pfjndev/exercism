package electionday

import "fmt"

// ElectionResult represents the data of a candidate's results.
/* type ElectionResult struct {
	Name  string
	Votes int
} */

// NewVoteCounter returns a pointer to the initial votes.
func NewVoteCounter(initialVotes int) *int {
	return &initialVotes
}

// VoteCount safely extracts the value from the pointer.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}
	return *counter
}

// IncrementVoteCount updates the value at the memory address.
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}

// NewElectionResult returns a pointer to a new struct instance.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	return &ElectionResult{
		Name:  candidateName,
		Votes: votes,
	}
}

// DisplayResult formats the output string.
func DisplayResult(result *ElectionResult) string {
	// result.Name is shorthand for (*result).Name in Go!
	return fmt.Sprintf("%s (%d)", result.Name, result.Votes)
}

// DecrementVotesOfCandidate modifies the map in place.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate]--
}
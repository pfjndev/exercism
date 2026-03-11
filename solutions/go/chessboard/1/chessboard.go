package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool
// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	sum := 0
    // if _, exists := cb[file]; !exists { return sum }

    for _, occupied := range cb[file] {
        if occupied { sum++ }
    }
    return sum
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 { return 0 }

    sum := 0
    for _, occupied := range cb {
        if occupied[rank - 1] { sum++ }
    }
    return sum
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	total := 0
    for _, file := range cb {
        total += len(file)
    }
    return total
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	acc := 0
    for file := range cb {
        acc += CountInFile(cb, file)
    }
    return acc
}

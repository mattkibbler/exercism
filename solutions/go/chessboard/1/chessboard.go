package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	fileContents, fileExists := cb[file]
    if !fileExists {
        return 0
    }
    result := 0
    for _, occupied := range fileContents {
        if occupied {
            result += 1
        }
    }
    return result
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    result := 0
	for _, file := range cb {
        for i, occupied := range file {
            if (i + 1) == rank && occupied {
                result += 1
            }
        }
    }
    return result
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    result := 0
	for _, file := range cb {
        for range file {
            result += 1
        }
    }
    return result
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	result := 0
	for _, file := range cb {
        for _, occupied := range file {
            if occupied {
            	result += 1
            }
        }
    }
    return result
}

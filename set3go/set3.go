// Placeholder
package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(eta("upd", "admu", routeMap))
	fmt.Println(eta("upd", "dlsu", routeMap))
	fmt.Println(eta("dlsu", "admu", routeMap))
}

// data for relationshipStatus
var socialGraph = map[string]map[string]string{
	"@bongolpoc": {
		"first_name": "Joselito",
		"last_name":  "Olpoc",
		"following":  "",
	},
	"@joaquin": {
		"first_name": "Joaquin",
		"last_name":  "Gonzales",
		"following":  "@chums,@jobenilagan",
	},
	"@chums": {
		"first_name": "Matthew",
		"last_name":  "Uy",
		"following":  "@bongolpoc,@miketan,@rudyang,@joeilagan",
	},
	"@jobenilagan": {
		"first_name": "Joben",
		"last_name":  "Ilagan",
		"following":  "@eeebeee,@joeilagan,@chums,@joaquin",
	},
	"@joeilagan": {
		"first_name": "Joe",
		"last_name":  "Ilagan",
		"following":  "@eeebeee,@jobenilagan,@chums",
	},
	"@eeebeee": {
		"first_name": "EB",
		"last_name":  "Ilagan",
		"following":  "@jobenilagan,@joeilagan",
	},
}

// Relationship status

func relationshipStatus(fromMember string, toMember string, socialGraph map[string]map[string]string) string {

	fromUser, fromExists := socialGraph[fromMember] //check if from exists
	toUser, toExists := socialGraph[toMember]       //check if to exists

	if !fromExists || !toExists {
		return "no relationship"
	}

	//check if from/to follow each other --> search through following list
	fromFollowsTo := strings.Contains(fromUser["following"], toMember)
	toFollowsFrom := strings.Contains(toUser["following"], fromMember)

	switch {
	//mutual
	case fromFollowsTo && toFollowsFrom:
		return "friends"
		//from --> to
	case fromFollowsTo:
		return "follower"
		//to --> from
	case toFollowsFrom:
		return "followed by"
		//neither follow eo
	default:
		return "no relationship"
	}

}

// Tic tac toe
// dataset for tic tac toe
var board [][]string = [][]string{
	{"", "X", "O"},
	{"X", "O", "X"},
	{"O", "X", ""},
}

// helper function to check if all elements in the slice are the same and not empty
func allSame(slice []string) (string, bool) {
	if len(slice) == 0 {
		return "", false //return empty string and false empty
	}

	first := slice[0]
	if first == "" { // Ignore empty cells
		return "", false //return empty string and false
	}

	for _, val := range slice { //found different value
		if val != first {
			return "", false //return empty string and false
		}
	}
	return first, true //returns x or o if true
}

// Tic-Tac-Toe winner check function
func ticTacToe(board [][]string) string {
	size := len(board)

	// Check rows -- horizontal
	for _, row := range board {
		if winner, ok := allSame(row); ok {
			return winner
		}
	}

	// Check columns -- vertical
	for col := 0; col < size; col++ {
		column := make([]string, size)
		for row := 0; row < size; row++ {
			column[row] = board[row][col]
		}
		if winner, ok := allSame(column); ok {
			return winner
		}
	}

	//  Check diagonal (top/left to bottom/right)
	mainDiagonal := make([]string, size)
	for i := 0; i < size; i++ {
		mainDiagonal[i] = board[i][i]
	}
	if winner, ok := allSame(mainDiagonal); ok {
		return winner
	}

	// check diagonal (top/right to bottom/left)
	antiDiagonal := make([]string, size)
	for i := 0; i < size; i++ {
		antiDiagonal[i] = board[i][size-1-i]
	}
	if winner, ok := allSame(antiDiagonal); ok {
		return winner
	}

	return "NO WINNER" //if no winning lines
}

// eta
// dataset for eta
var routeMap = map[string]map[string]int{
	"upd,admu": {
		"travel_time_mins": 10,
	},
	"admu,dlsu": {
		"travel_time_mins": 35,
	},
	"dlsu,upd": {
		"travel_time_mins": 55,
	},
}

func eta(firstStop string, secondStop string, routeMap map[string]map[string]int) int {
	// Direct route check
	routeKey := firstStop + "," + secondStop
	if time, exists := routeMap[routeKey]; exists {
		return time["travel_time_mins"]
	}

	//not direct --> calculate circular route
	totalTime := 0
	currentStop := firstStop

	// loop until in second stop/full circle
	for {
		//next stop
		var nextStop string
		var travelTime int

		// current --> next
		for route, time := range routeMap {
			stops := strings.Split(route, ",")
			if stops[0] == currentStop {
				nextStop = stops[1]
				travelTime = time["travel_time_mins"]
				break
			}
		}

		totalTime += travelTime

		//check if destination
		if nextStop == secondStop {
			return totalTime
		}

		//prevent infinite loop
		if nextStop == firstStop {
			break
		}

		currentStop = nextStop
	}

	return 0 //no valid route
}

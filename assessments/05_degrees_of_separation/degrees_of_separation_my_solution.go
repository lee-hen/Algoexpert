package degrees_of_separation

import "math"

// "Aaron": ["Paul"],
// "Akshay": [],
// "Alex": ["Antoine", "Clement", "Ryan", "Simon"],
// "Antoine": ["Alex", "Clement", "Simon"],
// "Ayushi": ["Lee"],
// "Changpeng": ["Kelly", "Sandeep"],
// "Clement": ["Alex", "Antoine", "Sandeep", "Simon"],
// "Hannah": ["Lexi", "Paul", "Sandeep"],
// "James": ["Paul"],
// "Kelly": ["Changpeng", "Molly"],
// "Lee": ["Ayushi", "Molly"],
// "Lexi": ["Hannah"],
// "Molly": ["Kelly", "Lee"],
// "Paul": ["Aaron", "James", "Hannah"],
// "Ryan": ["Alex"],
// "Sandeep": ["Changpeng", "Clement", "Hannah"],
// "Simon": ["Alex", "Antoine", "Clement"]

// Antoine
// Clement

// "Clement"
// Neither Antoine nor Clement have any connection to Akshay.
// Antoine is seven degrees of separation away from Ayushi.
// So Clement only has one person who is more than six degrees of
// separation away (Akshay), whereas Antoine has two (Akshay and Ayushi).

type FriendsGraph map[string][]string

func degreesOfSeparation(friendsLists FriendsGraph, personOne, personTwo string) string {
	var anyConnections int
	for person, friendsList := range friendsLists {
		if person == personOne || person == personTwo {
			continue
		}

		if len(friendsList) == 0 {
			anyConnections += 1
		}
	}

	numOfDegreeSeparation1 := getNumOfDegreeSeparation(personOne, friendsLists, anyConnections)
	numOfDegreeSeparation2 := getNumOfDegreeSeparation(personTwo, friendsLists, anyConnections)

	if numOfDegreeSeparation1 < numOfDegreeSeparation2 {
		return personOne
	} else if numOfDegreeSeparation1 > numOfDegreeSeparation2 {
		return personTwo
	}

	return ""
}

func getNumOfDegreeSeparation(person string, friendsLists FriendsGraph, anyConnections int) int {
	degrees := make(map[string]int)
	visited := make(map[string]bool)

	dfs(person, 0, friendsLists, degrees, visited)

	if len(friendsLists[person]) == 0 {
		return math.MaxInt32
	}

	numDegreesOverSix := 0
	for _, degree := range degrees {
		if degree > 6 {
			numDegreesOverSix++
		}
	}
	return numDegreesOverSix + anyConnections
}

// iddfs can not handle cycles, so we must calculate the min degree every time,
//  and when iterate to next vertex(friend) we must set it to false
func dfs(person string, depth int, friendsLists FriendsGraph, degrees map[string]int, visited map[string]bool) {
	if _, ok := degrees[person]; ok {
		degrees[person] = min(degrees[person], depth)
	} else {
		degrees[person] = depth
	}

	if len(friendsLists[person]) == 1 && visited[friendsLists[person][0]] {
		return
	}

	visited[person] = true
	for _, friend := range friendsLists[person] {
		if visited[friend] {
			continue
		}

		dfs(friend, depth+1, friendsLists, degrees, visited)
		visited[friend] = false
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package degrees_of_separation
// important question

import (
	"math"
)

// DegreesOfSeparation
// O(v + e) time | O(v) space - where v is the number of vertices (people) in the
// friends graph and e is the number of edges (total friends) in the friends graph
// https://stackoverflow.com/questions/18604803/why-is-the-complexity-of-bfs-ove-instead-of-ove
func DegreesOfSeparation(friendsLists FriendsGraph, personOne, personTwo string) string {
	degreesOne := getDegreesOfSeparation(friendsLists, personOne)
	degreesTwo := getDegreesOfSeparation(friendsLists, personTwo)
	numDegreesOverSixOne := getNumDegreesOverSix(friendsLists, degreesOne)
	numDegreesOverSixTwo := getNumDegreesOverSix(friendsLists, degreesTwo)
	if numDegreesOverSixOne == numDegreesOverSixTwo {
		return ""
	} else if numDegreesOverSixOne < numDegreesOverSixTwo {
		return personOne
	}
	return personTwo
}

type DegreePair struct {
	Person string
	Degree int
}

func getDegreesOfSeparation(friendsLists FriendsGraph, origin string) map[string]int {
	degrees := map[string]int{}
	visited := map[string]bool{}
	queue := []DegreePair{{Person: origin, Degree: 0}}
	for len(queue) > 0 {
		var pair DegreePair
		pair, queue = queue[0], queue[1:]

		person, degree := pair.Person, pair.Degree
		degrees[person] = degree

		for _, friend := range friendsLists[person] {
			if visited[friend] {
				continue
			}
			visited[friend] = true
			queue = append(queue, DegreePair{
				Person: friend, Degree: degree + 1,
			})
		}
	}

	for person := range friendsLists {
		if !visited[person] {
			degrees[person] = math.MaxInt32
		}
	}
	return degrees
}

func getNumDegreesOverSix(friendsLists FriendsGraph, degrees map[string]int) int {
	numDegreesOverSix := 0
	for person := range friendsLists {
		if degrees[person] > 6 {
			numDegreesOverSix++
		}
	}
	return numDegreesOverSix
}
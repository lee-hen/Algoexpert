package airport_connections
// important question

import (
	"sort"
)

type AirportNode struct {
	Airport                string
	Connections            []string
	IsReachable            bool
	UnreachableConnections []string
}

func NewAirportNode(airport string) *AirportNode {
	return &AirportNode{
		Airport:                airport,
		Connections:            []string{},
		IsReachable:            true,
		UnreachableConnections: []string{},
	}
}

// AirportConnections
// O(a * (a + r) + a + r + alog(a)) time | O(a + r) space - where a is the number of airports and r is the number of routes
func AirportConnections(airports []string, routes [][]string, startingAirport string) int {
	airportGraph := createAirportGraph(airports, routes)
	unreachableAirportNodes := getUnreachableAirportNodes(airportGraph, airports, startingAirport)
	markUnreachableConnections(airportGraph, unreachableAirportNodes)
	return getMinNumberOfNewConnections(airportGraph, unreachableAirportNodes)
}

// O(a + r) time | O(a + r) space
func createAirportGraph(airports []string, routes [][]string) map[string]*AirportNode {
	airportGraph := map[string]*AirportNode{}
	for _, airport := range airports {
		airportGraph[airport] = NewAirportNode(airport)
	}
	for _, route := range routes {
		airport, connection := route[0], route[1]
		airportGraph[airport].Connections = append(airportGraph[airport].Connections, connection)
	}
	return airportGraph
}

// O(a + r) time | O(a) space
func getUnreachableAirportNodes(
	airportGraph map[string]*AirportNode, airports []string, startingAirport string,
) []*AirportNode {
	visitedAirports := map[string]bool{}
	depthFirstTraverseAirports(airportGraph, startingAirport, visitedAirports)

	unreachableAirportNodes := []*AirportNode{}
	for _, airport := range airports {
		if _, found := visitedAirports[airport]; found {
			continue
		}
		airportNode := airportGraph[airport]
		airportNode.IsReachable = false
		unreachableAirportNodes = append(unreachableAirportNodes, airportNode)
	}
	return unreachableAirportNodes
}

func depthFirstTraverseAirports(
	airportGraph map[string]*AirportNode, airport string, visitedAirports map[string]bool,
) {
	if _, found := visitedAirports[airport]; found {
		return
	}
	visitedAirports[airport] = true
	connections := airportGraph[airport].Connections
	for _, connection := range connections {
		depthFirstTraverseAirports(airportGraph, connection, visitedAirports)
	}
}

// O(a * (a + r)) time | O(a) space
func markUnreachableConnections(
	airportGraph map[string]*AirportNode, unreachableAirportNodes []*AirportNode,
) {
	for _, airportNode := range unreachableAirportNodes {
		airport := airportNode.Airport
		unreachableConnections := make([]string, 0)
		visitedAirports := map[string]bool{}
		depthFirstAddUnreachableConnections(airportGraph, airport, &unreachableConnections, visitedAirports)
		airportNode.UnreachableConnections = unreachableConnections
	}
	return
}

func depthFirstAddUnreachableConnections(
	airportGraph map[string]*AirportNode, airport string,
	unreachableConnections *[]string, visitedAirports map[string]bool,
) {
	if airportGraph[airport].IsReachable {
		return
	} else if _, found := visitedAirports[airport]; found {
		return
	}
	visitedAirports[airport] = true
	*unreachableConnections = append(*unreachableConnections, airport)
	connections := airportGraph[airport].Connections
	for _, connection := range connections {
		depthFirstAddUnreachableConnections(airportGraph, connection, unreachableConnections, visitedAirports)
	}
}

// O(alog(a) + a + r) time | O(1) space
func getMinNumberOfNewConnections(
	airportGraph map[string]*AirportNode, unreachableAirportNodes []*AirportNode,
) int {
	sort.SliceStable(unreachableAirportNodes, func(i, j int) bool {
		a1, a2 := unreachableAirportNodes[i], unreachableAirportNodes[j]
		return len(a1.UnreachableConnections) > len(a2.UnreachableConnections)
	})
	numberOfNewConnections := 0
	for _, node := range unreachableAirportNodes {
		if node.IsReachable {
			continue
		}
		numberOfNewConnections++
		for _, connection := range node.UnreachableConnections {
			airportGraph[connection].IsReachable = true
		}
	}
	return numberOfNewConnections
}

// my solution
//
//airports = [
//	"BGI", "CDG", "DEL", "DOH", "DSM", "EWR", "EYW", "HND", "ICN",
//	"JFK", "LGA", "LHR", "ORD", "SAN", "SFO", "SIN", "TLV", "BUD",
//]
//routes = [
//	["DSM", "ORD"],
//	["ORD", "BGI"],
//	["BGI", "LGA"],
//	["SIN", "CDG"],
//	["CDG", "SIN"],
//	["CDG", "BUD"],
//	["DEL", "DOH"],
//	["DEL", "CDG"],
//	["TLV", "DEL"],
//	["EWR", "HND"],
//	["HND", "ICN"],
//	["HND", "JFK"],
//	["ICN", "JFK"],
//	["JFK", "LGA"],
//	["EYW", "LHR"],
//	["LHR", "SFO"],
//	["SFO", "SAN"],
//	["SFO", "DSM"],
//	["SAN", "EYW"],
//]
// startingAirport = "LGA"
// 3 ["LGA", "TLV"], ["LGA", "SFO"], and ["LGA", "EWR"]

type Node struct {
	Name        string
	Score       int
	IsReachable bool
}

func airportConnections(airports []string, routes [][]string, startingAirport string) int {
	airportNodes := make(map[string]*Node)
	graph := make(map[*Node][]*Node)

	var startingAirportNode *Node
	for _, airport := range airports {
		airportNodes[airport] = &Node{
			Name: airport,
		}

		if airport == startingAirport {
			startingAirportNode = airportNodes[airport]
		}

		graph[airportNodes[airport]] = make([]*Node, 0)
	}

	for _, route := range routes {
		airport, destination := route[0], route[1]
		airportNode := airportNodes[airport]

		graph[airportNode] = append(graph[airportNode], airportNodes[destination])
	}

	if startingAirportNode != nil {
		startingAirportNode.dfs(graph)
	}

	unReachableAirports := getUnReachableAirports(graph)

	dfs(unReachableAirports, graph)

	sort.Slice(unReachableAirports, func(i, j int) bool {
		return unReachableAirports[i].Score > unReachableAirports[j].Score
	})

	var airportConnections int
	for _, airportNode := range unReachableAirports {
		if airportNode.IsReachable {
			continue
		}
		airportConnections++
		airportNode.dfs(graph)
	}

	return airportConnections
}

// very hard to come up with this solution
// because the airportNodeGraph may have cycle
// so total destination of the airportNode visited is score of the airportNode
// we must reset score and visit for each airportNode
func dfs(airportNodes []*Node, graph map[*Node][]*Node) {
	for _, airportNode := range airportNodes {
		var score = -1
		visited := make(map[*Node]bool)
		airportNode.dfs2(&score, graph, visited)
		airportNode.Score = score
	}
}

func (airportNode *Node) dfs2(score *int, graph map[*Node][]*Node, visited map[*Node]bool) {
	if visited[airportNode] {
		return
	} else if airportNode.IsReachable {
		return
	}
	visited[airportNode] = true

	*score++
	for _, destination := range graph[airportNode] {
		destination.dfs2(score, graph, visited)
	}
}

func (airportNode *Node) dfs(graph map[*Node][]*Node) {
	if airportNode.IsReachable {
		return
	}

	airportNode.IsReachable = true
	for _, destination := range graph[airportNode] {
		destination.dfs(graph)
	}
}

func getUnReachableAirports(graph map[*Node][]*Node) []*Node {
	unReachableAirports := make([]*Node, 0)

	for airportNode := range graph {
		if !airportNode.IsReachable {
			unReachableAirports = append(unReachableAirports, airportNode)
		}
	}

	return unReachableAirports
}

package boggle_board

type LetterNode struct {
	Letter rune

	leftNode  *LetterNode
	rightNode *LetterNode
	upNode    *LetterNode
	downNode  *LetterNode

	diagonalLeftUpNode    *LetterNode
	diagonalRightUpNode   *LetterNode
	diagonalLeftDownNode  *LetterNode
	diagonalRightDownNode *LetterNode
}

type LetterKey struct {
	i, j int
}

type LetterGraph struct {
	LetterNodes []*LetterNode

	Graph map[LetterKey]*LetterNode
}

//   0     1    2    3    4     5    6
// ['t', 'h', 'i', 's', 'i', 's', 'a'], 0
// ['s', 'i', 'm', 'p', 'l', 'e', 'x'], 1
// ['b', 'x', 'x', 'x', 'x', 'e', 'b'], 2
// ['x', 'o', 'g', 'g', 'l', 'x', 'o'], 3
// ['x', 'x', 'x', 'D', 'T', 'r', 'a'], 4
// ['R', 'E', 'P', 'E', 'A', 'd', 'x'], 5
// ['x', 'x', 'x', 'x', 'x', 'x', 'x'], 6
// ['N', 'O', 'T', 'R', 'E', '-', 'P'], 7
// ['x', 'x', 'D', 'E', 'T', 'A', 'E'], 8

// "this", "is", "not", "a", "simple", "boggle", "board", "test", "REPEATED", "NOTRE-PEATED"

func boggleBoard(board [][]rune, words []string) []string {
	letterGraph := createLetterGraph(board)

	firstWordMap := make(map[rune][]string)
	for _, searchWord := range words {
		r := rune(searchWord[0])
		firstWordMap[r] = append(firstWordMap[r], searchWord)
	}

	boggleBoardWords := make([]string, 0)
	for _, node := range letterGraph.LetterNodes {
		if searchWords, found := firstWordMap[node.Letter]; found {

			notFoundWords := make([]string, 0)
			for len(searchWords) > 0 {
				searchWord := searchWords[len(searchWords)-1]

				seen := make(map[*LetterNode]struct{})
				foundLetter := make(map[int]struct{})
				marchedWord := traverseLetters(node, seen, 0, foundLetter, []rune(searchWord))

				searchWords = searchWords[:len(searchWords)-1]
				firstWordMap[node.Letter] = searchWords

				if marchedWord != nil && string(marchedWord) == searchWord {
					boggleBoardWords = append(boggleBoardWords, string(marchedWord))
				} else {
					notFoundWords = append(notFoundWords, searchWord)
				}
			}

			firstWordMap[node.Letter] = append(firstWordMap[node.Letter], notFoundWords...)
		}
	}

	return boggleBoardWords
}

func traverseLetters(node *LetterNode, seen map[*LetterNode]struct{}, idx int, foundLetter map[int]struct{}, searchWord []rune) (marchedWord []rune) {
	if idx > len(searchWord)-1 {
		return
	}

	letter := searchWord[idx]
	if letter == node.Letter {
		if _, found := foundLetter[idx]; !found {
			marchedWord = append(marchedWord, letter)

			foundLetter[idx] = struct{}{}
			seen[node] = struct{}{}
		}

		idx++
	} else {
		return
	}

	if node.diagonalLeftUpNode != nil {
		if _, saw := seen[node.diagonalLeftUpNode]; !saw {
			marchedWord = append(marchedWord, traverseLetters(node.diagonalLeftUpNode, seen, idx, foundLetter, searchWord)...)
		}
	}

	if node.upNode != nil {
		if _, saw := seen[node.upNode]; !saw {
			marchedWord = append(marchedWord, traverseLetters(node.upNode, seen, idx, foundLetter, searchWord)...)
		}
	}

	if node.diagonalRightUpNode != nil {
		if _, saw := seen[node.diagonalRightUpNode]; !saw {
			marchedWord = append(marchedWord, traverseLetters(node.diagonalRightUpNode, seen, idx, foundLetter, searchWord)...)
		}
	}

	if node.leftNode != nil {
		if _, saw := seen[node.leftNode]; !saw {
			marchedWord = append(marchedWord, traverseLetters(node.leftNode, seen, idx, foundLetter, searchWord)...)
		}
	}

	if node.rightNode != nil {
		if _, saw := seen[node.rightNode]; !saw {
			marchedWord = append(marchedWord, traverseLetters(node.rightNode, seen, idx, foundLetter, searchWord)...)
		}
	}

	if node.diagonalLeftDownNode != nil {
		if _, saw := seen[node.diagonalLeftDownNode]; !saw {
			marchedWord = append(marchedWord, traverseLetters(node.diagonalLeftDownNode, seen, idx, foundLetter, searchWord)...)
		}
	}

	if node.downNode != nil {
		if _, saw := seen[node.downNode]; !saw {
			marchedWord = append(marchedWord, traverseLetters(node.downNode, seen, idx, foundLetter, searchWord)...)
		}
	}

	if node.diagonalRightDownNode != nil {
		if _, saw := seen[node.diagonalRightDownNode]; !saw {
			marchedWord = append(marchedWord, traverseLetters(node.diagonalRightDownNode, seen, idx, foundLetter, searchWord)...)
		}
	}

	return marchedWord
}

func createLetterGraph(board [][]rune) *LetterGraph {
	graph := NewLetterGraph(board)

	for i := range board {
		for j := range board[i] {
			key := LetterKey{
				i: i,
				j: j,
			}

			left := j - 1
			right := j + 1
			up := i - 1
			down := i + 1

			touchedLeftBorder := isOutOfBounds(i, left, len(board)-1, len(board[i])-1)
			touchedRightBorder := isOutOfBounds(i, right, len(board)-1, len(board[i])-1)
			touchedUpBorder := isOutOfBounds(up, j, len(board)-1, len(board[i])-1)
			touchedDownBorder := isOutOfBounds(down, j, len(board)-1, len(board[i])-1)

			if !touchedLeftBorder {
				leftLetter := board[i][left]
				graph.Graph[key].leftNode = graph.GetNode(leftLetter, i, left)
			}

			if !touchedRightBorder {
				rightLetter := board[i][right]
				graph.Graph[key].rightNode = graph.GetNode(rightLetter, i, right)
			}

			if !touchedUpBorder {
				upLetter := board[up][j]
				graph.Graph[key].upNode = graph.GetNode(upLetter, up, j)
			}

			if !touchedDownBorder {
				downLetter := board[down][j]
				graph.Graph[key].downNode = graph.GetNode(downLetter, down, j)
			}

			if !touchedLeftBorder && !touchedUpBorder {
				diagonalLeftUpLetter := board[up][left]
				graph.Graph[key].diagonalLeftUpNode = graph.GetNode(diagonalLeftUpLetter, up, left)
			}

			if !touchedRightBorder && !touchedUpBorder {
				diagonalRightUpLetter := board[up][right]
				graph.Graph[key].diagonalRightUpNode = graph.GetNode(diagonalRightUpLetter, up, right)
			}

			if !touchedLeftBorder && !touchedDownBorder {
				diagonalLeftDownLetter := board[down][left]
				graph.Graph[key].diagonalLeftDownNode = graph.GetNode(diagonalLeftDownLetter, down, left)
			}

			if !touchedRightBorder && !touchedDownBorder {
				diagonalRightDownLetter := board[down][right]
				graph.Graph[key].diagonalRightDownNode = graph.GetNode(diagonalRightDownLetter, down, right)
			}
		}
	}

	return graph
}

func isOutOfBounds(row, col, height, width int) bool {
	return row < 0 || row > height || col < 0 || col > width
}

func NewLetterGraph(board [][]rune) *LetterGraph {
	g := &LetterGraph{
		Graph: map[LetterKey]*LetterNode{},
	}

	for row := range board {
		for col := range board[row] {
			g.AddNode(board[row][col], row, col)
		}
	}

	return g
}

func (g *LetterGraph) AddNode(letter rune, i, j int) {
	key := LetterKey{
		i: i,
		j: j,
	}

	g.Graph[key] = &LetterNode{Letter: letter}
	g.LetterNodes = append(g.LetterNodes, g.Graph[key])
}

func (g *LetterGraph) GetNode(letter rune, i, j int) *LetterNode {
	key := LetterKey{
		i: i,
		j: j,
	}

	if _, found := g.Graph[key]; !found {
		g.AddNode(letter, i, j)
	}

	return g.Graph[key]
}

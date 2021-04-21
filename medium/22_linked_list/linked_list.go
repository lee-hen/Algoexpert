package linked_list

type Node struct {
	Value      int
	Prev, Next *Node
}

type DoublyLinkedList struct {
	Head, Tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return nil
}

func (ll *DoublyLinkedList) SetHead(node *Node) {
}

func (ll *DoublyLinkedList) SetTail(node *Node) {
}

func (ll *DoublyLinkedList) InsertBefore(node, nodeToInsert *Node) {
}

func (ll *DoublyLinkedList) InsertAfter(node, nodeToInsert *Node) {
}

func (ll *DoublyLinkedList) InsertAtPosition(position int, nodeToInsert *Node) {
}

func (ll *DoublyLinkedList) RemoveNodesWithValue(value int) {
}

func (ll *DoublyLinkedList) Remove(node *Node) {
}

func (ll *DoublyLinkedList) ContainsNodeWithValue(value int) bool {
	return false
}

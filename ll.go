package main
import "fmt"


type Node struct {
    data int      // Data stored in the node
    next *Node    // Pointer to the next node
}

// LinkedList represents the linked list structure
type LinkedList struct {
    head *Node    // Pointer to the first node
}


// InsertAtEnd adds a new node at the end of the linked list
func (ll *LinkedList) InsertAtEnd(data int) {
    newNode := &Node{data: data, next: nil}

    if ll.head == nil {
        ll.head = newNode
        return
    }

    current := ll.head
    for current.next != nil {
        current = current.next
    }
    current.next = newNode
}


// PrintList traverses and prints all elements in the linked list
func (ll *LinkedList) PrintList() {
    current := ll.head
    for current != nil {
        fmt.Printf("%d -> ", current.data)
        current = current.next
    }
    fmt.Println("nil")
}


// DeleteNode removes the first occurrence of a node with the given data
func (ll *LinkedList) DeleteNode(data int) {
    if ll.head == nil {
        return
    }

    // If the head node is to be deleted
    if ll.head.data == data {
        ll.head = ll.head.next
        return
    }

    current := ll.head
    for current.next != nil {
        if current.next.data == data {
            current.next = current.next.next
            return
        }
        current = current.next
    }
}






func main() {
    ll := LinkedList{}

    // Insert elements
    ll.InsertAtEnd(10)
    ll.InsertAtEnd(20)
    ll.InsertAtEnd(30)
    ll.InsertAtEnd(40)

    fmt.Println("Original Linked List:")
    ll.PrintList()  // Output: 10 -> 20 -> 30 -> 40 -> nil

    // Delete a node
    ll.DeleteNode(20)
    fmt.Println("After deleting 20:")
    ll.PrintList()  // Output: 10 -> 30 -> 40 -> nil

    // Insert another node
    ll.InsertAtEnd(50)
    fmt.Println("After inserting 50:")
    ll.PrintList()  // Output: 10 -> 30 -> 40 -> 50 -> nil
}






package main

import "fmt"

func main() {
    list := makeList()
    fmt.Println("original list:             ", list)

    list = reverseIterativeInPlace(makeList())
    fmt.Println("reverse iterative in-place:", list)

    list = reverseRecursiveInPlace(makeList(), nil)
    fmt.Println("reverse recursive in-place:", list)

    list = reverseIterativeNew(makeList())
    fmt.Println("reverse iterative new:     ", list)

    list = reverseRecursiveNew(makeList(), nil)
    fmt.Println("reverse recursive new:     ", list)
}

func makeList() *List {
    return &List{"a", &List{"b", &List{"c", &List{"d", nil}}}}
}

// Reverses the given list in-place by changing the pointers in the existing
// list's nodes. Returns a pointer to the new head node.
func reverseIterativeInPlace(list *List) *List {
    var prev *List

    for list != nil {
        temp := list.next
        list.next = prev
        prev, list = list, temp
    }

    return prev
}

// Reverses the given list in-place by changing the pointers in the existing
// list's nodes. Returns a pointer to the new head node.
func reverseRecursiveInPlace(list *List, accum *List) *List {
    if list == nil {
        return accum
    }

    temp := list.next
    list.next = accum
    return reverseRecursiveInPlace(temp, list)
}

// Creates a new list that is the reverse of the given one. Returns a pointer
// to the new list's head node.
func reverseIterativeNew(old *List) *List {
    var list *List

    for old != nil {
        list = &List{old.value, list}
        old = old.next
    }

    return list
}

// Creates a new list that is the reverse of the given one. Returns a pointer
// to the new list's head node.
func reverseRecursiveNew(old *List, accum *List) *List {
    if (old == nil) {
        return accum
    }

    return reverseRecursiveNew(old.next, &List{old.value, accum})
}

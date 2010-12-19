package main

// Define a simple singly-linked list of strings data type.
type List struct {
    value string
    next *List
}

// Implement Stringer so we can print lists.
func (list *List) String() string {
    if list == nil {
        return "(empty)"
    }

    result := ""
    for list != nil {
        if len(result) > 0 {
            result += " > "
        }
        result += list.value
        list = list.next
    }

    return result
}

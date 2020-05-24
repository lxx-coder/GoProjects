package leetcode

import "unicode"

//Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func isListPalindrome(head *ListNode) bool {
    vals := []int{}
    node := head
    for {
        if node == nil {
            break
        }
        vals = append(vals, node.Val)
        node = node.Next
    }

    front := 0
    end := len(vals)-1
    for {
        if front >= end {
            break
        }
        if vals[front] != vals[end] {
            return false
        }
        front++
        end--
    }
    return true
}

func isStrPalindrome(s string) bool {
    front := 0
    end := len(s)-1
    for {
        if front >= end {
            return true
        }
        if !unicode.IsLetter(rune(s[front])) && !unicode.IsDigit(rune(s[front])){
            front++
            continue
        }
        if !unicode.IsLetter(rune(s[end])) && !unicode.IsDigit(rune(s[end])) {
            end--
            continue
        }
        if (s[front] == s[end]) {
            front++
            end--
            continue
        }else if unicode.IsLetter(rune(s[front])) && unicode.IsLetter(rune(s[end])) {
            if (s[front] == s[end]-32) || (s[front] == s[end]+32) {
                front++
                end--
                continue
            }
        }
        return false
    }
}

func validPalindrome(s string) bool {

    n := len(s)
    for i := 0; i < n; i++ {
        if s[i] == s[n-1-i] {
            continue
        }
        return helper(s[i+1:n-i]) || helper(s[i:n-1-i])
    }
    return true
}

func helper(s string) bool {
    n := len(s)
    for i := 0; i < n; i++ {
        if s[i] == s[n-1-i] {
            continue
        }
        return false
    }
    return true
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }

    var head *ListNode
    var tmp *ListNode
    if l1.Val < l2.Val {
        head = l1
        tmp = head
        l1 = l1.Next
    }else {
        head = l2
        tmp = l2
        l2 = l2.Next
    }
    loop:
    for {
        if l1 == nil {
            tmp.Next = l2
            break loop
        }else if l2 == nil {
            tmp.Next = l1
            break loop
        }
        if l1.Val < l2 .Val {
            tmp.Next = l1
            l1 = l1.Next
        } else {
            tmp.Next = l2
            l2 = l2.Next
        }
        tmp = tmp.Next
    }
    return head
}

type ListNodeTmp struct {
    Val   int
    Used  bool
    Next  *ListNodeTmp
}

func findOrder(numCourses int, prerequisites [][]int) []int {
    var order []int
    var depend []int
    var depend []
    for i := 0;i < len(prerequisites); i++ {

        tmp1 := &ListNodeTmp{
            Val:  prerequisites[i][1],
            Used: false,
            Next: nil,
        }
        tmp2 := &ListNodeTmp{
            Val: prerequisites[i][0],
            Used: false,
            Next: tmp1,
        }
    }
}
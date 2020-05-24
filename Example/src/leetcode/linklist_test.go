package leetcode

import (
    "testing"
)

func TestIsListPalindrome(t *testing.T) {

}

func TestIsStrPalindrome(t *testing.T) {
    inputs := []string{"race a car","A man, a plan, a canal: Panama","abc121cba","0P"}
    outputs := []bool{false,true,true,false}

    for i := 0;i < len(inputs);i++ {
        if result := isStrPalindrome(inputs[i]);result != outputs[i] {
            t.Errorf("isStrPalindrome: Expect result of %v is %v, but instand is %v\n",inputs[i], outputs[i], result)
        } else {
            t.Logf("isStrPalindrome: Pass %v : %v\n",inputs[i], result)
        }
    }
}

func TestValidPalindrome(t *testing.T) {
    inputs := []string{"aba","abcbda","0P"}
    outputs := []bool{true,true,true}

    for i := 0;i < len(inputs);i++ {
        if result := validPalindrome(inputs[i]);result != outputs[i] {
            t.Errorf("validPalindrome: Expect result of %v is %v, but instand is %v\n",inputs[i], outputs[i], result)
        } else {
            t.Logf("validPalindrome: Pass %v : %v\n",inputs[i], result)
        }
    }
}

func TestMergeTwoLists(t *testing.T) {
    l1 := &ListNode{
        1,
        &ListNode{
            2,
            &ListNode{
                4,
                nil,
            },
        },
    }
    l2 := &ListNode{
        1,
        &ListNode{
            2,
            &ListNode{
                3,
                nil,
            },
        },
    }
    //input := [][]*ListNode{{l1, l2}}
    mergeTwoLists(l1, l2)
    //Add := func(a,b int) int {
    //    return a+b
    //}(1,2)
    //fmt.Print(Add)
}
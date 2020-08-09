package leetcode

import "testing"

func TestRepeatedStringMatch(t *testing.T) {

}


func TestConvert(t *testing.T) {
    s := "LEETCODEISHIRING"
    if result := convert(s, 3); result != "LCIRETOESIIGEDHN" {
        t.Errorf("Convert: Expect result: %s, your result: %s", result,"")
    }
}


func TestLongestPalindrome(t *testing.T) {
    inputs := []string{"babad", "cbbd", "abcdefedcba"}

    outputs := []string{"bab", "bb", "abcdefedcba"}

    for i := 0;i < len(inputs); i++ {
        if result := LongestPalindrome1(inputs[i]); result != outputs[i] {
            t.Error("Expect result: " + outputs[i] + ", your result: " + result)
        }
    }
}

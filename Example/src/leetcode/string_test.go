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
package leetcode

import (
    "fmt"
    "testing"
)

func TestLowestCommonAncestor(t *testing.T) {

}

func TestLevelOrder(t *testing.T) {
    root := &TreeNode{
        Val:    3,
        Left:   &TreeNode{
            Val:   9,
        },
        Right:  &TreeNode{
            Val:   20,
            Left:  &TreeNode{
                Val:   15,
            },
            Right: &TreeNode{
                Val:   7,
            },
        },
    }
    level := levelOrder(root)
    fmt.Println(level)
}

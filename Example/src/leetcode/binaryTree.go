package leetcode

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var ans *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    dfs(root, p, q)
    return ans
}

func dfs(root, p, q *TreeNode) bool {
    if root == nil {
        return false
    }
    lson := dfs(root.Left, p, q)
    rson := dfs(root.Right, p, q)
    if (lson && rson) || ((root.Val == p.Val || root.Val == q.Val) && (lson || rson)) {
        ans = root
    }

    return lson|| rson || root.Val == p.Val || root.Val == q.Val
}

func levelOrder(root *TreeNode) [][]int {
    array := make([][]*TreeNode, 0)
    if root == nil {
        return [][]int{}
    }
    array = append(array, []*TreeNode{root})
    for i := 0; i < len(array); i++ {
        tmp := make([]*TreeNode, 0)
        for j := 0; j < len(array[i]); j++ {
            if array[i][j].Left != nil {
                tmp = append(tmp, array[i][j].Left)
            }
            if array[i][j].Right != nil {
                tmp = append(tmp, array[i][j].Right)
            }
        }
        if len(tmp) > 0 {
            array = append(array, tmp)
        }
    }
    var ret [][]int
    for i := 0; i < len(array); i++ {
        var tmp []int
        for j := 0; j < len(array[i]); j++ {
            tmp = append(tmp, array[i][j].Val)
        }
        ret = append(ret, tmp)
    }

    return ret
}

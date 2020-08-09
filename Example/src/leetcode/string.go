package leetcode

//func repeatedStringMatch(A string, B string) int {
//
//}

//func isSubString(A string, B string) bool {
    //front := 0
    //end := len(A)
    //for {
    //    if
    //}
//}

func convert(s string, numRows int) string {
    column := len(s) / numRows
    if len(s) % numRows > 0 {
        column++
    }

    conv := []byte(s)
    n := 0
    for k :=0; k < numRows; k++ {
        if k == 0 || k == numRows-1 {
            for i := k; i < len(s); i += 2*numRows-2 {
                conv[n] = s[i]
                n++
            }
        }

        for i := k;i < len(s); i += 2*numRows-2 {
            conv[n] = s[i]
            n++
            conv[n] = s[i + 2*(n-1-k)]
            n++
        }
    }

    return string(conv)
}

func LongestPalindrome(s string) string {
    n := len(s)
    ans := ""
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int ,n)
    }
    for k := 0; k < n;k++ {
        for i := 0; i+k < n; i++ {
            j := i+k
            if k == 0 {
                dp[i][j] = 1
            } else if k == 1 {
                if s[i] == s[j] {
                    dp[i][j] = 1
                }
            }else {
                if s[i] == s[j] {
                    dp[i][j] = dp[i+1][j-1]
                }
            }
            if dp[i][j] > 0 && k + 1 > len(ans) {
                ans = s[i:i+k+1]
            }
        }
    }
    return ans
}

func LongestPalindrome1(s string) string {
    if s == "" {
        return ""
    }
    start, end := 0, 0
    for i := 0; i < len(s); i++ {
        left1, right1 := expandAroundCenter(s, i, i)
        left2, right2 := expandAroundCenter(s, i, i + 1)
        if right1 - left1 > end - start {
            start, end = left1, right1
        }
        if right2 - left2 > end - start {
            start, end = left2, right2
        }
    }
    return s[start:end+1]
}
func expandAroundCenter(s string, left, right int) (int, int) {
    for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left - 1, right + 1 {}
    return left + 1, right - 1
}
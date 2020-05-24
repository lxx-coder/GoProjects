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
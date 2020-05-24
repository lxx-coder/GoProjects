package leetcode

import (
    "fmt"
    "math"
    "sort"
    "strings"
)

//使用二进制运算获取1
//如何获取二进制中最右边的 1：x & (-x)；-x补码运算中为x的反码加1，可获取最右边的1
//如何将二进制中最右边的 1 设置为 0：x & (x - 1)


func reverse(x int) int {
    const (
        INT32_MAX = int32(^uint32(0) >> 1)
        INT32_MIN = ^INT32_MAX
    )


    flag := x > 0
    if !flag {
        x = -1 * x
    }
    if x <= 0 {
        return 0
    }
    var result uint64 = 0
    for {
        if flag && result > uint64(INT32_MAX) {
            return 0
        } else if !flag && result > uint64(-1*int64(INT32_MIN)) {
            return 0
        }
        if x <= 0 {
            if flag {
                return int(result)
            } else {
                return -1*int(result)
            }
        }
        tmp := x % 10
        x = x / 10
        result = 10*result + uint64(tmp)
    }

    //return 0
}

func reverseBits(num uint32) uint32 {
    var result uint32 = 0
    var count uint = 0
    for {
        if num <= 0 {
            return result << (32-count)
        }
        b := num & 1
        num = num >> 1
        count++
        result = result << 1
        result = result + b
        //fmt.Printf("%b %b\n",num, result)
    }

}

func hammingWeight(num uint32) int {
    var result int = 0
    for {
        if num <= 0 {
            return result
        }
        b := num & 1
        num = num >> 1
        result = result + int(b)
    }
}

func isPalindrome(x int) bool {
    if x < 0 || (x % 10 == 0 && x != 0) {
        return false
    }
    var right int = 0
    for {
        b := x % 10
        x = x / 10
        right = right * 10 + b
        if x <= right {
            return x == right || x == right / 10
        }
    }
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    m := len(nums1)
    n := len(nums2)
    if m > n {
        nums1, nums2 = nums2, nums1
        m,n = n,m
    }
    iMin := 0
    iMax := m
    halfLen := (m+n+1)/2
    for {
        if iMin > iMax {
            return 0
        }
         i := (iMin+iMax)/2
         j := halfLen - i
        if i < iMax && nums2[j-1] > nums1[i] {
            iMin = i + 1
        }else if i > iMin && nums1[i-1] > nums2[j] {
            iMax = i - 1
        }else {
            maxLeft := 0
            if i == 0 {
                maxLeft = nums2[j-1]
            }else if j == 0 {
                maxLeft = nums1[i-1]
            }else {
                maxLeft = int(math.Max(float64(nums1[i-1]), float64(nums2[j-1])))
            }
            if (m+n)%2 == 1 {
                return float64(maxLeft)
            }

            minRight := 0
            if i == m {
                minRight = nums2[j]
            }else if j == n {
                minRight = nums1[i]
            }else {
                minRight = int(math.Min(float64(nums2[j]), float64(nums1[i])))
            }
            return float64(maxLeft+minRight) / 2.0
        }
    }
}

func arrayPairSum(nums []int) int {
    sort.Ints(nums)
    sum := 0
    for i := 0;i < len(nums); i = i + 2 {
        sum += nums[i]
    }
    return sum
}

func sum(a []int) int {
    var s = 0
    for _, x := range a {
        s += x
    }
    return s
}

func canTwoPartsEqualSum(A []int) (bool, int) {
    left := 0
    right := len(A)-1
    for {
        index := (left + right) / 2
        if index < left || index > right || left > right {
            return false, 0
        }
        sumLeft := sum(A[:index])
        sumRight := sum(A[index:])
        if sumLeft == sumRight {
            return true, sumLeft
        }else if sumLeft < sumRight {
            left = index + 1
        }else {
            right = index - 1
        }
    }
}

func canThreePartsEqualSum(A []int) bool {
    for i := 1 ;i < len(A)-1; i++{
        leftSum := sum(A[:i])
        flag, rightSum := canTwoPartsEqualSum(A[i:])
        if flag && rightSum == leftSum {
            return true
        }
    }
    return false
}

func moveZeros(nums []int) {
    //for i := len(nums)-2; i >= 0; i-- {
    //    if nums[i] == 0 {
    //        for j := i+1; j < len(nums); j++ {
    //            if nums[j] != 0 {
    //                nums[j-1], nums[j] = nums[j], nums[j-1]
    //            }
    //        }
    //    }
    //}

    i := 0
    for ;i < len(nums) && nums[i] != 0; i++ {

    }

    j := i
    i = i+1
    for ; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[j], nums[i] = nums[i], nums[j]
            j++
        }
    }
    fmt.Printf("%v", strings.Replace(fmt.Sprint(nums)," ", ",", -1))
}

func singleNumber(nums []int) int {
    if len(nums) <= 0 {
        return -1
    }
    single := nums[0]
    for i := 1; i < len(nums); i++ {
        single ^= nums[i]
    }
    return single
}

func subarraySum(nums []int, k int) int {
    count := 0
    for start := 0; start < len(nums); start++ {
        sum := 0
        for end := start; end < len(nums); end++ {
            sum += nums[end]
            if sum == k {
                count++
            }
        }
    }
    return count
}

func subarraySum1(nums []int, k int) int {
    count, pre := 0, 0
    m := map[int]int{}
    m[0]=1
    for i := 0; i < len(nums);i++ {
        pre += nums[i]
        m[pre] += 1
        if _, ok := m[pre - k];ok {
            count += m[pre-k]
        }
    }
    return count
}

func romanToInt(s string) int {
    m := map[rune]int{'I':1,'V':5,'X':10,'L':50,'C':100,'D':500,'M':1000}
    sum := 0
    for i := 0; i < len(s); i++ {
        if i + 1 < len(s) && m[rune(s[i])] < m[rune(s[i+1])] {
            sum += m[rune(s[i+1])] - m[rune(s[i])]
            i++
            continue
        }
        sum += m[rune(s[i])]
    }
    return sum
}
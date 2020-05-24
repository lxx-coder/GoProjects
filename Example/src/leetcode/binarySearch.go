package leetcode

func BinarySearch(i int, arr []int) int {
    length := len(arr)
    if i > arr[length-1] || i < arr[0] {
        return -1
    }

    first := 0
    last := length-1
    for {
        index := (first + last) / 2
        if arr[index] == i {
            return index
        }else if arr[index] > i {
            last = index-1
        }else {
            first = index+1
        }

        if last < first {
            return -1
        }
    }
}

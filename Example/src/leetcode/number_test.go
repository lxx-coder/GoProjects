package leetcode

import "testing"

func TestReverse(t *testing.T) {
	inputs := []int{123,-123,120,2147483647,2147483641,-2147483641,241000}
	outputs := []int{321,-321,21,0,1463847412,-1463847412,142}

	for i :=0;i < len(inputs);i++ {
		if result := reverse(inputs[i]);result != outputs[i] {
			t.Errorf("reverse: Expect result of %v is %v, but instand is %v\n",inputs[i], outputs[i], result)
		} else {
			t.Logf("reverse: Pass %v : %v\n",inputs[i], result)
		}
	}

}

func TestReverseBits(t *testing.T) {
	inputs := []uint32{4294967293, 43261596, 0, 1}
	outputs := []uint32{3221225471, 964176192, 0, 2147483648}

	for i := 0;i < len(inputs);i++ {
		if result := reverseBits(inputs[i]);result != outputs[i] {
			t.Errorf("reverseBits: Expect result of %v is %v, but instand is %v\n",inputs[i], outputs[i], result)
		} else {
			t.Logf("reverseBits: Pass %v : %v\n",inputs[i], result)
		}
	}
}

func TestHammingWeight(t *testing.T) {
	inputs := []uint32{11, 0, 4294967293}
	outputs := []int{3, 0, 31}

	for i := 0;i < len(inputs);i++ {
		if result := hammingWeight(inputs[i]);result != outputs[i] {
			t.Errorf("hammingWeight: Expect result of %v is %v, but instand is %v\n",inputs[i], outputs[i], result)
		} else {
			t.Logf("hammingWeight: Pass %v : %v\n",inputs[i], result)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	inputs := []int{121, -121, 10,123454321}
	outputs := []bool{true, false, false,true}

	for i := 0;i < len(inputs);i++ {
		if result := isPalindrome(inputs[i]);result != outputs[i] {
			t.Errorf("isPalindrome: Expect result of %v is %v, but instand is %v\n",inputs[i], outputs[i], result)
		} else {
			t.Logf("isPalindrome: Pass %v : %v\n",inputs[i], result)
		}
	}
}

func TestFindMedianSortedArrays(t *testing.T){
	inputs := [][][]int{{{1,3},{2}}, {{1,2},{3,4}}}
	outputs := []float64{2.0, 2.5}

	for i := 0;i < len(inputs);i++ {
		if result := findMedianSortedArrays(inputs[i][0],inputs[i][1]);result != outputs[i] {
			t.Errorf("isPalindrome: Expect result of %v is %v, but instand is %v\n",inputs[i], outputs[i], result)
		} else {
			t.Logf("isPalindrome: Pass %v : %v\n",inputs[i], result)
		}
	}
}
func TestArrayPairSum(t *testing.T) {
	inputs := [][]int{{1,4,3,2}}
	outputs := []int{4}

	for i := 0;i < len(inputs);i++ {
		if result := arrayPairSum(inputs[i]);result != outputs[i] {
			t.Errorf("isPalindrome: Expect result of %v is %v, but instand is %v\n",inputs[i], outputs[i], result)
		} else {
			t.Logf("isPalindrome: Pass %v : %v\n",inputs[i], result)
		}
	}
}
func TestCanThreePartsEqualSum(t *testing.T) {
	inputs := [][]int{{12,-4,16,-5,9,-3,3,8,0},{18,12,-18,18,-19,-1,10,10},{0,2,1,-6,6,-7,9,1,2,0,1},{0,2,1,-6,6,7,9,-1,2,0,1},{3,3,6,5,-2,2,5,1,-9,4}}
	outputs := []bool{true, true, true, false, true}

	for i := 0;i < len(inputs);i++ {
		if result := canThreePartsEqualSum(inputs[i]);result != outputs[i] {
			t.Errorf("canThreePartsEqualSum: Expect result of %v is %v, but instand is %v\n",inputs[i], outputs[i], result)
		} else {
			t.Logf("canThreePartsEqualSum: Pass %v : %v\n",inputs[i], result)
		}
	}
}

func TestMoveZeros(t *testing.T) {
	inputs := [][]int {
		{0,1,0,3,12},
		{0,0,0,0,1},
		{0,0,0,0,0},
	}
	//output := [][]int{
	//	{1,3,12,0,0},
	//}

	for i := 0; i < len(inputs); i++ {
		moveZeros(inputs[i])
	}
}

func TestSingleNumber(t *testing.T) {
	inputs := [][]int{
		{2,2,1},
		{4,1,2,1,2},
		{2,3,3,2,3},
	}

	outputs := []int {1,4,3}

	for i := 0; i < len(inputs); i++ {
		if result := singleNumber(inputs[i]); result != outputs[i] {
			t.Errorf("singleNumber: Expect result of %v is %v, but instand is %v\n",inputs[i], outputs[i], result)
		} else {
			t.Logf("canThreePartsEqualSum: Pass %v : %v\n",inputs[i], result)
		}
	}
}

func TestSubarraySum(t *testing.T) {
	input1 := [][]int{
		{1,1,1},

	}
	input2 := []int{2}

	output := []int{2}

	for i := 0; i < len(input1);i++ {
		if result := subarraySum(input1[i], input2[i]); result != output[i] {

		}
	}
}

func TestRomanToInt(t *testing.T) {
	inputs := []string{"III","IV","XXVII","MCMXCIV"}
	outputs := []int{3,4,27,1994}

	for i := 0;i < len(inputs); i++ {
		if result := romanToInt(inputs[i]); result != outputs[i] {
			t.Errorf("romanToInt: Expect result of %v is %v, but instand is %v\n",inputs[i], outputs[i], result)
		} else {
			t.Logf("romanToInt: Pass %v : %v\n",inputs[i], result)
		}
	}
}
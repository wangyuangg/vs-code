package main

import "fmt"

func main() {
	//二分查找
	//二分查找是一种在有序数组中查找某个元素的查找算法。
	//二分查找的基本思想是：
	//1.从数组的中间位置开始，如果中间位置的元素大于要查找的元素，则在数组的左半部分继续查找，如果中间位置的元素小于要查找的元素，则在数组的右半部分继续查找，如果中间位置的元素等于要查找的元素，则查找成功。
	//2.如果查找的元素大于数组的中间位置的元素，则在数组的右半部分继续查找，如果查找的元素小于数组的中间位置的元素，则在数组的左半部分继续查找。
	//3.如果查找的元素等于数组的中间位置的元素，则查找成功。
	arr := [6]int{1, 2, 3, 4, 5, 6}
	BinarySearch(&arr, 0, len(arr)-1, 6)
}

func BinarySearch(arr *[6]int, leftIndex int, rightIndex int, findval int) {
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findval {
		//在左半部分继续查找
		BinarySearch(arr, leftIndex, middle-1, findval)
	} else if (*arr)[middle] < findval {
		//在右半部分继续查找
		BinarySearch(arr, middle+1, rightIndex, findval)
	} else {
		//查找成功
		fmt.Println("find it",middle)
	}
}

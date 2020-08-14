package heap

/*堆排序
先把数组构建成一个大顶堆，此时堆顶就是最大值元素了，然后把堆顶元素放到数组的最后一个位置
把剩下的n-1个元素堆化(其实就是找出最大元素)，然后把堆顶放到数组的n-1的位置，重复这个过程，得到的数据就是从小到大的数组

如何把数组构建成一个大顶堆呢？
	从数组的n/2(第一个非叶子节点)开始往前进行堆化
*/

//BuildHeap 建堆，把给定的数组构建成一个堆
func BuildHeap(arr []int) {
	length := len(arr)

	for i := length / 2; i >= 1; i-- {
		heapify(arr, i, length)
	}
}

//HeapSort 堆排序
func HeapSort(arr []int, length int) {
	BuildHeap(arr)
	index := length
	for index >= 1 {
		swap(arr, 1, index)
		heapify(arr, 1, index-1)
		index--
	}
}

//堆化
func heapify(arr []int, top int, length int) {
	for i := top; i <= length/2; {
		position := i
		if arr[i] < arr[i*2] {
			position = i * 2
		}

		if i*2 <= length && arr[position] < arr[i*2+1] {
			position = i*2 + 1
		}
		if position == i {
			break
		}

		swap(arr, i, position)
		i = position
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

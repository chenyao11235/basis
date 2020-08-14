package hashtable

import (
	"fmt"
	"math"
	"sync"

	"github.com/OneOfOne/xxhash"
)

const expandFactor = 0.75

// 计算hash值的函数
var hashAlgorithm = func(key []byte) uint64 {
	h := xxhash.New64()
	h.Write(key)
	return h.Sum64()
}

// 键值对 其实一个链表
type keyPairs struct {
	key   string
	value interface{}
	next  *keyPairs
}

//HashMap hash表
type HashMap struct {
	array        []*keyPairs
	length       int
	capaticy     int
	capaticyMask int
	lock         sync.Mutex
}

//NewHashMap 创建hash表
func NewHashMap(capacity int) *HashMap {
	defaultCapacity := 1 << 4
	if capacity < defaultCapacity {
		capacity = defaultCapacity
	} else {
		capacity = 1 << (int(math.Ceil(math.Log2(float64(capacity)))))
	}

	m := new(HashMap)
	m.array = make([]*keyPairs, capacity, capacity)
	m.capaticy = capacity
	m.capaticyMask = capacity - 1

	return m
}

func (m *HashMap) hasIndex(key string, mask int) int {
	hash := hashAlgorithm([]byte(key))
	index := hash & uint64(mask)
	return int(index)
}

//Put 添加数据
func (m *HashMap) Put(key string, value interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()

	// 获取数据要存放的数组下标
	index := m.hasIndex(key, m.capaticyMask)
	// 获取元素
	element := m.array[index]

	// 元素为空，表示空链表，没有哈希冲突，直接赋值
	if element == nil {
		m.array[index] = &keyPairs{
			key:   key,
			value: value,
		}
	} else {
		var lastPairs *keyPairs
		// 遍历链表看元素是否存在，存在则替换，否则找到最后一个键值对
		for element != nil {
			if element.key == key {
				element.value = value
			}
			lastPairs = element
			element = element.next
		}
		// 找不到元素，就追加
		lastPairs.next = &keyPairs{
			value: value,
			key:   key,
		}
	}

	// 下面判断加载因子，看hash表是否需要扩容
	newLen := m.length + 1

	if float64(newLen)/float64(m.capaticy) >= expandFactor {
		// 新建一个两倍大于原来的hash表
		newM := new(HashMap)
		newM.array = make([]*keyPairs, 2*m.capaticy, 2*m.capaticy)
		newM.capaticy = 2 * m.capaticy
		newM.capaticyMask = 2*m.capaticy - 1

		// 遍历老的hash表，将键值对赋值给新的hash表
		for _, pairs := range m.array {
			for pairs != nil {
				newM.Put(pairs.key, pairs.value)
				pairs = pairs.next
			}
		}

		m.array = newM.array
		m.capaticy = newM.capaticy
		m.capaticyMask = newM.capaticyMask
	}

	m.length = newLen
}

//Get 查找
func (m *HashMap) Get(key string) (value interface{}, ok bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	index := m.hasIndex(key, m.capaticyMask)

	element := m.array[index]

	for element != nil {
		if element.key == key {
			return element.value, true
		}
		element = element.next
	}
	return
}

//Delete 删除
func (m *HashMap) Delete(key string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	index := m.hasIndex(key, m.capaticyMask)

	element := m.array[index]

	if element == nil {
		return
	}
	// 第一就是要删除的元素
	if element.key == key {
		m.array[index] = element.next
		m.length--
		return
	}

	nextElement := element.next
	for nextElement != nil {
		if nextElement.key == key {
			element.next = nextElement.next
			m.length = m.length - 1
			return
		}
		element = nextElement
		nextElement = nextElement.next
	}
}

//Range 遍历hash表
func (m *HashMap) Range() {
	m.lock.Lock()
	defer m.lock.Unlock()

	for _, pairs := range m.array {
		for pairs != nil {
			fmt.Printf("%v=%v, ", pairs.key, pairs.value)
			pairs = pairs.next
		}
	}
	fmt.Println()
}

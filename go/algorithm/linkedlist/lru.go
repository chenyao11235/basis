package linearlist

/* 通过链表实现lru算法
 */

//lruNode 节点
type lruNode struct {
	key   string
	value interface{}
	next  *lruNode
}

func newLruNode(key string, value interface{}) *lruNode {
	return &lruNode{
		key:   key,
		value: value,
	}
}

//LruCache 其实是链表
type LruCache struct {
	tail     *lruNode
	head     *lruNode
	length   int
	capacity int
}

//NewLruCache 新建
func NewLruCache(capacity int) *LruCache {
	return &LruCache{
		capacity: capacity,
	}
}

//Put 添加数据
func (l *LruCache) Put(key string, value interface{}) {
	element := newLruNode(key, value)
	// 链表中还没有任何数据
	if l.head == nil {
		l.head = element
		l.tail = element
		l.length++
		return
	}
	cur := l.head
	if l.length == 1 {
		if cur.key == key {
			return
		}
		l.head.next = element
		l.length++
		return
	}
	for cur.next != nil {
		// 找到了
		if cur.key == key {
			break
		}
		cur = cur.next
	}
	// 没有找到，在末尾添加新的
	if cur.next == nil {
		cur.next = element
		l.tail = element
		l.length++
	} else {
		cur.next = cur.next.next
		l.tail.next = element
	}
}

//Get 查询数据
func (l *LruCache) Get(key string) (value interface{}) {
	return
}

func (l *LruCache) delete() {

}

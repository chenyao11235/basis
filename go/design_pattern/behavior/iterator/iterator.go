package iterator

/*迭代器模式
迭代器模式一般用来遍历容器，很多语言都将迭代器作为一个基础的类库，直接提供出来
但是我们简单连了解一下它的原理

有一些点
	遍历的时候不要删除集合中的元素， 如果在遍历的时候增删了元素就抛出异常
*/

//Iterator 迭代器接口
type Iterator interface {
	hasNext() bool            // 是否还有下一个元素
	next()                    //向后移动一个元素
	currentItem() interface{} // 获取当前元素
}

package composite

//FileSystemNode 文件系统
type FileSystemNode struct {
	Path     string            // 路径
	IsFile   bool              // 是否是文件，还是目录
	SubNodes []*FileSystemNode // 子节点
}

func (n *FileSystemNode) getPath() string {
	return n.Path
}

//统计文件的总共的数目
func (n *FileSystemNode) countNumOfFile() int {
	return 0
}

//统计所有文件的总共的大小
func (n *FileSystemNode) countSizeOfFiles() int {
	return 0
}

func (n *FileSystemNode) addNode(node *FileSystemNode) {
	n.SubNodes = append(n.SubNodes, node)
}

func (n *FileSystemNode) removeNode(node *FileSystemNode) {
	length := len(n.SubNodes)

	var index int
	var item *FileSystemNode
	for index, item = range n.SubNodes {
		if node.getPath() == item.getPath() {
			break
		}
	}
	if index < length {
		n.SubNodes = append(n.SubNodes[:index], n.SubNodes[index+1:]...)
	}
}

/*
单纯从功能实现角度来说，上面的代码没有问题，已经实现了我们想要的功能。
但是，如果我们开发的是一个大型系统，从扩展性（文件或目录可能会对应不同的操作）、
业务建模（文件和目录从业务上是两个概念）、
代码的可读性（文件和目录区分对待更加符合人们对业务的认知）的角度来说，
我们最好对文件和目录进行区分设计，定义为 File 和 Directory 两个类。
*/

//IFileSystemNode 目录，文件通用接口
type IFileSystemNode interface {
	getPath() string
	countNumOfFile() int
	countSizeOfFiles() int
}

//File 文件
type File struct {
	path string
}

func (f *File) getPath() string {
	return f.path
}

//Directory 目录
type Directory struct {
	path string
}

func (d *Directory) getPath() string {
	return d.path
}

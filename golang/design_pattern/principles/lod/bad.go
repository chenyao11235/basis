package lod

/*
这个例子实现了简化版的搜索引擎爬取网页的功能。
代码中包含三个主要的类。其中，NetworkTransporter 类负责底层网络通信，根据请求获取数据；
HtmlDownloader 类用来通过 URL 获取网页；
Document 表示网页文档，后续的网页内容抽取、分词、索引都是以此为处理对象
*/

//NetworkTransporter 传输
type NetworkTransporter struct {
}

func (n *NetworkTransporter) send(r *HTMLRequest) []byte {
	return nil
}

//HTMLDownloader 获取网页
type HTMLDownloader struct {
	transporter NetworkTransporter
}

func (hd *HTMLDownloader) download(url string) *Html {
	r := &HTMLRequest{
		url: url,
	}
	rawHTML := hd.transporter.send(r)

	return &Html{rawHtml: rawHTML}
}

//Document 文档
type Document struct {
	html *Html
	url  string
}

// 从 document中提取有用的信息
func (d *Document) extract() {
}

//NewDocument 使用工厂函数
func NewDocument(url string, downloader *HTMLDownloader) *Document {
	html := downloader.download(url)
	return &Document{
		url:  url,
		html: html,
	}
}

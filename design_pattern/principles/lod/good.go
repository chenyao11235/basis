package lod

/*
这个例子实现了简化版的搜索引擎爬取网页的功能。
代码中包含三个主要的类。其中，NetworkTransporter 类负责底层网络通信，根据请求获取数据；
HtmlDownloader 类用来通过 URL 获取网页；
Document 表示网页文档，后续的网页内容抽取、分词、索引都是以此为处理对象
*/

//BetterNetworkTransporter 传输
type BetterNetworkTransporter struct {
}

// 作为底层类，不应该只服务于HTMLRequest
func (n *BetterNetworkTransporter) send(addr string, data []byte) []byte {
	return nil
}

//BetterHTMLDownloader 获取网页
type BetterHTMLDownloader struct {
	// transport 应该通过构造函数进行注入
	transporter BetterNetworkTransporter
}

func (hd *BetterHTMLDownloader) download(url string) *Html {
	r := &HTMLRequest{
		url: url,
	}
	rawHTML := hd.transporter.send(r.url, r.content)
	return &Html{rawHtml: rawHTML}
}

//BetterDocument 文档
type BetterDocument struct {
	html *Html
	url  string
}

// 从 document中提取有用的信息
func (d *BetterDocument) extract(url string) {
	d.url = url
	htmlDownloader := new(HTMLDownloader)
	d.html = htmlDownloader.download(url)
}

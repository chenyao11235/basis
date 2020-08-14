package lod

// 公用部分

//HTMLRequest 请求
type HTMLRequest struct {
	url     string
	content []byte
}

type Html struct {
	rawHtml []byte
}

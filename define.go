package wx

import "encoding/xml"

type Response struct {
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
type ParamNotify struct {
	MsgSignature string `json:"msg_signature"`
	Timestamp    string `json:"timestamp"`
	Nonce        string `json:"nonce"`
	Echostr      string `json:"echostr,omitempty"`
	EncryptType  string `json:"encrypt_type,omitempty"`
}

type ReturnResponseDataList struct {
	List  interface{} `json:"list"`
	Page  int         `json:"page"`
	Rows  int         `json:"rows"`
	Total int         `json:"total"`
	Count int         `json:"count"`
}

func (n *ParamNotify) IsSafeMode() bool {
	return n.EncryptType == "aes"
}

// CDATA  使用该类型,在序列化为 xml 文本时文本会被解析器忽略
type CDATA string

// MarshalXML 实现自己的序列化方法
func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

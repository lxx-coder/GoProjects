package xmlTest

import "testing"

func TestParseXML(t *testing.T) {
    input := `<xml><ToUserName><![CDATA[gh_b37a549f3a62]]></ToUserName><FromUserName><![CDATA[oln2dwMkNlfvWX1mU9-4yKqBjR3k]]></FromUserName><CreateTime>1549990748</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[你好]]></Content><MsgId>22190529027674590</MsgId></xml>`
    parseXML(input)
}

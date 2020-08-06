package dingtalk

import (
	"bytes"
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

const (
	MSG_TYPE_TEXT     = "text"
	MSG_TYPE_LINK     = "link"
	MSG_TYPE_MARKDOWN = "markdown"
)

type DingTalkMsg struct {
	MessageType string   `json:"msgtype"`
	Text        Text     `json:"text,omitempty"`
	Link        Link     `json:"link,omitempty"`
	MarkDown    MarkDown `json:"markdown,omitempty"`
	At          At       `json:"at,omitempty"`
}

type Text struct {
	Content string `json:"content"`
	At      At     `json:"at,omitempty"`
}

type Link struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	MessageURL string `json:"messageUrl"`
	PictureURL string `json:"picUrl,omitempty"`
}

type MarkDown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type At struct {
	IsAtAll   bool     `json:"isAtAll,omitempty"`
	AtMobiles []string `json:"atMobiles,omitempty"`
}

type dingTalkRsp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type dingTalkClient struct {
	dingRobotUrl string
}

func NewDingTalkClient(dingRobotUrl string) *dingTalkClient {
	return &dingTalkClient{
		dingRobotUrl: dingRobotUrl,
	}
}

func (d *dingTalkClient) SendText(text Text) error {
	return d.sendMsg(DingTalkMsg{
		MessageType: MSG_TYPE_TEXT,
		Text:        text,
	})
}

func (d *dingTalkClient) SendLink(link Link) error {
	return d.sendMsg(DingTalkMsg{
		MessageType: MSG_TYPE_LINK,
		Link:        link,
	})
}

func (d *dingTalkClient) SendMarkdown(markDown MarkDown) error {
	return d.sendMsg(DingTalkMsg{
		MessageType: MSG_TYPE_MARKDOWN,
		MarkDown:    markDown,
	})
}

func (d *dingTalkClient) sendMsg(dingMsg DingTalkMsg) error {
	dingMsgByte, err := jsoniter.Marshal(dingMsg)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, d.dingRobotUrl, bytes.NewReader(dingMsgByte))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	rsp, err := new(http.Client).Do(req)
	if err != nil {
		return err
	}

	defer rsp.Body.Close()
	rspBody, _ := ioutil.ReadAll(rsp.Body)
	if rsp.StatusCode != http.StatusOK {
		return d.newDingError(string(rspBody))
	} else {
		dingTalkRsp := new(dingTalkRsp)
		jsoniter.Unmarshal(rspBody, dingTalkRsp)
		if dingTalkRsp.ErrCode != 0 {
			return d.newDingError(string(rspBody))
		}
	}
	return nil
}

func (d *dingTalkClient) newDingError(content string) error {
	return fmt.Errorf("访问钉钉【%s】出错了: %s", d.dingRobotUrl, content)
}

package dingtalk

import "testing"

func TestDingTalkClient_SendMarkdown(t *testing.T) {
	dingRobotUrl := "https://oapi.dingtalk.com/robot/send?access_token=abc"
	title := "hello"
	text := "hello world"

	err := NewDingTalkClient(dingRobotUrl).SendMarkdown(MarkDown{
		Title: title,
		Text:  text,
	})
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

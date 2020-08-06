# go-dingtalk
golang 钉钉消息发送

# Usage
```
	dingRobotUrl := "https://oapi.dingtalk.com/robot/send?access_token=abc"
	title := "hello"
	text := "hello world"

	dingtalk.NewDingTalkClient(dingRobotUrl).SendMarkdown(dingtalk.MarkDown{
		Title: title,
		Text:  text,
	})
```
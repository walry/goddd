package remote

import "testing"

func TestCallAi(t *testing.T) {
	result := CallTongYi(AskMsg{
		System: "你是一位资深互联网产品经理，请你将用户的需求转化为标准用户故事",
		User:   "当我创建完工作项后希望能通知到经办人",
	})

	t.Log(result)
}

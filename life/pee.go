package life

import "WeChatWorkRobot/utils"

func SetupPee() {
	TimingWithDayAt(15, 0, func() {
		handlePee("水芳说，走，一起去尿尿，🐔")
	})

	TimingWithDayAt(17, 30, func() {
		handlePee("尿尿了啊，🐦")
	})

	TimingWithDayAt(19, 30, func() {
		handlePee("让我们，大手拉小手，一起去尿尿🐔🐦🐔🐦🐔🐦")
	})
}

func handlePee(s string)  {
	for _, m := range utils.GetLifeWeChatRobotURL() {
		if m.Pee {
			utils.PostTextData(m.WeChatRobotURL, s, true)
		}
	}
}
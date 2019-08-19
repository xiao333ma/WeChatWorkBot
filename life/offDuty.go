package life

import (
	"WeChatWorkRobot/utils"
)

func SetupOffDuty()  {

	TimingWithDayAt(18, 00, func() {
		handleOffDuty("哎呀，6 点了，要不我们，下班？")
	})

	TimingWithDayAt(20, 30, func() {
		handleOffDuty("如果没事的话，就下班吧😊O(∩_∩)O~~")
	})

	TimingWithDayAt(21, 30, func() {
		handleOffDuty("小可爱们还在么。打车的话，可以开始约车了。注意身体哦")
	})

	TimingWithDayAt(22, 00, func() {
		handleOffDuty("下班下班，命重要！！！！")
	})

	TimingWithDayAt(22, 30, func() {
		handleOffDuty("晚安(～﹃～)~zZ")
	})

	TimingWithFriday(19, 0, func() {
		handleOffDuty("周末愉快 😊")
	})
}

func handleOffDuty(s string) {
	for _, m := range utils.GetLifeWeChatRobotURL() {
		if m.OffDuty {
			utils.PostTextData(m.WeChatRobotURL, s, true)
		}
	}
}
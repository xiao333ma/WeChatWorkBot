package life

import (
	"WeChatWorkRobot/utils"
)

func SetUpOrderOrEatMeal()  {
	TimingWithDayAt(11, 0, func() {
		handleMeal("该定午饭啦 (๑•̀ㅁ•́ฅ)")
	})
	TimingWithDayAt(11, 30, func() {
		handleMeal("我的饭到了么")
	})
	TimingWithDayAt(12, 0, func() {
		handleMeal("吃饭吃饭🍚")
	})
	TimingWithDayAt(17, 0, func() {
		handleMeal("该定晚饭啦，吃个啥嘞 (๑•̀ㅁ•́ฅ)")
	})

	TimingWithDayAt(17, 30, func() {
		handleMeal("饭在不到，就饿死了，O__O …")
	})

	TimingWithDayAt(18, 0, func() {
		handleMeal("吃饭喽，吃饭喽🍚")
	})
}

func handleMeal(s string)  {
	for _, m := range utils.GetLifeWeChatRobotURL() {
		if m.OrderMeal {
			utils.PostTextData(m.WeChatRobotURL, s, true)
		}
	}
}
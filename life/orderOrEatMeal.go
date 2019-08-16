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
	TimingWithDayAt(17, 0, func() {
		handleMeal("该定完饭啦 (๑•̀ㅁ•́ฅ)")
	})
}

func handleMeal(s string)  {
	for _, url := range utils.GetLifeWeChatRobotURL() {
		utils.PostTextData(url, s, true)
	}
}
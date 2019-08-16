package life

import (
	"WeChatWorkRobot/utils"
)

func SetupDrinkWater()  {
	TimingWithDayAt(11, 00, func() {
		handleDrinkWater("喝水啦，渴死了渴死了 (╯‵□′)╯︵┻━┻")
	})

	TimingWithDayAt(14, 20, func() {
		handleDrinkWater("睡醒了，不如喝点水吧")
	})

	TimingWithDayAt(15, 20, func() {
		handleDrinkWater("走，去接点水喝")
	})

	TimingWithDayAt(16, 30, func() {
		handleDrinkWater("一天据说要喝八杯水，喝水喝水")
	})

	TimingWithDayAt(17, 40, func() {
		handleDrinkWater("一会就要吃饭了，要不先接点水")
	})

	TimingWithDayAt(19, 25, func() {
		handleDrinkWater("水，你懂得 (~￣▽￣)~[]")
	})

	TimingWithDayAt(20, 30, func() {
		handleDrinkWater("要不，喝点儿~")
	})

	TimingWithDayAt(21, 30, func() {
		handleDrinkWater("够了够了，最后一杯，整好八杯")
	})
}

func handleDrinkWater(s string)  {
	for _, url := range utils.GetLifeWeChatRobotURL() {
		utils.PostTextData(url, s, true)
	}
}
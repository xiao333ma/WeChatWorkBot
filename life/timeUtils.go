package life

import (
	"github.com/jasonlvhit/gocron"
	"strconv"
)

func TimingWithDayAt(hour int, min int, handler func())  {
	go func() {
		s := strconv.Itoa(hour) + ":" + strconv.Itoa(min)
		t := gocron.NewScheduler()
		t.Every(1).Monday().At(s).Do(handler)
		t.Every(1).Tuesday().At(s).Do(handler)
		t.Every(1).Wednesday().At(s).Do(handler)
		t.Every(1).Tuesday().At(s).Do(handler)
		t.Every(1).Friday().At(s).Do(handler)
		<- t.Start()
	}()
}

func TimingWithFriday(hour int, min int, handler func())  {
	go func() {
		s := strconv.Itoa(hour) + ":" + strconv.Itoa(min)
		t := gocron.NewScheduler()
		t.Every(1).Friday().At(s).Do(handler)
		<- t.Start()
	}()
}
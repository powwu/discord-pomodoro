package main

import (
	"time"
	"log"

	g "github.com/AllenDang/giu"
	"github.com/hugolgst/rich-go/client"
)

var cycles int32 = 4
var worktime int32 = 25
var breaktime int32 = 5

func main() {

	err := client.Login("970351555121397760")
	if err != nil {
		log.Fatalf("%+v", err)
	}

	wnd := g.NewMasterWindow("Pomodoro", 250, 150, g.MasterWindowFlagsNotResizable)
	wnd.Run(gui)



}

func gui() {
	g.SingleWindow().Layout(
		g.Label("Discord Pomodoro"),
		g.SliderInt(&cycles, 1, 20).Label("Cycles"),
		g.SliderInt(&worktime, 1, 60).Label("Work Time"),
		g.SliderInt(&breaktime, 1, 30).Label("Break Time"),
		g.Button("Start").OnClick(func() { go pomodoro(int(cycles), int(worktime), int(breaktime)) }),
	)
}

func work(worktime int) {
	now := time.Now()
	endtime := time.Now().Add(time.Minute * time.Duration(worktime))
	err := client.SetActivity(client.Activity{
		State:      "Working...",
		Timestamps: &client.Timestamps{
			Start: &now,
			End: &endtime,
		},
		Buttons: []*client.Button{
			&client.Button{
				Label: "GitHub",
				Url:   "https://github.com/powwu/discord-pomodoro",
			},
		},
	})


	if err != nil {
		log.Fatalf("%+v", err)
	}
	time.Sleep(time.Minute * time.Duration(worktime))

	return
}
func breakstart(breaktime int) {
	now := time.Now()
	endtime := time.Now().Add(time.Minute * time.Duration(breaktime))
	err := client.SetActivity(client.Activity{
		State:      "Break time!",
		Timestamps: &client.Timestamps{
			Start: &now,
			End: &endtime,
		},
		Buttons: []*client.Button{
			&client.Button{
				Label: "GitHub",
				Url:   "https://github.com/powwu/discord-pomodoro",
			},
		},
	})


	if err != nil {
		log.Fatalf("%+v", err)
	}
	time.Sleep(time.Minute * time.Duration(worktime))

	return
}

func pomodoro(cycles int, worktime int, breaktime int) {
	work(worktime)
	for i := 0; i < cycles; i++ {
		breakstart(breaktime)
		work(worktime)
	}
}

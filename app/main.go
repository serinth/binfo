package main

import (
	ui "github.com/gizak/termui"
	"github.com/serinth/binfo/buildInfoFetchers"
)

func main() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	bamboo := buildInfoFetchers.NewBamboo("../config/config.json")

	ui.Render(&bamboo.BuildTable)
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/timer/1s", func(e ui.Event) {
		t := e.Data.(ui.EvtTimer)

		if t.Count%5 == 0 {
			bamboo.Update()

			//ui.Render(&bamboo.BuildTable)
			for _, g := range bamboo.BuildsInProgressGauges {
				//		fmt.Println(g.BorderLabel)
				//	fmt.Println(g.BorderLabel, g.Percent)
				ui.Render(&g)
			}
		}
	})

	ui.Loop()
}

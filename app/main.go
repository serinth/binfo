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

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/timer/1s", func(e ui.Event) {
		t := e.Data.(ui.EvtTimer)

		if t.Count%2 == 0 {
			bamboo.Update()
			bufferers := []ui.Bufferer{bamboo.BuildTable}

			for _, b := range bamboo.ActiveBuildGauges {
				bufferers = append(bufferers, b)
			}
			ui.Clear()
			ui.Render(bufferers...)
		}
	})

	ui.Loop()
}

package main

import (
	ui "github.com/gizak/termui"
	"github.com/serinth/binfo/buildInfoFetchers"
	configuration "github.com/serinth/binfo/config"
)

func main() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	configPath := "../config/config.json"
	config, _ := configuration.GetConfig(configPath)
	bamboo := buildInfoFetchers.NewBamboo(configPath)

	ui.Render(bamboo.BuildTable)

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/timer/1s", func(e ui.Event) {
		t := e.Data.(ui.EvtTimer)

		var interval uint64

		if len(bamboo.ActiveBuildGauges) > 0 {
			interval = 5
		} else {
			interval = config.RefreshIntervalSecs
		}

		if t.Count%interval == 0 {
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

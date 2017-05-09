package main

import (
	"fmt"
	"os"

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
	config, err := configuration.GetConfig(configPath)

	if err != nil {
		fmt.Println("There was an error getting the configuration from ../config/config.json: ", err)
		os.Exit(-1)
	}

	bamboo := buildInfoFetchers.NewBamboo(configPath)

	drawBody(*bamboo)

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
			drawBody(*bamboo)
		}
	})

	ui.Loop()
}

func drawBody(data buildInfoFetchers.Bamboo) {
	ui.Clear()
	ui.Body.Rows = nil
	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 0, data.BuildTable)),
		ui.NewRow(
			ui.NewCol(6, 0, data.ActiveBuildGauges...)))
	ui.Body.Align()
	ui.Render(ui.Body)
}

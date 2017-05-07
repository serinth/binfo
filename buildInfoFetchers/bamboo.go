package buildInfoFetchers

import (
	"strings"

	"strconv"

	"github.com/gizak/termui"
	configuration "github.com/serinth/binfo/config"
	"github.com/serinth/binfo/constants"
	"github.com/serinth/binfo/util"
)

type Bamboo struct {
	Config            configuration.Config
	BuildTable        termui.GridBufferer
	ActiveBuildGauges []termui.GridBufferer
	tableHeight       int
}

func NewBamboo(configPath string) *Bamboo {
	config, _ := configuration.GetConfig(configPath)
	table, height := createStatusTable(config)
	return &Bamboo{Config: config, BuildTable: table, tableHeight: height}

}

func createStatusTable(config configuration.Config) (termui.GridBufferer, int) {
	table := termui.NewTable()
	table.FgColor = termui.ColorWhite
	table.BgColor = termui.ColorDefault
	table.Rows = populateInitialProjectState(config.BuildServer, config.Projects)
	table.Border = true
	table.Analysis()
	table.SetSize()
	colorBuildStates(*table)

	return termui.GridBufferer(table), table.Height
}

func (b *Bamboo) Update() {
	table, height := createStatusTable(b.Config)
	b.BuildTable = table
	b.tableHeight = height
	b.ActiveBuildGauges = createInProgressGauges(height, b.Config)

}

func populateInitialProjectState(server string, projects []string) [][]string {
	var rows = [][]string{
		[]string{"Build Plan", "Last Built", "Status"},
	}

	for _, projectKey := range projects {
		result := &BambooBuildResourceResponse{}
		err := util.GetJson(buildResourceURL(server, projectKey), result)
		if err == nil {
			rows = append(rows, []string{result.PlanName, result.BuildRelativeTime, result.State})
		}
	}

	return rows
}

func colorBuildStates(table termui.Table) {
	for i, row := range table.Rows {
		if strings.Contains(row[len(row)-1], constants.FAILED) {
			table.BgColors[i] = termui.ColorRed
			table.FgColors[i] = termui.ColorBlack
		}
	}
}

func buildResourceURL(server string, key string) string {
	return server + "/rest/api/latest/result/" + key + "/latest"
}

func buildNextResourceURL(server string, key string, currentBuildNumber int) string {
	return server + "/rest/api/latest/result/" + key + "/" + strconv.Itoa(currentBuildNumber+1)
}

func createInProgressGauges(tableHeight int, config configuration.Config) []termui.GridBufferer {
	var gauges []termui.GridBufferer
	var Y = tableHeight
	for _, projectKey := range config.Projects {
		currentResourceResponse := &BambooBuildResourceResponse{}
		err := util.GetJson(buildResourceURL(config.BuildServer, projectKey), currentResourceResponse)
		if err == nil {
			resourceBuildInProgressResponse := &BambooBuildInProgressResponse{}
			inProgressError := util.GetJson(buildNextResourceURL(config.BuildServer, projectKey, currentResourceResponse.BuildNumber), resourceBuildInProgressResponse)

			if inProgressError == nil && resourceBuildInProgressResponse.State == constants.UNKNOWN {
				gauge := termui.NewGauge()
				percentageCompleted := int(resourceBuildInProgressResponse.Progress.PercentageCompleted * 100)

				if percentageCompleted >= 100 {
					gauge.Percent = 100
					gauge.Label = resourceBuildInProgressResponse.Progress.PrettyTimeRemaining
					gauge.PercentColorHighlighted = termui.ColorBlack
				} else {
					gauge.Percent = percentageCompleted
				}
				gauge.BorderLabel = resourceBuildInProgressResponse.PlanName
				gauge.BarColor = termui.ColorYellow
				gauge.BorderFg = termui.ColorWhite
				gauge.Width = 50
				Y += 3
				gauge.Y = Y + 3
				gauge.Height = 3

				gauges = append(gauges, termui.GridBufferer(gauge))
			}
		}
	}
	return gauges
}

type BambooBuildResourceResponse struct {
	PlanName          string
	LifeCycleState    string
	Finished          bool
	BuildDuration     int
	State             string
	BuildNumber       int
	BuildReason       string
	BuildRelativeTime string
}

type BambooBuildInProgressResponse struct {
	PlanName       string
	LifeCycleState string
	BuildReason    string
	Progress       progress
	State          string
	BuildNumber    int
	Finished       bool
}

type progress struct {
	IsUnderAverageTime         bool
	PercentageCompleted        float64
	PercentageCompletedPretty  string
	PrettyTimeRemaining        string
	PrettyAverageBuildDuration string
	PrettyBuildTime            string
	PrettyStartedTime          string
}

package buildInfoFetchers

import (
	"strings"

	"github.com/gizak/termui"
	configuration "github.com/serinth/binfo/config"
	"github.com/serinth/binfo/util"
)

type Bamboo struct {
	Config     configuration.Config
	BuildTable termui.Table
}

func NewBamboo(configPath string) *Bamboo {
	config, _ := configuration.GetConfig(configPath)

	return &Bamboo{Config: config, BuildTable: createTable(config)}

}

func createTable(config configuration.Config) termui.Table {
	table := termui.NewTable()
	table.FgColor = termui.ColorWhite
	table.BgColor = termui.ColorDefault
	table.Width = 150
	table.Height = 5 * len(config.Projects)
	table.Rows = populateInitialProjectState(config.BuildServer, config.Projects)
	table.Border = true
	table.Analysis()
	table.SetSize()
	colorBuildStates(*table)

	return *table
}

func (b *Bamboo) Update() {
	b.BuildTable = createTable(b.Config)
}

func (b *Bamboo) IsBeingBuilt(key string) bool {
	//	resp, err := util.GetJson()
	return true
}

func (b *Bamboo) GetBuildPercentage(projectKey string) int {
	return -1
}

func GetBuildStatus(projectKey string) string {
	return "not implemented"
}

func (b *Bamboo) getLatestBuildInfo() (*BambooBuildResourceResponse, error) {

	result := &BambooBuildResourceResponse{}

	err := util.GetJson(b.Config.BuildServer+"/rest/api/latest/result/"+b.Config.Projects[0]+"/latest", result)
	if err != nil {
		return nil, err
	}

	return result, nil

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
		if strings.Contains(row[len(row)-1], "Failed") {
			table.BgColors[i] = termui.ColorRed
			table.FgColors[i] = termui.ColorBlack
		}
	}
}

func buildResourceURL(server string, key string) string {
	return server + "/rest/api/latest/result/" + key + "/latest"
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

type BambooNotFoundResponse struct {
	StatusCode int `json:"status-code"`
	Message    string
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
	PercentageCompletedPretty  string
	PrettyTimeRemaining        string
	PrettyAverageBuildDuration string
	PrettyBuildTime            string
	PrettyStartedTime          string
}

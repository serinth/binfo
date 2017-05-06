package client

type BuildInfoFetcher interface {
	IsBeingBuilt(projectKey string) bool
	GetBuildPercentage(projectKey string) int
	GetBuildStatus(projectKey string) string
}

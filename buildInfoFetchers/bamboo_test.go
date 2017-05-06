package buildInfoFetchers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/serinth/binfo/util"
)

func TestGetBambooResourceResponse(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, BambooResourceResponseStub)
	}))
	defer ts.Close()

	var result = &BambooBuildResourceResponse{}
	err := util.GetJson(ts.URL, result)

	t.Log("Object is: ", result)

	if err != nil ||
		result.BuildNumber != 1 ||
		result.State != "Successful" ||
		result.PlanName != "testPlan" {
		t.Error("Failed to get json into domain object with error: ", err)
	}

}

func TestGetBambooInProgressResponse(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, BambooInProgressResponseStub)
	}))
	defer ts.Close()

	var result = &BambooBuildInProgressResponse{}
	err := util.GetJson(ts.URL, result)

	t.Log("Object is: ", result)

	if err != nil ||
		result.Progress.PercentageCompletedPretty != "19826%" ||
		result.Progress.PrettyAverageBuildDuration != "< 1 sec" ||
		result.State != "Unknown" {
		t.Error("Failed to get json into domain object with error: ", err)
	}

}

func TestBambooNotFoundResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, BambooNotFoundResponseStub)
	}))
	defer ts.Close()

	var result = &BambooNotFoundResponse{}
	err := util.GetJson(ts.URL, result)

	t.Log("Object is: ", result)

	if err != nil ||
		result.Message != "Result 12 for TES-TES not found." ||
		result.StatusCode != 404 {
		t.Error("Failed to get json into domain object with error: ", err)
	}
}

func TestBuildResourceUrlShouldReturnCorrectUrl(t *testing.T) {
	url := buildResourceURL("http://myserver", "PLANKEY-BUILDKEY")

	if url != "http://myserver/rest/api/latest/result/PLANKEY-BUILDKEY/latest" {
		t.Error("URL for Bamboo resource was incorrect")
	}
}

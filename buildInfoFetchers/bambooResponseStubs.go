package buildInfoFetchers

const BambooNotFoundResponseStub = `{
  						"message": "Result 12 for TES-TES not found.",
  						"status-code": 404
						}`

const BambooResourceResponseStub = `{
  "expand": "changes,metadata,plan,artifacts,comments,labels,jiraIssues,stages",
  "link": {
    "href": "http://localhost:8085/rest/api/latest/result/TES-TES-1",
    "rel": "self"
  },
  "plan": {
    "shortName": "testPlan",
    "shortKey": "TES",
    "type": "chain",
    "enabled": true,
    "link": {
      "href": "http://localhost:8085/rest/api/latest/plan/TES-TES",
      "rel": "self"
    },
    "key": "TES-TES",
    "name": "testProject - testPlan",
    "planKey": {
      "key": "TES-TES"
    }
  },
  "planName": "testPlan",
  "projectName": "testProject",
  "buildResultKey": "TES-TES-1",
  "lifeCycleState": "Finished",
  "id": 524289,
  "buildStartedTime": "2017-05-05T11:54:47.990Z",
  "prettyBuildStartedTime": "Fri, 5 May, 11:54 AM",
  "buildCompletedTime": "2017-05-05T11:54:48.402Z",
  "buildCompletedDate": "2017-05-05T11:54:48.402Z",
  "prettyBuildCompletedTime": "Fri, 5 May, 11:54 AM",
  "buildDurationInSeconds": 0,
  "buildDuration": 412,
  "buildDurationDescription": "< 1 second",
  "buildRelativeTime": "2 minutes ago",
  "buildTestSummary": "No tests found",
  "successfulTestCount": 0,
  "failedTestCount": 0,
  "quarantinedTestCount": 0,
  "skippedTestCount": 0,
  "continuable": false,
  "onceOff": false,
  "restartable": false,
  "notRunYet": false,
  "finished": true,
  "successful": true,
  "buildReason": "Manual run by <a href=\"http://172.17.0.2:8085/browse/user/root\">root</a>",
  "reasonSummary": "Manual run by <a href=\"http://172.17.0.2:8085/browse/user/root\">root</a>",
  "artifacts": {
    "size": 0,
    "start-index": 0,
    "max-result": 0
  },
  "comments": {
    "size": 0,
    "start-index": 0,
    "max-result": 0
  },
  "labels": {
    "size": 0,
    "start-index": 0,
    "max-result": 0
  },
  "jiraIssues": {
    "size": 0,
    "start-index": 0,
    "max-result": 0
  },
  "stages": {
    "size": 1,
    "start-index": 0,
    "max-result": 1
  },
  "changes": {
    "size": 0,
    "start-index": 0,
    "max-result": 0
  },
  "metadata": {
    "size": 2,
    "start-index": 0,
    "max-result": 2
  },
  "key": "TES-TES-1",
  "planResultKey": {
    "key": "TES-TES-1",
    "entityKey": {
      "key": "TES-TES"
    },
    "resultNumber": 1
  },
  "state": "Successful",
  "buildState": "Successful",
  "number": 1,
  "buildNumber": 1
}`

const BambooInProgressResponseStub = `{
  "expand": "changes,metadata,plan,artifacts,comments,labels,jiraIssues,stages",
  "link": {
    "href": "http://localhost:8085/rest/api/latest/result/TES-TES-5",
    "rel": "self"
  },
  "plan": {
    "shortName": "testPlan",
    "shortKey": "TES",
    "type": "chain",
    "enabled": true,
    "link": {
      "href": "http://localhost:8085/rest/api/latest/plan/TES-TES",
      "rel": "self"
    },
    "key": "TES-TES",
    "name": "testProject - testPlan",
    "planKey": {
      "key": "TES-TES"
    }
  },
  "planName": "testPlan",
  "projectName": "testProject",
  "buildResultKey": "TES-TES-5",
  "lifeCycleState": "InProgress",
  "id": 524300,
  "buildStartedTime": "2017-05-05T12:12:24.827Z",
  "prettyBuildStartedTime": "Fri, 5 May, 12:12 PM",
  "buildDurationInSeconds": 0,
  "buildDuration": 0,
  "buildDurationDescription": "Unknown",
  "buildRelativeTime": "",
  "continuable": false,
  "onceOff": false,
  "restartable": false,
  "notRunYet": false,
  "finished": false,
  "successful": false,
  "buildReason": "Manual run by <a href=\"http://172.17.0.2:8085/browse/user/root\">root</a>",
  "reasonSummary": "Manual run by <a href=\"http://172.17.0.2:8085/browse/user/root\">root</a>",
  "artifacts": {
    "size": 0,
    "start-index": 0,
    "max-result": 0
  },
  "comments": {
    "size": 0,
    "start-index": 0,
    "max-result": 0
  },
  "labels": {
    "size": 0,
    "start-index": 0,
    "max-result": 0
  },
  "jiraIssues": {
    "size": 0,
    "start-index": 0,
    "max-result": 0
  },
  "stages": {
    "size": 1,
    "start-index": 0,
    "max-result": 1
  },
  "changes": {
    "size": 0,
    "start-index": 0,
    "max-result": 0
  },
  "metadata": {
    "size": 2,
    "start-index": 0,
    "max-result": 2
  },
  "progress": {
    "isValid": true,
    "isUnderAverageTime": false,
    "percentageCompleted": 198.26591760299627,
    "percentageCompletedPretty": "19826%",
    "prettyTimeRemaining": "52 secs slower than usual",
    "prettyTimeRemainingLong": "52 seconds slower than usual",
    "averageBuildDuration": 267,
    "prettyAverageBuildDuration": "< 1 sec",
    "buildTime": 52937,
    "prettyBuildTime": "52 secs",
    "startedTime": "05 May 2017, 12:12:24 PM",
    "startedTimeFormatted": "2017-05-05T12:12:24",
    "prettyStartedTime": "52 seconds ago"
  },
  "key": "TES-TES-5",
  "planResultKey": {
    "key": "TES-TES-5",
    "entityKey": {
      "key": "TES-TES"
    },
    "resultNumber": 5
  },
  "state": "Unknown",
  "buildState": "Unknown",
  "number": 5,
  "buildNumber": 5
}`

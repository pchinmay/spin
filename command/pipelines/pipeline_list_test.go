package pipelines

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/spinnaker/spin/command"
)

func TestPipelineList_basic(t *testing.T) {
	ts := testGatePipelineListSuccess()
	defer ts.Close()

	meta := command.ApiMeta{}
	args := []string{"--application", "app", "--gate-endpoint", ts.URL}
	cmd := PipelineListCommand{
		ApiMeta: meta,
	}
	ret := cmd.Run(args)
	if ret != 0 {
		t.Fatalf("Command failed with: %d", ret)
	}
}

func TestPipelineList_flags(t *testing.T) {
	ts := testGatePipelineListSuccess()
	defer ts.Close()

	meta := command.ApiMeta{}
	args := []string{"--gate-endpoint", ts.URL} // Missing application.
	cmd := PipelineListCommand{
		ApiMeta: meta,
	}
	ret := cmd.Run(args)
	if ret == 0 { // Success is actually failure here, flags are malformed.
		t.Fatalf("Command failed with: %d", ret)
	}
}

func TestPipelineList_malformed(t *testing.T) {
	ts := testGatePipelineListMalformed()
	defer ts.Close()

	meta := command.ApiMeta{}
	args := []string{"--application", "app", "--gate-endpoint", ts.URL}
	cmd := PipelineListCommand{
		ApiMeta: meta,
	}
	ret := cmd.Run(args)
	if ret == 0 { // Success is actually failure here, return payload is malformed.
		t.Fatalf("Command failed with: %d", ret)
	}
}

func TestPipelineList_fail(t *testing.T) {
	ts := GateServerFail()
	defer ts.Close()

	meta := command.ApiMeta{}
	args := []string{"--application", "app", "--gate-endpoint", ts.URL}
	cmd := PipelineListCommand{
		ApiMeta: meta,
	}
	ret := cmd.Run(args)
	if ret == 0 { // Success is actually failure here, internal server error.
		t.Fatalf("Command failed with: %d", ret)
	}
}

// testGatePipelineListSuccess spins up a local http server that we will configure the GateClient
// to direct requests to. Responds with a 200 and a well-formed pipeline list.
func testGatePipelineListSuccess() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, strings.TrimSpace(pipelineListJson))
	}))
}

// testGatePipelineListMalformed returns a malformed list of pipeline configs.
func testGatePipelineListMalformed() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, strings.TrimSpace(malformedPipelineListJson))
	}))
}

const malformedPipelineListJson = `
  {
    "application": "app",
    "id": "pipeline1",
    "index": 0,
    "keepWaitingPipelines": false,
    "lastModifiedBy": "jacobkiefer@google.com",
    "limitConcurrent": true,
    "name": "derp1",
    "parameterConfig": [
      {
        "default": "bar",
        "description": "A foo.",
        "name": "foo",
        "required": true
      }
    ],
    "stages": [
      {
        "comments": "${ parameters.derp }",
        "name": "Wait",
        "refId": "1",
        "requisiteStageRefIds": [],
        "type": "wait",
        "waitTime": 30
      }
    ],
    "triggers": [],
    "updateTs": "1526578883109"
  }
]
`

const pipelineListJson = `
[
  {
    "application": "app",
    "id": "pipeline1",
    "index": 0,
    "keepWaitingPipelines": false,
    "lastModifiedBy": "jacobkiefer@google.com",
    "limitConcurrent": true,
    "name": "derp1",
    "parameterConfig": [
      {
        "default": "bar",
        "description": "A foo.",
        "name": "foo",
        "required": true
      }
    ],
    "stages": [
      {
        "comments": "${ parameters.derp }",
        "name": "Wait",
        "refId": "1",
        "requisiteStageRefIds": [],
        "type": "wait",
        "waitTime": 30
      }
    ],
    "triggers": [],
    "updateTs": "1526578883109"
  }
]
`

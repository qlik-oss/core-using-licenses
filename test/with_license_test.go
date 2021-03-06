package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Metrics []struct {
	Name   string `json:"name"`
	Metric Metric `json:"metric"`
}

type Metric []struct {
	Gauge Gauge `json:"gauge"`
}

type Gauge struct {
	Value int `json:"value"`
}

var (
	client = &http.Client{}
)

func getNumberActiveQixSessions() int {
	req, _ := http.NewRequest("GET", fmt.Sprintf("http://%s:9090/metrics", host), nil)
	req.Header.Add("Accept", "application/json")
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	var metrics Metrics
	json.Unmarshal(body, &metrics)

	var activeSessions int
	for _, metric := range metrics {
		if metric.Name == "qix_active_sessions" {
			activeSessions = metric.Metric[0].Gauge.Value
		}
	}
	return activeSessions
}

func getLicensesMetrics() string {
	req, _ := http.NewRequest("GET", fmt.Sprintf("http://%s:9200/metrics", host), nil)
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func getLicenseTimeConsumed(licensesMetrics string) int {
	re := regexp.MustCompile(`\nlicense_time_consumption{.*}\s(.+)\n`)
	matched := re.FindStringSubmatch(licensesMetrics)
	timeConsumed, _ := strconv.ParseFloat(matched[1], 64)
	return int(timeConsumed)
}

func getLicenseTimeTotal(licensesMetrics string) int {
	re := regexp.MustCompile(`\nlicense_time_total{.*}\s(.+)\n`)
	matched := re.FindStringSubmatch(licensesMetrics)
	totalTime, _ := strconv.ParseFloat(matched[1], 64)
	return int(totalTime)
}

func TestThatMoreThanFiveSessionsWorkWithALicense(t *testing.T) {

	var nbrIterations = 10
	var costPerSession = 5       // Each session cost x nbr of analyzer minutes specified in the license
	var totalTimeLicense = 10000 // Total number of analyzer minutes specified in the license

	licensesMetrics := getLicensesMetrics()
	licenseTimeConsumedBeforeTest := getLicenseTimeConsumed(licensesMetrics)

	for i := 0; i < nbrIterations; i++ {
		message, err := ConnectToEngineAndReturnOnConnectedEventMessage(ctx, i)
		assert.Equal(t, "SESSION_CREATED", message)
		assert.Nil(t, err, "Connecting to engine should not give an error")
	}

	licensesMetrics = getLicensesMetrics()
	licenseTimeConsumedAfterTest := getLicenseTimeConsumed(licensesMetrics)

	// Verify that the license time consumed reported on Licenses metrics matches the expected time consumed
	assert.Equal(t, nbrIterations*costPerSession, licenseTimeConsumedAfterTest-licenseTimeConsumedBeforeTest)

	// Verify total license time reported in Licenses metrics is the expected value
	assert.Equal(t, totalTimeLicense, getLicenseTimeTotal(licensesMetrics))
}

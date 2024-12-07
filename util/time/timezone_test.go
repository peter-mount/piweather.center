//go:build linux

package time

import (
	"testing"
	"time"
)

/*
 * THIS FILE IS GENERATED, DO NOT EDIT!
 * Please refer to util/time/gentest/gentest.go for how this is generated.
 */

func TestTimeZones(t *testing.T) {
	testTimeZonesInner(t, testCases)
}

// testCaseHandler function that runs a specific test
type testCaseHandler func(t *testing.T, testCase testCase)

// testCase definition
type testCase struct {
	name         string          // Name of test
	children     []testCase      // testCase's that are under this entry
	zone         string          // Time Zone name
	dstHandover  string          // Time of DST/ST handover
	date         string          // Date of test
	expectedDate string          // Date of expected result
	testHandler  testCaseHandler // test handler
}

var (
	// TimeZone test cases generated at 2024-12-07T08:12:02Z
	testCases = []testCase{
		{
			name: "Africa",
			children: []testCase{
				{
					name: "Cairo",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-04-25", testHandler: testLocalMidnight,
									zone: "Africa/Cairo", dstHandover: "2024-04-26T01:00:00+03:00",
									date: "2024-04-25T01:00:00+02:00", expectedDate: "2024-04-25T01:00:00+02:00",
								},
								{
									name: "2024-04-26", testHandler: testLocalMidnight,
									zone: "Africa/Cairo", dstHandover: "2024-04-26T01:00:00+03:00",
									date: "2024-04-26T01:00:00+03:00", expectedDate: "2024-04-26T01:00:00+03:00",
								},
								{
									name: "2024-04-27", testHandler: testLocalMidnight,
									zone: "Africa/Cairo", dstHandover: "2024-04-26T01:00:00+03:00",
									date: "2024-04-27T01:00:00+03:00", expectedDate: "2024-04-27T01:00:00+03:00",
								},
								{
									name: "2024-10-31", testHandler: testLocalMidnight,
									zone: "Africa/Cairo", dstHandover: "2024-11-01T01:00:00+02:00",
									date: "2024-10-31T01:00:00+03:00", expectedDate: "2024-10-31T01:00:00+03:00",
								},
								{
									name: "2024-11-01", testHandler: testLocalMidnight,
									zone: "Africa/Cairo", dstHandover: "2024-11-01T01:00:00+02:00",
									date: "2024-11-01T01:00:00+02:00", expectedDate: "2024-11-01T01:00:00+02:00",
								},
								{
									name: "2024-11-02", testHandler: testLocalMidnight,
									zone: "Africa/Cairo", dstHandover: "2024-11-01T01:00:00+02:00",
									date: "2024-11-02T01:00:00+02:00", expectedDate: "2024-11-02T01:00:00+02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-04-25", testHandler: testYesterdayMidnight,
									zone: "Africa/Cairo", dstHandover: "2024-04-26T01:00:00+03:00",
									date: "2024-04-25T01:00:00+02:00", expectedDate: "2024-04-25T01:00:00+02:00",
								},
								{
									name: "2024-04-26", testHandler: testYesterdayMidnight,
									zone: "Africa/Cairo", dstHandover: "2024-04-26T01:00:00+03:00",
									date: "2024-04-26T01:00:00+03:00", expectedDate: "2024-04-26T01:00:00+03:00",
								},
								{
									name: "2024-04-27", testHandler: testYesterdayMidnight,
									zone: "Africa/Cairo", dstHandover: "2024-04-26T01:00:00+03:00",
									date: "2024-04-27T01:00:00+03:00", expectedDate: "2024-04-27T01:00:00+03:00",
								},
								{
									name: "2024-10-31", testHandler: testYesterdayMidnight,
									zone: "Africa/Cairo", dstHandover: "2024-11-01T01:00:00+02:00",
									date: "2024-10-31T01:00:00+03:00", expectedDate: "2024-10-31T01:00:00+03:00",
								},
								{
									name: "2024-11-01", testHandler: testYesterdayMidnight,
									zone: "Africa/Cairo", dstHandover: "2024-11-01T01:00:00+02:00",
									date: "2024-11-01T01:00:00+02:00", expectedDate: "2024-11-01T01:00:00+02:00",
								},
								{
									name: "2024-11-02", testHandler: testYesterdayMidnight,
									zone: "Africa/Cairo", dstHandover: "2024-11-01T01:00:00+02:00",
									date: "2024-11-02T01:00:00+02:00", expectedDate: "2024-11-02T01:00:00+02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-04-25", testHandler: testTomorrowMidnight,
									zone: "Africa/Cairo", dstHandover: "2024-04-26T01:00:00+03:00",
									date: "2024-04-25T01:00:00+02:00", expectedDate: "2024-04-25T01:00:00+02:00",
								},
								{
									name: "2024-04-26", testHandler: testTomorrowMidnight,
									zone: "Africa/Cairo", dstHandover: "2024-04-26T01:00:00+03:00",
									date: "2024-04-26T01:00:00+03:00", expectedDate: "2024-04-26T01:00:00+03:00",
								},
								{
									name: "2024-04-27", testHandler: testTomorrowMidnight,
									zone: "Africa/Cairo", dstHandover: "2024-04-26T01:00:00+03:00",
									date: "2024-04-27T01:00:00+03:00", expectedDate: "2024-04-27T01:00:00+03:00",
								},
								{
									name: "2024-10-31", testHandler: testTomorrowMidnight,
									zone: "Africa/Cairo", dstHandover: "2024-11-01T01:00:00+02:00",
									date: "2024-10-31T01:00:00+03:00", expectedDate: "2024-10-31T01:00:00+03:00",
								},
								{
									name: "2024-11-01", testHandler: testTomorrowMidnight,
									zone: "Africa/Cairo", dstHandover: "2024-11-01T01:00:00+02:00",
									date: "2024-11-01T01:00:00+02:00", expectedDate: "2024-11-01T01:00:00+02:00",
								},
								{
									name: "2024-11-02", testHandler: testTomorrowMidnight,
									zone: "Africa/Cairo", dstHandover: "2024-11-01T01:00:00+02:00",
									date: "2024-11-02T01:00:00+02:00", expectedDate: "2024-11-02T01:00:00+02:00",
								},
							}},
					},
				},
				{
					name: "Casablanca",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "Africa/Casablanca", dstHandover: "2024-03-11T00:00:00Z",
									date: "2024-03-10T00:00:00+01:00", expectedDate: "2024-03-10T00:00:00+01:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "Africa/Casablanca", dstHandover: "2024-03-11T00:00:00Z",
									date: "2024-03-11T00:00:00Z", expectedDate: "2024-03-11T00:00:00Z",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "Africa/Casablanca", dstHandover: "2024-03-11T00:00:00Z",
									date: "2024-03-12T00:00:00Z", expectedDate: "2024-03-12T00:00:00Z",
								},
								{
									name: "2024-04-14", testHandler: testLocalMidnight,
									zone: "Africa/Casablanca", dstHandover: "2024-04-15T00:00:00+01:00",
									date: "2024-04-14T00:00:00Z", expectedDate: "2024-04-14T00:00:00Z",
								},
								{
									name: "2024-04-15", testHandler: testLocalMidnight,
									zone: "Africa/Casablanca", dstHandover: "2024-04-15T00:00:00+01:00",
									date: "2024-04-15T00:00:00+01:00", expectedDate: "2024-04-15T00:00:00+01:00",
								},
								{
									name: "2024-04-16", testHandler: testLocalMidnight,
									zone: "Africa/Casablanca", dstHandover: "2024-04-15T00:00:00+01:00",
									date: "2024-04-16T00:00:00+01:00", expectedDate: "2024-04-16T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "Africa/Casablanca", dstHandover: "2024-03-11T00:00:00Z",
									date: "2024-03-10T00:00:00+01:00", expectedDate: "2024-03-10T00:00:00+01:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "Africa/Casablanca", dstHandover: "2024-03-11T00:00:00Z",
									date: "2024-03-11T00:00:00Z", expectedDate: "2024-03-11T00:00:00Z",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "Africa/Casablanca", dstHandover: "2024-03-11T00:00:00Z",
									date: "2024-03-12T00:00:00Z", expectedDate: "2024-03-12T00:00:00Z",
								},
								{
									name: "2024-04-14", testHandler: testYesterdayMidnight,
									zone: "Africa/Casablanca", dstHandover: "2024-04-15T00:00:00+01:00",
									date: "2024-04-14T00:00:00Z", expectedDate: "2024-04-14T00:00:00Z",
								},
								{
									name: "2024-04-15", testHandler: testYesterdayMidnight,
									zone: "Africa/Casablanca", dstHandover: "2024-04-15T00:00:00+01:00",
									date: "2024-04-15T00:00:00+01:00", expectedDate: "2024-04-15T00:00:00+01:00",
								},
								{
									name: "2024-04-16", testHandler: testYesterdayMidnight,
									zone: "Africa/Casablanca", dstHandover: "2024-04-15T00:00:00+01:00",
									date: "2024-04-16T00:00:00+01:00", expectedDate: "2024-04-16T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "Africa/Casablanca", dstHandover: "2024-03-11T00:00:00Z",
									date: "2024-03-10T00:00:00+01:00", expectedDate: "2024-03-10T00:00:00+01:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "Africa/Casablanca", dstHandover: "2024-03-11T00:00:00Z",
									date: "2024-03-11T00:00:00Z", expectedDate: "2024-03-11T00:00:00Z",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "Africa/Casablanca", dstHandover: "2024-03-11T00:00:00Z",
									date: "2024-03-12T00:00:00Z", expectedDate: "2024-03-12T00:00:00Z",
								},
								{
									name: "2024-04-14", testHandler: testTomorrowMidnight,
									zone: "Africa/Casablanca", dstHandover: "2024-04-15T00:00:00+01:00",
									date: "2024-04-14T00:00:00Z", expectedDate: "2024-04-14T00:00:00Z",
								},
								{
									name: "2024-04-15", testHandler: testTomorrowMidnight,
									zone: "Africa/Casablanca", dstHandover: "2024-04-15T00:00:00+01:00",
									date: "2024-04-15T00:00:00+01:00", expectedDate: "2024-04-15T00:00:00+01:00",
								},
								{
									name: "2024-04-16", testHandler: testTomorrowMidnight,
									zone: "Africa/Casablanca", dstHandover: "2024-04-15T00:00:00+01:00",
									date: "2024-04-16T00:00:00+01:00", expectedDate: "2024-04-16T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Ceuta",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Africa/Ceuta", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Africa/Ceuta", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Africa/Ceuta", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Africa/Ceuta", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Africa/Ceuta", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Africa/Ceuta", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Africa/Ceuta", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Africa/Ceuta", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Africa/Ceuta", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Africa/Ceuta", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Africa/Ceuta", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Africa/Ceuta", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Africa/Ceuta", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Africa/Ceuta", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Africa/Ceuta", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Africa/Ceuta", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Africa/Ceuta", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Africa/Ceuta", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "El_Aaiun",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "Africa/El_Aaiun", dstHandover: "2024-03-11T00:00:00Z",
									date: "2024-03-10T00:00:00+01:00", expectedDate: "2024-03-10T00:00:00+01:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "Africa/El_Aaiun", dstHandover: "2024-03-11T00:00:00Z",
									date: "2024-03-11T00:00:00Z", expectedDate: "2024-03-11T00:00:00Z",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "Africa/El_Aaiun", dstHandover: "2024-03-11T00:00:00Z",
									date: "2024-03-12T00:00:00Z", expectedDate: "2024-03-12T00:00:00Z",
								},
								{
									name: "2024-04-14", testHandler: testLocalMidnight,
									zone: "Africa/El_Aaiun", dstHandover: "2024-04-15T00:00:00+01:00",
									date: "2024-04-14T00:00:00Z", expectedDate: "2024-04-14T00:00:00Z",
								},
								{
									name: "2024-04-15", testHandler: testLocalMidnight,
									zone: "Africa/El_Aaiun", dstHandover: "2024-04-15T00:00:00+01:00",
									date: "2024-04-15T00:00:00+01:00", expectedDate: "2024-04-15T00:00:00+01:00",
								},
								{
									name: "2024-04-16", testHandler: testLocalMidnight,
									zone: "Africa/El_Aaiun", dstHandover: "2024-04-15T00:00:00+01:00",
									date: "2024-04-16T00:00:00+01:00", expectedDate: "2024-04-16T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "Africa/El_Aaiun", dstHandover: "2024-03-11T00:00:00Z",
									date: "2024-03-10T00:00:00+01:00", expectedDate: "2024-03-10T00:00:00+01:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "Africa/El_Aaiun", dstHandover: "2024-03-11T00:00:00Z",
									date: "2024-03-11T00:00:00Z", expectedDate: "2024-03-11T00:00:00Z",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "Africa/El_Aaiun", dstHandover: "2024-03-11T00:00:00Z",
									date: "2024-03-12T00:00:00Z", expectedDate: "2024-03-12T00:00:00Z",
								},
								{
									name: "2024-04-14", testHandler: testYesterdayMidnight,
									zone: "Africa/El_Aaiun", dstHandover: "2024-04-15T00:00:00+01:00",
									date: "2024-04-14T00:00:00Z", expectedDate: "2024-04-14T00:00:00Z",
								},
								{
									name: "2024-04-15", testHandler: testYesterdayMidnight,
									zone: "Africa/El_Aaiun", dstHandover: "2024-04-15T00:00:00+01:00",
									date: "2024-04-15T00:00:00+01:00", expectedDate: "2024-04-15T00:00:00+01:00",
								},
								{
									name: "2024-04-16", testHandler: testYesterdayMidnight,
									zone: "Africa/El_Aaiun", dstHandover: "2024-04-15T00:00:00+01:00",
									date: "2024-04-16T00:00:00+01:00", expectedDate: "2024-04-16T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "Africa/El_Aaiun", dstHandover: "2024-03-11T00:00:00Z",
									date: "2024-03-10T00:00:00+01:00", expectedDate: "2024-03-10T00:00:00+01:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "Africa/El_Aaiun", dstHandover: "2024-03-11T00:00:00Z",
									date: "2024-03-11T00:00:00Z", expectedDate: "2024-03-11T00:00:00Z",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "Africa/El_Aaiun", dstHandover: "2024-03-11T00:00:00Z",
									date: "2024-03-12T00:00:00Z", expectedDate: "2024-03-12T00:00:00Z",
								},
								{
									name: "2024-04-14", testHandler: testTomorrowMidnight,
									zone: "Africa/El_Aaiun", dstHandover: "2024-04-15T00:00:00+01:00",
									date: "2024-04-14T00:00:00Z", expectedDate: "2024-04-14T00:00:00Z",
								},
								{
									name: "2024-04-15", testHandler: testTomorrowMidnight,
									zone: "Africa/El_Aaiun", dstHandover: "2024-04-15T00:00:00+01:00",
									date: "2024-04-15T00:00:00+01:00", expectedDate: "2024-04-15T00:00:00+01:00",
								},
								{
									name: "2024-04-16", testHandler: testTomorrowMidnight,
									zone: "Africa/El_Aaiun", dstHandover: "2024-04-15T00:00:00+01:00",
									date: "2024-04-16T00:00:00+01:00", expectedDate: "2024-04-16T00:00:00+01:00",
								},
							}},
					},
				},
			},
		},
		{
			name: "America",
			children: []testCase{
				{
					name: "Adak",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Adak", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-10T00:00:00-10:00", expectedDate: "2024-03-10T00:00:00-10:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Adak", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-11T00:00:00-09:00", expectedDate: "2024-03-11T00:00:00-09:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Adak", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-12T00:00:00-09:00", expectedDate: "2024-03-12T00:00:00-09:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Adak", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-03T00:00:00-09:00", expectedDate: "2024-11-03T00:00:00-09:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Adak", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-04T00:00:00-10:00", expectedDate: "2024-11-04T00:00:00-10:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Adak", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-05T00:00:00-10:00", expectedDate: "2024-11-05T00:00:00-10:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Adak", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-10T00:00:00-10:00", expectedDate: "2024-03-10T00:00:00-10:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Adak", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-11T00:00:00-09:00", expectedDate: "2024-03-11T00:00:00-09:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Adak", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-12T00:00:00-09:00", expectedDate: "2024-03-12T00:00:00-09:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Adak", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-03T00:00:00-09:00", expectedDate: "2024-11-03T00:00:00-09:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Adak", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-04T00:00:00-10:00", expectedDate: "2024-11-04T00:00:00-10:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Adak", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-05T00:00:00-10:00", expectedDate: "2024-11-05T00:00:00-10:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Adak", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-10T00:00:00-10:00", expectedDate: "2024-03-10T00:00:00-10:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Adak", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-11T00:00:00-09:00", expectedDate: "2024-03-11T00:00:00-09:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Adak", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-12T00:00:00-09:00", expectedDate: "2024-03-12T00:00:00-09:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Adak", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-03T00:00:00-09:00", expectedDate: "2024-11-03T00:00:00-09:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Adak", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-04T00:00:00-10:00", expectedDate: "2024-11-04T00:00:00-10:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Adak", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-05T00:00:00-10:00", expectedDate: "2024-11-05T00:00:00-10:00",
								},
							}},
					},
				},
				{
					name: "Anchorage",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Anchorage", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-10T00:00:00-09:00", expectedDate: "2024-03-10T00:00:00-09:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Anchorage", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-11T00:00:00-08:00", expectedDate: "2024-03-11T00:00:00-08:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Anchorage", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-12T00:00:00-08:00", expectedDate: "2024-03-12T00:00:00-08:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Anchorage", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-03T00:00:00-08:00", expectedDate: "2024-11-03T00:00:00-08:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Anchorage", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-04T00:00:00-09:00", expectedDate: "2024-11-04T00:00:00-09:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Anchorage", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-05T00:00:00-09:00", expectedDate: "2024-11-05T00:00:00-09:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Anchorage", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-10T00:00:00-09:00", expectedDate: "2024-03-10T00:00:00-09:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Anchorage", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-11T00:00:00-08:00", expectedDate: "2024-03-11T00:00:00-08:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Anchorage", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-12T00:00:00-08:00", expectedDate: "2024-03-12T00:00:00-08:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Anchorage", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-03T00:00:00-08:00", expectedDate: "2024-11-03T00:00:00-08:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Anchorage", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-04T00:00:00-09:00", expectedDate: "2024-11-04T00:00:00-09:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Anchorage", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-05T00:00:00-09:00", expectedDate: "2024-11-05T00:00:00-09:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Anchorage", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-10T00:00:00-09:00", expectedDate: "2024-03-10T00:00:00-09:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Anchorage", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-11T00:00:00-08:00", expectedDate: "2024-03-11T00:00:00-08:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Anchorage", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-12T00:00:00-08:00", expectedDate: "2024-03-12T00:00:00-08:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Anchorage", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-03T00:00:00-08:00", expectedDate: "2024-11-03T00:00:00-08:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Anchorage", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-04T00:00:00-09:00", expectedDate: "2024-11-04T00:00:00-09:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Anchorage", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-05T00:00:00-09:00", expectedDate: "2024-11-05T00:00:00-09:00",
								},
							}},
					},
				},
				{
					name: "Asuncion",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-10-05", testHandler: testLocalMidnight,
									zone: "America/Asuncion", dstHandover: "2024-10-06T23:00:00-03:00",
									date: "2024-10-05T23:00:00-04:00", expectedDate: "2024-10-05T23:00:00-04:00",
								},
								{
									name: "2024-10-06", testHandler: testLocalMidnight,
									zone: "America/Asuncion", dstHandover: "2024-10-06T23:00:00-03:00",
									date: "2024-10-06T23:00:00-03:00", expectedDate: "2024-10-06T23:00:00-03:00",
								},
								{
									name: "2024-10-07", testHandler: testLocalMidnight,
									zone: "America/Asuncion", dstHandover: "2024-10-06T23:00:00-03:00",
									date: "2024-10-07T23:00:00-03:00", expectedDate: "2024-10-07T23:00:00-03:00",
								},
								{
									name: "2024-03-23", testHandler: testLocalMidnight,
									zone: "America/Asuncion", dstHandover: "2024-03-24T00:00:00-04:00",
									date: "2024-03-23T00:00:00-03:00", expectedDate: "2024-03-23T00:00:00-03:00",
								},
								{
									name: "2024-03-24", testHandler: testLocalMidnight,
									zone: "America/Asuncion", dstHandover: "2024-03-24T00:00:00-04:00",
									date: "2024-03-24T00:00:00-04:00", expectedDate: "2024-03-24T00:00:00-04:00",
								},
								{
									name: "2024-03-25", testHandler: testLocalMidnight,
									zone: "America/Asuncion", dstHandover: "2024-03-24T00:00:00-04:00",
									date: "2024-03-25T00:00:00-04:00", expectedDate: "2024-03-25T00:00:00-04:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-10-05", testHandler: testYesterdayMidnight,
									zone: "America/Asuncion", dstHandover: "2024-10-06T23:00:00-03:00",
									date: "2024-10-05T23:00:00-04:00", expectedDate: "2024-10-05T23:00:00-04:00",
								},
								{
									name: "2024-10-06", testHandler: testYesterdayMidnight,
									zone: "America/Asuncion", dstHandover: "2024-10-06T23:00:00-03:00",
									date: "2024-10-06T23:00:00-03:00", expectedDate: "2024-10-06T23:00:00-03:00",
								},
								{
									name: "2024-10-07", testHandler: testYesterdayMidnight,
									zone: "America/Asuncion", dstHandover: "2024-10-06T23:00:00-03:00",
									date: "2024-10-07T23:00:00-03:00", expectedDate: "2024-10-07T23:00:00-03:00",
								},
								{
									name: "2024-03-23", testHandler: testYesterdayMidnight,
									zone: "America/Asuncion", dstHandover: "2024-03-24T00:00:00-04:00",
									date: "2024-03-23T00:00:00-03:00", expectedDate: "2024-03-23T00:00:00-03:00",
								},
								{
									name: "2024-03-24", testHandler: testYesterdayMidnight,
									zone: "America/Asuncion", dstHandover: "2024-03-24T00:00:00-04:00",
									date: "2024-03-24T00:00:00-04:00", expectedDate: "2024-03-24T00:00:00-04:00",
								},
								{
									name: "2024-03-25", testHandler: testYesterdayMidnight,
									zone: "America/Asuncion", dstHandover: "2024-03-24T00:00:00-04:00",
									date: "2024-03-25T00:00:00-04:00", expectedDate: "2024-03-25T00:00:00-04:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-10-05", testHandler: testTomorrowMidnight,
									zone: "America/Asuncion", dstHandover: "2024-10-06T23:00:00-03:00",
									date: "2024-10-05T23:00:00-04:00", expectedDate: "2024-10-05T23:00:00-04:00",
								},
								{
									name: "2024-10-06", testHandler: testTomorrowMidnight,
									zone: "America/Asuncion", dstHandover: "2024-10-06T23:00:00-03:00",
									date: "2024-10-06T23:00:00-03:00", expectedDate: "2024-10-06T23:00:00-03:00",
								},
								{
									name: "2024-10-07", testHandler: testTomorrowMidnight,
									zone: "America/Asuncion", dstHandover: "2024-10-06T23:00:00-03:00",
									date: "2024-10-07T23:00:00-03:00", expectedDate: "2024-10-07T23:00:00-03:00",
								},
								{
									name: "2024-03-23", testHandler: testTomorrowMidnight,
									zone: "America/Asuncion", dstHandover: "2024-03-24T00:00:00-04:00",
									date: "2024-03-23T00:00:00-03:00", expectedDate: "2024-03-23T00:00:00-03:00",
								},
								{
									name: "2024-03-24", testHandler: testTomorrowMidnight,
									zone: "America/Asuncion", dstHandover: "2024-03-24T00:00:00-04:00",
									date: "2024-03-24T00:00:00-04:00", expectedDate: "2024-03-24T00:00:00-04:00",
								},
								{
									name: "2024-03-25", testHandler: testTomorrowMidnight,
									zone: "America/Asuncion", dstHandover: "2024-03-24T00:00:00-04:00",
									date: "2024-03-25T00:00:00-04:00", expectedDate: "2024-03-25T00:00:00-04:00",
								},
							}},
					},
				},
				{
					name: "Atka",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Atka", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-10T00:00:00-10:00", expectedDate: "2024-03-10T00:00:00-10:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Atka", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-11T00:00:00-09:00", expectedDate: "2024-03-11T00:00:00-09:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Atka", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-12T00:00:00-09:00", expectedDate: "2024-03-12T00:00:00-09:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Atka", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-03T00:00:00-09:00", expectedDate: "2024-11-03T00:00:00-09:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Atka", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-04T00:00:00-10:00", expectedDate: "2024-11-04T00:00:00-10:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Atka", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-05T00:00:00-10:00", expectedDate: "2024-11-05T00:00:00-10:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Atka", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-10T00:00:00-10:00", expectedDate: "2024-03-10T00:00:00-10:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Atka", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-11T00:00:00-09:00", expectedDate: "2024-03-11T00:00:00-09:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Atka", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-12T00:00:00-09:00", expectedDate: "2024-03-12T00:00:00-09:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Atka", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-03T00:00:00-09:00", expectedDate: "2024-11-03T00:00:00-09:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Atka", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-04T00:00:00-10:00", expectedDate: "2024-11-04T00:00:00-10:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Atka", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-05T00:00:00-10:00", expectedDate: "2024-11-05T00:00:00-10:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Atka", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-10T00:00:00-10:00", expectedDate: "2024-03-10T00:00:00-10:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Atka", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-11T00:00:00-09:00", expectedDate: "2024-03-11T00:00:00-09:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Atka", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-12T00:00:00-09:00", expectedDate: "2024-03-12T00:00:00-09:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Atka", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-03T00:00:00-09:00", expectedDate: "2024-11-03T00:00:00-09:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Atka", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-04T00:00:00-10:00", expectedDate: "2024-11-04T00:00:00-10:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Atka", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-05T00:00:00-10:00", expectedDate: "2024-11-05T00:00:00-10:00",
								},
							}},
					},
				},
				{
					name: "Boise",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Boise", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Boise", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Boise", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Boise", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Boise", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Boise", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Boise", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Boise", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Boise", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Boise", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Boise", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Boise", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Boise", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Boise", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Boise", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Boise", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Boise", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Boise", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
					},
				},
				{
					name: "Cambridge_Bay",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Cambridge_Bay", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Cambridge_Bay", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Cambridge_Bay", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Cambridge_Bay", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Cambridge_Bay", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Cambridge_Bay", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Cambridge_Bay", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Cambridge_Bay", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Cambridge_Bay", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Cambridge_Bay", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Cambridge_Bay", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Cambridge_Bay", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Cambridge_Bay", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Cambridge_Bay", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Cambridge_Bay", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Cambridge_Bay", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Cambridge_Bay", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Cambridge_Bay", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
					},
				},
				{
					name: "Chicago",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Chicago", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Chicago", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Chicago", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Chicago", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Chicago", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Chicago", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Chicago", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Chicago", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Chicago", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Chicago", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Chicago", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Chicago", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Chicago", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Chicago", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Chicago", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Chicago", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Chicago", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Chicago", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
					},
				},
				{
					name: "Ciudad_Juarez",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Ciudad_Juarez", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Ciudad_Juarez", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Ciudad_Juarez", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Ciudad_Juarez", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Ciudad_Juarez", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Ciudad_Juarez", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Ciudad_Juarez", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Ciudad_Juarez", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Ciudad_Juarez", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Ciudad_Juarez", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Ciudad_Juarez", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Ciudad_Juarez", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Ciudad_Juarez", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Ciudad_Juarez", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Ciudad_Juarez", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Ciudad_Juarez", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Ciudad_Juarez", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Ciudad_Juarez", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
					},
				},
				{
					name: "Denver",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Denver", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Denver", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Denver", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Denver", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Denver", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Denver", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Denver", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Denver", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Denver", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Denver", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Denver", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Denver", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Denver", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Denver", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Denver", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Denver", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Denver", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Denver", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
					},
				},
				{
					name: "Detroit",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Detroit", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Detroit", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Detroit", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Detroit", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Detroit", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Detroit", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Detroit", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Detroit", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Detroit", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Detroit", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Detroit", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Detroit", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Detroit", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Detroit", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Detroit", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Detroit", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Detroit", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Detroit", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
					},
				},
				{
					name: "Edmonton",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Edmonton", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Edmonton", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Edmonton", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Edmonton", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Edmonton", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Edmonton", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Edmonton", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Edmonton", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Edmonton", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Edmonton", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Edmonton", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Edmonton", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Edmonton", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Edmonton", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Edmonton", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Edmonton", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Edmonton", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Edmonton", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
					},
				},
				{
					name: "Ensenada",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Ensenada", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Ensenada", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Ensenada", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Ensenada", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Ensenada", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Ensenada", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Ensenada", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Ensenada", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Ensenada", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Ensenada", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Ensenada", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Ensenada", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Ensenada", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Ensenada", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Ensenada", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Ensenada", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Ensenada", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Ensenada", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
					},
				},
				{
					name: "Fort_Wayne",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Fort_Wayne", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Fort_Wayne", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Fort_Wayne", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Fort_Wayne", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Fort_Wayne", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Fort_Wayne", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Fort_Wayne", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Fort_Wayne", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Fort_Wayne", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Fort_Wayne", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Fort_Wayne", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Fort_Wayne", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Fort_Wayne", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Fort_Wayne", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Fort_Wayne", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Fort_Wayne", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Fort_Wayne", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Fort_Wayne", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
					},
				},
				{
					name: "Glace_Bay",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Glace_Bay", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-10T00:00:00-04:00", expectedDate: "2024-03-10T00:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Glace_Bay", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-11T00:00:00-03:00", expectedDate: "2024-03-11T00:00:00-03:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Glace_Bay", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-12T00:00:00-03:00", expectedDate: "2024-03-12T00:00:00-03:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Glace_Bay", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-03T00:00:00-03:00", expectedDate: "2024-11-03T00:00:00-03:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Glace_Bay", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-04T00:00:00-04:00", expectedDate: "2024-11-04T00:00:00-04:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Glace_Bay", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-05T00:00:00-04:00", expectedDate: "2024-11-05T00:00:00-04:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Glace_Bay", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-10T00:00:00-04:00", expectedDate: "2024-03-10T00:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Glace_Bay", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-11T00:00:00-03:00", expectedDate: "2024-03-11T00:00:00-03:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Glace_Bay", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-12T00:00:00-03:00", expectedDate: "2024-03-12T00:00:00-03:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Glace_Bay", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-03T00:00:00-03:00", expectedDate: "2024-11-03T00:00:00-03:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Glace_Bay", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-04T00:00:00-04:00", expectedDate: "2024-11-04T00:00:00-04:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Glace_Bay", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-05T00:00:00-04:00", expectedDate: "2024-11-05T00:00:00-04:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Glace_Bay", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-10T00:00:00-04:00", expectedDate: "2024-03-10T00:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Glace_Bay", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-11T00:00:00-03:00", expectedDate: "2024-03-11T00:00:00-03:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Glace_Bay", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-12T00:00:00-03:00", expectedDate: "2024-03-12T00:00:00-03:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Glace_Bay", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-03T00:00:00-03:00", expectedDate: "2024-11-03T00:00:00-03:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Glace_Bay", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-04T00:00:00-04:00", expectedDate: "2024-11-04T00:00:00-04:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Glace_Bay", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-05T00:00:00-04:00", expectedDate: "2024-11-05T00:00:00-04:00",
								},
							}},
					},
				},
				{
					name: "Godthab",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-30", testHandler: testLocalMidnight,
									zone: "America/Godthab", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-03-30T00:00:00-02:00", expectedDate: "2024-03-30T00:00:00-02:00",
								},
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "America/Godthab", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-03-31T00:00:00-01:00", expectedDate: "2024-03-31T00:00:00-01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "America/Godthab", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-04-01T00:00:00-01:00", expectedDate: "2024-04-01T00:00:00-01:00",
								},
								{
									name: "2024-10-26", testHandler: testLocalMidnight,
									zone: "America/Godthab", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-26T00:00:00-01:00", expectedDate: "2024-10-26T00:00:00-01:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "America/Godthab", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-27T00:00:00-02:00", expectedDate: "2024-10-27T00:00:00-02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "America/Godthab", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-28T00:00:00-02:00", expectedDate: "2024-10-28T00:00:00-02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-30", testHandler: testYesterdayMidnight,
									zone: "America/Godthab", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-03-30T00:00:00-02:00", expectedDate: "2024-03-30T00:00:00-02:00",
								},
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "America/Godthab", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-03-31T00:00:00-01:00", expectedDate: "2024-03-31T00:00:00-01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "America/Godthab", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-04-01T00:00:00-01:00", expectedDate: "2024-04-01T00:00:00-01:00",
								},
								{
									name: "2024-10-26", testHandler: testYesterdayMidnight,
									zone: "America/Godthab", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-26T00:00:00-01:00", expectedDate: "2024-10-26T00:00:00-01:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "America/Godthab", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-27T00:00:00-02:00", expectedDate: "2024-10-27T00:00:00-02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "America/Godthab", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-28T00:00:00-02:00", expectedDate: "2024-10-28T00:00:00-02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-30", testHandler: testTomorrowMidnight,
									zone: "America/Godthab", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-03-30T00:00:00-02:00", expectedDate: "2024-03-30T00:00:00-02:00",
								},
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "America/Godthab", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-03-31T00:00:00-01:00", expectedDate: "2024-03-31T00:00:00-01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "America/Godthab", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-04-01T00:00:00-01:00", expectedDate: "2024-04-01T00:00:00-01:00",
								},
								{
									name: "2024-10-26", testHandler: testTomorrowMidnight,
									zone: "America/Godthab", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-26T00:00:00-01:00", expectedDate: "2024-10-26T00:00:00-01:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "America/Godthab", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-27T00:00:00-02:00", expectedDate: "2024-10-27T00:00:00-02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "America/Godthab", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-28T00:00:00-02:00", expectedDate: "2024-10-28T00:00:00-02:00",
								},
							}},
					},
				},
				{
					name: "Goose_Bay",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Goose_Bay", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-10T00:00:00-04:00", expectedDate: "2024-03-10T00:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Goose_Bay", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-11T00:00:00-03:00", expectedDate: "2024-03-11T00:00:00-03:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Goose_Bay", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-12T00:00:00-03:00", expectedDate: "2024-03-12T00:00:00-03:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Goose_Bay", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-03T00:00:00-03:00", expectedDate: "2024-11-03T00:00:00-03:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Goose_Bay", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-04T00:00:00-04:00", expectedDate: "2024-11-04T00:00:00-04:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Goose_Bay", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-05T00:00:00-04:00", expectedDate: "2024-11-05T00:00:00-04:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Goose_Bay", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-10T00:00:00-04:00", expectedDate: "2024-03-10T00:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Goose_Bay", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-11T00:00:00-03:00", expectedDate: "2024-03-11T00:00:00-03:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Goose_Bay", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-12T00:00:00-03:00", expectedDate: "2024-03-12T00:00:00-03:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Goose_Bay", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-03T00:00:00-03:00", expectedDate: "2024-11-03T00:00:00-03:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Goose_Bay", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-04T00:00:00-04:00", expectedDate: "2024-11-04T00:00:00-04:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Goose_Bay", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-05T00:00:00-04:00", expectedDate: "2024-11-05T00:00:00-04:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Goose_Bay", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-10T00:00:00-04:00", expectedDate: "2024-03-10T00:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Goose_Bay", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-11T00:00:00-03:00", expectedDate: "2024-03-11T00:00:00-03:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Goose_Bay", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-12T00:00:00-03:00", expectedDate: "2024-03-12T00:00:00-03:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Goose_Bay", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-03T00:00:00-03:00", expectedDate: "2024-11-03T00:00:00-03:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Goose_Bay", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-04T00:00:00-04:00", expectedDate: "2024-11-04T00:00:00-04:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Goose_Bay", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-05T00:00:00-04:00", expectedDate: "2024-11-05T00:00:00-04:00",
								},
							}},
					},
				},
				{
					name: "Grand_Turk",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Grand_Turk", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Grand_Turk", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Grand_Turk", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Grand_Turk", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Grand_Turk", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Grand_Turk", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Grand_Turk", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Grand_Turk", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Grand_Turk", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Grand_Turk", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Grand_Turk", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Grand_Turk", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Grand_Turk", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Grand_Turk", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Grand_Turk", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Grand_Turk", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Grand_Turk", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Grand_Turk", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
					},
				},
				{
					name: "Halifax",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Halifax", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-10T00:00:00-04:00", expectedDate: "2024-03-10T00:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Halifax", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-11T00:00:00-03:00", expectedDate: "2024-03-11T00:00:00-03:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Halifax", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-12T00:00:00-03:00", expectedDate: "2024-03-12T00:00:00-03:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Halifax", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-03T00:00:00-03:00", expectedDate: "2024-11-03T00:00:00-03:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Halifax", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-04T00:00:00-04:00", expectedDate: "2024-11-04T00:00:00-04:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Halifax", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-05T00:00:00-04:00", expectedDate: "2024-11-05T00:00:00-04:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Halifax", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-10T00:00:00-04:00", expectedDate: "2024-03-10T00:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Halifax", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-11T00:00:00-03:00", expectedDate: "2024-03-11T00:00:00-03:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Halifax", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-12T00:00:00-03:00", expectedDate: "2024-03-12T00:00:00-03:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Halifax", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-03T00:00:00-03:00", expectedDate: "2024-11-03T00:00:00-03:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Halifax", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-04T00:00:00-04:00", expectedDate: "2024-11-04T00:00:00-04:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Halifax", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-05T00:00:00-04:00", expectedDate: "2024-11-05T00:00:00-04:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Halifax", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-10T00:00:00-04:00", expectedDate: "2024-03-10T00:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Halifax", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-11T00:00:00-03:00", expectedDate: "2024-03-11T00:00:00-03:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Halifax", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-12T00:00:00-03:00", expectedDate: "2024-03-12T00:00:00-03:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Halifax", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-03T00:00:00-03:00", expectedDate: "2024-11-03T00:00:00-03:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Halifax", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-04T00:00:00-04:00", expectedDate: "2024-11-04T00:00:00-04:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Halifax", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-05T00:00:00-04:00", expectedDate: "2024-11-05T00:00:00-04:00",
								},
							}},
					},
				},
				{
					name: "Havana",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-09", testHandler: testLocalMidnight,
									zone: "America/Havana", dstHandover: "2024-03-10T23:00:00-04:00",
									date: "2024-03-09T23:00:00-05:00", expectedDate: "2024-03-09T23:00:00-05:00",
								},
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Havana", dstHandover: "2024-03-10T23:00:00-04:00",
									date: "2024-03-10T23:00:00-04:00", expectedDate: "2024-03-10T23:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Havana", dstHandover: "2024-03-10T23:00:00-04:00",
									date: "2024-03-11T23:00:00-04:00", expectedDate: "2024-03-11T23:00:00-04:00",
								},
								{
									name: "2024-11-02", testHandler: testLocalMidnight,
									zone: "America/Havana", dstHandover: "2024-11-03T23:00:00-05:00",
									date: "2024-11-02T23:00:00-04:00", expectedDate: "2024-11-02T23:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Havana", dstHandover: "2024-11-03T23:00:00-05:00",
									date: "2024-11-03T23:00:00-05:00", expectedDate: "2024-11-03T23:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Havana", dstHandover: "2024-11-03T23:00:00-05:00",
									date: "2024-11-04T23:00:00-05:00", expectedDate: "2024-11-04T23:00:00-05:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-09", testHandler: testYesterdayMidnight,
									zone: "America/Havana", dstHandover: "2024-03-10T23:00:00-04:00",
									date: "2024-03-09T23:00:00-05:00", expectedDate: "2024-03-09T23:00:00-05:00",
								},
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Havana", dstHandover: "2024-03-10T23:00:00-04:00",
									date: "2024-03-10T23:00:00-04:00", expectedDate: "2024-03-10T23:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Havana", dstHandover: "2024-03-10T23:00:00-04:00",
									date: "2024-03-11T23:00:00-04:00", expectedDate: "2024-03-11T23:00:00-04:00",
								},
								{
									name: "2024-11-02", testHandler: testYesterdayMidnight,
									zone: "America/Havana", dstHandover: "2024-11-03T23:00:00-05:00",
									date: "2024-11-02T23:00:00-04:00", expectedDate: "2024-11-02T23:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Havana", dstHandover: "2024-11-03T23:00:00-05:00",
									date: "2024-11-03T23:00:00-05:00", expectedDate: "2024-11-03T23:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Havana", dstHandover: "2024-11-03T23:00:00-05:00",
									date: "2024-11-04T23:00:00-05:00", expectedDate: "2024-11-04T23:00:00-05:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-09", testHandler: testTomorrowMidnight,
									zone: "America/Havana", dstHandover: "2024-03-10T23:00:00-04:00",
									date: "2024-03-09T23:00:00-05:00", expectedDate: "2024-03-09T23:00:00-05:00",
								},
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Havana", dstHandover: "2024-03-10T23:00:00-04:00",
									date: "2024-03-10T23:00:00-04:00", expectedDate: "2024-03-10T23:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Havana", dstHandover: "2024-03-10T23:00:00-04:00",
									date: "2024-03-11T23:00:00-04:00", expectedDate: "2024-03-11T23:00:00-04:00",
								},
								{
									name: "2024-11-02", testHandler: testTomorrowMidnight,
									zone: "America/Havana", dstHandover: "2024-11-03T23:00:00-05:00",
									date: "2024-11-02T23:00:00-04:00", expectedDate: "2024-11-02T23:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Havana", dstHandover: "2024-11-03T23:00:00-05:00",
									date: "2024-11-03T23:00:00-05:00", expectedDate: "2024-11-03T23:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Havana", dstHandover: "2024-11-03T23:00:00-05:00",
									date: "2024-11-04T23:00:00-05:00", expectedDate: "2024-11-04T23:00:00-05:00",
								},
							}},
					},
				},
				{
					name: "Indiana",
					children: []testCase{
						{
							name: "Indianapolis",
							children: []testCase{
								{
									name: "LocalMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testLocalMidnight,
											zone: "America/Indiana/Indianapolis", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testLocalMidnight,
											zone: "America/Indiana/Indianapolis", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testLocalMidnight,
											zone: "America/Indiana/Indianapolis", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testLocalMidnight,
											zone: "America/Indiana/Indianapolis", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testLocalMidnight,
											zone: "America/Indiana/Indianapolis", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testLocalMidnight,
											zone: "America/Indiana/Indianapolis", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
								{
									name: "YesterdayMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Indianapolis", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Indianapolis", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Indianapolis", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Indianapolis", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Indianapolis", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Indianapolis", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
								{
									name: "TomorrowMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Indianapolis", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Indianapolis", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Indianapolis", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Indianapolis", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Indianapolis", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Indianapolis", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
							},
						},
						{
							name: "Knox",
							children: []testCase{
								{
									name: "LocalMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testLocalMidnight,
											zone: "America/Indiana/Knox", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
										},
										{
											name: "2024-03-11", testHandler: testLocalMidnight,
											zone: "America/Indiana/Knox", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
										},
										{
											name: "2024-03-12", testHandler: testLocalMidnight,
											zone: "America/Indiana/Knox", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
										},
										{
											name: "2024-11-03", testHandler: testLocalMidnight,
											zone: "America/Indiana/Knox", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
										},
										{
											name: "2024-11-04", testHandler: testLocalMidnight,
											zone: "America/Indiana/Knox", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
										},
										{
											name: "2024-11-05", testHandler: testLocalMidnight,
											zone: "America/Indiana/Knox", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
										},
									}},
								{
									name: "YesterdayMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Knox", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
										},
										{
											name: "2024-03-11", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Knox", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
										},
										{
											name: "2024-03-12", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Knox", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
										},
										{
											name: "2024-11-03", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Knox", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
										},
										{
											name: "2024-11-04", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Knox", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
										},
										{
											name: "2024-11-05", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Knox", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
										},
									}},
								{
									name: "TomorrowMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Knox", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
										},
										{
											name: "2024-03-11", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Knox", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
										},
										{
											name: "2024-03-12", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Knox", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
										},
										{
											name: "2024-11-03", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Knox", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
										},
										{
											name: "2024-11-04", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Knox", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
										},
										{
											name: "2024-11-05", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Knox", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
										},
									}},
							},
						},
						{
							name: "Marengo",
							children: []testCase{
								{
									name: "LocalMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testLocalMidnight,
											zone: "America/Indiana/Marengo", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testLocalMidnight,
											zone: "America/Indiana/Marengo", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testLocalMidnight,
											zone: "America/Indiana/Marengo", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testLocalMidnight,
											zone: "America/Indiana/Marengo", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testLocalMidnight,
											zone: "America/Indiana/Marengo", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testLocalMidnight,
											zone: "America/Indiana/Marengo", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
								{
									name: "YesterdayMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Marengo", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Marengo", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Marengo", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Marengo", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Marengo", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Marengo", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
								{
									name: "TomorrowMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Marengo", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Marengo", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Marengo", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Marengo", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Marengo", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Marengo", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
							},
						},
						{
							name: "Petersburg",
							children: []testCase{
								{
									name: "LocalMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testLocalMidnight,
											zone: "America/Indiana/Petersburg", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testLocalMidnight,
											zone: "America/Indiana/Petersburg", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testLocalMidnight,
											zone: "America/Indiana/Petersburg", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testLocalMidnight,
											zone: "America/Indiana/Petersburg", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testLocalMidnight,
											zone: "America/Indiana/Petersburg", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testLocalMidnight,
											zone: "America/Indiana/Petersburg", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
								{
									name: "YesterdayMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Petersburg", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Petersburg", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Petersburg", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Petersburg", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Petersburg", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Petersburg", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
								{
									name: "TomorrowMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Petersburg", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Petersburg", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Petersburg", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Petersburg", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Petersburg", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Petersburg", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
							},
						},
						{
							name: "Tell_City",
							children: []testCase{
								{
									name: "LocalMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testLocalMidnight,
											zone: "America/Indiana/Tell_City", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
										},
										{
											name: "2024-03-11", testHandler: testLocalMidnight,
											zone: "America/Indiana/Tell_City", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
										},
										{
											name: "2024-03-12", testHandler: testLocalMidnight,
											zone: "America/Indiana/Tell_City", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
										},
										{
											name: "2024-11-03", testHandler: testLocalMidnight,
											zone: "America/Indiana/Tell_City", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
										},
										{
											name: "2024-11-04", testHandler: testLocalMidnight,
											zone: "America/Indiana/Tell_City", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
										},
										{
											name: "2024-11-05", testHandler: testLocalMidnight,
											zone: "America/Indiana/Tell_City", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
										},
									}},
								{
									name: "YesterdayMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Tell_City", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
										},
										{
											name: "2024-03-11", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Tell_City", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
										},
										{
											name: "2024-03-12", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Tell_City", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
										},
										{
											name: "2024-11-03", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Tell_City", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
										},
										{
											name: "2024-11-04", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Tell_City", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
										},
										{
											name: "2024-11-05", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Tell_City", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
										},
									}},
								{
									name: "TomorrowMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Tell_City", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
										},
										{
											name: "2024-03-11", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Tell_City", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
										},
										{
											name: "2024-03-12", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Tell_City", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
										},
										{
											name: "2024-11-03", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Tell_City", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
										},
										{
											name: "2024-11-04", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Tell_City", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
										},
										{
											name: "2024-11-05", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Tell_City", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
										},
									}},
							},
						},
						{
							name: "Vevay",
							children: []testCase{
								{
									name: "LocalMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testLocalMidnight,
											zone: "America/Indiana/Vevay", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testLocalMidnight,
											zone: "America/Indiana/Vevay", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testLocalMidnight,
											zone: "America/Indiana/Vevay", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testLocalMidnight,
											zone: "America/Indiana/Vevay", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testLocalMidnight,
											zone: "America/Indiana/Vevay", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testLocalMidnight,
											zone: "America/Indiana/Vevay", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
								{
									name: "YesterdayMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Vevay", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Vevay", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Vevay", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Vevay", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Vevay", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Vevay", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
								{
									name: "TomorrowMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Vevay", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Vevay", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Vevay", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Vevay", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Vevay", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Vevay", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
							},
						},
						{
							name: "Vincennes",
							children: []testCase{
								{
									name: "LocalMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testLocalMidnight,
											zone: "America/Indiana/Vincennes", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testLocalMidnight,
											zone: "America/Indiana/Vincennes", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testLocalMidnight,
											zone: "America/Indiana/Vincennes", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testLocalMidnight,
											zone: "America/Indiana/Vincennes", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testLocalMidnight,
											zone: "America/Indiana/Vincennes", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testLocalMidnight,
											zone: "America/Indiana/Vincennes", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
								{
									name: "YesterdayMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Vincennes", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Vincennes", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Vincennes", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Vincennes", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Vincennes", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Vincennes", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
								{
									name: "TomorrowMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Vincennes", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Vincennes", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Vincennes", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Vincennes", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Vincennes", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Vincennes", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
							},
						},
						{
							name: "Winamac",
							children: []testCase{
								{
									name: "LocalMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testLocalMidnight,
											zone: "America/Indiana/Winamac", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testLocalMidnight,
											zone: "America/Indiana/Winamac", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testLocalMidnight,
											zone: "America/Indiana/Winamac", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testLocalMidnight,
											zone: "America/Indiana/Winamac", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testLocalMidnight,
											zone: "America/Indiana/Winamac", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testLocalMidnight,
											zone: "America/Indiana/Winamac", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
								{
									name: "YesterdayMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Winamac", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Winamac", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Winamac", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Winamac", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Winamac", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testYesterdayMidnight,
											zone: "America/Indiana/Winamac", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
								{
									name: "TomorrowMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Winamac", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Winamac", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Winamac", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Winamac", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Winamac", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testTomorrowMidnight,
											zone: "America/Indiana/Winamac", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
							},
						},
					},
				},
				{
					name: "Indianapolis",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Indianapolis", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Indianapolis", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Indianapolis", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Indianapolis", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Indianapolis", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Indianapolis", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Indianapolis", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Indianapolis", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Indianapolis", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Indianapolis", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Indianapolis", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Indianapolis", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Indianapolis", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Indianapolis", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Indianapolis", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Indianapolis", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Indianapolis", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Indianapolis", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
					},
				},
				{
					name: "Inuvik",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Inuvik", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Inuvik", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Inuvik", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Inuvik", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Inuvik", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Inuvik", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Inuvik", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Inuvik", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Inuvik", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Inuvik", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Inuvik", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Inuvik", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Inuvik", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Inuvik", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Inuvik", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Inuvik", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Inuvik", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Inuvik", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
					},
				},
				{
					name: "Iqaluit",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Iqaluit", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Iqaluit", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Iqaluit", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Iqaluit", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Iqaluit", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Iqaluit", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Iqaluit", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Iqaluit", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Iqaluit", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Iqaluit", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Iqaluit", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Iqaluit", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Iqaluit", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Iqaluit", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Iqaluit", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Iqaluit", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Iqaluit", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Iqaluit", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
					},
				},
				{
					name: "Juneau",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Juneau", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-10T00:00:00-09:00", expectedDate: "2024-03-10T00:00:00-09:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Juneau", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-11T00:00:00-08:00", expectedDate: "2024-03-11T00:00:00-08:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Juneau", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-12T00:00:00-08:00", expectedDate: "2024-03-12T00:00:00-08:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Juneau", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-03T00:00:00-08:00", expectedDate: "2024-11-03T00:00:00-08:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Juneau", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-04T00:00:00-09:00", expectedDate: "2024-11-04T00:00:00-09:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Juneau", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-05T00:00:00-09:00", expectedDate: "2024-11-05T00:00:00-09:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Juneau", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-10T00:00:00-09:00", expectedDate: "2024-03-10T00:00:00-09:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Juneau", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-11T00:00:00-08:00", expectedDate: "2024-03-11T00:00:00-08:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Juneau", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-12T00:00:00-08:00", expectedDate: "2024-03-12T00:00:00-08:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Juneau", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-03T00:00:00-08:00", expectedDate: "2024-11-03T00:00:00-08:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Juneau", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-04T00:00:00-09:00", expectedDate: "2024-11-04T00:00:00-09:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Juneau", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-05T00:00:00-09:00", expectedDate: "2024-11-05T00:00:00-09:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Juneau", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-10T00:00:00-09:00", expectedDate: "2024-03-10T00:00:00-09:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Juneau", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-11T00:00:00-08:00", expectedDate: "2024-03-11T00:00:00-08:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Juneau", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-12T00:00:00-08:00", expectedDate: "2024-03-12T00:00:00-08:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Juneau", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-03T00:00:00-08:00", expectedDate: "2024-11-03T00:00:00-08:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Juneau", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-04T00:00:00-09:00", expectedDate: "2024-11-04T00:00:00-09:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Juneau", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-05T00:00:00-09:00", expectedDate: "2024-11-05T00:00:00-09:00",
								},
							}},
					},
				},
				{
					name: "Kentucky",
					children: []testCase{
						{
							name: "Louisville",
							children: []testCase{
								{
									name: "LocalMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testLocalMidnight,
											zone: "America/Kentucky/Louisville", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testLocalMidnight,
											zone: "America/Kentucky/Louisville", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testLocalMidnight,
											zone: "America/Kentucky/Louisville", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testLocalMidnight,
											zone: "America/Kentucky/Louisville", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testLocalMidnight,
											zone: "America/Kentucky/Louisville", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testLocalMidnight,
											zone: "America/Kentucky/Louisville", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
								{
									name: "YesterdayMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testYesterdayMidnight,
											zone: "America/Kentucky/Louisville", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testYesterdayMidnight,
											zone: "America/Kentucky/Louisville", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testYesterdayMidnight,
											zone: "America/Kentucky/Louisville", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testYesterdayMidnight,
											zone: "America/Kentucky/Louisville", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testYesterdayMidnight,
											zone: "America/Kentucky/Louisville", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testYesterdayMidnight,
											zone: "America/Kentucky/Louisville", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
								{
									name: "TomorrowMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testTomorrowMidnight,
											zone: "America/Kentucky/Louisville", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testTomorrowMidnight,
											zone: "America/Kentucky/Louisville", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testTomorrowMidnight,
											zone: "America/Kentucky/Louisville", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testTomorrowMidnight,
											zone: "America/Kentucky/Louisville", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testTomorrowMidnight,
											zone: "America/Kentucky/Louisville", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testTomorrowMidnight,
											zone: "America/Kentucky/Louisville", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
							},
						},
						{
							name: "Monticello",
							children: []testCase{
								{
									name: "LocalMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testLocalMidnight,
											zone: "America/Kentucky/Monticello", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testLocalMidnight,
											zone: "America/Kentucky/Monticello", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testLocalMidnight,
											zone: "America/Kentucky/Monticello", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testLocalMidnight,
											zone: "America/Kentucky/Monticello", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testLocalMidnight,
											zone: "America/Kentucky/Monticello", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testLocalMidnight,
											zone: "America/Kentucky/Monticello", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
								{
									name: "YesterdayMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testYesterdayMidnight,
											zone: "America/Kentucky/Monticello", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testYesterdayMidnight,
											zone: "America/Kentucky/Monticello", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testYesterdayMidnight,
											zone: "America/Kentucky/Monticello", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testYesterdayMidnight,
											zone: "America/Kentucky/Monticello", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testYesterdayMidnight,
											zone: "America/Kentucky/Monticello", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testYesterdayMidnight,
											zone: "America/Kentucky/Monticello", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
								{
									name: "TomorrowMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testTomorrowMidnight,
											zone: "America/Kentucky/Monticello", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
										},
										{
											name: "2024-03-11", testHandler: testTomorrowMidnight,
											zone: "America/Kentucky/Monticello", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
										},
										{
											name: "2024-03-12", testHandler: testTomorrowMidnight,
											zone: "America/Kentucky/Monticello", dstHandover: "2024-03-11T00:00:00-04:00",
											date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
										},
										{
											name: "2024-11-03", testHandler: testTomorrowMidnight,
											zone: "America/Kentucky/Monticello", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
										},
										{
											name: "2024-11-04", testHandler: testTomorrowMidnight,
											zone: "America/Kentucky/Monticello", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
										},
										{
											name: "2024-11-05", testHandler: testTomorrowMidnight,
											zone: "America/Kentucky/Monticello", dstHandover: "2024-11-04T00:00:00-05:00",
											date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
										},
									}},
							},
						},
					},
				},
				{
					name: "Knox_IN",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Knox_IN", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Knox_IN", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Knox_IN", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Knox_IN", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Knox_IN", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Knox_IN", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Knox_IN", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Knox_IN", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Knox_IN", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Knox_IN", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Knox_IN", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Knox_IN", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Knox_IN", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Knox_IN", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Knox_IN", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Knox_IN", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Knox_IN", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Knox_IN", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
					},
				},
				{
					name: "Los_Angeles",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Los_Angeles", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Los_Angeles", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Los_Angeles", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Los_Angeles", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Los_Angeles", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Los_Angeles", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Los_Angeles", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Los_Angeles", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Los_Angeles", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Los_Angeles", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Los_Angeles", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Los_Angeles", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Los_Angeles", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Los_Angeles", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Los_Angeles", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Los_Angeles", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Los_Angeles", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Los_Angeles", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
					},
				},
				{
					name: "Louisville",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Louisville", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Louisville", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Louisville", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Louisville", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Louisville", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Louisville", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Louisville", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Louisville", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Louisville", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Louisville", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Louisville", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Louisville", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Louisville", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Louisville", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Louisville", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Louisville", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Louisville", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Louisville", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
					},
				},
				{
					name: "Matamoros",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Matamoros", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Matamoros", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Matamoros", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Matamoros", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Matamoros", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Matamoros", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Matamoros", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Matamoros", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Matamoros", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Matamoros", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Matamoros", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Matamoros", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Matamoros", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Matamoros", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Matamoros", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Matamoros", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Matamoros", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Matamoros", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
					},
				},
				{
					name: "Menominee",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Menominee", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Menominee", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Menominee", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Menominee", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Menominee", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Menominee", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Menominee", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Menominee", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Menominee", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Menominee", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Menominee", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Menominee", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Menominee", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Menominee", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Menominee", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Menominee", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Menominee", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Menominee", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
					},
				},
				{
					name: "Metlakatla",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Metlakatla", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-10T00:00:00-09:00", expectedDate: "2024-03-10T00:00:00-09:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Metlakatla", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-11T00:00:00-08:00", expectedDate: "2024-03-11T00:00:00-08:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Metlakatla", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-12T00:00:00-08:00", expectedDate: "2024-03-12T00:00:00-08:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Metlakatla", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-03T00:00:00-08:00", expectedDate: "2024-11-03T00:00:00-08:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Metlakatla", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-04T00:00:00-09:00", expectedDate: "2024-11-04T00:00:00-09:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Metlakatla", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-05T00:00:00-09:00", expectedDate: "2024-11-05T00:00:00-09:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Metlakatla", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-10T00:00:00-09:00", expectedDate: "2024-03-10T00:00:00-09:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Metlakatla", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-11T00:00:00-08:00", expectedDate: "2024-03-11T00:00:00-08:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Metlakatla", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-12T00:00:00-08:00", expectedDate: "2024-03-12T00:00:00-08:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Metlakatla", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-03T00:00:00-08:00", expectedDate: "2024-11-03T00:00:00-08:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Metlakatla", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-04T00:00:00-09:00", expectedDate: "2024-11-04T00:00:00-09:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Metlakatla", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-05T00:00:00-09:00", expectedDate: "2024-11-05T00:00:00-09:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Metlakatla", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-10T00:00:00-09:00", expectedDate: "2024-03-10T00:00:00-09:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Metlakatla", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-11T00:00:00-08:00", expectedDate: "2024-03-11T00:00:00-08:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Metlakatla", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-12T00:00:00-08:00", expectedDate: "2024-03-12T00:00:00-08:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Metlakatla", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-03T00:00:00-08:00", expectedDate: "2024-11-03T00:00:00-08:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Metlakatla", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-04T00:00:00-09:00", expectedDate: "2024-11-04T00:00:00-09:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Metlakatla", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-05T00:00:00-09:00", expectedDate: "2024-11-05T00:00:00-09:00",
								},
							}},
					},
				},
				{
					name: "Miquelon",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Miquelon", dstHandover: "2024-03-11T00:00:00-02:00",
									date: "2024-03-10T00:00:00-03:00", expectedDate: "2024-03-10T00:00:00-03:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Miquelon", dstHandover: "2024-03-11T00:00:00-02:00",
									date: "2024-03-11T00:00:00-02:00", expectedDate: "2024-03-11T00:00:00-02:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Miquelon", dstHandover: "2024-03-11T00:00:00-02:00",
									date: "2024-03-12T00:00:00-02:00", expectedDate: "2024-03-12T00:00:00-02:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Miquelon", dstHandover: "2024-11-04T00:00:00-03:00",
									date: "2024-11-03T00:00:00-02:00", expectedDate: "2024-11-03T00:00:00-02:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Miquelon", dstHandover: "2024-11-04T00:00:00-03:00",
									date: "2024-11-04T00:00:00-03:00", expectedDate: "2024-11-04T00:00:00-03:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Miquelon", dstHandover: "2024-11-04T00:00:00-03:00",
									date: "2024-11-05T00:00:00-03:00", expectedDate: "2024-11-05T00:00:00-03:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Miquelon", dstHandover: "2024-03-11T00:00:00-02:00",
									date: "2024-03-10T00:00:00-03:00", expectedDate: "2024-03-10T00:00:00-03:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Miquelon", dstHandover: "2024-03-11T00:00:00-02:00",
									date: "2024-03-11T00:00:00-02:00", expectedDate: "2024-03-11T00:00:00-02:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Miquelon", dstHandover: "2024-03-11T00:00:00-02:00",
									date: "2024-03-12T00:00:00-02:00", expectedDate: "2024-03-12T00:00:00-02:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Miquelon", dstHandover: "2024-11-04T00:00:00-03:00",
									date: "2024-11-03T00:00:00-02:00", expectedDate: "2024-11-03T00:00:00-02:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Miquelon", dstHandover: "2024-11-04T00:00:00-03:00",
									date: "2024-11-04T00:00:00-03:00", expectedDate: "2024-11-04T00:00:00-03:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Miquelon", dstHandover: "2024-11-04T00:00:00-03:00",
									date: "2024-11-05T00:00:00-03:00", expectedDate: "2024-11-05T00:00:00-03:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Miquelon", dstHandover: "2024-03-11T00:00:00-02:00",
									date: "2024-03-10T00:00:00-03:00", expectedDate: "2024-03-10T00:00:00-03:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Miquelon", dstHandover: "2024-03-11T00:00:00-02:00",
									date: "2024-03-11T00:00:00-02:00", expectedDate: "2024-03-11T00:00:00-02:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Miquelon", dstHandover: "2024-03-11T00:00:00-02:00",
									date: "2024-03-12T00:00:00-02:00", expectedDate: "2024-03-12T00:00:00-02:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Miquelon", dstHandover: "2024-11-04T00:00:00-03:00",
									date: "2024-11-03T00:00:00-02:00", expectedDate: "2024-11-03T00:00:00-02:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Miquelon", dstHandover: "2024-11-04T00:00:00-03:00",
									date: "2024-11-04T00:00:00-03:00", expectedDate: "2024-11-04T00:00:00-03:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Miquelon", dstHandover: "2024-11-04T00:00:00-03:00",
									date: "2024-11-05T00:00:00-03:00", expectedDate: "2024-11-05T00:00:00-03:00",
								},
							}},
					},
				},
				{
					name: "Moncton",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Moncton", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-10T00:00:00-04:00", expectedDate: "2024-03-10T00:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Moncton", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-11T00:00:00-03:00", expectedDate: "2024-03-11T00:00:00-03:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Moncton", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-12T00:00:00-03:00", expectedDate: "2024-03-12T00:00:00-03:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Moncton", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-03T00:00:00-03:00", expectedDate: "2024-11-03T00:00:00-03:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Moncton", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-04T00:00:00-04:00", expectedDate: "2024-11-04T00:00:00-04:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Moncton", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-05T00:00:00-04:00", expectedDate: "2024-11-05T00:00:00-04:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Moncton", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-10T00:00:00-04:00", expectedDate: "2024-03-10T00:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Moncton", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-11T00:00:00-03:00", expectedDate: "2024-03-11T00:00:00-03:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Moncton", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-12T00:00:00-03:00", expectedDate: "2024-03-12T00:00:00-03:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Moncton", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-03T00:00:00-03:00", expectedDate: "2024-11-03T00:00:00-03:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Moncton", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-04T00:00:00-04:00", expectedDate: "2024-11-04T00:00:00-04:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Moncton", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-05T00:00:00-04:00", expectedDate: "2024-11-05T00:00:00-04:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Moncton", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-10T00:00:00-04:00", expectedDate: "2024-03-10T00:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Moncton", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-11T00:00:00-03:00", expectedDate: "2024-03-11T00:00:00-03:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Moncton", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-12T00:00:00-03:00", expectedDate: "2024-03-12T00:00:00-03:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Moncton", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-03T00:00:00-03:00", expectedDate: "2024-11-03T00:00:00-03:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Moncton", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-04T00:00:00-04:00", expectedDate: "2024-11-04T00:00:00-04:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Moncton", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-05T00:00:00-04:00", expectedDate: "2024-11-05T00:00:00-04:00",
								},
							}},
					},
				},
				{
					name: "Montreal",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Montreal", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Montreal", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Montreal", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Montreal", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Montreal", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Montreal", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Montreal", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Montreal", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Montreal", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Montreal", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Montreal", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Montreal", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Montreal", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Montreal", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Montreal", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Montreal", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Montreal", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Montreal", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
					},
				},
				{
					name: "Nassau",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Nassau", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Nassau", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Nassau", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Nassau", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Nassau", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Nassau", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Nassau", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Nassau", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Nassau", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Nassau", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Nassau", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Nassau", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Nassau", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Nassau", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Nassau", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Nassau", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Nassau", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Nassau", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
					},
				},
				{
					name: "New_York",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/New_York", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/New_York", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/New_York", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/New_York", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/New_York", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/New_York", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/New_York", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/New_York", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/New_York", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/New_York", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/New_York", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/New_York", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/New_York", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/New_York", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/New_York", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/New_York", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/New_York", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/New_York", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
					},
				},
				{
					name: "Nipigon",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Nipigon", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Nipigon", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Nipigon", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Nipigon", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Nipigon", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Nipigon", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Nipigon", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Nipigon", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Nipigon", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Nipigon", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Nipigon", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Nipigon", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Nipigon", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Nipigon", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Nipigon", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Nipigon", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Nipigon", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Nipigon", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
					},
				},
				{
					name: "Nome",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Nome", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-10T00:00:00-09:00", expectedDate: "2024-03-10T00:00:00-09:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Nome", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-11T00:00:00-08:00", expectedDate: "2024-03-11T00:00:00-08:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Nome", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-12T00:00:00-08:00", expectedDate: "2024-03-12T00:00:00-08:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Nome", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-03T00:00:00-08:00", expectedDate: "2024-11-03T00:00:00-08:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Nome", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-04T00:00:00-09:00", expectedDate: "2024-11-04T00:00:00-09:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Nome", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-05T00:00:00-09:00", expectedDate: "2024-11-05T00:00:00-09:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Nome", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-10T00:00:00-09:00", expectedDate: "2024-03-10T00:00:00-09:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Nome", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-11T00:00:00-08:00", expectedDate: "2024-03-11T00:00:00-08:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Nome", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-12T00:00:00-08:00", expectedDate: "2024-03-12T00:00:00-08:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Nome", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-03T00:00:00-08:00", expectedDate: "2024-11-03T00:00:00-08:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Nome", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-04T00:00:00-09:00", expectedDate: "2024-11-04T00:00:00-09:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Nome", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-05T00:00:00-09:00", expectedDate: "2024-11-05T00:00:00-09:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Nome", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-10T00:00:00-09:00", expectedDate: "2024-03-10T00:00:00-09:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Nome", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-11T00:00:00-08:00", expectedDate: "2024-03-11T00:00:00-08:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Nome", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-12T00:00:00-08:00", expectedDate: "2024-03-12T00:00:00-08:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Nome", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-03T00:00:00-08:00", expectedDate: "2024-11-03T00:00:00-08:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Nome", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-04T00:00:00-09:00", expectedDate: "2024-11-04T00:00:00-09:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Nome", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-05T00:00:00-09:00", expectedDate: "2024-11-05T00:00:00-09:00",
								},
							}},
					},
				},
				{
					name: "North_Dakota",
					children: []testCase{
						{
							name: "Beulah",
							children: []testCase{
								{
									name: "LocalMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testLocalMidnight,
											zone: "America/North_Dakota/Beulah", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
										},
										{
											name: "2024-03-11", testHandler: testLocalMidnight,
											zone: "America/North_Dakota/Beulah", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
										},
										{
											name: "2024-03-12", testHandler: testLocalMidnight,
											zone: "America/North_Dakota/Beulah", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
										},
										{
											name: "2024-11-03", testHandler: testLocalMidnight,
											zone: "America/North_Dakota/Beulah", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
										},
										{
											name: "2024-11-04", testHandler: testLocalMidnight,
											zone: "America/North_Dakota/Beulah", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
										},
										{
											name: "2024-11-05", testHandler: testLocalMidnight,
											zone: "America/North_Dakota/Beulah", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
										},
									}},
								{
									name: "YesterdayMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testYesterdayMidnight,
											zone: "America/North_Dakota/Beulah", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
										},
										{
											name: "2024-03-11", testHandler: testYesterdayMidnight,
											zone: "America/North_Dakota/Beulah", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
										},
										{
											name: "2024-03-12", testHandler: testYesterdayMidnight,
											zone: "America/North_Dakota/Beulah", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
										},
										{
											name: "2024-11-03", testHandler: testYesterdayMidnight,
											zone: "America/North_Dakota/Beulah", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
										},
										{
											name: "2024-11-04", testHandler: testYesterdayMidnight,
											zone: "America/North_Dakota/Beulah", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
										},
										{
											name: "2024-11-05", testHandler: testYesterdayMidnight,
											zone: "America/North_Dakota/Beulah", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
										},
									}},
								{
									name: "TomorrowMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testTomorrowMidnight,
											zone: "America/North_Dakota/Beulah", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
										},
										{
											name: "2024-03-11", testHandler: testTomorrowMidnight,
											zone: "America/North_Dakota/Beulah", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
										},
										{
											name: "2024-03-12", testHandler: testTomorrowMidnight,
											zone: "America/North_Dakota/Beulah", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
										},
										{
											name: "2024-11-03", testHandler: testTomorrowMidnight,
											zone: "America/North_Dakota/Beulah", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
										},
										{
											name: "2024-11-04", testHandler: testTomorrowMidnight,
											zone: "America/North_Dakota/Beulah", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
										},
										{
											name: "2024-11-05", testHandler: testTomorrowMidnight,
											zone: "America/North_Dakota/Beulah", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
										},
									}},
							},
						},
						{
							name: "Center",
							children: []testCase{
								{
									name: "LocalMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testLocalMidnight,
											zone: "America/North_Dakota/Center", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
										},
										{
											name: "2024-03-11", testHandler: testLocalMidnight,
											zone: "America/North_Dakota/Center", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
										},
										{
											name: "2024-03-12", testHandler: testLocalMidnight,
											zone: "America/North_Dakota/Center", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
										},
										{
											name: "2024-11-03", testHandler: testLocalMidnight,
											zone: "America/North_Dakota/Center", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
										},
										{
											name: "2024-11-04", testHandler: testLocalMidnight,
											zone: "America/North_Dakota/Center", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
										},
										{
											name: "2024-11-05", testHandler: testLocalMidnight,
											zone: "America/North_Dakota/Center", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
										},
									}},
								{
									name: "YesterdayMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testYesterdayMidnight,
											zone: "America/North_Dakota/Center", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
										},
										{
											name: "2024-03-11", testHandler: testYesterdayMidnight,
											zone: "America/North_Dakota/Center", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
										},
										{
											name: "2024-03-12", testHandler: testYesterdayMidnight,
											zone: "America/North_Dakota/Center", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
										},
										{
											name: "2024-11-03", testHandler: testYesterdayMidnight,
											zone: "America/North_Dakota/Center", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
										},
										{
											name: "2024-11-04", testHandler: testYesterdayMidnight,
											zone: "America/North_Dakota/Center", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
										},
										{
											name: "2024-11-05", testHandler: testYesterdayMidnight,
											zone: "America/North_Dakota/Center", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
										},
									}},
								{
									name: "TomorrowMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testTomorrowMidnight,
											zone: "America/North_Dakota/Center", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
										},
										{
											name: "2024-03-11", testHandler: testTomorrowMidnight,
											zone: "America/North_Dakota/Center", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
										},
										{
											name: "2024-03-12", testHandler: testTomorrowMidnight,
											zone: "America/North_Dakota/Center", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
										},
										{
											name: "2024-11-03", testHandler: testTomorrowMidnight,
											zone: "America/North_Dakota/Center", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
										},
										{
											name: "2024-11-04", testHandler: testTomorrowMidnight,
											zone: "America/North_Dakota/Center", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
										},
										{
											name: "2024-11-05", testHandler: testTomorrowMidnight,
											zone: "America/North_Dakota/Center", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
										},
									}},
							},
						},
						{
							name: "New_Salem",
							children: []testCase{
								{
									name: "LocalMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testLocalMidnight,
											zone: "America/North_Dakota/New_Salem", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
										},
										{
											name: "2024-03-11", testHandler: testLocalMidnight,
											zone: "America/North_Dakota/New_Salem", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
										},
										{
											name: "2024-03-12", testHandler: testLocalMidnight,
											zone: "America/North_Dakota/New_Salem", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
										},
										{
											name: "2024-11-03", testHandler: testLocalMidnight,
											zone: "America/North_Dakota/New_Salem", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
										},
										{
											name: "2024-11-04", testHandler: testLocalMidnight,
											zone: "America/North_Dakota/New_Salem", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
										},
										{
											name: "2024-11-05", testHandler: testLocalMidnight,
											zone: "America/North_Dakota/New_Salem", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
										},
									}},
								{
									name: "YesterdayMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testYesterdayMidnight,
											zone: "America/North_Dakota/New_Salem", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
										},
										{
											name: "2024-03-11", testHandler: testYesterdayMidnight,
											zone: "America/North_Dakota/New_Salem", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
										},
										{
											name: "2024-03-12", testHandler: testYesterdayMidnight,
											zone: "America/North_Dakota/New_Salem", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
										},
										{
											name: "2024-11-03", testHandler: testYesterdayMidnight,
											zone: "America/North_Dakota/New_Salem", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
										},
										{
											name: "2024-11-04", testHandler: testYesterdayMidnight,
											zone: "America/North_Dakota/New_Salem", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
										},
										{
											name: "2024-11-05", testHandler: testYesterdayMidnight,
											zone: "America/North_Dakota/New_Salem", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
										},
									}},
								{
									name: "TomorrowMidnight", children: []testCase{
										{
											name: "2024-03-10", testHandler: testTomorrowMidnight,
											zone: "America/North_Dakota/New_Salem", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
										},
										{
											name: "2024-03-11", testHandler: testTomorrowMidnight,
											zone: "America/North_Dakota/New_Salem", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
										},
										{
											name: "2024-03-12", testHandler: testTomorrowMidnight,
											zone: "America/North_Dakota/New_Salem", dstHandover: "2024-03-11T00:00:00-05:00",
											date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
										},
										{
											name: "2024-11-03", testHandler: testTomorrowMidnight,
											zone: "America/North_Dakota/New_Salem", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
										},
										{
											name: "2024-11-04", testHandler: testTomorrowMidnight,
											zone: "America/North_Dakota/New_Salem", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
										},
										{
											name: "2024-11-05", testHandler: testTomorrowMidnight,
											zone: "America/North_Dakota/New_Salem", dstHandover: "2024-11-04T00:00:00-06:00",
											date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
										},
									}},
							},
						},
					},
				},
				{
					name: "Nuuk",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-30", testHandler: testLocalMidnight,
									zone: "America/Nuuk", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-03-30T00:00:00-02:00", expectedDate: "2024-03-30T00:00:00-02:00",
								},
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "America/Nuuk", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-03-31T00:00:00-01:00", expectedDate: "2024-03-31T00:00:00-01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "America/Nuuk", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-04-01T00:00:00-01:00", expectedDate: "2024-04-01T00:00:00-01:00",
								},
								{
									name: "2024-10-26", testHandler: testLocalMidnight,
									zone: "America/Nuuk", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-26T00:00:00-01:00", expectedDate: "2024-10-26T00:00:00-01:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "America/Nuuk", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-27T00:00:00-02:00", expectedDate: "2024-10-27T00:00:00-02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "America/Nuuk", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-28T00:00:00-02:00", expectedDate: "2024-10-28T00:00:00-02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-30", testHandler: testYesterdayMidnight,
									zone: "America/Nuuk", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-03-30T00:00:00-02:00", expectedDate: "2024-03-30T00:00:00-02:00",
								},
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "America/Nuuk", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-03-31T00:00:00-01:00", expectedDate: "2024-03-31T00:00:00-01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "America/Nuuk", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-04-01T00:00:00-01:00", expectedDate: "2024-04-01T00:00:00-01:00",
								},
								{
									name: "2024-10-26", testHandler: testYesterdayMidnight,
									zone: "America/Nuuk", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-26T00:00:00-01:00", expectedDate: "2024-10-26T00:00:00-01:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "America/Nuuk", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-27T00:00:00-02:00", expectedDate: "2024-10-27T00:00:00-02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "America/Nuuk", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-28T00:00:00-02:00", expectedDate: "2024-10-28T00:00:00-02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-30", testHandler: testTomorrowMidnight,
									zone: "America/Nuuk", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-03-30T00:00:00-02:00", expectedDate: "2024-03-30T00:00:00-02:00",
								},
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "America/Nuuk", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-03-31T00:00:00-01:00", expectedDate: "2024-03-31T00:00:00-01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "America/Nuuk", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-04-01T00:00:00-01:00", expectedDate: "2024-04-01T00:00:00-01:00",
								},
								{
									name: "2024-10-26", testHandler: testTomorrowMidnight,
									zone: "America/Nuuk", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-26T00:00:00-01:00", expectedDate: "2024-10-26T00:00:00-01:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "America/Nuuk", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-27T00:00:00-02:00", expectedDate: "2024-10-27T00:00:00-02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "America/Nuuk", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-28T00:00:00-02:00", expectedDate: "2024-10-28T00:00:00-02:00",
								},
							}},
					},
				},
				{
					name: "Ojinaga",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Ojinaga", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Ojinaga", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Ojinaga", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Ojinaga", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Ojinaga", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Ojinaga", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Ojinaga", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Ojinaga", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Ojinaga", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Ojinaga", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Ojinaga", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Ojinaga", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Ojinaga", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Ojinaga", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Ojinaga", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Ojinaga", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Ojinaga", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Ojinaga", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
					},
				},
				{
					name: "Pangnirtung",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Pangnirtung", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Pangnirtung", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Pangnirtung", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Pangnirtung", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Pangnirtung", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Pangnirtung", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Pangnirtung", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Pangnirtung", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Pangnirtung", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Pangnirtung", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Pangnirtung", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Pangnirtung", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Pangnirtung", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Pangnirtung", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Pangnirtung", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Pangnirtung", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Pangnirtung", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Pangnirtung", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
					},
				},
				{
					name: "Port-au-Prince",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Port-au-Prince", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Port-au-Prince", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Port-au-Prince", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Port-au-Prince", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Port-au-Prince", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Port-au-Prince", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Port-au-Prince", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Port-au-Prince", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Port-au-Prince", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Port-au-Prince", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Port-au-Prince", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Port-au-Prince", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Port-au-Prince", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Port-au-Prince", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Port-au-Prince", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Port-au-Prince", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Port-au-Prince", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Port-au-Prince", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
					},
				},
				{
					name: "Rainy_River",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Rainy_River", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Rainy_River", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Rainy_River", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Rainy_River", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Rainy_River", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Rainy_River", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Rainy_River", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Rainy_River", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Rainy_River", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Rainy_River", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Rainy_River", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Rainy_River", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Rainy_River", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Rainy_River", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Rainy_River", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Rainy_River", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Rainy_River", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Rainy_River", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
					},
				},
				{
					name: "Rankin_Inlet",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Rankin_Inlet", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Rankin_Inlet", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Rankin_Inlet", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Rankin_Inlet", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Rankin_Inlet", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Rankin_Inlet", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Rankin_Inlet", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Rankin_Inlet", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Rankin_Inlet", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Rankin_Inlet", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Rankin_Inlet", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Rankin_Inlet", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Rankin_Inlet", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Rankin_Inlet", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Rankin_Inlet", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Rankin_Inlet", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Rankin_Inlet", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Rankin_Inlet", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
					},
				},
				{
					name: "Resolute",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Resolute", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Resolute", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Resolute", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Resolute", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Resolute", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Resolute", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Resolute", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Resolute", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Resolute", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Resolute", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Resolute", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Resolute", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Resolute", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Resolute", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Resolute", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Resolute", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Resolute", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Resolute", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
					},
				},
				{
					name: "Santa_Isabel",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Santa_Isabel", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Santa_Isabel", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Santa_Isabel", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Santa_Isabel", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Santa_Isabel", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Santa_Isabel", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Santa_Isabel", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Santa_Isabel", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Santa_Isabel", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Santa_Isabel", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Santa_Isabel", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Santa_Isabel", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Santa_Isabel", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Santa_Isabel", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Santa_Isabel", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Santa_Isabel", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Santa_Isabel", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Santa_Isabel", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
					},
				},
				{
					name: "Santiago",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-09-07", testHandler: testLocalMidnight,
									zone: "America/Santiago", dstHandover: "2024-09-08T23:00:00-03:00",
									date: "2024-09-07T23:00:00-04:00", expectedDate: "2024-09-07T23:00:00-04:00",
								},
								{
									name: "2024-09-08", testHandler: testLocalMidnight,
									zone: "America/Santiago", dstHandover: "2024-09-08T23:00:00-03:00",
									date: "2024-09-08T23:00:00-03:00", expectedDate: "2024-09-08T23:00:00-03:00",
								},
								{
									name: "2024-09-09", testHandler: testLocalMidnight,
									zone: "America/Santiago", dstHandover: "2024-09-08T23:00:00-03:00",
									date: "2024-09-09T23:00:00-03:00", expectedDate: "2024-09-09T23:00:00-03:00",
								},
								{
									name: "2024-04-06", testHandler: testLocalMidnight,
									zone: "America/Santiago", dstHandover: "2024-04-07T00:00:00-04:00",
									date: "2024-04-06T00:00:00-03:00", expectedDate: "2024-04-06T00:00:00-03:00",
								},
								{
									name: "2024-04-07", testHandler: testLocalMidnight,
									zone: "America/Santiago", dstHandover: "2024-04-07T00:00:00-04:00",
									date: "2024-04-07T00:00:00-04:00", expectedDate: "2024-04-07T00:00:00-04:00",
								},
								{
									name: "2024-04-08", testHandler: testLocalMidnight,
									zone: "America/Santiago", dstHandover: "2024-04-07T00:00:00-04:00",
									date: "2024-04-08T00:00:00-04:00", expectedDate: "2024-04-08T00:00:00-04:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-09-07", testHandler: testYesterdayMidnight,
									zone: "America/Santiago", dstHandover: "2024-09-08T23:00:00-03:00",
									date: "2024-09-07T23:00:00-04:00", expectedDate: "2024-09-07T23:00:00-04:00",
								},
								{
									name: "2024-09-08", testHandler: testYesterdayMidnight,
									zone: "America/Santiago", dstHandover: "2024-09-08T23:00:00-03:00",
									date: "2024-09-08T23:00:00-03:00", expectedDate: "2024-09-08T23:00:00-03:00",
								},
								{
									name: "2024-09-09", testHandler: testYesterdayMidnight,
									zone: "America/Santiago", dstHandover: "2024-09-08T23:00:00-03:00",
									date: "2024-09-09T23:00:00-03:00", expectedDate: "2024-09-09T23:00:00-03:00",
								},
								{
									name: "2024-04-06", testHandler: testYesterdayMidnight,
									zone: "America/Santiago", dstHandover: "2024-04-07T00:00:00-04:00",
									date: "2024-04-06T00:00:00-03:00", expectedDate: "2024-04-06T00:00:00-03:00",
								},
								{
									name: "2024-04-07", testHandler: testYesterdayMidnight,
									zone: "America/Santiago", dstHandover: "2024-04-07T00:00:00-04:00",
									date: "2024-04-07T00:00:00-04:00", expectedDate: "2024-04-07T00:00:00-04:00",
								},
								{
									name: "2024-04-08", testHandler: testYesterdayMidnight,
									zone: "America/Santiago", dstHandover: "2024-04-07T00:00:00-04:00",
									date: "2024-04-08T00:00:00-04:00", expectedDate: "2024-04-08T00:00:00-04:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-09-07", testHandler: testTomorrowMidnight,
									zone: "America/Santiago", dstHandover: "2024-09-08T23:00:00-03:00",
									date: "2024-09-07T23:00:00-04:00", expectedDate: "2024-09-07T23:00:00-04:00",
								},
								{
									name: "2024-09-08", testHandler: testTomorrowMidnight,
									zone: "America/Santiago", dstHandover: "2024-09-08T23:00:00-03:00",
									date: "2024-09-08T23:00:00-03:00", expectedDate: "2024-09-08T23:00:00-03:00",
								},
								{
									name: "2024-09-09", testHandler: testTomorrowMidnight,
									zone: "America/Santiago", dstHandover: "2024-09-08T23:00:00-03:00",
									date: "2024-09-09T23:00:00-03:00", expectedDate: "2024-09-09T23:00:00-03:00",
								},
								{
									name: "2024-04-06", testHandler: testTomorrowMidnight,
									zone: "America/Santiago", dstHandover: "2024-04-07T00:00:00-04:00",
									date: "2024-04-06T00:00:00-03:00", expectedDate: "2024-04-06T00:00:00-03:00",
								},
								{
									name: "2024-04-07", testHandler: testTomorrowMidnight,
									zone: "America/Santiago", dstHandover: "2024-04-07T00:00:00-04:00",
									date: "2024-04-07T00:00:00-04:00", expectedDate: "2024-04-07T00:00:00-04:00",
								},
								{
									name: "2024-04-08", testHandler: testTomorrowMidnight,
									zone: "America/Santiago", dstHandover: "2024-04-07T00:00:00-04:00",
									date: "2024-04-08T00:00:00-04:00", expectedDate: "2024-04-08T00:00:00-04:00",
								},
							}},
					},
				},
				{
					name: "Scoresbysund",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-30", testHandler: testLocalMidnight,
									zone: "America/Scoresbysund", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-03-30T00:00:00-01:00", expectedDate: "2024-03-30T00:00:00-01:00",
								},
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "America/Scoresbysund", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-03-31T00:00:00-01:00", expectedDate: "2024-03-31T00:00:00-01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "America/Scoresbysund", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-04-01T00:00:00-01:00", expectedDate: "2024-04-01T00:00:00-01:00",
								},
								{
									name: "2024-10-26", testHandler: testLocalMidnight,
									zone: "America/Scoresbysund", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-26T00:00:00-01:00", expectedDate: "2024-10-26T00:00:00-01:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "America/Scoresbysund", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-27T00:00:00-02:00", expectedDate: "2024-10-27T00:00:00-02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "America/Scoresbysund", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-28T00:00:00-02:00", expectedDate: "2024-10-28T00:00:00-02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-30", testHandler: testYesterdayMidnight,
									zone: "America/Scoresbysund", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-03-30T00:00:00-01:00", expectedDate: "2024-03-30T00:00:00-01:00",
								},
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "America/Scoresbysund", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-03-31T00:00:00-01:00", expectedDate: "2024-03-31T00:00:00-01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "America/Scoresbysund", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-04-01T00:00:00-01:00", expectedDate: "2024-04-01T00:00:00-01:00",
								},
								{
									name: "2024-10-26", testHandler: testYesterdayMidnight,
									zone: "America/Scoresbysund", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-26T00:00:00-01:00", expectedDate: "2024-10-26T00:00:00-01:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "America/Scoresbysund", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-27T00:00:00-02:00", expectedDate: "2024-10-27T00:00:00-02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "America/Scoresbysund", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-28T00:00:00-02:00", expectedDate: "2024-10-28T00:00:00-02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-30", testHandler: testTomorrowMidnight,
									zone: "America/Scoresbysund", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-03-30T00:00:00-01:00", expectedDate: "2024-03-30T00:00:00-01:00",
								},
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "America/Scoresbysund", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-03-31T00:00:00-01:00", expectedDate: "2024-03-31T00:00:00-01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "America/Scoresbysund", dstHandover: "2024-03-31T00:00:00-01:00",
									date: "2024-04-01T00:00:00-01:00", expectedDate: "2024-04-01T00:00:00-01:00",
								},
								{
									name: "2024-10-26", testHandler: testTomorrowMidnight,
									zone: "America/Scoresbysund", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-26T00:00:00-01:00", expectedDate: "2024-10-26T00:00:00-01:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "America/Scoresbysund", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-27T00:00:00-02:00", expectedDate: "2024-10-27T00:00:00-02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "America/Scoresbysund", dstHandover: "2024-10-27T00:00:00-02:00",
									date: "2024-10-28T00:00:00-02:00", expectedDate: "2024-10-28T00:00:00-02:00",
								},
							}},
					},
				},
				{
					name: "Shiprock",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Shiprock", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Shiprock", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Shiprock", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Shiprock", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Shiprock", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Shiprock", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Shiprock", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Shiprock", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Shiprock", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Shiprock", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Shiprock", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Shiprock", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Shiprock", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Shiprock", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Shiprock", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Shiprock", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Shiprock", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Shiprock", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
					},
				},
				{
					name: "Sitka",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Sitka", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-10T00:00:00-09:00", expectedDate: "2024-03-10T00:00:00-09:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Sitka", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-11T00:00:00-08:00", expectedDate: "2024-03-11T00:00:00-08:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Sitka", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-12T00:00:00-08:00", expectedDate: "2024-03-12T00:00:00-08:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Sitka", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-03T00:00:00-08:00", expectedDate: "2024-11-03T00:00:00-08:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Sitka", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-04T00:00:00-09:00", expectedDate: "2024-11-04T00:00:00-09:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Sitka", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-05T00:00:00-09:00", expectedDate: "2024-11-05T00:00:00-09:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Sitka", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-10T00:00:00-09:00", expectedDate: "2024-03-10T00:00:00-09:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Sitka", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-11T00:00:00-08:00", expectedDate: "2024-03-11T00:00:00-08:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Sitka", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-12T00:00:00-08:00", expectedDate: "2024-03-12T00:00:00-08:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Sitka", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-03T00:00:00-08:00", expectedDate: "2024-11-03T00:00:00-08:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Sitka", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-04T00:00:00-09:00", expectedDate: "2024-11-04T00:00:00-09:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Sitka", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-05T00:00:00-09:00", expectedDate: "2024-11-05T00:00:00-09:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Sitka", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-10T00:00:00-09:00", expectedDate: "2024-03-10T00:00:00-09:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Sitka", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-11T00:00:00-08:00", expectedDate: "2024-03-11T00:00:00-08:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Sitka", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-12T00:00:00-08:00", expectedDate: "2024-03-12T00:00:00-08:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Sitka", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-03T00:00:00-08:00", expectedDate: "2024-11-03T00:00:00-08:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Sitka", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-04T00:00:00-09:00", expectedDate: "2024-11-04T00:00:00-09:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Sitka", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-05T00:00:00-09:00", expectedDate: "2024-11-05T00:00:00-09:00",
								},
							}},
					},
				},
				{
					name: "St_Johns",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/St_Johns", dstHandover: "2024-03-11T00:00:00-02:30",
									date: "2024-03-10T00:00:00-03:30", expectedDate: "2024-03-10T00:00:00-03:30",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/St_Johns", dstHandover: "2024-03-11T00:00:00-02:30",
									date: "2024-03-11T00:00:00-02:30", expectedDate: "2024-03-11T00:00:00-02:30",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/St_Johns", dstHandover: "2024-03-11T00:00:00-02:30",
									date: "2024-03-12T00:00:00-02:30", expectedDate: "2024-03-12T00:00:00-02:30",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/St_Johns", dstHandover: "2024-11-04T00:00:00-03:30",
									date: "2024-11-03T00:00:00-02:30", expectedDate: "2024-11-03T00:00:00-02:30",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/St_Johns", dstHandover: "2024-11-04T00:00:00-03:30",
									date: "2024-11-04T00:00:00-03:30", expectedDate: "2024-11-04T00:00:00-03:30",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/St_Johns", dstHandover: "2024-11-04T00:00:00-03:30",
									date: "2024-11-05T00:00:00-03:30", expectedDate: "2024-11-05T00:00:00-03:30",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/St_Johns", dstHandover: "2024-03-11T00:00:00-02:30",
									date: "2024-03-10T00:00:00-03:30", expectedDate: "2024-03-10T00:00:00-03:30",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/St_Johns", dstHandover: "2024-03-11T00:00:00-02:30",
									date: "2024-03-11T00:00:00-02:30", expectedDate: "2024-03-11T00:00:00-02:30",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/St_Johns", dstHandover: "2024-03-11T00:00:00-02:30",
									date: "2024-03-12T00:00:00-02:30", expectedDate: "2024-03-12T00:00:00-02:30",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/St_Johns", dstHandover: "2024-11-04T00:00:00-03:30",
									date: "2024-11-03T00:00:00-02:30", expectedDate: "2024-11-03T00:00:00-02:30",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/St_Johns", dstHandover: "2024-11-04T00:00:00-03:30",
									date: "2024-11-04T00:00:00-03:30", expectedDate: "2024-11-04T00:00:00-03:30",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/St_Johns", dstHandover: "2024-11-04T00:00:00-03:30",
									date: "2024-11-05T00:00:00-03:30", expectedDate: "2024-11-05T00:00:00-03:30",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/St_Johns", dstHandover: "2024-03-11T00:00:00-02:30",
									date: "2024-03-10T00:00:00-03:30", expectedDate: "2024-03-10T00:00:00-03:30",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/St_Johns", dstHandover: "2024-03-11T00:00:00-02:30",
									date: "2024-03-11T00:00:00-02:30", expectedDate: "2024-03-11T00:00:00-02:30",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/St_Johns", dstHandover: "2024-03-11T00:00:00-02:30",
									date: "2024-03-12T00:00:00-02:30", expectedDate: "2024-03-12T00:00:00-02:30",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/St_Johns", dstHandover: "2024-11-04T00:00:00-03:30",
									date: "2024-11-03T00:00:00-02:30", expectedDate: "2024-11-03T00:00:00-02:30",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/St_Johns", dstHandover: "2024-11-04T00:00:00-03:30",
									date: "2024-11-04T00:00:00-03:30", expectedDate: "2024-11-04T00:00:00-03:30",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/St_Johns", dstHandover: "2024-11-04T00:00:00-03:30",
									date: "2024-11-05T00:00:00-03:30", expectedDate: "2024-11-05T00:00:00-03:30",
								},
							}},
					},
				},
				{
					name: "Thule",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Thule", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-10T00:00:00-04:00", expectedDate: "2024-03-10T00:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Thule", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-11T00:00:00-03:00", expectedDate: "2024-03-11T00:00:00-03:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Thule", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-12T00:00:00-03:00", expectedDate: "2024-03-12T00:00:00-03:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Thule", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-03T00:00:00-03:00", expectedDate: "2024-11-03T00:00:00-03:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Thule", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-04T00:00:00-04:00", expectedDate: "2024-11-04T00:00:00-04:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Thule", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-05T00:00:00-04:00", expectedDate: "2024-11-05T00:00:00-04:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Thule", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-10T00:00:00-04:00", expectedDate: "2024-03-10T00:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Thule", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-11T00:00:00-03:00", expectedDate: "2024-03-11T00:00:00-03:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Thule", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-12T00:00:00-03:00", expectedDate: "2024-03-12T00:00:00-03:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Thule", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-03T00:00:00-03:00", expectedDate: "2024-11-03T00:00:00-03:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Thule", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-04T00:00:00-04:00", expectedDate: "2024-11-04T00:00:00-04:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Thule", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-05T00:00:00-04:00", expectedDate: "2024-11-05T00:00:00-04:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Thule", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-10T00:00:00-04:00", expectedDate: "2024-03-10T00:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Thule", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-11T00:00:00-03:00", expectedDate: "2024-03-11T00:00:00-03:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Thule", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-12T00:00:00-03:00", expectedDate: "2024-03-12T00:00:00-03:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Thule", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-03T00:00:00-03:00", expectedDate: "2024-11-03T00:00:00-03:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Thule", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-04T00:00:00-04:00", expectedDate: "2024-11-04T00:00:00-04:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Thule", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-05T00:00:00-04:00", expectedDate: "2024-11-05T00:00:00-04:00",
								},
							}},
					},
				},
				{
					name: "Thunder_Bay",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Thunder_Bay", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Thunder_Bay", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Thunder_Bay", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Thunder_Bay", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Thunder_Bay", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Thunder_Bay", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Thunder_Bay", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Thunder_Bay", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Thunder_Bay", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Thunder_Bay", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Thunder_Bay", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Thunder_Bay", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Thunder_Bay", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Thunder_Bay", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Thunder_Bay", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Thunder_Bay", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Thunder_Bay", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Thunder_Bay", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
					},
				},
				{
					name: "Tijuana",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Tijuana", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Tijuana", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Tijuana", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Tijuana", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Tijuana", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Tijuana", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Tijuana", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Tijuana", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Tijuana", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Tijuana", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Tijuana", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Tijuana", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Tijuana", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Tijuana", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Tijuana", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Tijuana", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Tijuana", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Tijuana", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
					},
				},
				{
					name: "Toronto",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Toronto", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Toronto", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Toronto", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Toronto", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Toronto", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Toronto", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Toronto", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Toronto", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Toronto", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Toronto", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Toronto", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Toronto", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Toronto", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Toronto", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Toronto", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Toronto", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Toronto", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Toronto", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
					},
				},
				{
					name: "Vancouver",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Vancouver", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Vancouver", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Vancouver", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Vancouver", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Vancouver", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Vancouver", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Vancouver", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Vancouver", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Vancouver", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Vancouver", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Vancouver", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Vancouver", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Vancouver", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Vancouver", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Vancouver", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Vancouver", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Vancouver", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Vancouver", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
					},
				},
				{
					name: "Winnipeg",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Winnipeg", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Winnipeg", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Winnipeg", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Winnipeg", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Winnipeg", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Winnipeg", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Winnipeg", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Winnipeg", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Winnipeg", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Winnipeg", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Winnipeg", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Winnipeg", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Winnipeg", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Winnipeg", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Winnipeg", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Winnipeg", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Winnipeg", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Winnipeg", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
					},
				},
				{
					name: "Yakutat",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Yakutat", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-10T00:00:00-09:00", expectedDate: "2024-03-10T00:00:00-09:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Yakutat", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-11T00:00:00-08:00", expectedDate: "2024-03-11T00:00:00-08:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Yakutat", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-12T00:00:00-08:00", expectedDate: "2024-03-12T00:00:00-08:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Yakutat", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-03T00:00:00-08:00", expectedDate: "2024-11-03T00:00:00-08:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Yakutat", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-04T00:00:00-09:00", expectedDate: "2024-11-04T00:00:00-09:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Yakutat", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-05T00:00:00-09:00", expectedDate: "2024-11-05T00:00:00-09:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Yakutat", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-10T00:00:00-09:00", expectedDate: "2024-03-10T00:00:00-09:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Yakutat", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-11T00:00:00-08:00", expectedDate: "2024-03-11T00:00:00-08:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Yakutat", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-12T00:00:00-08:00", expectedDate: "2024-03-12T00:00:00-08:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Yakutat", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-03T00:00:00-08:00", expectedDate: "2024-11-03T00:00:00-08:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Yakutat", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-04T00:00:00-09:00", expectedDate: "2024-11-04T00:00:00-09:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Yakutat", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-05T00:00:00-09:00", expectedDate: "2024-11-05T00:00:00-09:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Yakutat", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-10T00:00:00-09:00", expectedDate: "2024-03-10T00:00:00-09:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Yakutat", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-11T00:00:00-08:00", expectedDate: "2024-03-11T00:00:00-08:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Yakutat", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-12T00:00:00-08:00", expectedDate: "2024-03-12T00:00:00-08:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Yakutat", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-03T00:00:00-08:00", expectedDate: "2024-11-03T00:00:00-08:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Yakutat", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-04T00:00:00-09:00", expectedDate: "2024-11-04T00:00:00-09:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Yakutat", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-05T00:00:00-09:00", expectedDate: "2024-11-05T00:00:00-09:00",
								},
							}},
					},
				},
				{
					name: "Yellowknife",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "America/Yellowknife", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "America/Yellowknife", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "America/Yellowknife", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "America/Yellowknife", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "America/Yellowknife", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "America/Yellowknife", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "America/Yellowknife", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "America/Yellowknife", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "America/Yellowknife", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "America/Yellowknife", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "America/Yellowknife", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "America/Yellowknife", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "America/Yellowknife", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "America/Yellowknife", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "America/Yellowknife", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "America/Yellowknife", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "America/Yellowknife", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "America/Yellowknife", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
					},
				},
			},
		},
		{
			name: "Antarctica",
			children: []testCase{
				{
					name: "Macquarie",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testLocalMidnight,
									zone: "Antarctica/Macquarie", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testLocalMidnight,
									zone: "Antarctica/Macquarie", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testLocalMidnight,
									zone: "Antarctica/Macquarie", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testLocalMidnight,
									zone: "Antarctica/Macquarie", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testLocalMidnight,
									zone: "Antarctica/Macquarie", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testLocalMidnight,
									zone: "Antarctica/Macquarie", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testYesterdayMidnight,
									zone: "Antarctica/Macquarie", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testYesterdayMidnight,
									zone: "Antarctica/Macquarie", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testYesterdayMidnight,
									zone: "Antarctica/Macquarie", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testYesterdayMidnight,
									zone: "Antarctica/Macquarie", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testYesterdayMidnight,
									zone: "Antarctica/Macquarie", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testYesterdayMidnight,
									zone: "Antarctica/Macquarie", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testTomorrowMidnight,
									zone: "Antarctica/Macquarie", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testTomorrowMidnight,
									zone: "Antarctica/Macquarie", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testTomorrowMidnight,
									zone: "Antarctica/Macquarie", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testTomorrowMidnight,
									zone: "Antarctica/Macquarie", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testTomorrowMidnight,
									zone: "Antarctica/Macquarie", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testTomorrowMidnight,
									zone: "Antarctica/Macquarie", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
					},
				},
				{
					name: "McMurdo",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-09-29", testHandler: testLocalMidnight,
									zone: "Antarctica/McMurdo", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-09-29T00:00:00+12:00", expectedDate: "2024-09-29T00:00:00+12:00",
								},
								{
									name: "2024-09-30", testHandler: testLocalMidnight,
									zone: "Antarctica/McMurdo", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-09-30T00:00:00+13:00", expectedDate: "2024-09-30T00:00:00+13:00",
								},
								{
									name: "2024-10-01", testHandler: testLocalMidnight,
									zone: "Antarctica/McMurdo", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-10-01T00:00:00+13:00", expectedDate: "2024-10-01T00:00:00+13:00",
								},
								{
									name: "2024-04-07", testHandler: testLocalMidnight,
									zone: "Antarctica/McMurdo", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-07T00:00:00+13:00", expectedDate: "2024-04-07T00:00:00+13:00",
								},
								{
									name: "2024-04-08", testHandler: testLocalMidnight,
									zone: "Antarctica/McMurdo", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-08T00:00:00+12:00", expectedDate: "2024-04-08T00:00:00+12:00",
								},
								{
									name: "2024-04-09", testHandler: testLocalMidnight,
									zone: "Antarctica/McMurdo", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-09T00:00:00+12:00", expectedDate: "2024-04-09T00:00:00+12:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-09-29", testHandler: testYesterdayMidnight,
									zone: "Antarctica/McMurdo", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-09-29T00:00:00+12:00", expectedDate: "2024-09-29T00:00:00+12:00",
								},
								{
									name: "2024-09-30", testHandler: testYesterdayMidnight,
									zone: "Antarctica/McMurdo", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-09-30T00:00:00+13:00", expectedDate: "2024-09-30T00:00:00+13:00",
								},
								{
									name: "2024-10-01", testHandler: testYesterdayMidnight,
									zone: "Antarctica/McMurdo", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-10-01T00:00:00+13:00", expectedDate: "2024-10-01T00:00:00+13:00",
								},
								{
									name: "2024-04-07", testHandler: testYesterdayMidnight,
									zone: "Antarctica/McMurdo", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-07T00:00:00+13:00", expectedDate: "2024-04-07T00:00:00+13:00",
								},
								{
									name: "2024-04-08", testHandler: testYesterdayMidnight,
									zone: "Antarctica/McMurdo", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-08T00:00:00+12:00", expectedDate: "2024-04-08T00:00:00+12:00",
								},
								{
									name: "2024-04-09", testHandler: testYesterdayMidnight,
									zone: "Antarctica/McMurdo", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-09T00:00:00+12:00", expectedDate: "2024-04-09T00:00:00+12:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-09-29", testHandler: testTomorrowMidnight,
									zone: "Antarctica/McMurdo", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-09-29T00:00:00+12:00", expectedDate: "2024-09-29T00:00:00+12:00",
								},
								{
									name: "2024-09-30", testHandler: testTomorrowMidnight,
									zone: "Antarctica/McMurdo", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-09-30T00:00:00+13:00", expectedDate: "2024-09-30T00:00:00+13:00",
								},
								{
									name: "2024-10-01", testHandler: testTomorrowMidnight,
									zone: "Antarctica/McMurdo", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-10-01T00:00:00+13:00", expectedDate: "2024-10-01T00:00:00+13:00",
								},
								{
									name: "2024-04-07", testHandler: testTomorrowMidnight,
									zone: "Antarctica/McMurdo", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-07T00:00:00+13:00", expectedDate: "2024-04-07T00:00:00+13:00",
								},
								{
									name: "2024-04-08", testHandler: testTomorrowMidnight,
									zone: "Antarctica/McMurdo", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-08T00:00:00+12:00", expectedDate: "2024-04-08T00:00:00+12:00",
								},
								{
									name: "2024-04-09", testHandler: testTomorrowMidnight,
									zone: "Antarctica/McMurdo", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-09T00:00:00+12:00", expectedDate: "2024-04-09T00:00:00+12:00",
								},
							}},
					},
				},
				{
					name: "South_Pole",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-09-29", testHandler: testLocalMidnight,
									zone: "Antarctica/South_Pole", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-09-29T00:00:00+12:00", expectedDate: "2024-09-29T00:00:00+12:00",
								},
								{
									name: "2024-09-30", testHandler: testLocalMidnight,
									zone: "Antarctica/South_Pole", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-09-30T00:00:00+13:00", expectedDate: "2024-09-30T00:00:00+13:00",
								},
								{
									name: "2024-10-01", testHandler: testLocalMidnight,
									zone: "Antarctica/South_Pole", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-10-01T00:00:00+13:00", expectedDate: "2024-10-01T00:00:00+13:00",
								},
								{
									name: "2024-04-07", testHandler: testLocalMidnight,
									zone: "Antarctica/South_Pole", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-07T00:00:00+13:00", expectedDate: "2024-04-07T00:00:00+13:00",
								},
								{
									name: "2024-04-08", testHandler: testLocalMidnight,
									zone: "Antarctica/South_Pole", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-08T00:00:00+12:00", expectedDate: "2024-04-08T00:00:00+12:00",
								},
								{
									name: "2024-04-09", testHandler: testLocalMidnight,
									zone: "Antarctica/South_Pole", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-09T00:00:00+12:00", expectedDate: "2024-04-09T00:00:00+12:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-09-29", testHandler: testYesterdayMidnight,
									zone: "Antarctica/South_Pole", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-09-29T00:00:00+12:00", expectedDate: "2024-09-29T00:00:00+12:00",
								},
								{
									name: "2024-09-30", testHandler: testYesterdayMidnight,
									zone: "Antarctica/South_Pole", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-09-30T00:00:00+13:00", expectedDate: "2024-09-30T00:00:00+13:00",
								},
								{
									name: "2024-10-01", testHandler: testYesterdayMidnight,
									zone: "Antarctica/South_Pole", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-10-01T00:00:00+13:00", expectedDate: "2024-10-01T00:00:00+13:00",
								},
								{
									name: "2024-04-07", testHandler: testYesterdayMidnight,
									zone: "Antarctica/South_Pole", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-07T00:00:00+13:00", expectedDate: "2024-04-07T00:00:00+13:00",
								},
								{
									name: "2024-04-08", testHandler: testYesterdayMidnight,
									zone: "Antarctica/South_Pole", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-08T00:00:00+12:00", expectedDate: "2024-04-08T00:00:00+12:00",
								},
								{
									name: "2024-04-09", testHandler: testYesterdayMidnight,
									zone: "Antarctica/South_Pole", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-09T00:00:00+12:00", expectedDate: "2024-04-09T00:00:00+12:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-09-29", testHandler: testTomorrowMidnight,
									zone: "Antarctica/South_Pole", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-09-29T00:00:00+12:00", expectedDate: "2024-09-29T00:00:00+12:00",
								},
								{
									name: "2024-09-30", testHandler: testTomorrowMidnight,
									zone: "Antarctica/South_Pole", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-09-30T00:00:00+13:00", expectedDate: "2024-09-30T00:00:00+13:00",
								},
								{
									name: "2024-10-01", testHandler: testTomorrowMidnight,
									zone: "Antarctica/South_Pole", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-10-01T00:00:00+13:00", expectedDate: "2024-10-01T00:00:00+13:00",
								},
								{
									name: "2024-04-07", testHandler: testTomorrowMidnight,
									zone: "Antarctica/South_Pole", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-07T00:00:00+13:00", expectedDate: "2024-04-07T00:00:00+13:00",
								},
								{
									name: "2024-04-08", testHandler: testTomorrowMidnight,
									zone: "Antarctica/South_Pole", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-08T00:00:00+12:00", expectedDate: "2024-04-08T00:00:00+12:00",
								},
								{
									name: "2024-04-09", testHandler: testTomorrowMidnight,
									zone: "Antarctica/South_Pole", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-09T00:00:00+12:00", expectedDate: "2024-04-09T00:00:00+12:00",
								},
							}},
					},
				},
				{
					name: "Troll",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Antarctica/Troll", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Antarctica/Troll", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Antarctica/Troll", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Antarctica/Troll", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Antarctica/Troll", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Antarctica/Troll", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Antarctica/Troll", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Antarctica/Troll", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Antarctica/Troll", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Antarctica/Troll", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Antarctica/Troll", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Antarctica/Troll", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Antarctica/Troll", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Antarctica/Troll", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Antarctica/Troll", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Antarctica/Troll", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Antarctica/Troll", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Antarctica/Troll", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
					},
				},
			},
		},
		{
			name: "Arctic",
			children: []testCase{
				{
					name: "Longyearbyen",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Arctic/Longyearbyen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Arctic/Longyearbyen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Arctic/Longyearbyen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Arctic/Longyearbyen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Arctic/Longyearbyen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Arctic/Longyearbyen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Arctic/Longyearbyen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Arctic/Longyearbyen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Arctic/Longyearbyen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Arctic/Longyearbyen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Arctic/Longyearbyen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Arctic/Longyearbyen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Arctic/Longyearbyen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Arctic/Longyearbyen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Arctic/Longyearbyen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Arctic/Longyearbyen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Arctic/Longyearbyen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Arctic/Longyearbyen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
			},
		},
		{
			name: "Asia",
			children: []testCase{
				{
					name: "Beirut",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-30", testHandler: testLocalMidnight,
									zone: "Asia/Beirut", dstHandover: "2024-03-31T01:00:00+03:00",
									date: "2024-03-30T01:00:00+02:00", expectedDate: "2024-03-30T01:00:00+02:00",
								},
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Asia/Beirut", dstHandover: "2024-03-31T01:00:00+03:00",
									date: "2024-03-31T01:00:00+03:00", expectedDate: "2024-03-31T01:00:00+03:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Asia/Beirut", dstHandover: "2024-03-31T01:00:00+03:00",
									date: "2024-04-01T01:00:00+03:00", expectedDate: "2024-04-01T01:00:00+03:00",
								},
								{
									name: "2024-10-26", testHandler: testLocalMidnight,
									zone: "Asia/Beirut", dstHandover: "2024-10-27T01:00:00+02:00",
									date: "2024-10-26T01:00:00+03:00", expectedDate: "2024-10-26T01:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Asia/Beirut", dstHandover: "2024-10-27T01:00:00+02:00",
									date: "2024-10-27T01:00:00+02:00", expectedDate: "2024-10-27T01:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Asia/Beirut", dstHandover: "2024-10-27T01:00:00+02:00",
									date: "2024-10-28T01:00:00+02:00", expectedDate: "2024-10-28T01:00:00+02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-30", testHandler: testYesterdayMidnight,
									zone: "Asia/Beirut", dstHandover: "2024-03-31T01:00:00+03:00",
									date: "2024-03-30T01:00:00+02:00", expectedDate: "2024-03-30T01:00:00+02:00",
								},
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Asia/Beirut", dstHandover: "2024-03-31T01:00:00+03:00",
									date: "2024-03-31T01:00:00+03:00", expectedDate: "2024-03-31T01:00:00+03:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Asia/Beirut", dstHandover: "2024-03-31T01:00:00+03:00",
									date: "2024-04-01T01:00:00+03:00", expectedDate: "2024-04-01T01:00:00+03:00",
								},
								{
									name: "2024-10-26", testHandler: testYesterdayMidnight,
									zone: "Asia/Beirut", dstHandover: "2024-10-27T01:00:00+02:00",
									date: "2024-10-26T01:00:00+03:00", expectedDate: "2024-10-26T01:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Asia/Beirut", dstHandover: "2024-10-27T01:00:00+02:00",
									date: "2024-10-27T01:00:00+02:00", expectedDate: "2024-10-27T01:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Asia/Beirut", dstHandover: "2024-10-27T01:00:00+02:00",
									date: "2024-10-28T01:00:00+02:00", expectedDate: "2024-10-28T01:00:00+02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-30", testHandler: testTomorrowMidnight,
									zone: "Asia/Beirut", dstHandover: "2024-03-31T01:00:00+03:00",
									date: "2024-03-30T01:00:00+02:00", expectedDate: "2024-03-30T01:00:00+02:00",
								},
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Asia/Beirut", dstHandover: "2024-03-31T01:00:00+03:00",
									date: "2024-03-31T01:00:00+03:00", expectedDate: "2024-03-31T01:00:00+03:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Asia/Beirut", dstHandover: "2024-03-31T01:00:00+03:00",
									date: "2024-04-01T01:00:00+03:00", expectedDate: "2024-04-01T01:00:00+03:00",
								},
								{
									name: "2024-10-26", testHandler: testTomorrowMidnight,
									zone: "Asia/Beirut", dstHandover: "2024-10-27T01:00:00+02:00",
									date: "2024-10-26T01:00:00+03:00", expectedDate: "2024-10-26T01:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Asia/Beirut", dstHandover: "2024-10-27T01:00:00+02:00",
									date: "2024-10-27T01:00:00+02:00", expectedDate: "2024-10-27T01:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Asia/Beirut", dstHandover: "2024-10-27T01:00:00+02:00",
									date: "2024-10-28T01:00:00+02:00", expectedDate: "2024-10-28T01:00:00+02:00",
								},
							}},
					},
				},
				{
					name: "Famagusta",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Asia/Famagusta", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Asia/Famagusta", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Asia/Famagusta", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Asia/Famagusta", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Asia/Famagusta", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Asia/Famagusta", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Asia/Famagusta", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Asia/Famagusta", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Asia/Famagusta", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Asia/Famagusta", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Asia/Famagusta", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Asia/Famagusta", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Asia/Famagusta", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Asia/Famagusta", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Asia/Famagusta", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Asia/Famagusta", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Asia/Famagusta", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Asia/Famagusta", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
					},
				},
				{
					name: "Gaza",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-04-20", testHandler: testLocalMidnight,
									zone: "Asia/Gaza", dstHandover: "2024-04-21T00:00:00+03:00",
									date: "2024-04-20T00:00:00+02:00", expectedDate: "2024-04-20T00:00:00+02:00",
								},
								{
									name: "2024-04-21", testHandler: testLocalMidnight,
									zone: "Asia/Gaza", dstHandover: "2024-04-21T00:00:00+03:00",
									date: "2024-04-21T00:00:00+03:00", expectedDate: "2024-04-21T00:00:00+03:00",
								},
								{
									name: "2024-04-22", testHandler: testLocalMidnight,
									zone: "Asia/Gaza", dstHandover: "2024-04-21T00:00:00+03:00",
									date: "2024-04-22T00:00:00+03:00", expectedDate: "2024-04-22T00:00:00+03:00",
								},
								{
									name: "2024-10-26", testHandler: testLocalMidnight,
									zone: "Asia/Gaza", dstHandover: "2024-10-27T00:00:00+02:00",
									date: "2024-10-26T00:00:00+03:00", expectedDate: "2024-10-26T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Asia/Gaza", dstHandover: "2024-10-27T00:00:00+02:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Asia/Gaza", dstHandover: "2024-10-27T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-04-20", testHandler: testYesterdayMidnight,
									zone: "Asia/Gaza", dstHandover: "2024-04-21T00:00:00+03:00",
									date: "2024-04-20T00:00:00+02:00", expectedDate: "2024-04-20T00:00:00+02:00",
								},
								{
									name: "2024-04-21", testHandler: testYesterdayMidnight,
									zone: "Asia/Gaza", dstHandover: "2024-04-21T00:00:00+03:00",
									date: "2024-04-21T00:00:00+03:00", expectedDate: "2024-04-21T00:00:00+03:00",
								},
								{
									name: "2024-04-22", testHandler: testYesterdayMidnight,
									zone: "Asia/Gaza", dstHandover: "2024-04-21T00:00:00+03:00",
									date: "2024-04-22T00:00:00+03:00", expectedDate: "2024-04-22T00:00:00+03:00",
								},
								{
									name: "2024-10-26", testHandler: testYesterdayMidnight,
									zone: "Asia/Gaza", dstHandover: "2024-10-27T00:00:00+02:00",
									date: "2024-10-26T00:00:00+03:00", expectedDate: "2024-10-26T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Asia/Gaza", dstHandover: "2024-10-27T00:00:00+02:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Asia/Gaza", dstHandover: "2024-10-27T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-04-20", testHandler: testTomorrowMidnight,
									zone: "Asia/Gaza", dstHandover: "2024-04-21T00:00:00+03:00",
									date: "2024-04-20T00:00:00+02:00", expectedDate: "2024-04-20T00:00:00+02:00",
								},
								{
									name: "2024-04-21", testHandler: testTomorrowMidnight,
									zone: "Asia/Gaza", dstHandover: "2024-04-21T00:00:00+03:00",
									date: "2024-04-21T00:00:00+03:00", expectedDate: "2024-04-21T00:00:00+03:00",
								},
								{
									name: "2024-04-22", testHandler: testTomorrowMidnight,
									zone: "Asia/Gaza", dstHandover: "2024-04-21T00:00:00+03:00",
									date: "2024-04-22T00:00:00+03:00", expectedDate: "2024-04-22T00:00:00+03:00",
								},
								{
									name: "2024-10-26", testHandler: testTomorrowMidnight,
									zone: "Asia/Gaza", dstHandover: "2024-10-27T00:00:00+02:00",
									date: "2024-10-26T00:00:00+03:00", expectedDate: "2024-10-26T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Asia/Gaza", dstHandover: "2024-10-27T00:00:00+02:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Asia/Gaza", dstHandover: "2024-10-27T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
							}},
					},
				},
				{
					name: "Hebron",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-04-20", testHandler: testLocalMidnight,
									zone: "Asia/Hebron", dstHandover: "2024-04-21T00:00:00+03:00",
									date: "2024-04-20T00:00:00+02:00", expectedDate: "2024-04-20T00:00:00+02:00",
								},
								{
									name: "2024-04-21", testHandler: testLocalMidnight,
									zone: "Asia/Hebron", dstHandover: "2024-04-21T00:00:00+03:00",
									date: "2024-04-21T00:00:00+03:00", expectedDate: "2024-04-21T00:00:00+03:00",
								},
								{
									name: "2024-04-22", testHandler: testLocalMidnight,
									zone: "Asia/Hebron", dstHandover: "2024-04-21T00:00:00+03:00",
									date: "2024-04-22T00:00:00+03:00", expectedDate: "2024-04-22T00:00:00+03:00",
								},
								{
									name: "2024-10-26", testHandler: testLocalMidnight,
									zone: "Asia/Hebron", dstHandover: "2024-10-27T00:00:00+02:00",
									date: "2024-10-26T00:00:00+03:00", expectedDate: "2024-10-26T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Asia/Hebron", dstHandover: "2024-10-27T00:00:00+02:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Asia/Hebron", dstHandover: "2024-10-27T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-04-20", testHandler: testYesterdayMidnight,
									zone: "Asia/Hebron", dstHandover: "2024-04-21T00:00:00+03:00",
									date: "2024-04-20T00:00:00+02:00", expectedDate: "2024-04-20T00:00:00+02:00",
								},
								{
									name: "2024-04-21", testHandler: testYesterdayMidnight,
									zone: "Asia/Hebron", dstHandover: "2024-04-21T00:00:00+03:00",
									date: "2024-04-21T00:00:00+03:00", expectedDate: "2024-04-21T00:00:00+03:00",
								},
								{
									name: "2024-04-22", testHandler: testYesterdayMidnight,
									zone: "Asia/Hebron", dstHandover: "2024-04-21T00:00:00+03:00",
									date: "2024-04-22T00:00:00+03:00", expectedDate: "2024-04-22T00:00:00+03:00",
								},
								{
									name: "2024-10-26", testHandler: testYesterdayMidnight,
									zone: "Asia/Hebron", dstHandover: "2024-10-27T00:00:00+02:00",
									date: "2024-10-26T00:00:00+03:00", expectedDate: "2024-10-26T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Asia/Hebron", dstHandover: "2024-10-27T00:00:00+02:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Asia/Hebron", dstHandover: "2024-10-27T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-04-20", testHandler: testTomorrowMidnight,
									zone: "Asia/Hebron", dstHandover: "2024-04-21T00:00:00+03:00",
									date: "2024-04-20T00:00:00+02:00", expectedDate: "2024-04-20T00:00:00+02:00",
								},
								{
									name: "2024-04-21", testHandler: testTomorrowMidnight,
									zone: "Asia/Hebron", dstHandover: "2024-04-21T00:00:00+03:00",
									date: "2024-04-21T00:00:00+03:00", expectedDate: "2024-04-21T00:00:00+03:00",
								},
								{
									name: "2024-04-22", testHandler: testTomorrowMidnight,
									zone: "Asia/Hebron", dstHandover: "2024-04-21T00:00:00+03:00",
									date: "2024-04-22T00:00:00+03:00", expectedDate: "2024-04-22T00:00:00+03:00",
								},
								{
									name: "2024-10-26", testHandler: testTomorrowMidnight,
									zone: "Asia/Hebron", dstHandover: "2024-10-27T00:00:00+02:00",
									date: "2024-10-26T00:00:00+03:00", expectedDate: "2024-10-26T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Asia/Hebron", dstHandover: "2024-10-27T00:00:00+02:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Asia/Hebron", dstHandover: "2024-10-27T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
							}},
					},
				},
				{
					name: "Jerusalem",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-29", testHandler: testLocalMidnight,
									zone: "Asia/Jerusalem", dstHandover: "2024-03-30T00:00:00+03:00",
									date: "2024-03-29T00:00:00+02:00", expectedDate: "2024-03-29T00:00:00+02:00",
								},
								{
									name: "2024-03-30", testHandler: testLocalMidnight,
									zone: "Asia/Jerusalem", dstHandover: "2024-03-30T00:00:00+03:00",
									date: "2024-03-30T00:00:00+03:00", expectedDate: "2024-03-30T00:00:00+03:00",
								},
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Asia/Jerusalem", dstHandover: "2024-03-30T00:00:00+03:00",
									date: "2024-03-31T00:00:00+03:00", expectedDate: "2024-03-31T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Asia/Jerusalem", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Asia/Jerusalem", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Asia/Jerusalem", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-29", testHandler: testYesterdayMidnight,
									zone: "Asia/Jerusalem", dstHandover: "2024-03-30T00:00:00+03:00",
									date: "2024-03-29T00:00:00+02:00", expectedDate: "2024-03-29T00:00:00+02:00",
								},
								{
									name: "2024-03-30", testHandler: testYesterdayMidnight,
									zone: "Asia/Jerusalem", dstHandover: "2024-03-30T00:00:00+03:00",
									date: "2024-03-30T00:00:00+03:00", expectedDate: "2024-03-30T00:00:00+03:00",
								},
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Asia/Jerusalem", dstHandover: "2024-03-30T00:00:00+03:00",
									date: "2024-03-31T00:00:00+03:00", expectedDate: "2024-03-31T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Asia/Jerusalem", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Asia/Jerusalem", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Asia/Jerusalem", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-29", testHandler: testTomorrowMidnight,
									zone: "Asia/Jerusalem", dstHandover: "2024-03-30T00:00:00+03:00",
									date: "2024-03-29T00:00:00+02:00", expectedDate: "2024-03-29T00:00:00+02:00",
								},
								{
									name: "2024-03-30", testHandler: testTomorrowMidnight,
									zone: "Asia/Jerusalem", dstHandover: "2024-03-30T00:00:00+03:00",
									date: "2024-03-30T00:00:00+03:00", expectedDate: "2024-03-30T00:00:00+03:00",
								},
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Asia/Jerusalem", dstHandover: "2024-03-30T00:00:00+03:00",
									date: "2024-03-31T00:00:00+03:00", expectedDate: "2024-03-31T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Asia/Jerusalem", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Asia/Jerusalem", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Asia/Jerusalem", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
					},
				},
				{
					name: "Nicosia",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Asia/Nicosia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Asia/Nicosia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Asia/Nicosia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Asia/Nicosia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Asia/Nicosia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Asia/Nicosia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Asia/Nicosia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Asia/Nicosia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Asia/Nicosia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Asia/Nicosia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Asia/Nicosia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Asia/Nicosia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Asia/Nicosia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Asia/Nicosia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Asia/Nicosia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Asia/Nicosia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Asia/Nicosia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Asia/Nicosia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
					},
				},
				{
					name: "Tel_Aviv",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-29", testHandler: testLocalMidnight,
									zone: "Asia/Tel_Aviv", dstHandover: "2024-03-30T00:00:00+03:00",
									date: "2024-03-29T00:00:00+02:00", expectedDate: "2024-03-29T00:00:00+02:00",
								},
								{
									name: "2024-03-30", testHandler: testLocalMidnight,
									zone: "Asia/Tel_Aviv", dstHandover: "2024-03-30T00:00:00+03:00",
									date: "2024-03-30T00:00:00+03:00", expectedDate: "2024-03-30T00:00:00+03:00",
								},
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Asia/Tel_Aviv", dstHandover: "2024-03-30T00:00:00+03:00",
									date: "2024-03-31T00:00:00+03:00", expectedDate: "2024-03-31T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Asia/Tel_Aviv", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Asia/Tel_Aviv", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Asia/Tel_Aviv", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-29", testHandler: testYesterdayMidnight,
									zone: "Asia/Tel_Aviv", dstHandover: "2024-03-30T00:00:00+03:00",
									date: "2024-03-29T00:00:00+02:00", expectedDate: "2024-03-29T00:00:00+02:00",
								},
								{
									name: "2024-03-30", testHandler: testYesterdayMidnight,
									zone: "Asia/Tel_Aviv", dstHandover: "2024-03-30T00:00:00+03:00",
									date: "2024-03-30T00:00:00+03:00", expectedDate: "2024-03-30T00:00:00+03:00",
								},
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Asia/Tel_Aviv", dstHandover: "2024-03-30T00:00:00+03:00",
									date: "2024-03-31T00:00:00+03:00", expectedDate: "2024-03-31T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Asia/Tel_Aviv", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Asia/Tel_Aviv", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Asia/Tel_Aviv", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-29", testHandler: testTomorrowMidnight,
									zone: "Asia/Tel_Aviv", dstHandover: "2024-03-30T00:00:00+03:00",
									date: "2024-03-29T00:00:00+02:00", expectedDate: "2024-03-29T00:00:00+02:00",
								},
								{
									name: "2024-03-30", testHandler: testTomorrowMidnight,
									zone: "Asia/Tel_Aviv", dstHandover: "2024-03-30T00:00:00+03:00",
									date: "2024-03-30T00:00:00+03:00", expectedDate: "2024-03-30T00:00:00+03:00",
								},
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Asia/Tel_Aviv", dstHandover: "2024-03-30T00:00:00+03:00",
									date: "2024-03-31T00:00:00+03:00", expectedDate: "2024-03-31T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Asia/Tel_Aviv", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Asia/Tel_Aviv", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Asia/Tel_Aviv", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
					},
				},
			},
		},
		{
			name: "Atlantic",
			children: []testCase{
				{
					name: "Azores",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-30", testHandler: testLocalMidnight,
									zone: "Atlantic/Azores", dstHandover: "2024-03-31T23:00:00Z",
									date: "2024-03-30T23:00:00-01:00", expectedDate: "2024-03-30T23:00:00-01:00",
								},
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Atlantic/Azores", dstHandover: "2024-03-31T23:00:00Z",
									date: "2024-03-31T23:00:00Z", expectedDate: "2024-03-31T23:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Atlantic/Azores", dstHandover: "2024-03-31T23:00:00Z",
									date: "2024-04-01T23:00:00Z", expectedDate: "2024-04-01T23:00:00Z",
								},
								{
									name: "2024-10-26", testHandler: testLocalMidnight,
									zone: "Atlantic/Azores", dstHandover: "2024-10-27T23:00:00-01:00",
									date: "2024-10-26T23:00:00Z", expectedDate: "2024-10-26T23:00:00Z",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Atlantic/Azores", dstHandover: "2024-10-27T23:00:00-01:00",
									date: "2024-10-27T23:00:00-01:00", expectedDate: "2024-10-27T23:00:00-01:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Atlantic/Azores", dstHandover: "2024-10-27T23:00:00-01:00",
									date: "2024-10-28T23:00:00-01:00", expectedDate: "2024-10-28T23:00:00-01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-30", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Azores", dstHandover: "2024-03-31T23:00:00Z",
									date: "2024-03-30T23:00:00-01:00", expectedDate: "2024-03-30T23:00:00-01:00",
								},
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Azores", dstHandover: "2024-03-31T23:00:00Z",
									date: "2024-03-31T23:00:00Z", expectedDate: "2024-03-31T23:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Azores", dstHandover: "2024-03-31T23:00:00Z",
									date: "2024-04-01T23:00:00Z", expectedDate: "2024-04-01T23:00:00Z",
								},
								{
									name: "2024-10-26", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Azores", dstHandover: "2024-10-27T23:00:00-01:00",
									date: "2024-10-26T23:00:00Z", expectedDate: "2024-10-26T23:00:00Z",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Azores", dstHandover: "2024-10-27T23:00:00-01:00",
									date: "2024-10-27T23:00:00-01:00", expectedDate: "2024-10-27T23:00:00-01:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Azores", dstHandover: "2024-10-27T23:00:00-01:00",
									date: "2024-10-28T23:00:00-01:00", expectedDate: "2024-10-28T23:00:00-01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-30", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Azores", dstHandover: "2024-03-31T23:00:00Z",
									date: "2024-03-30T23:00:00-01:00", expectedDate: "2024-03-30T23:00:00-01:00",
								},
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Azores", dstHandover: "2024-03-31T23:00:00Z",
									date: "2024-03-31T23:00:00Z", expectedDate: "2024-03-31T23:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Azores", dstHandover: "2024-03-31T23:00:00Z",
									date: "2024-04-01T23:00:00Z", expectedDate: "2024-04-01T23:00:00Z",
								},
								{
									name: "2024-10-26", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Azores", dstHandover: "2024-10-27T23:00:00-01:00",
									date: "2024-10-26T23:00:00Z", expectedDate: "2024-10-26T23:00:00Z",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Azores", dstHandover: "2024-10-27T23:00:00-01:00",
									date: "2024-10-27T23:00:00-01:00", expectedDate: "2024-10-27T23:00:00-01:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Azores", dstHandover: "2024-10-27T23:00:00-01:00",
									date: "2024-10-28T23:00:00-01:00", expectedDate: "2024-10-28T23:00:00-01:00",
								},
							}},
					},
				},
				{
					name: "Bermuda",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "Atlantic/Bermuda", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-10T00:00:00-04:00", expectedDate: "2024-03-10T00:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "Atlantic/Bermuda", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-11T00:00:00-03:00", expectedDate: "2024-03-11T00:00:00-03:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "Atlantic/Bermuda", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-12T00:00:00-03:00", expectedDate: "2024-03-12T00:00:00-03:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "Atlantic/Bermuda", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-03T00:00:00-03:00", expectedDate: "2024-11-03T00:00:00-03:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "Atlantic/Bermuda", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-04T00:00:00-04:00", expectedDate: "2024-11-04T00:00:00-04:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "Atlantic/Bermuda", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-05T00:00:00-04:00", expectedDate: "2024-11-05T00:00:00-04:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Bermuda", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-10T00:00:00-04:00", expectedDate: "2024-03-10T00:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Bermuda", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-11T00:00:00-03:00", expectedDate: "2024-03-11T00:00:00-03:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Bermuda", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-12T00:00:00-03:00", expectedDate: "2024-03-12T00:00:00-03:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Bermuda", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-03T00:00:00-03:00", expectedDate: "2024-11-03T00:00:00-03:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Bermuda", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-04T00:00:00-04:00", expectedDate: "2024-11-04T00:00:00-04:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Bermuda", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-05T00:00:00-04:00", expectedDate: "2024-11-05T00:00:00-04:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Bermuda", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-10T00:00:00-04:00", expectedDate: "2024-03-10T00:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Bermuda", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-11T00:00:00-03:00", expectedDate: "2024-03-11T00:00:00-03:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Bermuda", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-12T00:00:00-03:00", expectedDate: "2024-03-12T00:00:00-03:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Bermuda", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-03T00:00:00-03:00", expectedDate: "2024-11-03T00:00:00-03:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Bermuda", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-04T00:00:00-04:00", expectedDate: "2024-11-04T00:00:00-04:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Bermuda", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-05T00:00:00-04:00", expectedDate: "2024-11-05T00:00:00-04:00",
								},
							}},
					},
				},
				{
					name: "Canary",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Atlantic/Canary", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Atlantic/Canary", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Atlantic/Canary", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Atlantic/Canary", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Atlantic/Canary", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Atlantic/Canary", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Canary", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Canary", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Canary", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Canary", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Canary", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Canary", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Canary", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Canary", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Canary", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Canary", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Canary", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Canary", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
					},
				},
				{
					name: "Faeroe",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Atlantic/Faeroe", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Atlantic/Faeroe", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Atlantic/Faeroe", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Atlantic/Faeroe", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Atlantic/Faeroe", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Atlantic/Faeroe", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Faeroe", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Faeroe", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Faeroe", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Faeroe", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Faeroe", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Faeroe", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Faeroe", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Faeroe", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Faeroe", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Faeroe", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Faeroe", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Faeroe", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
					},
				},
				{
					name: "Faroe",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Atlantic/Faroe", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Atlantic/Faroe", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Atlantic/Faroe", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Atlantic/Faroe", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Atlantic/Faroe", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Atlantic/Faroe", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Faroe", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Faroe", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Faroe", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Faroe", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Faroe", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Faroe", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Faroe", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Faroe", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Faroe", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Faroe", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Faroe", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Faroe", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
					},
				},
				{
					name: "Jan_Mayen",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Atlantic/Jan_Mayen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Atlantic/Jan_Mayen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Atlantic/Jan_Mayen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Atlantic/Jan_Mayen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Atlantic/Jan_Mayen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Atlantic/Jan_Mayen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Jan_Mayen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Jan_Mayen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Jan_Mayen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Jan_Mayen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Jan_Mayen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Jan_Mayen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Jan_Mayen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Jan_Mayen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Jan_Mayen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Jan_Mayen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Jan_Mayen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Jan_Mayen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Madeira",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Atlantic/Madeira", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Atlantic/Madeira", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Atlantic/Madeira", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Atlantic/Madeira", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Atlantic/Madeira", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Atlantic/Madeira", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Madeira", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Madeira", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Madeira", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Madeira", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Madeira", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Atlantic/Madeira", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Madeira", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Madeira", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Madeira", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Madeira", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Madeira", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Atlantic/Madeira", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
					},
				},
			},
		},
		{
			name: "Australia",
			children: []testCase{
				{
					name: "ACT",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testLocalMidnight,
									zone: "Australia/ACT", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testLocalMidnight,
									zone: "Australia/ACT", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testLocalMidnight,
									zone: "Australia/ACT", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testLocalMidnight,
									zone: "Australia/ACT", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testLocalMidnight,
									zone: "Australia/ACT", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testLocalMidnight,
									zone: "Australia/ACT", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testYesterdayMidnight,
									zone: "Australia/ACT", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testYesterdayMidnight,
									zone: "Australia/ACT", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testYesterdayMidnight,
									zone: "Australia/ACT", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testYesterdayMidnight,
									zone: "Australia/ACT", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testYesterdayMidnight,
									zone: "Australia/ACT", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testYesterdayMidnight,
									zone: "Australia/ACT", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testTomorrowMidnight,
									zone: "Australia/ACT", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testTomorrowMidnight,
									zone: "Australia/ACT", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testTomorrowMidnight,
									zone: "Australia/ACT", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testTomorrowMidnight,
									zone: "Australia/ACT", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testTomorrowMidnight,
									zone: "Australia/ACT", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testTomorrowMidnight,
									zone: "Australia/ACT", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
					},
				},
				{
					name: "Adelaide",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testLocalMidnight,
									zone: "Australia/Adelaide", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-06T00:00:00+09:30", expectedDate: "2024-10-06T00:00:00+09:30",
								},
								{
									name: "2024-10-07", testHandler: testLocalMidnight,
									zone: "Australia/Adelaide", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-07T00:00:00+10:30", expectedDate: "2024-10-07T00:00:00+10:30",
								},
								{
									name: "2024-10-08", testHandler: testLocalMidnight,
									zone: "Australia/Adelaide", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-08T00:00:00+10:30", expectedDate: "2024-10-08T00:00:00+10:30",
								},
								{
									name: "2024-04-07", testHandler: testLocalMidnight,
									zone: "Australia/Adelaide", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-07T00:00:00+10:30", expectedDate: "2024-04-07T00:00:00+10:30",
								},
								{
									name: "2024-04-08", testHandler: testLocalMidnight,
									zone: "Australia/Adelaide", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-08T00:00:00+09:30", expectedDate: "2024-04-08T00:00:00+09:30",
								},
								{
									name: "2024-04-09", testHandler: testLocalMidnight,
									zone: "Australia/Adelaide", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-09T00:00:00+09:30", expectedDate: "2024-04-09T00:00:00+09:30",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testYesterdayMidnight,
									zone: "Australia/Adelaide", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-06T00:00:00+09:30", expectedDate: "2024-10-06T00:00:00+09:30",
								},
								{
									name: "2024-10-07", testHandler: testYesterdayMidnight,
									zone: "Australia/Adelaide", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-07T00:00:00+10:30", expectedDate: "2024-10-07T00:00:00+10:30",
								},
								{
									name: "2024-10-08", testHandler: testYesterdayMidnight,
									zone: "Australia/Adelaide", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-08T00:00:00+10:30", expectedDate: "2024-10-08T00:00:00+10:30",
								},
								{
									name: "2024-04-07", testHandler: testYesterdayMidnight,
									zone: "Australia/Adelaide", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-07T00:00:00+10:30", expectedDate: "2024-04-07T00:00:00+10:30",
								},
								{
									name: "2024-04-08", testHandler: testYesterdayMidnight,
									zone: "Australia/Adelaide", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-08T00:00:00+09:30", expectedDate: "2024-04-08T00:00:00+09:30",
								},
								{
									name: "2024-04-09", testHandler: testYesterdayMidnight,
									zone: "Australia/Adelaide", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-09T00:00:00+09:30", expectedDate: "2024-04-09T00:00:00+09:30",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testTomorrowMidnight,
									zone: "Australia/Adelaide", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-06T00:00:00+09:30", expectedDate: "2024-10-06T00:00:00+09:30",
								},
								{
									name: "2024-10-07", testHandler: testTomorrowMidnight,
									zone: "Australia/Adelaide", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-07T00:00:00+10:30", expectedDate: "2024-10-07T00:00:00+10:30",
								},
								{
									name: "2024-10-08", testHandler: testTomorrowMidnight,
									zone: "Australia/Adelaide", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-08T00:00:00+10:30", expectedDate: "2024-10-08T00:00:00+10:30",
								},
								{
									name: "2024-04-07", testHandler: testTomorrowMidnight,
									zone: "Australia/Adelaide", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-07T00:00:00+10:30", expectedDate: "2024-04-07T00:00:00+10:30",
								},
								{
									name: "2024-04-08", testHandler: testTomorrowMidnight,
									zone: "Australia/Adelaide", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-08T00:00:00+09:30", expectedDate: "2024-04-08T00:00:00+09:30",
								},
								{
									name: "2024-04-09", testHandler: testTomorrowMidnight,
									zone: "Australia/Adelaide", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-09T00:00:00+09:30", expectedDate: "2024-04-09T00:00:00+09:30",
								},
							}},
					},
				},
				{
					name: "Broken_Hill",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testLocalMidnight,
									zone: "Australia/Broken_Hill", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-06T00:00:00+09:30", expectedDate: "2024-10-06T00:00:00+09:30",
								},
								{
									name: "2024-10-07", testHandler: testLocalMidnight,
									zone: "Australia/Broken_Hill", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-07T00:00:00+10:30", expectedDate: "2024-10-07T00:00:00+10:30",
								},
								{
									name: "2024-10-08", testHandler: testLocalMidnight,
									zone: "Australia/Broken_Hill", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-08T00:00:00+10:30", expectedDate: "2024-10-08T00:00:00+10:30",
								},
								{
									name: "2024-04-07", testHandler: testLocalMidnight,
									zone: "Australia/Broken_Hill", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-07T00:00:00+10:30", expectedDate: "2024-04-07T00:00:00+10:30",
								},
								{
									name: "2024-04-08", testHandler: testLocalMidnight,
									zone: "Australia/Broken_Hill", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-08T00:00:00+09:30", expectedDate: "2024-04-08T00:00:00+09:30",
								},
								{
									name: "2024-04-09", testHandler: testLocalMidnight,
									zone: "Australia/Broken_Hill", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-09T00:00:00+09:30", expectedDate: "2024-04-09T00:00:00+09:30",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testYesterdayMidnight,
									zone: "Australia/Broken_Hill", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-06T00:00:00+09:30", expectedDate: "2024-10-06T00:00:00+09:30",
								},
								{
									name: "2024-10-07", testHandler: testYesterdayMidnight,
									zone: "Australia/Broken_Hill", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-07T00:00:00+10:30", expectedDate: "2024-10-07T00:00:00+10:30",
								},
								{
									name: "2024-10-08", testHandler: testYesterdayMidnight,
									zone: "Australia/Broken_Hill", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-08T00:00:00+10:30", expectedDate: "2024-10-08T00:00:00+10:30",
								},
								{
									name: "2024-04-07", testHandler: testYesterdayMidnight,
									zone: "Australia/Broken_Hill", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-07T00:00:00+10:30", expectedDate: "2024-04-07T00:00:00+10:30",
								},
								{
									name: "2024-04-08", testHandler: testYesterdayMidnight,
									zone: "Australia/Broken_Hill", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-08T00:00:00+09:30", expectedDate: "2024-04-08T00:00:00+09:30",
								},
								{
									name: "2024-04-09", testHandler: testYesterdayMidnight,
									zone: "Australia/Broken_Hill", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-09T00:00:00+09:30", expectedDate: "2024-04-09T00:00:00+09:30",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testTomorrowMidnight,
									zone: "Australia/Broken_Hill", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-06T00:00:00+09:30", expectedDate: "2024-10-06T00:00:00+09:30",
								},
								{
									name: "2024-10-07", testHandler: testTomorrowMidnight,
									zone: "Australia/Broken_Hill", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-07T00:00:00+10:30", expectedDate: "2024-10-07T00:00:00+10:30",
								},
								{
									name: "2024-10-08", testHandler: testTomorrowMidnight,
									zone: "Australia/Broken_Hill", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-08T00:00:00+10:30", expectedDate: "2024-10-08T00:00:00+10:30",
								},
								{
									name: "2024-04-07", testHandler: testTomorrowMidnight,
									zone: "Australia/Broken_Hill", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-07T00:00:00+10:30", expectedDate: "2024-04-07T00:00:00+10:30",
								},
								{
									name: "2024-04-08", testHandler: testTomorrowMidnight,
									zone: "Australia/Broken_Hill", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-08T00:00:00+09:30", expectedDate: "2024-04-08T00:00:00+09:30",
								},
								{
									name: "2024-04-09", testHandler: testTomorrowMidnight,
									zone: "Australia/Broken_Hill", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-09T00:00:00+09:30", expectedDate: "2024-04-09T00:00:00+09:30",
								},
							}},
					},
				},
				{
					name: "Canberra",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testLocalMidnight,
									zone: "Australia/Canberra", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testLocalMidnight,
									zone: "Australia/Canberra", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testLocalMidnight,
									zone: "Australia/Canberra", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testLocalMidnight,
									zone: "Australia/Canberra", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testLocalMidnight,
									zone: "Australia/Canberra", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testLocalMidnight,
									zone: "Australia/Canberra", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testYesterdayMidnight,
									zone: "Australia/Canberra", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testYesterdayMidnight,
									zone: "Australia/Canberra", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testYesterdayMidnight,
									zone: "Australia/Canberra", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testYesterdayMidnight,
									zone: "Australia/Canberra", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testYesterdayMidnight,
									zone: "Australia/Canberra", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testYesterdayMidnight,
									zone: "Australia/Canberra", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testTomorrowMidnight,
									zone: "Australia/Canberra", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testTomorrowMidnight,
									zone: "Australia/Canberra", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testTomorrowMidnight,
									zone: "Australia/Canberra", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testTomorrowMidnight,
									zone: "Australia/Canberra", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testTomorrowMidnight,
									zone: "Australia/Canberra", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testTomorrowMidnight,
									zone: "Australia/Canberra", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
					},
				},
				{
					name: "Currie",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testLocalMidnight,
									zone: "Australia/Currie", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testLocalMidnight,
									zone: "Australia/Currie", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testLocalMidnight,
									zone: "Australia/Currie", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testLocalMidnight,
									zone: "Australia/Currie", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testLocalMidnight,
									zone: "Australia/Currie", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testLocalMidnight,
									zone: "Australia/Currie", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testYesterdayMidnight,
									zone: "Australia/Currie", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testYesterdayMidnight,
									zone: "Australia/Currie", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testYesterdayMidnight,
									zone: "Australia/Currie", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testYesterdayMidnight,
									zone: "Australia/Currie", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testYesterdayMidnight,
									zone: "Australia/Currie", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testYesterdayMidnight,
									zone: "Australia/Currie", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testTomorrowMidnight,
									zone: "Australia/Currie", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testTomorrowMidnight,
									zone: "Australia/Currie", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testTomorrowMidnight,
									zone: "Australia/Currie", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testTomorrowMidnight,
									zone: "Australia/Currie", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testTomorrowMidnight,
									zone: "Australia/Currie", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testTomorrowMidnight,
									zone: "Australia/Currie", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
					},
				},
				{
					name: "Hobart",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testLocalMidnight,
									zone: "Australia/Hobart", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testLocalMidnight,
									zone: "Australia/Hobart", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testLocalMidnight,
									zone: "Australia/Hobart", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testLocalMidnight,
									zone: "Australia/Hobart", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testLocalMidnight,
									zone: "Australia/Hobart", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testLocalMidnight,
									zone: "Australia/Hobart", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testYesterdayMidnight,
									zone: "Australia/Hobart", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testYesterdayMidnight,
									zone: "Australia/Hobart", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testYesterdayMidnight,
									zone: "Australia/Hobart", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testYesterdayMidnight,
									zone: "Australia/Hobart", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testYesterdayMidnight,
									zone: "Australia/Hobart", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testYesterdayMidnight,
									zone: "Australia/Hobart", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testTomorrowMidnight,
									zone: "Australia/Hobart", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testTomorrowMidnight,
									zone: "Australia/Hobart", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testTomorrowMidnight,
									zone: "Australia/Hobart", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testTomorrowMidnight,
									zone: "Australia/Hobart", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testTomorrowMidnight,
									zone: "Australia/Hobart", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testTomorrowMidnight,
									zone: "Australia/Hobart", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
					},
				},
				{
					name: "Melbourne",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testLocalMidnight,
									zone: "Australia/Melbourne", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testLocalMidnight,
									zone: "Australia/Melbourne", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testLocalMidnight,
									zone: "Australia/Melbourne", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testLocalMidnight,
									zone: "Australia/Melbourne", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testLocalMidnight,
									zone: "Australia/Melbourne", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testLocalMidnight,
									zone: "Australia/Melbourne", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testYesterdayMidnight,
									zone: "Australia/Melbourne", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testYesterdayMidnight,
									zone: "Australia/Melbourne", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testYesterdayMidnight,
									zone: "Australia/Melbourne", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testYesterdayMidnight,
									zone: "Australia/Melbourne", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testYesterdayMidnight,
									zone: "Australia/Melbourne", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testYesterdayMidnight,
									zone: "Australia/Melbourne", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testTomorrowMidnight,
									zone: "Australia/Melbourne", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testTomorrowMidnight,
									zone: "Australia/Melbourne", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testTomorrowMidnight,
									zone: "Australia/Melbourne", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testTomorrowMidnight,
									zone: "Australia/Melbourne", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testTomorrowMidnight,
									zone: "Australia/Melbourne", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testTomorrowMidnight,
									zone: "Australia/Melbourne", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
					},
				},
				{
					name: "NSW",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testLocalMidnight,
									zone: "Australia/NSW", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testLocalMidnight,
									zone: "Australia/NSW", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testLocalMidnight,
									zone: "Australia/NSW", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testLocalMidnight,
									zone: "Australia/NSW", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testLocalMidnight,
									zone: "Australia/NSW", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testLocalMidnight,
									zone: "Australia/NSW", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testYesterdayMidnight,
									zone: "Australia/NSW", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testYesterdayMidnight,
									zone: "Australia/NSW", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testYesterdayMidnight,
									zone: "Australia/NSW", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testYesterdayMidnight,
									zone: "Australia/NSW", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testYesterdayMidnight,
									zone: "Australia/NSW", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testYesterdayMidnight,
									zone: "Australia/NSW", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testTomorrowMidnight,
									zone: "Australia/NSW", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testTomorrowMidnight,
									zone: "Australia/NSW", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testTomorrowMidnight,
									zone: "Australia/NSW", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testTomorrowMidnight,
									zone: "Australia/NSW", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testTomorrowMidnight,
									zone: "Australia/NSW", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testTomorrowMidnight,
									zone: "Australia/NSW", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
					},
				},
				{
					name: "South",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testLocalMidnight,
									zone: "Australia/South", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-06T00:00:00+09:30", expectedDate: "2024-10-06T00:00:00+09:30",
								},
								{
									name: "2024-10-07", testHandler: testLocalMidnight,
									zone: "Australia/South", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-07T00:00:00+10:30", expectedDate: "2024-10-07T00:00:00+10:30",
								},
								{
									name: "2024-10-08", testHandler: testLocalMidnight,
									zone: "Australia/South", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-08T00:00:00+10:30", expectedDate: "2024-10-08T00:00:00+10:30",
								},
								{
									name: "2024-04-07", testHandler: testLocalMidnight,
									zone: "Australia/South", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-07T00:00:00+10:30", expectedDate: "2024-04-07T00:00:00+10:30",
								},
								{
									name: "2024-04-08", testHandler: testLocalMidnight,
									zone: "Australia/South", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-08T00:00:00+09:30", expectedDate: "2024-04-08T00:00:00+09:30",
								},
								{
									name: "2024-04-09", testHandler: testLocalMidnight,
									zone: "Australia/South", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-09T00:00:00+09:30", expectedDate: "2024-04-09T00:00:00+09:30",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testYesterdayMidnight,
									zone: "Australia/South", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-06T00:00:00+09:30", expectedDate: "2024-10-06T00:00:00+09:30",
								},
								{
									name: "2024-10-07", testHandler: testYesterdayMidnight,
									zone: "Australia/South", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-07T00:00:00+10:30", expectedDate: "2024-10-07T00:00:00+10:30",
								},
								{
									name: "2024-10-08", testHandler: testYesterdayMidnight,
									zone: "Australia/South", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-08T00:00:00+10:30", expectedDate: "2024-10-08T00:00:00+10:30",
								},
								{
									name: "2024-04-07", testHandler: testYesterdayMidnight,
									zone: "Australia/South", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-07T00:00:00+10:30", expectedDate: "2024-04-07T00:00:00+10:30",
								},
								{
									name: "2024-04-08", testHandler: testYesterdayMidnight,
									zone: "Australia/South", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-08T00:00:00+09:30", expectedDate: "2024-04-08T00:00:00+09:30",
								},
								{
									name: "2024-04-09", testHandler: testYesterdayMidnight,
									zone: "Australia/South", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-09T00:00:00+09:30", expectedDate: "2024-04-09T00:00:00+09:30",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testTomorrowMidnight,
									zone: "Australia/South", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-06T00:00:00+09:30", expectedDate: "2024-10-06T00:00:00+09:30",
								},
								{
									name: "2024-10-07", testHandler: testTomorrowMidnight,
									zone: "Australia/South", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-07T00:00:00+10:30", expectedDate: "2024-10-07T00:00:00+10:30",
								},
								{
									name: "2024-10-08", testHandler: testTomorrowMidnight,
									zone: "Australia/South", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-08T00:00:00+10:30", expectedDate: "2024-10-08T00:00:00+10:30",
								},
								{
									name: "2024-04-07", testHandler: testTomorrowMidnight,
									zone: "Australia/South", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-07T00:00:00+10:30", expectedDate: "2024-04-07T00:00:00+10:30",
								},
								{
									name: "2024-04-08", testHandler: testTomorrowMidnight,
									zone: "Australia/South", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-08T00:00:00+09:30", expectedDate: "2024-04-08T00:00:00+09:30",
								},
								{
									name: "2024-04-09", testHandler: testTomorrowMidnight,
									zone: "Australia/South", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-09T00:00:00+09:30", expectedDate: "2024-04-09T00:00:00+09:30",
								},
							}},
					},
				},
				{
					name: "Sydney",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testLocalMidnight,
									zone: "Australia/Sydney", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testLocalMidnight,
									zone: "Australia/Sydney", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testLocalMidnight,
									zone: "Australia/Sydney", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testLocalMidnight,
									zone: "Australia/Sydney", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testLocalMidnight,
									zone: "Australia/Sydney", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testLocalMidnight,
									zone: "Australia/Sydney", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testYesterdayMidnight,
									zone: "Australia/Sydney", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testYesterdayMidnight,
									zone: "Australia/Sydney", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testYesterdayMidnight,
									zone: "Australia/Sydney", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testYesterdayMidnight,
									zone: "Australia/Sydney", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testYesterdayMidnight,
									zone: "Australia/Sydney", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testYesterdayMidnight,
									zone: "Australia/Sydney", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testTomorrowMidnight,
									zone: "Australia/Sydney", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testTomorrowMidnight,
									zone: "Australia/Sydney", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testTomorrowMidnight,
									zone: "Australia/Sydney", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testTomorrowMidnight,
									zone: "Australia/Sydney", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testTomorrowMidnight,
									zone: "Australia/Sydney", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testTomorrowMidnight,
									zone: "Australia/Sydney", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
					},
				},
				{
					name: "Tasmania",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testLocalMidnight,
									zone: "Australia/Tasmania", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testLocalMidnight,
									zone: "Australia/Tasmania", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testLocalMidnight,
									zone: "Australia/Tasmania", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testLocalMidnight,
									zone: "Australia/Tasmania", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testLocalMidnight,
									zone: "Australia/Tasmania", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testLocalMidnight,
									zone: "Australia/Tasmania", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testYesterdayMidnight,
									zone: "Australia/Tasmania", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testYesterdayMidnight,
									zone: "Australia/Tasmania", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testYesterdayMidnight,
									zone: "Australia/Tasmania", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testYesterdayMidnight,
									zone: "Australia/Tasmania", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testYesterdayMidnight,
									zone: "Australia/Tasmania", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testYesterdayMidnight,
									zone: "Australia/Tasmania", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testTomorrowMidnight,
									zone: "Australia/Tasmania", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testTomorrowMidnight,
									zone: "Australia/Tasmania", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testTomorrowMidnight,
									zone: "Australia/Tasmania", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testTomorrowMidnight,
									zone: "Australia/Tasmania", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testTomorrowMidnight,
									zone: "Australia/Tasmania", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testTomorrowMidnight,
									zone: "Australia/Tasmania", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
					},
				},
				{
					name: "Victoria",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testLocalMidnight,
									zone: "Australia/Victoria", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testLocalMidnight,
									zone: "Australia/Victoria", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testLocalMidnight,
									zone: "Australia/Victoria", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testLocalMidnight,
									zone: "Australia/Victoria", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testLocalMidnight,
									zone: "Australia/Victoria", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testLocalMidnight,
									zone: "Australia/Victoria", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testYesterdayMidnight,
									zone: "Australia/Victoria", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testYesterdayMidnight,
									zone: "Australia/Victoria", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testYesterdayMidnight,
									zone: "Australia/Victoria", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testYesterdayMidnight,
									zone: "Australia/Victoria", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testYesterdayMidnight,
									zone: "Australia/Victoria", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testYesterdayMidnight,
									zone: "Australia/Victoria", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testTomorrowMidnight,
									zone: "Australia/Victoria", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-06T00:00:00+10:00", expectedDate: "2024-10-06T00:00:00+10:00",
								},
								{
									name: "2024-10-07", testHandler: testTomorrowMidnight,
									zone: "Australia/Victoria", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-07T00:00:00+11:00", expectedDate: "2024-10-07T00:00:00+11:00",
								},
								{
									name: "2024-10-08", testHandler: testTomorrowMidnight,
									zone: "Australia/Victoria", dstHandover: "2024-10-07T00:00:00+11:00",
									date: "2024-10-08T00:00:00+11:00", expectedDate: "2024-10-08T00:00:00+11:00",
								},
								{
									name: "2024-04-07", testHandler: testTomorrowMidnight,
									zone: "Australia/Victoria", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-07T00:00:00+11:00", expectedDate: "2024-04-07T00:00:00+11:00",
								},
								{
									name: "2024-04-08", testHandler: testTomorrowMidnight,
									zone: "Australia/Victoria", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-08T00:00:00+10:00", expectedDate: "2024-04-08T00:00:00+10:00",
								},
								{
									name: "2024-04-09", testHandler: testTomorrowMidnight,
									zone: "Australia/Victoria", dstHandover: "2024-04-08T00:00:00+10:00",
									date: "2024-04-09T00:00:00+10:00", expectedDate: "2024-04-09T00:00:00+10:00",
								},
							}},
					},
				},
				{
					name: "Yancowinna",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testLocalMidnight,
									zone: "Australia/Yancowinna", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-06T00:00:00+09:30", expectedDate: "2024-10-06T00:00:00+09:30",
								},
								{
									name: "2024-10-07", testHandler: testLocalMidnight,
									zone: "Australia/Yancowinna", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-07T00:00:00+10:30", expectedDate: "2024-10-07T00:00:00+10:30",
								},
								{
									name: "2024-10-08", testHandler: testLocalMidnight,
									zone: "Australia/Yancowinna", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-08T00:00:00+10:30", expectedDate: "2024-10-08T00:00:00+10:30",
								},
								{
									name: "2024-04-07", testHandler: testLocalMidnight,
									zone: "Australia/Yancowinna", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-07T00:00:00+10:30", expectedDate: "2024-04-07T00:00:00+10:30",
								},
								{
									name: "2024-04-08", testHandler: testLocalMidnight,
									zone: "Australia/Yancowinna", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-08T00:00:00+09:30", expectedDate: "2024-04-08T00:00:00+09:30",
								},
								{
									name: "2024-04-09", testHandler: testLocalMidnight,
									zone: "Australia/Yancowinna", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-09T00:00:00+09:30", expectedDate: "2024-04-09T00:00:00+09:30",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testYesterdayMidnight,
									zone: "Australia/Yancowinna", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-06T00:00:00+09:30", expectedDate: "2024-10-06T00:00:00+09:30",
								},
								{
									name: "2024-10-07", testHandler: testYesterdayMidnight,
									zone: "Australia/Yancowinna", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-07T00:00:00+10:30", expectedDate: "2024-10-07T00:00:00+10:30",
								},
								{
									name: "2024-10-08", testHandler: testYesterdayMidnight,
									zone: "Australia/Yancowinna", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-08T00:00:00+10:30", expectedDate: "2024-10-08T00:00:00+10:30",
								},
								{
									name: "2024-04-07", testHandler: testYesterdayMidnight,
									zone: "Australia/Yancowinna", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-07T00:00:00+10:30", expectedDate: "2024-04-07T00:00:00+10:30",
								},
								{
									name: "2024-04-08", testHandler: testYesterdayMidnight,
									zone: "Australia/Yancowinna", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-08T00:00:00+09:30", expectedDate: "2024-04-08T00:00:00+09:30",
								},
								{
									name: "2024-04-09", testHandler: testYesterdayMidnight,
									zone: "Australia/Yancowinna", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-09T00:00:00+09:30", expectedDate: "2024-04-09T00:00:00+09:30",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testTomorrowMidnight,
									zone: "Australia/Yancowinna", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-06T00:00:00+09:30", expectedDate: "2024-10-06T00:00:00+09:30",
								},
								{
									name: "2024-10-07", testHandler: testTomorrowMidnight,
									zone: "Australia/Yancowinna", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-07T00:00:00+10:30", expectedDate: "2024-10-07T00:00:00+10:30",
								},
								{
									name: "2024-10-08", testHandler: testTomorrowMidnight,
									zone: "Australia/Yancowinna", dstHandover: "2024-10-07T00:00:00+10:30",
									date: "2024-10-08T00:00:00+10:30", expectedDate: "2024-10-08T00:00:00+10:30",
								},
								{
									name: "2024-04-07", testHandler: testTomorrowMidnight,
									zone: "Australia/Yancowinna", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-07T00:00:00+10:30", expectedDate: "2024-04-07T00:00:00+10:30",
								},
								{
									name: "2024-04-08", testHandler: testTomorrowMidnight,
									zone: "Australia/Yancowinna", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-08T00:00:00+09:30", expectedDate: "2024-04-08T00:00:00+09:30",
								},
								{
									name: "2024-04-09", testHandler: testTomorrowMidnight,
									zone: "Australia/Yancowinna", dstHandover: "2024-04-08T00:00:00+09:30",
									date: "2024-04-09T00:00:00+09:30", expectedDate: "2024-04-09T00:00:00+09:30",
								},
							}},
					},
				},
			},
		},
		{
			name: "CET",
			children: []testCase{
				{
					name: "LocalMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testLocalMidnight,
							zone: "CET", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
						},
						{
							name: "2024-04-01", testHandler: testLocalMidnight,
							zone: "CET", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
						},
						{
							name: "2024-04-02", testHandler: testLocalMidnight,
							zone: "CET", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
						},
						{
							name: "2024-10-27", testHandler: testLocalMidnight,
							zone: "CET", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
						},
						{
							name: "2024-10-28", testHandler: testLocalMidnight,
							zone: "CET", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
						},
						{
							name: "2024-10-29", testHandler: testLocalMidnight,
							zone: "CET", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
						},
					}},
				{
					name: "YesterdayMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testYesterdayMidnight,
							zone: "CET", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
						},
						{
							name: "2024-04-01", testHandler: testYesterdayMidnight,
							zone: "CET", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
						},
						{
							name: "2024-04-02", testHandler: testYesterdayMidnight,
							zone: "CET", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
						},
						{
							name: "2024-10-27", testHandler: testYesterdayMidnight,
							zone: "CET", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
						},
						{
							name: "2024-10-28", testHandler: testYesterdayMidnight,
							zone: "CET", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
						},
						{
							name: "2024-10-29", testHandler: testYesterdayMidnight,
							zone: "CET", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
						},
					}},
				{
					name: "TomorrowMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testTomorrowMidnight,
							zone: "CET", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
						},
						{
							name: "2024-04-01", testHandler: testTomorrowMidnight,
							zone: "CET", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
						},
						{
							name: "2024-04-02", testHandler: testTomorrowMidnight,
							zone: "CET", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
						},
						{
							name: "2024-10-27", testHandler: testTomorrowMidnight,
							zone: "CET", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
						},
						{
							name: "2024-10-28", testHandler: testTomorrowMidnight,
							zone: "CET", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
						},
						{
							name: "2024-10-29", testHandler: testTomorrowMidnight,
							zone: "CET", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
						},
					}},
			},
		},
		{
			name: "CST6CDT",
			children: []testCase{
				{
					name: "LocalMidnight", children: []testCase{
						{
							name: "2024-03-10", testHandler: testLocalMidnight,
							zone: "CST6CDT", dstHandover: "2024-03-11T00:00:00-05:00",
							date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
						},
						{
							name: "2024-03-11", testHandler: testLocalMidnight,
							zone: "CST6CDT", dstHandover: "2024-03-11T00:00:00-05:00",
							date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
						},
						{
							name: "2024-03-12", testHandler: testLocalMidnight,
							zone: "CST6CDT", dstHandover: "2024-03-11T00:00:00-05:00",
							date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
						},
						{
							name: "2024-11-03", testHandler: testLocalMidnight,
							zone: "CST6CDT", dstHandover: "2024-11-04T00:00:00-06:00",
							date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
						},
						{
							name: "2024-11-04", testHandler: testLocalMidnight,
							zone: "CST6CDT", dstHandover: "2024-11-04T00:00:00-06:00",
							date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
						},
						{
							name: "2024-11-05", testHandler: testLocalMidnight,
							zone: "CST6CDT", dstHandover: "2024-11-04T00:00:00-06:00",
							date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
						},
					}},
				{
					name: "YesterdayMidnight", children: []testCase{
						{
							name: "2024-03-10", testHandler: testYesterdayMidnight,
							zone: "CST6CDT", dstHandover: "2024-03-11T00:00:00-05:00",
							date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
						},
						{
							name: "2024-03-11", testHandler: testYesterdayMidnight,
							zone: "CST6CDT", dstHandover: "2024-03-11T00:00:00-05:00",
							date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
						},
						{
							name: "2024-03-12", testHandler: testYesterdayMidnight,
							zone: "CST6CDT", dstHandover: "2024-03-11T00:00:00-05:00",
							date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
						},
						{
							name: "2024-11-03", testHandler: testYesterdayMidnight,
							zone: "CST6CDT", dstHandover: "2024-11-04T00:00:00-06:00",
							date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
						},
						{
							name: "2024-11-04", testHandler: testYesterdayMidnight,
							zone: "CST6CDT", dstHandover: "2024-11-04T00:00:00-06:00",
							date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
						},
						{
							name: "2024-11-05", testHandler: testYesterdayMidnight,
							zone: "CST6CDT", dstHandover: "2024-11-04T00:00:00-06:00",
							date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
						},
					}},
				{
					name: "TomorrowMidnight", children: []testCase{
						{
							name: "2024-03-10", testHandler: testTomorrowMidnight,
							zone: "CST6CDT", dstHandover: "2024-03-11T00:00:00-05:00",
							date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
						},
						{
							name: "2024-03-11", testHandler: testTomorrowMidnight,
							zone: "CST6CDT", dstHandover: "2024-03-11T00:00:00-05:00",
							date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
						},
						{
							name: "2024-03-12", testHandler: testTomorrowMidnight,
							zone: "CST6CDT", dstHandover: "2024-03-11T00:00:00-05:00",
							date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
						},
						{
							name: "2024-11-03", testHandler: testTomorrowMidnight,
							zone: "CST6CDT", dstHandover: "2024-11-04T00:00:00-06:00",
							date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
						},
						{
							name: "2024-11-04", testHandler: testTomorrowMidnight,
							zone: "CST6CDT", dstHandover: "2024-11-04T00:00:00-06:00",
							date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
						},
						{
							name: "2024-11-05", testHandler: testTomorrowMidnight,
							zone: "CST6CDT", dstHandover: "2024-11-04T00:00:00-06:00",
							date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
						},
					}},
			},
		},
		{
			name: "Canada",
			children: []testCase{
				{
					name: "Atlantic",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "Canada/Atlantic", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-10T00:00:00-04:00", expectedDate: "2024-03-10T00:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "Canada/Atlantic", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-11T00:00:00-03:00", expectedDate: "2024-03-11T00:00:00-03:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "Canada/Atlantic", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-12T00:00:00-03:00", expectedDate: "2024-03-12T00:00:00-03:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "Canada/Atlantic", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-03T00:00:00-03:00", expectedDate: "2024-11-03T00:00:00-03:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "Canada/Atlantic", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-04T00:00:00-04:00", expectedDate: "2024-11-04T00:00:00-04:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "Canada/Atlantic", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-05T00:00:00-04:00", expectedDate: "2024-11-05T00:00:00-04:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "Canada/Atlantic", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-10T00:00:00-04:00", expectedDate: "2024-03-10T00:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "Canada/Atlantic", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-11T00:00:00-03:00", expectedDate: "2024-03-11T00:00:00-03:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "Canada/Atlantic", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-12T00:00:00-03:00", expectedDate: "2024-03-12T00:00:00-03:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "Canada/Atlantic", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-03T00:00:00-03:00", expectedDate: "2024-11-03T00:00:00-03:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "Canada/Atlantic", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-04T00:00:00-04:00", expectedDate: "2024-11-04T00:00:00-04:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "Canada/Atlantic", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-05T00:00:00-04:00", expectedDate: "2024-11-05T00:00:00-04:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "Canada/Atlantic", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-10T00:00:00-04:00", expectedDate: "2024-03-10T00:00:00-04:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "Canada/Atlantic", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-11T00:00:00-03:00", expectedDate: "2024-03-11T00:00:00-03:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "Canada/Atlantic", dstHandover: "2024-03-11T00:00:00-03:00",
									date: "2024-03-12T00:00:00-03:00", expectedDate: "2024-03-12T00:00:00-03:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "Canada/Atlantic", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-03T00:00:00-03:00", expectedDate: "2024-11-03T00:00:00-03:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "Canada/Atlantic", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-04T00:00:00-04:00", expectedDate: "2024-11-04T00:00:00-04:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "Canada/Atlantic", dstHandover: "2024-11-04T00:00:00-04:00",
									date: "2024-11-05T00:00:00-04:00", expectedDate: "2024-11-05T00:00:00-04:00",
								},
							}},
					},
				},
				{
					name: "Central",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "Canada/Central", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "Canada/Central", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "Canada/Central", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "Canada/Central", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "Canada/Central", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "Canada/Central", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "Canada/Central", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "Canada/Central", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "Canada/Central", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "Canada/Central", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "Canada/Central", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "Canada/Central", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "Canada/Central", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "Canada/Central", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "Canada/Central", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "Canada/Central", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "Canada/Central", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "Canada/Central", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
					},
				},
				{
					name: "Eastern",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "Canada/Eastern", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "Canada/Eastern", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "Canada/Eastern", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "Canada/Eastern", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "Canada/Eastern", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "Canada/Eastern", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "Canada/Eastern", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "Canada/Eastern", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "Canada/Eastern", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "Canada/Eastern", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "Canada/Eastern", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "Canada/Eastern", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "Canada/Eastern", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "Canada/Eastern", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "Canada/Eastern", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "Canada/Eastern", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "Canada/Eastern", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "Canada/Eastern", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
					},
				},
				{
					name: "Mountain",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "Canada/Mountain", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "Canada/Mountain", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "Canada/Mountain", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "Canada/Mountain", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "Canada/Mountain", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "Canada/Mountain", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "Canada/Mountain", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "Canada/Mountain", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "Canada/Mountain", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "Canada/Mountain", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "Canada/Mountain", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "Canada/Mountain", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "Canada/Mountain", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "Canada/Mountain", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "Canada/Mountain", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "Canada/Mountain", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "Canada/Mountain", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "Canada/Mountain", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
					},
				},
				{
					name: "Newfoundland",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "Canada/Newfoundland", dstHandover: "2024-03-11T00:00:00-02:30",
									date: "2024-03-10T00:00:00-03:30", expectedDate: "2024-03-10T00:00:00-03:30",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "Canada/Newfoundland", dstHandover: "2024-03-11T00:00:00-02:30",
									date: "2024-03-11T00:00:00-02:30", expectedDate: "2024-03-11T00:00:00-02:30",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "Canada/Newfoundland", dstHandover: "2024-03-11T00:00:00-02:30",
									date: "2024-03-12T00:00:00-02:30", expectedDate: "2024-03-12T00:00:00-02:30",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "Canada/Newfoundland", dstHandover: "2024-11-04T00:00:00-03:30",
									date: "2024-11-03T00:00:00-02:30", expectedDate: "2024-11-03T00:00:00-02:30",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "Canada/Newfoundland", dstHandover: "2024-11-04T00:00:00-03:30",
									date: "2024-11-04T00:00:00-03:30", expectedDate: "2024-11-04T00:00:00-03:30",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "Canada/Newfoundland", dstHandover: "2024-11-04T00:00:00-03:30",
									date: "2024-11-05T00:00:00-03:30", expectedDate: "2024-11-05T00:00:00-03:30",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "Canada/Newfoundland", dstHandover: "2024-03-11T00:00:00-02:30",
									date: "2024-03-10T00:00:00-03:30", expectedDate: "2024-03-10T00:00:00-03:30",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "Canada/Newfoundland", dstHandover: "2024-03-11T00:00:00-02:30",
									date: "2024-03-11T00:00:00-02:30", expectedDate: "2024-03-11T00:00:00-02:30",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "Canada/Newfoundland", dstHandover: "2024-03-11T00:00:00-02:30",
									date: "2024-03-12T00:00:00-02:30", expectedDate: "2024-03-12T00:00:00-02:30",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "Canada/Newfoundland", dstHandover: "2024-11-04T00:00:00-03:30",
									date: "2024-11-03T00:00:00-02:30", expectedDate: "2024-11-03T00:00:00-02:30",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "Canada/Newfoundland", dstHandover: "2024-11-04T00:00:00-03:30",
									date: "2024-11-04T00:00:00-03:30", expectedDate: "2024-11-04T00:00:00-03:30",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "Canada/Newfoundland", dstHandover: "2024-11-04T00:00:00-03:30",
									date: "2024-11-05T00:00:00-03:30", expectedDate: "2024-11-05T00:00:00-03:30",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "Canada/Newfoundland", dstHandover: "2024-03-11T00:00:00-02:30",
									date: "2024-03-10T00:00:00-03:30", expectedDate: "2024-03-10T00:00:00-03:30",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "Canada/Newfoundland", dstHandover: "2024-03-11T00:00:00-02:30",
									date: "2024-03-11T00:00:00-02:30", expectedDate: "2024-03-11T00:00:00-02:30",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "Canada/Newfoundland", dstHandover: "2024-03-11T00:00:00-02:30",
									date: "2024-03-12T00:00:00-02:30", expectedDate: "2024-03-12T00:00:00-02:30",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "Canada/Newfoundland", dstHandover: "2024-11-04T00:00:00-03:30",
									date: "2024-11-03T00:00:00-02:30", expectedDate: "2024-11-03T00:00:00-02:30",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "Canada/Newfoundland", dstHandover: "2024-11-04T00:00:00-03:30",
									date: "2024-11-04T00:00:00-03:30", expectedDate: "2024-11-04T00:00:00-03:30",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "Canada/Newfoundland", dstHandover: "2024-11-04T00:00:00-03:30",
									date: "2024-11-05T00:00:00-03:30", expectedDate: "2024-11-05T00:00:00-03:30",
								},
							}},
					},
				},
				{
					name: "Pacific",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "Canada/Pacific", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "Canada/Pacific", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "Canada/Pacific", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "Canada/Pacific", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "Canada/Pacific", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "Canada/Pacific", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "Canada/Pacific", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "Canada/Pacific", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "Canada/Pacific", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "Canada/Pacific", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "Canada/Pacific", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "Canada/Pacific", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "Canada/Pacific", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "Canada/Pacific", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "Canada/Pacific", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "Canada/Pacific", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "Canada/Pacific", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "Canada/Pacific", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
					},
				},
			},
		},
		{
			name: "Chile",
			children: []testCase{
				{
					name: "Continental",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-09-07", testHandler: testLocalMidnight,
									zone: "Chile/Continental", dstHandover: "2024-09-08T23:00:00-03:00",
									date: "2024-09-07T23:00:00-04:00", expectedDate: "2024-09-07T23:00:00-04:00",
								},
								{
									name: "2024-09-08", testHandler: testLocalMidnight,
									zone: "Chile/Continental", dstHandover: "2024-09-08T23:00:00-03:00",
									date: "2024-09-08T23:00:00-03:00", expectedDate: "2024-09-08T23:00:00-03:00",
								},
								{
									name: "2024-09-09", testHandler: testLocalMidnight,
									zone: "Chile/Continental", dstHandover: "2024-09-08T23:00:00-03:00",
									date: "2024-09-09T23:00:00-03:00", expectedDate: "2024-09-09T23:00:00-03:00",
								},
								{
									name: "2024-04-06", testHandler: testLocalMidnight,
									zone: "Chile/Continental", dstHandover: "2024-04-07T00:00:00-04:00",
									date: "2024-04-06T00:00:00-03:00", expectedDate: "2024-04-06T00:00:00-03:00",
								},
								{
									name: "2024-04-07", testHandler: testLocalMidnight,
									zone: "Chile/Continental", dstHandover: "2024-04-07T00:00:00-04:00",
									date: "2024-04-07T00:00:00-04:00", expectedDate: "2024-04-07T00:00:00-04:00",
								},
								{
									name: "2024-04-08", testHandler: testLocalMidnight,
									zone: "Chile/Continental", dstHandover: "2024-04-07T00:00:00-04:00",
									date: "2024-04-08T00:00:00-04:00", expectedDate: "2024-04-08T00:00:00-04:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-09-07", testHandler: testYesterdayMidnight,
									zone: "Chile/Continental", dstHandover: "2024-09-08T23:00:00-03:00",
									date: "2024-09-07T23:00:00-04:00", expectedDate: "2024-09-07T23:00:00-04:00",
								},
								{
									name: "2024-09-08", testHandler: testYesterdayMidnight,
									zone: "Chile/Continental", dstHandover: "2024-09-08T23:00:00-03:00",
									date: "2024-09-08T23:00:00-03:00", expectedDate: "2024-09-08T23:00:00-03:00",
								},
								{
									name: "2024-09-09", testHandler: testYesterdayMidnight,
									zone: "Chile/Continental", dstHandover: "2024-09-08T23:00:00-03:00",
									date: "2024-09-09T23:00:00-03:00", expectedDate: "2024-09-09T23:00:00-03:00",
								},
								{
									name: "2024-04-06", testHandler: testYesterdayMidnight,
									zone: "Chile/Continental", dstHandover: "2024-04-07T00:00:00-04:00",
									date: "2024-04-06T00:00:00-03:00", expectedDate: "2024-04-06T00:00:00-03:00",
								},
								{
									name: "2024-04-07", testHandler: testYesterdayMidnight,
									zone: "Chile/Continental", dstHandover: "2024-04-07T00:00:00-04:00",
									date: "2024-04-07T00:00:00-04:00", expectedDate: "2024-04-07T00:00:00-04:00",
								},
								{
									name: "2024-04-08", testHandler: testYesterdayMidnight,
									zone: "Chile/Continental", dstHandover: "2024-04-07T00:00:00-04:00",
									date: "2024-04-08T00:00:00-04:00", expectedDate: "2024-04-08T00:00:00-04:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-09-07", testHandler: testTomorrowMidnight,
									zone: "Chile/Continental", dstHandover: "2024-09-08T23:00:00-03:00",
									date: "2024-09-07T23:00:00-04:00", expectedDate: "2024-09-07T23:00:00-04:00",
								},
								{
									name: "2024-09-08", testHandler: testTomorrowMidnight,
									zone: "Chile/Continental", dstHandover: "2024-09-08T23:00:00-03:00",
									date: "2024-09-08T23:00:00-03:00", expectedDate: "2024-09-08T23:00:00-03:00",
								},
								{
									name: "2024-09-09", testHandler: testTomorrowMidnight,
									zone: "Chile/Continental", dstHandover: "2024-09-08T23:00:00-03:00",
									date: "2024-09-09T23:00:00-03:00", expectedDate: "2024-09-09T23:00:00-03:00",
								},
								{
									name: "2024-04-06", testHandler: testTomorrowMidnight,
									zone: "Chile/Continental", dstHandover: "2024-04-07T00:00:00-04:00",
									date: "2024-04-06T00:00:00-03:00", expectedDate: "2024-04-06T00:00:00-03:00",
								},
								{
									name: "2024-04-07", testHandler: testTomorrowMidnight,
									zone: "Chile/Continental", dstHandover: "2024-04-07T00:00:00-04:00",
									date: "2024-04-07T00:00:00-04:00", expectedDate: "2024-04-07T00:00:00-04:00",
								},
								{
									name: "2024-04-08", testHandler: testTomorrowMidnight,
									zone: "Chile/Continental", dstHandover: "2024-04-07T00:00:00-04:00",
									date: "2024-04-08T00:00:00-04:00", expectedDate: "2024-04-08T00:00:00-04:00",
								},
							}},
					},
				},
				{
					name: "EasterIsland",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-09-07", testHandler: testLocalMidnight,
									zone: "Chile/EasterIsland", dstHandover: "2024-09-08T00:00:00-05:00",
									date: "2024-09-07T00:00:00-06:00", expectedDate: "2024-09-07T00:00:00-06:00",
								},
								{
									name: "2024-09-08", testHandler: testLocalMidnight,
									zone: "Chile/EasterIsland", dstHandover: "2024-09-08T00:00:00-05:00",
									date: "2024-09-08T00:00:00-05:00", expectedDate: "2024-09-08T00:00:00-05:00",
								},
								{
									name: "2024-09-09", testHandler: testLocalMidnight,
									zone: "Chile/EasterIsland", dstHandover: "2024-09-08T00:00:00-05:00",
									date: "2024-09-09T00:00:00-05:00", expectedDate: "2024-09-09T00:00:00-05:00",
								},
								{
									name: "2024-04-06", testHandler: testLocalMidnight,
									zone: "Chile/EasterIsland", dstHandover: "2024-04-07T00:00:00-06:00",
									date: "2024-04-06T00:00:00-05:00", expectedDate: "2024-04-06T00:00:00-05:00",
								},
								{
									name: "2024-04-07", testHandler: testLocalMidnight,
									zone: "Chile/EasterIsland", dstHandover: "2024-04-07T00:00:00-06:00",
									date: "2024-04-07T00:00:00-06:00", expectedDate: "2024-04-07T00:00:00-06:00",
								},
								{
									name: "2024-04-08", testHandler: testLocalMidnight,
									zone: "Chile/EasterIsland", dstHandover: "2024-04-07T00:00:00-06:00",
									date: "2024-04-08T00:00:00-06:00", expectedDate: "2024-04-08T00:00:00-06:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-09-07", testHandler: testYesterdayMidnight,
									zone: "Chile/EasterIsland", dstHandover: "2024-09-08T00:00:00-05:00",
									date: "2024-09-07T00:00:00-06:00", expectedDate: "2024-09-07T00:00:00-06:00",
								},
								{
									name: "2024-09-08", testHandler: testYesterdayMidnight,
									zone: "Chile/EasterIsland", dstHandover: "2024-09-08T00:00:00-05:00",
									date: "2024-09-08T00:00:00-05:00", expectedDate: "2024-09-08T00:00:00-05:00",
								},
								{
									name: "2024-09-09", testHandler: testYesterdayMidnight,
									zone: "Chile/EasterIsland", dstHandover: "2024-09-08T00:00:00-05:00",
									date: "2024-09-09T00:00:00-05:00", expectedDate: "2024-09-09T00:00:00-05:00",
								},
								{
									name: "2024-04-06", testHandler: testYesterdayMidnight,
									zone: "Chile/EasterIsland", dstHandover: "2024-04-07T00:00:00-06:00",
									date: "2024-04-06T00:00:00-05:00", expectedDate: "2024-04-06T00:00:00-05:00",
								},
								{
									name: "2024-04-07", testHandler: testYesterdayMidnight,
									zone: "Chile/EasterIsland", dstHandover: "2024-04-07T00:00:00-06:00",
									date: "2024-04-07T00:00:00-06:00", expectedDate: "2024-04-07T00:00:00-06:00",
								},
								{
									name: "2024-04-08", testHandler: testYesterdayMidnight,
									zone: "Chile/EasterIsland", dstHandover: "2024-04-07T00:00:00-06:00",
									date: "2024-04-08T00:00:00-06:00", expectedDate: "2024-04-08T00:00:00-06:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-09-07", testHandler: testTomorrowMidnight,
									zone: "Chile/EasterIsland", dstHandover: "2024-09-08T00:00:00-05:00",
									date: "2024-09-07T00:00:00-06:00", expectedDate: "2024-09-07T00:00:00-06:00",
								},
								{
									name: "2024-09-08", testHandler: testTomorrowMidnight,
									zone: "Chile/EasterIsland", dstHandover: "2024-09-08T00:00:00-05:00",
									date: "2024-09-08T00:00:00-05:00", expectedDate: "2024-09-08T00:00:00-05:00",
								},
								{
									name: "2024-09-09", testHandler: testTomorrowMidnight,
									zone: "Chile/EasterIsland", dstHandover: "2024-09-08T00:00:00-05:00",
									date: "2024-09-09T00:00:00-05:00", expectedDate: "2024-09-09T00:00:00-05:00",
								},
								{
									name: "2024-04-06", testHandler: testTomorrowMidnight,
									zone: "Chile/EasterIsland", dstHandover: "2024-04-07T00:00:00-06:00",
									date: "2024-04-06T00:00:00-05:00", expectedDate: "2024-04-06T00:00:00-05:00",
								},
								{
									name: "2024-04-07", testHandler: testTomorrowMidnight,
									zone: "Chile/EasterIsland", dstHandover: "2024-04-07T00:00:00-06:00",
									date: "2024-04-07T00:00:00-06:00", expectedDate: "2024-04-07T00:00:00-06:00",
								},
								{
									name: "2024-04-08", testHandler: testTomorrowMidnight,
									zone: "Chile/EasterIsland", dstHandover: "2024-04-07T00:00:00-06:00",
									date: "2024-04-08T00:00:00-06:00", expectedDate: "2024-04-08T00:00:00-06:00",
								},
							}},
					},
				},
			},
		},
		{
			name: "Cuba",
			children: []testCase{
				{
					name: "LocalMidnight", children: []testCase{
						{
							name: "2024-03-09", testHandler: testLocalMidnight,
							zone: "Cuba", dstHandover: "2024-03-10T23:00:00-04:00",
							date: "2024-03-09T23:00:00-05:00", expectedDate: "2024-03-09T23:00:00-05:00",
						},
						{
							name: "2024-03-10", testHandler: testLocalMidnight,
							zone: "Cuba", dstHandover: "2024-03-10T23:00:00-04:00",
							date: "2024-03-10T23:00:00-04:00", expectedDate: "2024-03-10T23:00:00-04:00",
						},
						{
							name: "2024-03-11", testHandler: testLocalMidnight,
							zone: "Cuba", dstHandover: "2024-03-10T23:00:00-04:00",
							date: "2024-03-11T23:00:00-04:00", expectedDate: "2024-03-11T23:00:00-04:00",
						},
						{
							name: "2024-11-02", testHandler: testLocalMidnight,
							zone: "Cuba", dstHandover: "2024-11-03T23:00:00-05:00",
							date: "2024-11-02T23:00:00-04:00", expectedDate: "2024-11-02T23:00:00-04:00",
						},
						{
							name: "2024-11-03", testHandler: testLocalMidnight,
							zone: "Cuba", dstHandover: "2024-11-03T23:00:00-05:00",
							date: "2024-11-03T23:00:00-05:00", expectedDate: "2024-11-03T23:00:00-05:00",
						},
						{
							name: "2024-11-04", testHandler: testLocalMidnight,
							zone: "Cuba", dstHandover: "2024-11-03T23:00:00-05:00",
							date: "2024-11-04T23:00:00-05:00", expectedDate: "2024-11-04T23:00:00-05:00",
						},
					}},
				{
					name: "YesterdayMidnight", children: []testCase{
						{
							name: "2024-03-09", testHandler: testYesterdayMidnight,
							zone: "Cuba", dstHandover: "2024-03-10T23:00:00-04:00",
							date: "2024-03-09T23:00:00-05:00", expectedDate: "2024-03-09T23:00:00-05:00",
						},
						{
							name: "2024-03-10", testHandler: testYesterdayMidnight,
							zone: "Cuba", dstHandover: "2024-03-10T23:00:00-04:00",
							date: "2024-03-10T23:00:00-04:00", expectedDate: "2024-03-10T23:00:00-04:00",
						},
						{
							name: "2024-03-11", testHandler: testYesterdayMidnight,
							zone: "Cuba", dstHandover: "2024-03-10T23:00:00-04:00",
							date: "2024-03-11T23:00:00-04:00", expectedDate: "2024-03-11T23:00:00-04:00",
						},
						{
							name: "2024-11-02", testHandler: testYesterdayMidnight,
							zone: "Cuba", dstHandover: "2024-11-03T23:00:00-05:00",
							date: "2024-11-02T23:00:00-04:00", expectedDate: "2024-11-02T23:00:00-04:00",
						},
						{
							name: "2024-11-03", testHandler: testYesterdayMidnight,
							zone: "Cuba", dstHandover: "2024-11-03T23:00:00-05:00",
							date: "2024-11-03T23:00:00-05:00", expectedDate: "2024-11-03T23:00:00-05:00",
						},
						{
							name: "2024-11-04", testHandler: testYesterdayMidnight,
							zone: "Cuba", dstHandover: "2024-11-03T23:00:00-05:00",
							date: "2024-11-04T23:00:00-05:00", expectedDate: "2024-11-04T23:00:00-05:00",
						},
					}},
				{
					name: "TomorrowMidnight", children: []testCase{
						{
							name: "2024-03-09", testHandler: testTomorrowMidnight,
							zone: "Cuba", dstHandover: "2024-03-10T23:00:00-04:00",
							date: "2024-03-09T23:00:00-05:00", expectedDate: "2024-03-09T23:00:00-05:00",
						},
						{
							name: "2024-03-10", testHandler: testTomorrowMidnight,
							zone: "Cuba", dstHandover: "2024-03-10T23:00:00-04:00",
							date: "2024-03-10T23:00:00-04:00", expectedDate: "2024-03-10T23:00:00-04:00",
						},
						{
							name: "2024-03-11", testHandler: testTomorrowMidnight,
							zone: "Cuba", dstHandover: "2024-03-10T23:00:00-04:00",
							date: "2024-03-11T23:00:00-04:00", expectedDate: "2024-03-11T23:00:00-04:00",
						},
						{
							name: "2024-11-02", testHandler: testTomorrowMidnight,
							zone: "Cuba", dstHandover: "2024-11-03T23:00:00-05:00",
							date: "2024-11-02T23:00:00-04:00", expectedDate: "2024-11-02T23:00:00-04:00",
						},
						{
							name: "2024-11-03", testHandler: testTomorrowMidnight,
							zone: "Cuba", dstHandover: "2024-11-03T23:00:00-05:00",
							date: "2024-11-03T23:00:00-05:00", expectedDate: "2024-11-03T23:00:00-05:00",
						},
						{
							name: "2024-11-04", testHandler: testTomorrowMidnight,
							zone: "Cuba", dstHandover: "2024-11-03T23:00:00-05:00",
							date: "2024-11-04T23:00:00-05:00", expectedDate: "2024-11-04T23:00:00-05:00",
						},
					}},
			},
		},
		{
			name: "EET",
			children: []testCase{
				{
					name: "LocalMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testLocalMidnight,
							zone: "EET", dstHandover: "2024-04-01T00:00:00+03:00",
							date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
						},
						{
							name: "2024-04-01", testHandler: testLocalMidnight,
							zone: "EET", dstHandover: "2024-04-01T00:00:00+03:00",
							date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
						},
						{
							name: "2024-04-02", testHandler: testLocalMidnight,
							zone: "EET", dstHandover: "2024-04-01T00:00:00+03:00",
							date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
						},
						{
							name: "2024-10-27", testHandler: testLocalMidnight,
							zone: "EET", dstHandover: "2024-10-28T00:00:00+02:00",
							date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
						},
						{
							name: "2024-10-28", testHandler: testLocalMidnight,
							zone: "EET", dstHandover: "2024-10-28T00:00:00+02:00",
							date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
						},
						{
							name: "2024-10-29", testHandler: testLocalMidnight,
							zone: "EET", dstHandover: "2024-10-28T00:00:00+02:00",
							date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
						},
					}},
				{
					name: "YesterdayMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testYesterdayMidnight,
							zone: "EET", dstHandover: "2024-04-01T00:00:00+03:00",
							date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
						},
						{
							name: "2024-04-01", testHandler: testYesterdayMidnight,
							zone: "EET", dstHandover: "2024-04-01T00:00:00+03:00",
							date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
						},
						{
							name: "2024-04-02", testHandler: testYesterdayMidnight,
							zone: "EET", dstHandover: "2024-04-01T00:00:00+03:00",
							date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
						},
						{
							name: "2024-10-27", testHandler: testYesterdayMidnight,
							zone: "EET", dstHandover: "2024-10-28T00:00:00+02:00",
							date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
						},
						{
							name: "2024-10-28", testHandler: testYesterdayMidnight,
							zone: "EET", dstHandover: "2024-10-28T00:00:00+02:00",
							date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
						},
						{
							name: "2024-10-29", testHandler: testYesterdayMidnight,
							zone: "EET", dstHandover: "2024-10-28T00:00:00+02:00",
							date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
						},
					}},
				{
					name: "TomorrowMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testTomorrowMidnight,
							zone: "EET", dstHandover: "2024-04-01T00:00:00+03:00",
							date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
						},
						{
							name: "2024-04-01", testHandler: testTomorrowMidnight,
							zone: "EET", dstHandover: "2024-04-01T00:00:00+03:00",
							date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
						},
						{
							name: "2024-04-02", testHandler: testTomorrowMidnight,
							zone: "EET", dstHandover: "2024-04-01T00:00:00+03:00",
							date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
						},
						{
							name: "2024-10-27", testHandler: testTomorrowMidnight,
							zone: "EET", dstHandover: "2024-10-28T00:00:00+02:00",
							date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
						},
						{
							name: "2024-10-28", testHandler: testTomorrowMidnight,
							zone: "EET", dstHandover: "2024-10-28T00:00:00+02:00",
							date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
						},
						{
							name: "2024-10-29", testHandler: testTomorrowMidnight,
							zone: "EET", dstHandover: "2024-10-28T00:00:00+02:00",
							date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
						},
					}},
			},
		},
		{
			name: "EST5EDT",
			children: []testCase{
				{
					name: "LocalMidnight", children: []testCase{
						{
							name: "2024-03-10", testHandler: testLocalMidnight,
							zone: "EST5EDT", dstHandover: "2024-03-11T00:00:00-04:00",
							date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
						},
						{
							name: "2024-03-11", testHandler: testLocalMidnight,
							zone: "EST5EDT", dstHandover: "2024-03-11T00:00:00-04:00",
							date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
						},
						{
							name: "2024-03-12", testHandler: testLocalMidnight,
							zone: "EST5EDT", dstHandover: "2024-03-11T00:00:00-04:00",
							date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
						},
						{
							name: "2024-11-03", testHandler: testLocalMidnight,
							zone: "EST5EDT", dstHandover: "2024-11-04T00:00:00-05:00",
							date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
						},
						{
							name: "2024-11-04", testHandler: testLocalMidnight,
							zone: "EST5EDT", dstHandover: "2024-11-04T00:00:00-05:00",
							date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
						},
						{
							name: "2024-11-05", testHandler: testLocalMidnight,
							zone: "EST5EDT", dstHandover: "2024-11-04T00:00:00-05:00",
							date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
						},
					}},
				{
					name: "YesterdayMidnight", children: []testCase{
						{
							name: "2024-03-10", testHandler: testYesterdayMidnight,
							zone: "EST5EDT", dstHandover: "2024-03-11T00:00:00-04:00",
							date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
						},
						{
							name: "2024-03-11", testHandler: testYesterdayMidnight,
							zone: "EST5EDT", dstHandover: "2024-03-11T00:00:00-04:00",
							date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
						},
						{
							name: "2024-03-12", testHandler: testYesterdayMidnight,
							zone: "EST5EDT", dstHandover: "2024-03-11T00:00:00-04:00",
							date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
						},
						{
							name: "2024-11-03", testHandler: testYesterdayMidnight,
							zone: "EST5EDT", dstHandover: "2024-11-04T00:00:00-05:00",
							date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
						},
						{
							name: "2024-11-04", testHandler: testYesterdayMidnight,
							zone: "EST5EDT", dstHandover: "2024-11-04T00:00:00-05:00",
							date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
						},
						{
							name: "2024-11-05", testHandler: testYesterdayMidnight,
							zone: "EST5EDT", dstHandover: "2024-11-04T00:00:00-05:00",
							date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
						},
					}},
				{
					name: "TomorrowMidnight", children: []testCase{
						{
							name: "2024-03-10", testHandler: testTomorrowMidnight,
							zone: "EST5EDT", dstHandover: "2024-03-11T00:00:00-04:00",
							date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
						},
						{
							name: "2024-03-11", testHandler: testTomorrowMidnight,
							zone: "EST5EDT", dstHandover: "2024-03-11T00:00:00-04:00",
							date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
						},
						{
							name: "2024-03-12", testHandler: testTomorrowMidnight,
							zone: "EST5EDT", dstHandover: "2024-03-11T00:00:00-04:00",
							date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
						},
						{
							name: "2024-11-03", testHandler: testTomorrowMidnight,
							zone: "EST5EDT", dstHandover: "2024-11-04T00:00:00-05:00",
							date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
						},
						{
							name: "2024-11-04", testHandler: testTomorrowMidnight,
							zone: "EST5EDT", dstHandover: "2024-11-04T00:00:00-05:00",
							date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
						},
						{
							name: "2024-11-05", testHandler: testTomorrowMidnight,
							zone: "EST5EDT", dstHandover: "2024-11-04T00:00:00-05:00",
							date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
						},
					}},
			},
		},
		{
			name: "Egypt",
			children: []testCase{
				{
					name: "LocalMidnight", children: []testCase{
						{
							name: "2024-04-25", testHandler: testLocalMidnight,
							zone: "Egypt", dstHandover: "2024-04-26T01:00:00+03:00",
							date: "2024-04-25T01:00:00+02:00", expectedDate: "2024-04-25T01:00:00+02:00",
						},
						{
							name: "2024-04-26", testHandler: testLocalMidnight,
							zone: "Egypt", dstHandover: "2024-04-26T01:00:00+03:00",
							date: "2024-04-26T01:00:00+03:00", expectedDate: "2024-04-26T01:00:00+03:00",
						},
						{
							name: "2024-04-27", testHandler: testLocalMidnight,
							zone: "Egypt", dstHandover: "2024-04-26T01:00:00+03:00",
							date: "2024-04-27T01:00:00+03:00", expectedDate: "2024-04-27T01:00:00+03:00",
						},
						{
							name: "2024-10-31", testHandler: testLocalMidnight,
							zone: "Egypt", dstHandover: "2024-11-01T01:00:00+02:00",
							date: "2024-10-31T01:00:00+03:00", expectedDate: "2024-10-31T01:00:00+03:00",
						},
						{
							name: "2024-11-01", testHandler: testLocalMidnight,
							zone: "Egypt", dstHandover: "2024-11-01T01:00:00+02:00",
							date: "2024-11-01T01:00:00+02:00", expectedDate: "2024-11-01T01:00:00+02:00",
						},
						{
							name: "2024-11-02", testHandler: testLocalMidnight,
							zone: "Egypt", dstHandover: "2024-11-01T01:00:00+02:00",
							date: "2024-11-02T01:00:00+02:00", expectedDate: "2024-11-02T01:00:00+02:00",
						},
					}},
				{
					name: "YesterdayMidnight", children: []testCase{
						{
							name: "2024-04-25", testHandler: testYesterdayMidnight,
							zone: "Egypt", dstHandover: "2024-04-26T01:00:00+03:00",
							date: "2024-04-25T01:00:00+02:00", expectedDate: "2024-04-25T01:00:00+02:00",
						},
						{
							name: "2024-04-26", testHandler: testYesterdayMidnight,
							zone: "Egypt", dstHandover: "2024-04-26T01:00:00+03:00",
							date: "2024-04-26T01:00:00+03:00", expectedDate: "2024-04-26T01:00:00+03:00",
						},
						{
							name: "2024-04-27", testHandler: testYesterdayMidnight,
							zone: "Egypt", dstHandover: "2024-04-26T01:00:00+03:00",
							date: "2024-04-27T01:00:00+03:00", expectedDate: "2024-04-27T01:00:00+03:00",
						},
						{
							name: "2024-10-31", testHandler: testYesterdayMidnight,
							zone: "Egypt", dstHandover: "2024-11-01T01:00:00+02:00",
							date: "2024-10-31T01:00:00+03:00", expectedDate: "2024-10-31T01:00:00+03:00",
						},
						{
							name: "2024-11-01", testHandler: testYesterdayMidnight,
							zone: "Egypt", dstHandover: "2024-11-01T01:00:00+02:00",
							date: "2024-11-01T01:00:00+02:00", expectedDate: "2024-11-01T01:00:00+02:00",
						},
						{
							name: "2024-11-02", testHandler: testYesterdayMidnight,
							zone: "Egypt", dstHandover: "2024-11-01T01:00:00+02:00",
							date: "2024-11-02T01:00:00+02:00", expectedDate: "2024-11-02T01:00:00+02:00",
						},
					}},
				{
					name: "TomorrowMidnight", children: []testCase{
						{
							name: "2024-04-25", testHandler: testTomorrowMidnight,
							zone: "Egypt", dstHandover: "2024-04-26T01:00:00+03:00",
							date: "2024-04-25T01:00:00+02:00", expectedDate: "2024-04-25T01:00:00+02:00",
						},
						{
							name: "2024-04-26", testHandler: testTomorrowMidnight,
							zone: "Egypt", dstHandover: "2024-04-26T01:00:00+03:00",
							date: "2024-04-26T01:00:00+03:00", expectedDate: "2024-04-26T01:00:00+03:00",
						},
						{
							name: "2024-04-27", testHandler: testTomorrowMidnight,
							zone: "Egypt", dstHandover: "2024-04-26T01:00:00+03:00",
							date: "2024-04-27T01:00:00+03:00", expectedDate: "2024-04-27T01:00:00+03:00",
						},
						{
							name: "2024-10-31", testHandler: testTomorrowMidnight,
							zone: "Egypt", dstHandover: "2024-11-01T01:00:00+02:00",
							date: "2024-10-31T01:00:00+03:00", expectedDate: "2024-10-31T01:00:00+03:00",
						},
						{
							name: "2024-11-01", testHandler: testTomorrowMidnight,
							zone: "Egypt", dstHandover: "2024-11-01T01:00:00+02:00",
							date: "2024-11-01T01:00:00+02:00", expectedDate: "2024-11-01T01:00:00+02:00",
						},
						{
							name: "2024-11-02", testHandler: testTomorrowMidnight,
							zone: "Egypt", dstHandover: "2024-11-01T01:00:00+02:00",
							date: "2024-11-02T01:00:00+02:00", expectedDate: "2024-11-02T01:00:00+02:00",
						},
					}},
			},
		},
		{
			name: "Eire",
			children: []testCase{
				{
					name: "LocalMidnight", children: []testCase{
						{
							name: "2024-10-27", testHandler: testLocalMidnight,
							zone: "Eire", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
						},
						{
							name: "2024-10-28", testHandler: testLocalMidnight,
							zone: "Eire", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
						},
						{
							name: "2024-10-29", testHandler: testLocalMidnight,
							zone: "Eire", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
						},
						{
							name: "2024-03-31", testHandler: testLocalMidnight,
							zone: "Eire", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
						},
						{
							name: "2024-04-01", testHandler: testLocalMidnight,
							zone: "Eire", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
						},
						{
							name: "2024-04-02", testHandler: testLocalMidnight,
							zone: "Eire", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
						},
					}},
				{
					name: "YesterdayMidnight", children: []testCase{
						{
							name: "2024-10-27", testHandler: testYesterdayMidnight,
							zone: "Eire", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
						},
						{
							name: "2024-10-28", testHandler: testYesterdayMidnight,
							zone: "Eire", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
						},
						{
							name: "2024-10-29", testHandler: testYesterdayMidnight,
							zone: "Eire", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
						},
						{
							name: "2024-03-31", testHandler: testYesterdayMidnight,
							zone: "Eire", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
						},
						{
							name: "2024-04-01", testHandler: testYesterdayMidnight,
							zone: "Eire", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
						},
						{
							name: "2024-04-02", testHandler: testYesterdayMidnight,
							zone: "Eire", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
						},
					}},
				{
					name: "TomorrowMidnight", children: []testCase{
						{
							name: "2024-10-27", testHandler: testTomorrowMidnight,
							zone: "Eire", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
						},
						{
							name: "2024-10-28", testHandler: testTomorrowMidnight,
							zone: "Eire", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
						},
						{
							name: "2024-10-29", testHandler: testTomorrowMidnight,
							zone: "Eire", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
						},
						{
							name: "2024-03-31", testHandler: testTomorrowMidnight,
							zone: "Eire", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
						},
						{
							name: "2024-04-01", testHandler: testTomorrowMidnight,
							zone: "Eire", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
						},
						{
							name: "2024-04-02", testHandler: testTomorrowMidnight,
							zone: "Eire", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
						},
					}},
			},
		},
		{
			name: "Europe",
			children: []testCase{
				{
					name: "Amsterdam",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Amsterdam", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Amsterdam", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Amsterdam", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Amsterdam", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Amsterdam", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Amsterdam", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Amsterdam", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Amsterdam", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Amsterdam", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Amsterdam", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Amsterdam", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Amsterdam", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Amsterdam", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Amsterdam", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Amsterdam", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Amsterdam", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Amsterdam", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Amsterdam", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Andorra",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Andorra", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Andorra", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Andorra", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Andorra", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Andorra", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Andorra", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Andorra", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Andorra", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Andorra", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Andorra", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Andorra", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Andorra", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Andorra", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Andorra", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Andorra", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Andorra", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Andorra", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Andorra", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Athens",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Athens", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Athens", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Athens", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Athens", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Athens", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Athens", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Athens", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Athens", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Athens", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Athens", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Athens", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Athens", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Athens", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Athens", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Athens", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Athens", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Athens", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Athens", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
					},
				},
				{
					name: "Belfast",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Belfast", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Belfast", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Belfast", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Belfast", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Belfast", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Belfast", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Belfast", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Belfast", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Belfast", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Belfast", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Belfast", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Belfast", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Belfast", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Belfast", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Belfast", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Belfast", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Belfast", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Belfast", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
					},
				},
				{
					name: "Belgrade",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Belgrade", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Belgrade", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Belgrade", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Belgrade", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Belgrade", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Belgrade", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Belgrade", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Belgrade", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Belgrade", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Belgrade", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Belgrade", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Belgrade", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Belgrade", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Belgrade", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Belgrade", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Belgrade", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Belgrade", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Belgrade", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Berlin",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Berlin", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Berlin", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Berlin", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Berlin", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Berlin", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Berlin", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Berlin", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Berlin", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Berlin", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Berlin", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Berlin", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Berlin", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Berlin", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Berlin", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Berlin", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Berlin", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Berlin", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Berlin", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Bratislava",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Bratislava", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Bratislava", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Bratislava", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Bratislava", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Bratislava", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Bratislava", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Bratislava", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Bratislava", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Bratislava", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Bratislava", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Bratislava", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Bratislava", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Bratislava", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Bratislava", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Bratislava", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Bratislava", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Bratislava", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Bratislava", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Brussels",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Brussels", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Brussels", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Brussels", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Brussels", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Brussels", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Brussels", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Brussels", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Brussels", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Brussels", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Brussels", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Brussels", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Brussels", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Brussels", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Brussels", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Brussels", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Brussels", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Brussels", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Brussels", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Bucharest",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Bucharest", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Bucharest", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Bucharest", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Bucharest", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Bucharest", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Bucharest", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Bucharest", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Bucharest", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Bucharest", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Bucharest", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Bucharest", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Bucharest", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Bucharest", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Bucharest", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Bucharest", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Bucharest", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Bucharest", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Bucharest", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
					},
				},
				{
					name: "Budapest",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Budapest", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Budapest", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Budapest", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Budapest", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Budapest", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Budapest", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Budapest", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Budapest", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Budapest", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Budapest", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Budapest", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Budapest", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Budapest", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Budapest", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Budapest", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Budapest", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Budapest", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Budapest", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Busingen",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Busingen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Busingen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Busingen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Busingen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Busingen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Busingen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Busingen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Busingen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Busingen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Busingen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Busingen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Busingen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Busingen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Busingen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Busingen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Busingen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Busingen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Busingen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Chisinau",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Chisinau", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Chisinau", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Chisinau", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Chisinau", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Chisinau", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Chisinau", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Chisinau", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Chisinau", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Chisinau", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Chisinau", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Chisinau", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Chisinau", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Chisinau", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Chisinau", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Chisinau", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Chisinau", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Chisinau", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Chisinau", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
					},
				},
				{
					name: "Copenhagen",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Copenhagen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Copenhagen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Copenhagen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Copenhagen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Copenhagen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Copenhagen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Copenhagen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Copenhagen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Copenhagen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Copenhagen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Copenhagen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Copenhagen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Copenhagen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Copenhagen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Copenhagen", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Copenhagen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Copenhagen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Copenhagen", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Dublin",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Dublin", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Dublin", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Dublin", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Dublin", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Dublin", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Dublin", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Dublin", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Dublin", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Dublin", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Dublin", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Dublin", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Dublin", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Dublin", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Dublin", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Dublin", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Dublin", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Dublin", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Dublin", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Gibraltar",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Gibraltar", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Gibraltar", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Gibraltar", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Gibraltar", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Gibraltar", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Gibraltar", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Gibraltar", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Gibraltar", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Gibraltar", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Gibraltar", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Gibraltar", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Gibraltar", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Gibraltar", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Gibraltar", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Gibraltar", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Gibraltar", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Gibraltar", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Gibraltar", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Guernsey",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Guernsey", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Guernsey", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Guernsey", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Guernsey", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Guernsey", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Guernsey", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Guernsey", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Guernsey", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Guernsey", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Guernsey", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Guernsey", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Guernsey", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Guernsey", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Guernsey", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Guernsey", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Guernsey", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Guernsey", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Guernsey", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
					},
				},
				{
					name: "Helsinki",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Helsinki", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Helsinki", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Helsinki", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Helsinki", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Helsinki", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Helsinki", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Helsinki", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Helsinki", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Helsinki", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Helsinki", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Helsinki", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Helsinki", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Helsinki", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Helsinki", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Helsinki", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Helsinki", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Helsinki", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Helsinki", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
					},
				},
				{
					name: "Isle_of_Man",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Isle_of_Man", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Isle_of_Man", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Isle_of_Man", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Isle_of_Man", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Isle_of_Man", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Isle_of_Man", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Isle_of_Man", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Isle_of_Man", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Isle_of_Man", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Isle_of_Man", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Isle_of_Man", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Isle_of_Man", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Isle_of_Man", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Isle_of_Man", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Isle_of_Man", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Isle_of_Man", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Isle_of_Man", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Isle_of_Man", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
					},
				},
				{
					name: "Jersey",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Jersey", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Jersey", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Jersey", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Jersey", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Jersey", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Jersey", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Jersey", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Jersey", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Jersey", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Jersey", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Jersey", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Jersey", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Jersey", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Jersey", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Jersey", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Jersey", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Jersey", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Jersey", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
					},
				},
				{
					name: "Kiev",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Kiev", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Kiev", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Kiev", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Kiev", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Kiev", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Kiev", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Kiev", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Kiev", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Kiev", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Kiev", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Kiev", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Kiev", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Kiev", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Kiev", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Kiev", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Kiev", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Kiev", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Kiev", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
					},
				},
				{
					name: "Kyiv",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Kyiv", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Kyiv", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Kyiv", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Kyiv", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Kyiv", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Kyiv", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Kyiv", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Kyiv", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Kyiv", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Kyiv", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Kyiv", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Kyiv", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Kyiv", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Kyiv", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Kyiv", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Kyiv", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Kyiv", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Kyiv", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
					},
				},
				{
					name: "Lisbon",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Lisbon", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Lisbon", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Lisbon", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Lisbon", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Lisbon", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Lisbon", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Lisbon", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Lisbon", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Lisbon", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Lisbon", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Lisbon", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Lisbon", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Lisbon", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Lisbon", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Lisbon", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Lisbon", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Lisbon", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Lisbon", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
					},
				},
				{
					name: "Ljubljana",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Ljubljana", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Ljubljana", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Ljubljana", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Ljubljana", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Ljubljana", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Ljubljana", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Ljubljana", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Ljubljana", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Ljubljana", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Ljubljana", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Ljubljana", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Ljubljana", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Ljubljana", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Ljubljana", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Ljubljana", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Ljubljana", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Ljubljana", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Ljubljana", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "London",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/London", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/London", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/London", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/London", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/London", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/London", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/London", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/London", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/London", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/London", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/London", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/London", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/London", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/London", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/London", dstHandover: "2024-04-01T00:00:00+01:00",
									date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/London", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/London", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/London", dstHandover: "2024-10-28T00:00:00Z",
									date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
								},
							}},
					},
				},
				{
					name: "Luxembourg",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Luxembourg", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Luxembourg", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Luxembourg", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Luxembourg", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Luxembourg", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Luxembourg", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Luxembourg", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Luxembourg", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Luxembourg", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Luxembourg", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Luxembourg", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Luxembourg", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Luxembourg", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Luxembourg", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Luxembourg", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Luxembourg", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Luxembourg", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Luxembourg", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Madrid",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Madrid", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Madrid", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Madrid", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Madrid", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Madrid", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Madrid", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Madrid", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Madrid", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Madrid", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Madrid", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Madrid", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Madrid", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Madrid", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Madrid", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Madrid", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Madrid", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Madrid", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Madrid", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Malta",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Malta", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Malta", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Malta", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Malta", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Malta", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Malta", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Malta", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Malta", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Malta", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Malta", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Malta", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Malta", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Malta", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Malta", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Malta", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Malta", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Malta", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Malta", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Mariehamn",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Mariehamn", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Mariehamn", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Mariehamn", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Mariehamn", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Mariehamn", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Mariehamn", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Mariehamn", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Mariehamn", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Mariehamn", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Mariehamn", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Mariehamn", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Mariehamn", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Mariehamn", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Mariehamn", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Mariehamn", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Mariehamn", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Mariehamn", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Mariehamn", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
					},
				},
				{
					name: "Monaco",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Monaco", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Monaco", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Monaco", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Monaco", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Monaco", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Monaco", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Monaco", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Monaco", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Monaco", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Monaco", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Monaco", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Monaco", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Monaco", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Monaco", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Monaco", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Monaco", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Monaco", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Monaco", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Nicosia",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Nicosia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Nicosia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Nicosia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Nicosia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Nicosia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Nicosia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Nicosia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Nicosia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Nicosia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Nicosia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Nicosia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Nicosia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Nicosia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Nicosia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Nicosia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Nicosia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Nicosia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Nicosia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
					},
				},
				{
					name: "Oslo",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Oslo", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Oslo", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Oslo", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Oslo", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Oslo", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Oslo", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Oslo", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Oslo", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Oslo", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Oslo", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Oslo", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Oslo", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Oslo", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Oslo", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Oslo", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Oslo", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Oslo", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Oslo", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Paris",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Paris", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Paris", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Paris", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Paris", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Paris", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Paris", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Paris", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Paris", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Paris", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Paris", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Paris", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Paris", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Paris", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Paris", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Paris", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Paris", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Paris", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Paris", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Podgorica",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Podgorica", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Podgorica", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Podgorica", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Podgorica", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Podgorica", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Podgorica", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Podgorica", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Podgorica", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Podgorica", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Podgorica", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Podgorica", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Podgorica", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Podgorica", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Podgorica", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Podgorica", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Podgorica", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Podgorica", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Podgorica", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Prague",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Prague", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Prague", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Prague", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Prague", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Prague", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Prague", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Prague", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Prague", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Prague", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Prague", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Prague", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Prague", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Prague", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Prague", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Prague", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Prague", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Prague", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Prague", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Riga",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Riga", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Riga", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Riga", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Riga", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Riga", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Riga", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Riga", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Riga", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Riga", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Riga", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Riga", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Riga", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Riga", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Riga", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Riga", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Riga", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Riga", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Riga", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
					},
				},
				{
					name: "Rome",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Rome", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Rome", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Rome", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Rome", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Rome", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Rome", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Rome", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Rome", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Rome", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Rome", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Rome", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Rome", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Rome", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Rome", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Rome", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Rome", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Rome", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Rome", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "San_Marino",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/San_Marino", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/San_Marino", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/San_Marino", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/San_Marino", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/San_Marino", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/San_Marino", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/San_Marino", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/San_Marino", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/San_Marino", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/San_Marino", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/San_Marino", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/San_Marino", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/San_Marino", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/San_Marino", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/San_Marino", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/San_Marino", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/San_Marino", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/San_Marino", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Sarajevo",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Sarajevo", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Sarajevo", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Sarajevo", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Sarajevo", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Sarajevo", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Sarajevo", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Sarajevo", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Sarajevo", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Sarajevo", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Sarajevo", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Sarajevo", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Sarajevo", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Sarajevo", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Sarajevo", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Sarajevo", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Sarajevo", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Sarajevo", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Sarajevo", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Skopje",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Skopje", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Skopje", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Skopje", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Skopje", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Skopje", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Skopje", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Skopje", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Skopje", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Skopje", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Skopje", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Skopje", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Skopje", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Skopje", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Skopje", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Skopje", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Skopje", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Skopje", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Skopje", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Sofia",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Sofia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Sofia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Sofia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Sofia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Sofia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Sofia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Sofia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Sofia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Sofia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Sofia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Sofia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Sofia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Sofia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Sofia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Sofia", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Sofia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Sofia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Sofia", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
					},
				},
				{
					name: "Stockholm",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Stockholm", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Stockholm", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Stockholm", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Stockholm", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Stockholm", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Stockholm", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Stockholm", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Stockholm", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Stockholm", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Stockholm", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Stockholm", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Stockholm", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Stockholm", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Stockholm", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Stockholm", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Stockholm", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Stockholm", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Stockholm", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Tallinn",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Tallinn", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Tallinn", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Tallinn", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Tallinn", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Tallinn", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Tallinn", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Tallinn", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Tallinn", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Tallinn", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Tallinn", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Tallinn", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Tallinn", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Tallinn", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Tallinn", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Tallinn", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Tallinn", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Tallinn", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Tallinn", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
					},
				},
				{
					name: "Tirane",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Tirane", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Tirane", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Tirane", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Tirane", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Tirane", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Tirane", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Tirane", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Tirane", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Tirane", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Tirane", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Tirane", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Tirane", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Tirane", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Tirane", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Tirane", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Tirane", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Tirane", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Tirane", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Tiraspol",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Tiraspol", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Tiraspol", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Tiraspol", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Tiraspol", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Tiraspol", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Tiraspol", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Tiraspol", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Tiraspol", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Tiraspol", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Tiraspol", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Tiraspol", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Tiraspol", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Tiraspol", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Tiraspol", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Tiraspol", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Tiraspol", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Tiraspol", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Tiraspol", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
					},
				},
				{
					name: "Uzhgorod",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Uzhgorod", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Uzhgorod", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Uzhgorod", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Uzhgorod", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Uzhgorod", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Uzhgorod", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Uzhgorod", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Uzhgorod", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Uzhgorod", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Uzhgorod", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Uzhgorod", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Uzhgorod", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Uzhgorod", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Uzhgorod", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Uzhgorod", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Uzhgorod", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Uzhgorod", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Uzhgorod", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
					},
				},
				{
					name: "Vaduz",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Vaduz", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Vaduz", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Vaduz", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Vaduz", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Vaduz", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Vaduz", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Vaduz", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Vaduz", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Vaduz", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Vaduz", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Vaduz", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Vaduz", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Vaduz", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Vaduz", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Vaduz", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Vaduz", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Vaduz", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Vaduz", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Vatican",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Vatican", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Vatican", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Vatican", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Vatican", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Vatican", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Vatican", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Vatican", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Vatican", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Vatican", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Vatican", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Vatican", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Vatican", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Vatican", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Vatican", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Vatican", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Vatican", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Vatican", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Vatican", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Vienna",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Vienna", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Vienna", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Vienna", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Vienna", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Vienna", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Vienna", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Vienna", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Vienna", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Vienna", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Vienna", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Vienna", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Vienna", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Vienna", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Vienna", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Vienna", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Vienna", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Vienna", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Vienna", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Vilnius",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Vilnius", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Vilnius", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Vilnius", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Vilnius", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Vilnius", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Vilnius", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Vilnius", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Vilnius", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Vilnius", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Vilnius", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Vilnius", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Vilnius", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Vilnius", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Vilnius", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Vilnius", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Vilnius", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Vilnius", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Vilnius", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
					},
				},
				{
					name: "Warsaw",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Warsaw", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Warsaw", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Warsaw", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Warsaw", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Warsaw", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Warsaw", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Warsaw", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Warsaw", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Warsaw", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Warsaw", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Warsaw", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Warsaw", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Warsaw", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Warsaw", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Warsaw", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Warsaw", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Warsaw", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Warsaw", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Zagreb",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Zagreb", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Zagreb", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Zagreb", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Zagreb", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Zagreb", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Zagreb", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Zagreb", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Zagreb", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Zagreb", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Zagreb", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Zagreb", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Zagreb", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Zagreb", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Zagreb", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Zagreb", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Zagreb", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Zagreb", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Zagreb", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
				{
					name: "Zaporozhye",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Zaporozhye", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Zaporozhye", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Zaporozhye", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Zaporozhye", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Zaporozhye", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Zaporozhye", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Zaporozhye", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Zaporozhye", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Zaporozhye", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Zaporozhye", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Zaporozhye", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Zaporozhye", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Zaporozhye", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-03-31T00:00:00+02:00", expectedDate: "2024-03-31T00:00:00+02:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Zaporozhye", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-01T00:00:00+03:00", expectedDate: "2024-04-01T00:00:00+03:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Zaporozhye", dstHandover: "2024-04-01T00:00:00+03:00",
									date: "2024-04-02T00:00:00+03:00", expectedDate: "2024-04-02T00:00:00+03:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Zaporozhye", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Zaporozhye", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Zaporozhye", dstHandover: "2024-10-28T00:00:00+02:00",
									date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
								},
							}},
					},
				},
				{
					name: "Zurich",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testLocalMidnight,
									zone: "Europe/Zurich", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testLocalMidnight,
									zone: "Europe/Zurich", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testLocalMidnight,
									zone: "Europe/Zurich", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testLocalMidnight,
									zone: "Europe/Zurich", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testLocalMidnight,
									zone: "Europe/Zurich", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testLocalMidnight,
									zone: "Europe/Zurich", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testYesterdayMidnight,
									zone: "Europe/Zurich", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testYesterdayMidnight,
									zone: "Europe/Zurich", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testYesterdayMidnight,
									zone: "Europe/Zurich", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testYesterdayMidnight,
									zone: "Europe/Zurich", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testYesterdayMidnight,
									zone: "Europe/Zurich", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testYesterdayMidnight,
									zone: "Europe/Zurich", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-31", testHandler: testTomorrowMidnight,
									zone: "Europe/Zurich", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
								},
								{
									name: "2024-04-01", testHandler: testTomorrowMidnight,
									zone: "Europe/Zurich", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
								},
								{
									name: "2024-04-02", testHandler: testTomorrowMidnight,
									zone: "Europe/Zurich", dstHandover: "2024-04-01T00:00:00+02:00",
									date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
								},
								{
									name: "2024-10-27", testHandler: testTomorrowMidnight,
									zone: "Europe/Zurich", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
								},
								{
									name: "2024-10-28", testHandler: testTomorrowMidnight,
									zone: "Europe/Zurich", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
								},
								{
									name: "2024-10-29", testHandler: testTomorrowMidnight,
									zone: "Europe/Zurich", dstHandover: "2024-10-28T00:00:00+01:00",
									date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
								},
							}},
					},
				},
			},
		},
		{
			name: "GB",
			children: []testCase{
				{
					name: "LocalMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testLocalMidnight,
							zone: "GB", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
						},
						{
							name: "2024-04-01", testHandler: testLocalMidnight,
							zone: "GB", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
						},
						{
							name: "2024-04-02", testHandler: testLocalMidnight,
							zone: "GB", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
						},
						{
							name: "2024-10-27", testHandler: testLocalMidnight,
							zone: "GB", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
						},
						{
							name: "2024-10-28", testHandler: testLocalMidnight,
							zone: "GB", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
						},
						{
							name: "2024-10-29", testHandler: testLocalMidnight,
							zone: "GB", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
						},
					}},
				{
					name: "YesterdayMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testYesterdayMidnight,
							zone: "GB", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
						},
						{
							name: "2024-04-01", testHandler: testYesterdayMidnight,
							zone: "GB", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
						},
						{
							name: "2024-04-02", testHandler: testYesterdayMidnight,
							zone: "GB", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
						},
						{
							name: "2024-10-27", testHandler: testYesterdayMidnight,
							zone: "GB", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
						},
						{
							name: "2024-10-28", testHandler: testYesterdayMidnight,
							zone: "GB", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
						},
						{
							name: "2024-10-29", testHandler: testYesterdayMidnight,
							zone: "GB", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
						},
					}},
				{
					name: "TomorrowMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testTomorrowMidnight,
							zone: "GB", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
						},
						{
							name: "2024-04-01", testHandler: testTomorrowMidnight,
							zone: "GB", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
						},
						{
							name: "2024-04-02", testHandler: testTomorrowMidnight,
							zone: "GB", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
						},
						{
							name: "2024-10-27", testHandler: testTomorrowMidnight,
							zone: "GB", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
						},
						{
							name: "2024-10-28", testHandler: testTomorrowMidnight,
							zone: "GB", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
						},
						{
							name: "2024-10-29", testHandler: testTomorrowMidnight,
							zone: "GB", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
						},
					}},
			},
		},
		{
			name: "GB-Eire",
			children: []testCase{
				{
					name: "LocalMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testLocalMidnight,
							zone: "GB-Eire", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
						},
						{
							name: "2024-04-01", testHandler: testLocalMidnight,
							zone: "GB-Eire", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
						},
						{
							name: "2024-04-02", testHandler: testLocalMidnight,
							zone: "GB-Eire", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
						},
						{
							name: "2024-10-27", testHandler: testLocalMidnight,
							zone: "GB-Eire", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
						},
						{
							name: "2024-10-28", testHandler: testLocalMidnight,
							zone: "GB-Eire", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
						},
						{
							name: "2024-10-29", testHandler: testLocalMidnight,
							zone: "GB-Eire", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
						},
					}},
				{
					name: "YesterdayMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testYesterdayMidnight,
							zone: "GB-Eire", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
						},
						{
							name: "2024-04-01", testHandler: testYesterdayMidnight,
							zone: "GB-Eire", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
						},
						{
							name: "2024-04-02", testHandler: testYesterdayMidnight,
							zone: "GB-Eire", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
						},
						{
							name: "2024-10-27", testHandler: testYesterdayMidnight,
							zone: "GB-Eire", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
						},
						{
							name: "2024-10-28", testHandler: testYesterdayMidnight,
							zone: "GB-Eire", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
						},
						{
							name: "2024-10-29", testHandler: testYesterdayMidnight,
							zone: "GB-Eire", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
						},
					}},
				{
					name: "TomorrowMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testTomorrowMidnight,
							zone: "GB-Eire", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
						},
						{
							name: "2024-04-01", testHandler: testTomorrowMidnight,
							zone: "GB-Eire", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
						},
						{
							name: "2024-04-02", testHandler: testTomorrowMidnight,
							zone: "GB-Eire", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
						},
						{
							name: "2024-10-27", testHandler: testTomorrowMidnight,
							zone: "GB-Eire", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
						},
						{
							name: "2024-10-28", testHandler: testTomorrowMidnight,
							zone: "GB-Eire", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
						},
						{
							name: "2024-10-29", testHandler: testTomorrowMidnight,
							zone: "GB-Eire", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
						},
					}},
			},
		},
		{
			name: "Israel",
			children: []testCase{
				{
					name: "LocalMidnight", children: []testCase{
						{
							name: "2024-03-29", testHandler: testLocalMidnight,
							zone: "Israel", dstHandover: "2024-03-30T00:00:00+03:00",
							date: "2024-03-29T00:00:00+02:00", expectedDate: "2024-03-29T00:00:00+02:00",
						},
						{
							name: "2024-03-30", testHandler: testLocalMidnight,
							zone: "Israel", dstHandover: "2024-03-30T00:00:00+03:00",
							date: "2024-03-30T00:00:00+03:00", expectedDate: "2024-03-30T00:00:00+03:00",
						},
						{
							name: "2024-03-31", testHandler: testLocalMidnight,
							zone: "Israel", dstHandover: "2024-03-30T00:00:00+03:00",
							date: "2024-03-31T00:00:00+03:00", expectedDate: "2024-03-31T00:00:00+03:00",
						},
						{
							name: "2024-10-27", testHandler: testLocalMidnight,
							zone: "Israel", dstHandover: "2024-10-28T00:00:00+02:00",
							date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
						},
						{
							name: "2024-10-28", testHandler: testLocalMidnight,
							zone: "Israel", dstHandover: "2024-10-28T00:00:00+02:00",
							date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
						},
						{
							name: "2024-10-29", testHandler: testLocalMidnight,
							zone: "Israel", dstHandover: "2024-10-28T00:00:00+02:00",
							date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
						},
					}},
				{
					name: "YesterdayMidnight", children: []testCase{
						{
							name: "2024-03-29", testHandler: testYesterdayMidnight,
							zone: "Israel", dstHandover: "2024-03-30T00:00:00+03:00",
							date: "2024-03-29T00:00:00+02:00", expectedDate: "2024-03-29T00:00:00+02:00",
						},
						{
							name: "2024-03-30", testHandler: testYesterdayMidnight,
							zone: "Israel", dstHandover: "2024-03-30T00:00:00+03:00",
							date: "2024-03-30T00:00:00+03:00", expectedDate: "2024-03-30T00:00:00+03:00",
						},
						{
							name: "2024-03-31", testHandler: testYesterdayMidnight,
							zone: "Israel", dstHandover: "2024-03-30T00:00:00+03:00",
							date: "2024-03-31T00:00:00+03:00", expectedDate: "2024-03-31T00:00:00+03:00",
						},
						{
							name: "2024-10-27", testHandler: testYesterdayMidnight,
							zone: "Israel", dstHandover: "2024-10-28T00:00:00+02:00",
							date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
						},
						{
							name: "2024-10-28", testHandler: testYesterdayMidnight,
							zone: "Israel", dstHandover: "2024-10-28T00:00:00+02:00",
							date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
						},
						{
							name: "2024-10-29", testHandler: testYesterdayMidnight,
							zone: "Israel", dstHandover: "2024-10-28T00:00:00+02:00",
							date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
						},
					}},
				{
					name: "TomorrowMidnight", children: []testCase{
						{
							name: "2024-03-29", testHandler: testTomorrowMidnight,
							zone: "Israel", dstHandover: "2024-03-30T00:00:00+03:00",
							date: "2024-03-29T00:00:00+02:00", expectedDate: "2024-03-29T00:00:00+02:00",
						},
						{
							name: "2024-03-30", testHandler: testTomorrowMidnight,
							zone: "Israel", dstHandover: "2024-03-30T00:00:00+03:00",
							date: "2024-03-30T00:00:00+03:00", expectedDate: "2024-03-30T00:00:00+03:00",
						},
						{
							name: "2024-03-31", testHandler: testTomorrowMidnight,
							zone: "Israel", dstHandover: "2024-03-30T00:00:00+03:00",
							date: "2024-03-31T00:00:00+03:00", expectedDate: "2024-03-31T00:00:00+03:00",
						},
						{
							name: "2024-10-27", testHandler: testTomorrowMidnight,
							zone: "Israel", dstHandover: "2024-10-28T00:00:00+02:00",
							date: "2024-10-27T00:00:00+03:00", expectedDate: "2024-10-27T00:00:00+03:00",
						},
						{
							name: "2024-10-28", testHandler: testTomorrowMidnight,
							zone: "Israel", dstHandover: "2024-10-28T00:00:00+02:00",
							date: "2024-10-28T00:00:00+02:00", expectedDate: "2024-10-28T00:00:00+02:00",
						},
						{
							name: "2024-10-29", testHandler: testTomorrowMidnight,
							zone: "Israel", dstHandover: "2024-10-28T00:00:00+02:00",
							date: "2024-10-29T00:00:00+02:00", expectedDate: "2024-10-29T00:00:00+02:00",
						},
					}},
			},
		},
		{
			name: "MET",
			children: []testCase{
				{
					name: "LocalMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testLocalMidnight,
							zone: "MET", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
						},
						{
							name: "2024-04-01", testHandler: testLocalMidnight,
							zone: "MET", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
						},
						{
							name: "2024-04-02", testHandler: testLocalMidnight,
							zone: "MET", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
						},
						{
							name: "2024-10-27", testHandler: testLocalMidnight,
							zone: "MET", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
						},
						{
							name: "2024-10-28", testHandler: testLocalMidnight,
							zone: "MET", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
						},
						{
							name: "2024-10-29", testHandler: testLocalMidnight,
							zone: "MET", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
						},
					}},
				{
					name: "YesterdayMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testYesterdayMidnight,
							zone: "MET", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
						},
						{
							name: "2024-04-01", testHandler: testYesterdayMidnight,
							zone: "MET", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
						},
						{
							name: "2024-04-02", testHandler: testYesterdayMidnight,
							zone: "MET", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
						},
						{
							name: "2024-10-27", testHandler: testYesterdayMidnight,
							zone: "MET", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
						},
						{
							name: "2024-10-28", testHandler: testYesterdayMidnight,
							zone: "MET", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
						},
						{
							name: "2024-10-29", testHandler: testYesterdayMidnight,
							zone: "MET", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
						},
					}},
				{
					name: "TomorrowMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testTomorrowMidnight,
							zone: "MET", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
						},
						{
							name: "2024-04-01", testHandler: testTomorrowMidnight,
							zone: "MET", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
						},
						{
							name: "2024-04-02", testHandler: testTomorrowMidnight,
							zone: "MET", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
						},
						{
							name: "2024-10-27", testHandler: testTomorrowMidnight,
							zone: "MET", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
						},
						{
							name: "2024-10-28", testHandler: testTomorrowMidnight,
							zone: "MET", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
						},
						{
							name: "2024-10-29", testHandler: testTomorrowMidnight,
							zone: "MET", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
						},
					}},
			},
		},
		{
			name: "MST7MDT",
			children: []testCase{
				{
					name: "LocalMidnight", children: []testCase{
						{
							name: "2024-03-10", testHandler: testLocalMidnight,
							zone: "MST7MDT", dstHandover: "2024-03-11T00:00:00-06:00",
							date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
						},
						{
							name: "2024-03-11", testHandler: testLocalMidnight,
							zone: "MST7MDT", dstHandover: "2024-03-11T00:00:00-06:00",
							date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
						},
						{
							name: "2024-03-12", testHandler: testLocalMidnight,
							zone: "MST7MDT", dstHandover: "2024-03-11T00:00:00-06:00",
							date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
						},
						{
							name: "2024-11-03", testHandler: testLocalMidnight,
							zone: "MST7MDT", dstHandover: "2024-11-04T00:00:00-07:00",
							date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
						},
						{
							name: "2024-11-04", testHandler: testLocalMidnight,
							zone: "MST7MDT", dstHandover: "2024-11-04T00:00:00-07:00",
							date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
						},
						{
							name: "2024-11-05", testHandler: testLocalMidnight,
							zone: "MST7MDT", dstHandover: "2024-11-04T00:00:00-07:00",
							date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
						},
					}},
				{
					name: "YesterdayMidnight", children: []testCase{
						{
							name: "2024-03-10", testHandler: testYesterdayMidnight,
							zone: "MST7MDT", dstHandover: "2024-03-11T00:00:00-06:00",
							date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
						},
						{
							name: "2024-03-11", testHandler: testYesterdayMidnight,
							zone: "MST7MDT", dstHandover: "2024-03-11T00:00:00-06:00",
							date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
						},
						{
							name: "2024-03-12", testHandler: testYesterdayMidnight,
							zone: "MST7MDT", dstHandover: "2024-03-11T00:00:00-06:00",
							date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
						},
						{
							name: "2024-11-03", testHandler: testYesterdayMidnight,
							zone: "MST7MDT", dstHandover: "2024-11-04T00:00:00-07:00",
							date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
						},
						{
							name: "2024-11-04", testHandler: testYesterdayMidnight,
							zone: "MST7MDT", dstHandover: "2024-11-04T00:00:00-07:00",
							date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
						},
						{
							name: "2024-11-05", testHandler: testYesterdayMidnight,
							zone: "MST7MDT", dstHandover: "2024-11-04T00:00:00-07:00",
							date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
						},
					}},
				{
					name: "TomorrowMidnight", children: []testCase{
						{
							name: "2024-03-10", testHandler: testTomorrowMidnight,
							zone: "MST7MDT", dstHandover: "2024-03-11T00:00:00-06:00",
							date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
						},
						{
							name: "2024-03-11", testHandler: testTomorrowMidnight,
							zone: "MST7MDT", dstHandover: "2024-03-11T00:00:00-06:00",
							date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
						},
						{
							name: "2024-03-12", testHandler: testTomorrowMidnight,
							zone: "MST7MDT", dstHandover: "2024-03-11T00:00:00-06:00",
							date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
						},
						{
							name: "2024-11-03", testHandler: testTomorrowMidnight,
							zone: "MST7MDT", dstHandover: "2024-11-04T00:00:00-07:00",
							date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
						},
						{
							name: "2024-11-04", testHandler: testTomorrowMidnight,
							zone: "MST7MDT", dstHandover: "2024-11-04T00:00:00-07:00",
							date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
						},
						{
							name: "2024-11-05", testHandler: testTomorrowMidnight,
							zone: "MST7MDT", dstHandover: "2024-11-04T00:00:00-07:00",
							date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
						},
					}},
			},
		},
		{
			name: "Mexico",
			children: []testCase{
				{
					name: "BajaNorte",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "Mexico/BajaNorte", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "Mexico/BajaNorte", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "Mexico/BajaNorte", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "Mexico/BajaNorte", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "Mexico/BajaNorte", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "Mexico/BajaNorte", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "Mexico/BajaNorte", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "Mexico/BajaNorte", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "Mexico/BajaNorte", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "Mexico/BajaNorte", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "Mexico/BajaNorte", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "Mexico/BajaNorte", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "Mexico/BajaNorte", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "Mexico/BajaNorte", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "Mexico/BajaNorte", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "Mexico/BajaNorte", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "Mexico/BajaNorte", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "Mexico/BajaNorte", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
					},
				},
			},
		},
		{
			name: "NZ",
			children: []testCase{
				{
					name: "LocalMidnight", children: []testCase{
						{
							name: "2024-09-29", testHandler: testLocalMidnight,
							zone: "NZ", dstHandover: "2024-09-30T00:00:00+13:00",
							date: "2024-09-29T00:00:00+12:00", expectedDate: "2024-09-29T00:00:00+12:00",
						},
						{
							name: "2024-09-30", testHandler: testLocalMidnight,
							zone: "NZ", dstHandover: "2024-09-30T00:00:00+13:00",
							date: "2024-09-30T00:00:00+13:00", expectedDate: "2024-09-30T00:00:00+13:00",
						},
						{
							name: "2024-10-01", testHandler: testLocalMidnight,
							zone: "NZ", dstHandover: "2024-09-30T00:00:00+13:00",
							date: "2024-10-01T00:00:00+13:00", expectedDate: "2024-10-01T00:00:00+13:00",
						},
						{
							name: "2024-04-07", testHandler: testLocalMidnight,
							zone: "NZ", dstHandover: "2024-04-08T00:00:00+12:00",
							date: "2024-04-07T00:00:00+13:00", expectedDate: "2024-04-07T00:00:00+13:00",
						},
						{
							name: "2024-04-08", testHandler: testLocalMidnight,
							zone: "NZ", dstHandover: "2024-04-08T00:00:00+12:00",
							date: "2024-04-08T00:00:00+12:00", expectedDate: "2024-04-08T00:00:00+12:00",
						},
						{
							name: "2024-04-09", testHandler: testLocalMidnight,
							zone: "NZ", dstHandover: "2024-04-08T00:00:00+12:00",
							date: "2024-04-09T00:00:00+12:00", expectedDate: "2024-04-09T00:00:00+12:00",
						},
					}},
				{
					name: "YesterdayMidnight", children: []testCase{
						{
							name: "2024-09-29", testHandler: testYesterdayMidnight,
							zone: "NZ", dstHandover: "2024-09-30T00:00:00+13:00",
							date: "2024-09-29T00:00:00+12:00", expectedDate: "2024-09-29T00:00:00+12:00",
						},
						{
							name: "2024-09-30", testHandler: testYesterdayMidnight,
							zone: "NZ", dstHandover: "2024-09-30T00:00:00+13:00",
							date: "2024-09-30T00:00:00+13:00", expectedDate: "2024-09-30T00:00:00+13:00",
						},
						{
							name: "2024-10-01", testHandler: testYesterdayMidnight,
							zone: "NZ", dstHandover: "2024-09-30T00:00:00+13:00",
							date: "2024-10-01T00:00:00+13:00", expectedDate: "2024-10-01T00:00:00+13:00",
						},
						{
							name: "2024-04-07", testHandler: testYesterdayMidnight,
							zone: "NZ", dstHandover: "2024-04-08T00:00:00+12:00",
							date: "2024-04-07T00:00:00+13:00", expectedDate: "2024-04-07T00:00:00+13:00",
						},
						{
							name: "2024-04-08", testHandler: testYesterdayMidnight,
							zone: "NZ", dstHandover: "2024-04-08T00:00:00+12:00",
							date: "2024-04-08T00:00:00+12:00", expectedDate: "2024-04-08T00:00:00+12:00",
						},
						{
							name: "2024-04-09", testHandler: testYesterdayMidnight,
							zone: "NZ", dstHandover: "2024-04-08T00:00:00+12:00",
							date: "2024-04-09T00:00:00+12:00", expectedDate: "2024-04-09T00:00:00+12:00",
						},
					}},
				{
					name: "TomorrowMidnight", children: []testCase{
						{
							name: "2024-09-29", testHandler: testTomorrowMidnight,
							zone: "NZ", dstHandover: "2024-09-30T00:00:00+13:00",
							date: "2024-09-29T00:00:00+12:00", expectedDate: "2024-09-29T00:00:00+12:00",
						},
						{
							name: "2024-09-30", testHandler: testTomorrowMidnight,
							zone: "NZ", dstHandover: "2024-09-30T00:00:00+13:00",
							date: "2024-09-30T00:00:00+13:00", expectedDate: "2024-09-30T00:00:00+13:00",
						},
						{
							name: "2024-10-01", testHandler: testTomorrowMidnight,
							zone: "NZ", dstHandover: "2024-09-30T00:00:00+13:00",
							date: "2024-10-01T00:00:00+13:00", expectedDate: "2024-10-01T00:00:00+13:00",
						},
						{
							name: "2024-04-07", testHandler: testTomorrowMidnight,
							zone: "NZ", dstHandover: "2024-04-08T00:00:00+12:00",
							date: "2024-04-07T00:00:00+13:00", expectedDate: "2024-04-07T00:00:00+13:00",
						},
						{
							name: "2024-04-08", testHandler: testTomorrowMidnight,
							zone: "NZ", dstHandover: "2024-04-08T00:00:00+12:00",
							date: "2024-04-08T00:00:00+12:00", expectedDate: "2024-04-08T00:00:00+12:00",
						},
						{
							name: "2024-04-09", testHandler: testTomorrowMidnight,
							zone: "NZ", dstHandover: "2024-04-08T00:00:00+12:00",
							date: "2024-04-09T00:00:00+12:00", expectedDate: "2024-04-09T00:00:00+12:00",
						},
					}},
			},
		},
		{
			name: "NZ-CHAT",
			children: []testCase{
				{
					name: "LocalMidnight", children: []testCase{
						{
							name: "2024-09-29", testHandler: testLocalMidnight,
							zone: "NZ-CHAT", dstHandover: "2024-09-30T00:00:00+13:45",
							date: "2024-09-29T00:00:00+12:45", expectedDate: "2024-09-29T00:00:00+12:45",
						},
						{
							name: "2024-09-30", testHandler: testLocalMidnight,
							zone: "NZ-CHAT", dstHandover: "2024-09-30T00:00:00+13:45",
							date: "2024-09-30T00:00:00+13:45", expectedDate: "2024-09-30T00:00:00+13:45",
						},
						{
							name: "2024-10-01", testHandler: testLocalMidnight,
							zone: "NZ-CHAT", dstHandover: "2024-09-30T00:00:00+13:45",
							date: "2024-10-01T00:00:00+13:45", expectedDate: "2024-10-01T00:00:00+13:45",
						},
						{
							name: "2024-04-07", testHandler: testLocalMidnight,
							zone: "NZ-CHAT", dstHandover: "2024-04-08T00:00:00+12:45",
							date: "2024-04-07T00:00:00+13:45", expectedDate: "2024-04-07T00:00:00+13:45",
						},
						{
							name: "2024-04-08", testHandler: testLocalMidnight,
							zone: "NZ-CHAT", dstHandover: "2024-04-08T00:00:00+12:45",
							date: "2024-04-08T00:00:00+12:45", expectedDate: "2024-04-08T00:00:00+12:45",
						},
						{
							name: "2024-04-09", testHandler: testLocalMidnight,
							zone: "NZ-CHAT", dstHandover: "2024-04-08T00:00:00+12:45",
							date: "2024-04-09T00:00:00+12:45", expectedDate: "2024-04-09T00:00:00+12:45",
						},
					}},
				{
					name: "YesterdayMidnight", children: []testCase{
						{
							name: "2024-09-29", testHandler: testYesterdayMidnight,
							zone: "NZ-CHAT", dstHandover: "2024-09-30T00:00:00+13:45",
							date: "2024-09-29T00:00:00+12:45", expectedDate: "2024-09-29T00:00:00+12:45",
						},
						{
							name: "2024-09-30", testHandler: testYesterdayMidnight,
							zone: "NZ-CHAT", dstHandover: "2024-09-30T00:00:00+13:45",
							date: "2024-09-30T00:00:00+13:45", expectedDate: "2024-09-30T00:00:00+13:45",
						},
						{
							name: "2024-10-01", testHandler: testYesterdayMidnight,
							zone: "NZ-CHAT", dstHandover: "2024-09-30T00:00:00+13:45",
							date: "2024-10-01T00:00:00+13:45", expectedDate: "2024-10-01T00:00:00+13:45",
						},
						{
							name: "2024-04-07", testHandler: testYesterdayMidnight,
							zone: "NZ-CHAT", dstHandover: "2024-04-08T00:00:00+12:45",
							date: "2024-04-07T00:00:00+13:45", expectedDate: "2024-04-07T00:00:00+13:45",
						},
						{
							name: "2024-04-08", testHandler: testYesterdayMidnight,
							zone: "NZ-CHAT", dstHandover: "2024-04-08T00:00:00+12:45",
							date: "2024-04-08T00:00:00+12:45", expectedDate: "2024-04-08T00:00:00+12:45",
						},
						{
							name: "2024-04-09", testHandler: testYesterdayMidnight,
							zone: "NZ-CHAT", dstHandover: "2024-04-08T00:00:00+12:45",
							date: "2024-04-09T00:00:00+12:45", expectedDate: "2024-04-09T00:00:00+12:45",
						},
					}},
				{
					name: "TomorrowMidnight", children: []testCase{
						{
							name: "2024-09-29", testHandler: testTomorrowMidnight,
							zone: "NZ-CHAT", dstHandover: "2024-09-30T00:00:00+13:45",
							date: "2024-09-29T00:00:00+12:45", expectedDate: "2024-09-29T00:00:00+12:45",
						},
						{
							name: "2024-09-30", testHandler: testTomorrowMidnight,
							zone: "NZ-CHAT", dstHandover: "2024-09-30T00:00:00+13:45",
							date: "2024-09-30T00:00:00+13:45", expectedDate: "2024-09-30T00:00:00+13:45",
						},
						{
							name: "2024-10-01", testHandler: testTomorrowMidnight,
							zone: "NZ-CHAT", dstHandover: "2024-09-30T00:00:00+13:45",
							date: "2024-10-01T00:00:00+13:45", expectedDate: "2024-10-01T00:00:00+13:45",
						},
						{
							name: "2024-04-07", testHandler: testTomorrowMidnight,
							zone: "NZ-CHAT", dstHandover: "2024-04-08T00:00:00+12:45",
							date: "2024-04-07T00:00:00+13:45", expectedDate: "2024-04-07T00:00:00+13:45",
						},
						{
							name: "2024-04-08", testHandler: testTomorrowMidnight,
							zone: "NZ-CHAT", dstHandover: "2024-04-08T00:00:00+12:45",
							date: "2024-04-08T00:00:00+12:45", expectedDate: "2024-04-08T00:00:00+12:45",
						},
						{
							name: "2024-04-09", testHandler: testTomorrowMidnight,
							zone: "NZ-CHAT", dstHandover: "2024-04-08T00:00:00+12:45",
							date: "2024-04-09T00:00:00+12:45", expectedDate: "2024-04-09T00:00:00+12:45",
						},
					}},
			},
		},
		{
			name: "Navajo",
			children: []testCase{
				{
					name: "LocalMidnight", children: []testCase{
						{
							name: "2024-03-10", testHandler: testLocalMidnight,
							zone: "Navajo", dstHandover: "2024-03-11T00:00:00-06:00",
							date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
						},
						{
							name: "2024-03-11", testHandler: testLocalMidnight,
							zone: "Navajo", dstHandover: "2024-03-11T00:00:00-06:00",
							date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
						},
						{
							name: "2024-03-12", testHandler: testLocalMidnight,
							zone: "Navajo", dstHandover: "2024-03-11T00:00:00-06:00",
							date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
						},
						{
							name: "2024-11-03", testHandler: testLocalMidnight,
							zone: "Navajo", dstHandover: "2024-11-04T00:00:00-07:00",
							date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
						},
						{
							name: "2024-11-04", testHandler: testLocalMidnight,
							zone: "Navajo", dstHandover: "2024-11-04T00:00:00-07:00",
							date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
						},
						{
							name: "2024-11-05", testHandler: testLocalMidnight,
							zone: "Navajo", dstHandover: "2024-11-04T00:00:00-07:00",
							date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
						},
					}},
				{
					name: "YesterdayMidnight", children: []testCase{
						{
							name: "2024-03-10", testHandler: testYesterdayMidnight,
							zone: "Navajo", dstHandover: "2024-03-11T00:00:00-06:00",
							date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
						},
						{
							name: "2024-03-11", testHandler: testYesterdayMidnight,
							zone: "Navajo", dstHandover: "2024-03-11T00:00:00-06:00",
							date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
						},
						{
							name: "2024-03-12", testHandler: testYesterdayMidnight,
							zone: "Navajo", dstHandover: "2024-03-11T00:00:00-06:00",
							date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
						},
						{
							name: "2024-11-03", testHandler: testYesterdayMidnight,
							zone: "Navajo", dstHandover: "2024-11-04T00:00:00-07:00",
							date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
						},
						{
							name: "2024-11-04", testHandler: testYesterdayMidnight,
							zone: "Navajo", dstHandover: "2024-11-04T00:00:00-07:00",
							date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
						},
						{
							name: "2024-11-05", testHandler: testYesterdayMidnight,
							zone: "Navajo", dstHandover: "2024-11-04T00:00:00-07:00",
							date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
						},
					}},
				{
					name: "TomorrowMidnight", children: []testCase{
						{
							name: "2024-03-10", testHandler: testTomorrowMidnight,
							zone: "Navajo", dstHandover: "2024-03-11T00:00:00-06:00",
							date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
						},
						{
							name: "2024-03-11", testHandler: testTomorrowMidnight,
							zone: "Navajo", dstHandover: "2024-03-11T00:00:00-06:00",
							date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
						},
						{
							name: "2024-03-12", testHandler: testTomorrowMidnight,
							zone: "Navajo", dstHandover: "2024-03-11T00:00:00-06:00",
							date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
						},
						{
							name: "2024-11-03", testHandler: testTomorrowMidnight,
							zone: "Navajo", dstHandover: "2024-11-04T00:00:00-07:00",
							date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
						},
						{
							name: "2024-11-04", testHandler: testTomorrowMidnight,
							zone: "Navajo", dstHandover: "2024-11-04T00:00:00-07:00",
							date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
						},
						{
							name: "2024-11-05", testHandler: testTomorrowMidnight,
							zone: "Navajo", dstHandover: "2024-11-04T00:00:00-07:00",
							date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
						},
					}},
			},
		},
		{
			name: "PST8PDT",
			children: []testCase{
				{
					name: "LocalMidnight", children: []testCase{
						{
							name: "2024-03-10", testHandler: testLocalMidnight,
							zone: "PST8PDT", dstHandover: "2024-03-11T00:00:00-07:00",
							date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
						},
						{
							name: "2024-03-11", testHandler: testLocalMidnight,
							zone: "PST8PDT", dstHandover: "2024-03-11T00:00:00-07:00",
							date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
						},
						{
							name: "2024-03-12", testHandler: testLocalMidnight,
							zone: "PST8PDT", dstHandover: "2024-03-11T00:00:00-07:00",
							date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
						},
						{
							name: "2024-11-03", testHandler: testLocalMidnight,
							zone: "PST8PDT", dstHandover: "2024-11-04T00:00:00-08:00",
							date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
						},
						{
							name: "2024-11-04", testHandler: testLocalMidnight,
							zone: "PST8PDT", dstHandover: "2024-11-04T00:00:00-08:00",
							date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
						},
						{
							name: "2024-11-05", testHandler: testLocalMidnight,
							zone: "PST8PDT", dstHandover: "2024-11-04T00:00:00-08:00",
							date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
						},
					}},
				{
					name: "YesterdayMidnight", children: []testCase{
						{
							name: "2024-03-10", testHandler: testYesterdayMidnight,
							zone: "PST8PDT", dstHandover: "2024-03-11T00:00:00-07:00",
							date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
						},
						{
							name: "2024-03-11", testHandler: testYesterdayMidnight,
							zone: "PST8PDT", dstHandover: "2024-03-11T00:00:00-07:00",
							date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
						},
						{
							name: "2024-03-12", testHandler: testYesterdayMidnight,
							zone: "PST8PDT", dstHandover: "2024-03-11T00:00:00-07:00",
							date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
						},
						{
							name: "2024-11-03", testHandler: testYesterdayMidnight,
							zone: "PST8PDT", dstHandover: "2024-11-04T00:00:00-08:00",
							date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
						},
						{
							name: "2024-11-04", testHandler: testYesterdayMidnight,
							zone: "PST8PDT", dstHandover: "2024-11-04T00:00:00-08:00",
							date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
						},
						{
							name: "2024-11-05", testHandler: testYesterdayMidnight,
							zone: "PST8PDT", dstHandover: "2024-11-04T00:00:00-08:00",
							date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
						},
					}},
				{
					name: "TomorrowMidnight", children: []testCase{
						{
							name: "2024-03-10", testHandler: testTomorrowMidnight,
							zone: "PST8PDT", dstHandover: "2024-03-11T00:00:00-07:00",
							date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
						},
						{
							name: "2024-03-11", testHandler: testTomorrowMidnight,
							zone: "PST8PDT", dstHandover: "2024-03-11T00:00:00-07:00",
							date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
						},
						{
							name: "2024-03-12", testHandler: testTomorrowMidnight,
							zone: "PST8PDT", dstHandover: "2024-03-11T00:00:00-07:00",
							date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
						},
						{
							name: "2024-11-03", testHandler: testTomorrowMidnight,
							zone: "PST8PDT", dstHandover: "2024-11-04T00:00:00-08:00",
							date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
						},
						{
							name: "2024-11-04", testHandler: testTomorrowMidnight,
							zone: "PST8PDT", dstHandover: "2024-11-04T00:00:00-08:00",
							date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
						},
						{
							name: "2024-11-05", testHandler: testTomorrowMidnight,
							zone: "PST8PDT", dstHandover: "2024-11-04T00:00:00-08:00",
							date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
						},
					}},
			},
		},
		{
			name: "Pacific",
			children: []testCase{
				{
					name: "Auckland",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-09-29", testHandler: testLocalMidnight,
									zone: "Pacific/Auckland", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-09-29T00:00:00+12:00", expectedDate: "2024-09-29T00:00:00+12:00",
								},
								{
									name: "2024-09-30", testHandler: testLocalMidnight,
									zone: "Pacific/Auckland", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-09-30T00:00:00+13:00", expectedDate: "2024-09-30T00:00:00+13:00",
								},
								{
									name: "2024-10-01", testHandler: testLocalMidnight,
									zone: "Pacific/Auckland", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-10-01T00:00:00+13:00", expectedDate: "2024-10-01T00:00:00+13:00",
								},
								{
									name: "2024-04-07", testHandler: testLocalMidnight,
									zone: "Pacific/Auckland", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-07T00:00:00+13:00", expectedDate: "2024-04-07T00:00:00+13:00",
								},
								{
									name: "2024-04-08", testHandler: testLocalMidnight,
									zone: "Pacific/Auckland", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-08T00:00:00+12:00", expectedDate: "2024-04-08T00:00:00+12:00",
								},
								{
									name: "2024-04-09", testHandler: testLocalMidnight,
									zone: "Pacific/Auckland", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-09T00:00:00+12:00", expectedDate: "2024-04-09T00:00:00+12:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-09-29", testHandler: testYesterdayMidnight,
									zone: "Pacific/Auckland", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-09-29T00:00:00+12:00", expectedDate: "2024-09-29T00:00:00+12:00",
								},
								{
									name: "2024-09-30", testHandler: testYesterdayMidnight,
									zone: "Pacific/Auckland", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-09-30T00:00:00+13:00", expectedDate: "2024-09-30T00:00:00+13:00",
								},
								{
									name: "2024-10-01", testHandler: testYesterdayMidnight,
									zone: "Pacific/Auckland", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-10-01T00:00:00+13:00", expectedDate: "2024-10-01T00:00:00+13:00",
								},
								{
									name: "2024-04-07", testHandler: testYesterdayMidnight,
									zone: "Pacific/Auckland", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-07T00:00:00+13:00", expectedDate: "2024-04-07T00:00:00+13:00",
								},
								{
									name: "2024-04-08", testHandler: testYesterdayMidnight,
									zone: "Pacific/Auckland", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-08T00:00:00+12:00", expectedDate: "2024-04-08T00:00:00+12:00",
								},
								{
									name: "2024-04-09", testHandler: testYesterdayMidnight,
									zone: "Pacific/Auckland", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-09T00:00:00+12:00", expectedDate: "2024-04-09T00:00:00+12:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-09-29", testHandler: testTomorrowMidnight,
									zone: "Pacific/Auckland", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-09-29T00:00:00+12:00", expectedDate: "2024-09-29T00:00:00+12:00",
								},
								{
									name: "2024-09-30", testHandler: testTomorrowMidnight,
									zone: "Pacific/Auckland", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-09-30T00:00:00+13:00", expectedDate: "2024-09-30T00:00:00+13:00",
								},
								{
									name: "2024-10-01", testHandler: testTomorrowMidnight,
									zone: "Pacific/Auckland", dstHandover: "2024-09-30T00:00:00+13:00",
									date: "2024-10-01T00:00:00+13:00", expectedDate: "2024-10-01T00:00:00+13:00",
								},
								{
									name: "2024-04-07", testHandler: testTomorrowMidnight,
									zone: "Pacific/Auckland", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-07T00:00:00+13:00", expectedDate: "2024-04-07T00:00:00+13:00",
								},
								{
									name: "2024-04-08", testHandler: testTomorrowMidnight,
									zone: "Pacific/Auckland", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-08T00:00:00+12:00", expectedDate: "2024-04-08T00:00:00+12:00",
								},
								{
									name: "2024-04-09", testHandler: testTomorrowMidnight,
									zone: "Pacific/Auckland", dstHandover: "2024-04-08T00:00:00+12:00",
									date: "2024-04-09T00:00:00+12:00", expectedDate: "2024-04-09T00:00:00+12:00",
								},
							}},
					},
				},
				{
					name: "Chatham",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-09-29", testHandler: testLocalMidnight,
									zone: "Pacific/Chatham", dstHandover: "2024-09-30T00:00:00+13:45",
									date: "2024-09-29T00:00:00+12:45", expectedDate: "2024-09-29T00:00:00+12:45",
								},
								{
									name: "2024-09-30", testHandler: testLocalMidnight,
									zone: "Pacific/Chatham", dstHandover: "2024-09-30T00:00:00+13:45",
									date: "2024-09-30T00:00:00+13:45", expectedDate: "2024-09-30T00:00:00+13:45",
								},
								{
									name: "2024-10-01", testHandler: testLocalMidnight,
									zone: "Pacific/Chatham", dstHandover: "2024-09-30T00:00:00+13:45",
									date: "2024-10-01T00:00:00+13:45", expectedDate: "2024-10-01T00:00:00+13:45",
								},
								{
									name: "2024-04-07", testHandler: testLocalMidnight,
									zone: "Pacific/Chatham", dstHandover: "2024-04-08T00:00:00+12:45",
									date: "2024-04-07T00:00:00+13:45", expectedDate: "2024-04-07T00:00:00+13:45",
								},
								{
									name: "2024-04-08", testHandler: testLocalMidnight,
									zone: "Pacific/Chatham", dstHandover: "2024-04-08T00:00:00+12:45",
									date: "2024-04-08T00:00:00+12:45", expectedDate: "2024-04-08T00:00:00+12:45",
								},
								{
									name: "2024-04-09", testHandler: testLocalMidnight,
									zone: "Pacific/Chatham", dstHandover: "2024-04-08T00:00:00+12:45",
									date: "2024-04-09T00:00:00+12:45", expectedDate: "2024-04-09T00:00:00+12:45",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-09-29", testHandler: testYesterdayMidnight,
									zone: "Pacific/Chatham", dstHandover: "2024-09-30T00:00:00+13:45",
									date: "2024-09-29T00:00:00+12:45", expectedDate: "2024-09-29T00:00:00+12:45",
								},
								{
									name: "2024-09-30", testHandler: testYesterdayMidnight,
									zone: "Pacific/Chatham", dstHandover: "2024-09-30T00:00:00+13:45",
									date: "2024-09-30T00:00:00+13:45", expectedDate: "2024-09-30T00:00:00+13:45",
								},
								{
									name: "2024-10-01", testHandler: testYesterdayMidnight,
									zone: "Pacific/Chatham", dstHandover: "2024-09-30T00:00:00+13:45",
									date: "2024-10-01T00:00:00+13:45", expectedDate: "2024-10-01T00:00:00+13:45",
								},
								{
									name: "2024-04-07", testHandler: testYesterdayMidnight,
									zone: "Pacific/Chatham", dstHandover: "2024-04-08T00:00:00+12:45",
									date: "2024-04-07T00:00:00+13:45", expectedDate: "2024-04-07T00:00:00+13:45",
								},
								{
									name: "2024-04-08", testHandler: testYesterdayMidnight,
									zone: "Pacific/Chatham", dstHandover: "2024-04-08T00:00:00+12:45",
									date: "2024-04-08T00:00:00+12:45", expectedDate: "2024-04-08T00:00:00+12:45",
								},
								{
									name: "2024-04-09", testHandler: testYesterdayMidnight,
									zone: "Pacific/Chatham", dstHandover: "2024-04-08T00:00:00+12:45",
									date: "2024-04-09T00:00:00+12:45", expectedDate: "2024-04-09T00:00:00+12:45",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-09-29", testHandler: testTomorrowMidnight,
									zone: "Pacific/Chatham", dstHandover: "2024-09-30T00:00:00+13:45",
									date: "2024-09-29T00:00:00+12:45", expectedDate: "2024-09-29T00:00:00+12:45",
								},
								{
									name: "2024-09-30", testHandler: testTomorrowMidnight,
									zone: "Pacific/Chatham", dstHandover: "2024-09-30T00:00:00+13:45",
									date: "2024-09-30T00:00:00+13:45", expectedDate: "2024-09-30T00:00:00+13:45",
								},
								{
									name: "2024-10-01", testHandler: testTomorrowMidnight,
									zone: "Pacific/Chatham", dstHandover: "2024-09-30T00:00:00+13:45",
									date: "2024-10-01T00:00:00+13:45", expectedDate: "2024-10-01T00:00:00+13:45",
								},
								{
									name: "2024-04-07", testHandler: testTomorrowMidnight,
									zone: "Pacific/Chatham", dstHandover: "2024-04-08T00:00:00+12:45",
									date: "2024-04-07T00:00:00+13:45", expectedDate: "2024-04-07T00:00:00+13:45",
								},
								{
									name: "2024-04-08", testHandler: testTomorrowMidnight,
									zone: "Pacific/Chatham", dstHandover: "2024-04-08T00:00:00+12:45",
									date: "2024-04-08T00:00:00+12:45", expectedDate: "2024-04-08T00:00:00+12:45",
								},
								{
									name: "2024-04-09", testHandler: testTomorrowMidnight,
									zone: "Pacific/Chatham", dstHandover: "2024-04-08T00:00:00+12:45",
									date: "2024-04-09T00:00:00+12:45", expectedDate: "2024-04-09T00:00:00+12:45",
								},
							}},
					},
				},
				{
					name: "Easter",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-09-07", testHandler: testLocalMidnight,
									zone: "Pacific/Easter", dstHandover: "2024-09-08T00:00:00-05:00",
									date: "2024-09-07T00:00:00-06:00", expectedDate: "2024-09-07T00:00:00-06:00",
								},
								{
									name: "2024-09-08", testHandler: testLocalMidnight,
									zone: "Pacific/Easter", dstHandover: "2024-09-08T00:00:00-05:00",
									date: "2024-09-08T00:00:00-05:00", expectedDate: "2024-09-08T00:00:00-05:00",
								},
								{
									name: "2024-09-09", testHandler: testLocalMidnight,
									zone: "Pacific/Easter", dstHandover: "2024-09-08T00:00:00-05:00",
									date: "2024-09-09T00:00:00-05:00", expectedDate: "2024-09-09T00:00:00-05:00",
								},
								{
									name: "2024-04-06", testHandler: testLocalMidnight,
									zone: "Pacific/Easter", dstHandover: "2024-04-07T00:00:00-06:00",
									date: "2024-04-06T00:00:00-05:00", expectedDate: "2024-04-06T00:00:00-05:00",
								},
								{
									name: "2024-04-07", testHandler: testLocalMidnight,
									zone: "Pacific/Easter", dstHandover: "2024-04-07T00:00:00-06:00",
									date: "2024-04-07T00:00:00-06:00", expectedDate: "2024-04-07T00:00:00-06:00",
								},
								{
									name: "2024-04-08", testHandler: testLocalMidnight,
									zone: "Pacific/Easter", dstHandover: "2024-04-07T00:00:00-06:00",
									date: "2024-04-08T00:00:00-06:00", expectedDate: "2024-04-08T00:00:00-06:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-09-07", testHandler: testYesterdayMidnight,
									zone: "Pacific/Easter", dstHandover: "2024-09-08T00:00:00-05:00",
									date: "2024-09-07T00:00:00-06:00", expectedDate: "2024-09-07T00:00:00-06:00",
								},
								{
									name: "2024-09-08", testHandler: testYesterdayMidnight,
									zone: "Pacific/Easter", dstHandover: "2024-09-08T00:00:00-05:00",
									date: "2024-09-08T00:00:00-05:00", expectedDate: "2024-09-08T00:00:00-05:00",
								},
								{
									name: "2024-09-09", testHandler: testYesterdayMidnight,
									zone: "Pacific/Easter", dstHandover: "2024-09-08T00:00:00-05:00",
									date: "2024-09-09T00:00:00-05:00", expectedDate: "2024-09-09T00:00:00-05:00",
								},
								{
									name: "2024-04-06", testHandler: testYesterdayMidnight,
									zone: "Pacific/Easter", dstHandover: "2024-04-07T00:00:00-06:00",
									date: "2024-04-06T00:00:00-05:00", expectedDate: "2024-04-06T00:00:00-05:00",
								},
								{
									name: "2024-04-07", testHandler: testYesterdayMidnight,
									zone: "Pacific/Easter", dstHandover: "2024-04-07T00:00:00-06:00",
									date: "2024-04-07T00:00:00-06:00", expectedDate: "2024-04-07T00:00:00-06:00",
								},
								{
									name: "2024-04-08", testHandler: testYesterdayMidnight,
									zone: "Pacific/Easter", dstHandover: "2024-04-07T00:00:00-06:00",
									date: "2024-04-08T00:00:00-06:00", expectedDate: "2024-04-08T00:00:00-06:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-09-07", testHandler: testTomorrowMidnight,
									zone: "Pacific/Easter", dstHandover: "2024-09-08T00:00:00-05:00",
									date: "2024-09-07T00:00:00-06:00", expectedDate: "2024-09-07T00:00:00-06:00",
								},
								{
									name: "2024-09-08", testHandler: testTomorrowMidnight,
									zone: "Pacific/Easter", dstHandover: "2024-09-08T00:00:00-05:00",
									date: "2024-09-08T00:00:00-05:00", expectedDate: "2024-09-08T00:00:00-05:00",
								},
								{
									name: "2024-09-09", testHandler: testTomorrowMidnight,
									zone: "Pacific/Easter", dstHandover: "2024-09-08T00:00:00-05:00",
									date: "2024-09-09T00:00:00-05:00", expectedDate: "2024-09-09T00:00:00-05:00",
								},
								{
									name: "2024-04-06", testHandler: testTomorrowMidnight,
									zone: "Pacific/Easter", dstHandover: "2024-04-07T00:00:00-06:00",
									date: "2024-04-06T00:00:00-05:00", expectedDate: "2024-04-06T00:00:00-05:00",
								},
								{
									name: "2024-04-07", testHandler: testTomorrowMidnight,
									zone: "Pacific/Easter", dstHandover: "2024-04-07T00:00:00-06:00",
									date: "2024-04-07T00:00:00-06:00", expectedDate: "2024-04-07T00:00:00-06:00",
								},
								{
									name: "2024-04-08", testHandler: testTomorrowMidnight,
									zone: "Pacific/Easter", dstHandover: "2024-04-07T00:00:00-06:00",
									date: "2024-04-08T00:00:00-06:00", expectedDate: "2024-04-08T00:00:00-06:00",
								},
							}},
					},
				},
				{
					name: "Norfolk",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testLocalMidnight,
									zone: "Pacific/Norfolk", dstHandover: "2024-10-07T00:00:00+12:00",
									date: "2024-10-06T00:00:00+11:00", expectedDate: "2024-10-06T00:00:00+11:00",
								},
								{
									name: "2024-10-07", testHandler: testLocalMidnight,
									zone: "Pacific/Norfolk", dstHandover: "2024-10-07T00:00:00+12:00",
									date: "2024-10-07T00:00:00+12:00", expectedDate: "2024-10-07T00:00:00+12:00",
								},
								{
									name: "2024-10-08", testHandler: testLocalMidnight,
									zone: "Pacific/Norfolk", dstHandover: "2024-10-07T00:00:00+12:00",
									date: "2024-10-08T00:00:00+12:00", expectedDate: "2024-10-08T00:00:00+12:00",
								},
								{
									name: "2024-04-07", testHandler: testLocalMidnight,
									zone: "Pacific/Norfolk", dstHandover: "2024-04-08T00:00:00+11:00",
									date: "2024-04-07T00:00:00+12:00", expectedDate: "2024-04-07T00:00:00+12:00",
								},
								{
									name: "2024-04-08", testHandler: testLocalMidnight,
									zone: "Pacific/Norfolk", dstHandover: "2024-04-08T00:00:00+11:00",
									date: "2024-04-08T00:00:00+11:00", expectedDate: "2024-04-08T00:00:00+11:00",
								},
								{
									name: "2024-04-09", testHandler: testLocalMidnight,
									zone: "Pacific/Norfolk", dstHandover: "2024-04-08T00:00:00+11:00",
									date: "2024-04-09T00:00:00+11:00", expectedDate: "2024-04-09T00:00:00+11:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testYesterdayMidnight,
									zone: "Pacific/Norfolk", dstHandover: "2024-10-07T00:00:00+12:00",
									date: "2024-10-06T00:00:00+11:00", expectedDate: "2024-10-06T00:00:00+11:00",
								},
								{
									name: "2024-10-07", testHandler: testYesterdayMidnight,
									zone: "Pacific/Norfolk", dstHandover: "2024-10-07T00:00:00+12:00",
									date: "2024-10-07T00:00:00+12:00", expectedDate: "2024-10-07T00:00:00+12:00",
								},
								{
									name: "2024-10-08", testHandler: testYesterdayMidnight,
									zone: "Pacific/Norfolk", dstHandover: "2024-10-07T00:00:00+12:00",
									date: "2024-10-08T00:00:00+12:00", expectedDate: "2024-10-08T00:00:00+12:00",
								},
								{
									name: "2024-04-07", testHandler: testYesterdayMidnight,
									zone: "Pacific/Norfolk", dstHandover: "2024-04-08T00:00:00+11:00",
									date: "2024-04-07T00:00:00+12:00", expectedDate: "2024-04-07T00:00:00+12:00",
								},
								{
									name: "2024-04-08", testHandler: testYesterdayMidnight,
									zone: "Pacific/Norfolk", dstHandover: "2024-04-08T00:00:00+11:00",
									date: "2024-04-08T00:00:00+11:00", expectedDate: "2024-04-08T00:00:00+11:00",
								},
								{
									name: "2024-04-09", testHandler: testYesterdayMidnight,
									zone: "Pacific/Norfolk", dstHandover: "2024-04-08T00:00:00+11:00",
									date: "2024-04-09T00:00:00+11:00", expectedDate: "2024-04-09T00:00:00+11:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-10-06", testHandler: testTomorrowMidnight,
									zone: "Pacific/Norfolk", dstHandover: "2024-10-07T00:00:00+12:00",
									date: "2024-10-06T00:00:00+11:00", expectedDate: "2024-10-06T00:00:00+11:00",
								},
								{
									name: "2024-10-07", testHandler: testTomorrowMidnight,
									zone: "Pacific/Norfolk", dstHandover: "2024-10-07T00:00:00+12:00",
									date: "2024-10-07T00:00:00+12:00", expectedDate: "2024-10-07T00:00:00+12:00",
								},
								{
									name: "2024-10-08", testHandler: testTomorrowMidnight,
									zone: "Pacific/Norfolk", dstHandover: "2024-10-07T00:00:00+12:00",
									date: "2024-10-08T00:00:00+12:00", expectedDate: "2024-10-08T00:00:00+12:00",
								},
								{
									name: "2024-04-07", testHandler: testTomorrowMidnight,
									zone: "Pacific/Norfolk", dstHandover: "2024-04-08T00:00:00+11:00",
									date: "2024-04-07T00:00:00+12:00", expectedDate: "2024-04-07T00:00:00+12:00",
								},
								{
									name: "2024-04-08", testHandler: testTomorrowMidnight,
									zone: "Pacific/Norfolk", dstHandover: "2024-04-08T00:00:00+11:00",
									date: "2024-04-08T00:00:00+11:00", expectedDate: "2024-04-08T00:00:00+11:00",
								},
								{
									name: "2024-04-09", testHandler: testTomorrowMidnight,
									zone: "Pacific/Norfolk", dstHandover: "2024-04-08T00:00:00+11:00",
									date: "2024-04-09T00:00:00+11:00", expectedDate: "2024-04-09T00:00:00+11:00",
								},
							}},
					},
				},
			},
		},
		{
			name: "Poland",
			children: []testCase{
				{
					name: "LocalMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testLocalMidnight,
							zone: "Poland", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
						},
						{
							name: "2024-04-01", testHandler: testLocalMidnight,
							zone: "Poland", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
						},
						{
							name: "2024-04-02", testHandler: testLocalMidnight,
							zone: "Poland", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
						},
						{
							name: "2024-10-27", testHandler: testLocalMidnight,
							zone: "Poland", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
						},
						{
							name: "2024-10-28", testHandler: testLocalMidnight,
							zone: "Poland", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
						},
						{
							name: "2024-10-29", testHandler: testLocalMidnight,
							zone: "Poland", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
						},
					}},
				{
					name: "YesterdayMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testYesterdayMidnight,
							zone: "Poland", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
						},
						{
							name: "2024-04-01", testHandler: testYesterdayMidnight,
							zone: "Poland", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
						},
						{
							name: "2024-04-02", testHandler: testYesterdayMidnight,
							zone: "Poland", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
						},
						{
							name: "2024-10-27", testHandler: testYesterdayMidnight,
							zone: "Poland", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
						},
						{
							name: "2024-10-28", testHandler: testYesterdayMidnight,
							zone: "Poland", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
						},
						{
							name: "2024-10-29", testHandler: testYesterdayMidnight,
							zone: "Poland", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
						},
					}},
				{
					name: "TomorrowMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testTomorrowMidnight,
							zone: "Poland", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-03-31T00:00:00+01:00", expectedDate: "2024-03-31T00:00:00+01:00",
						},
						{
							name: "2024-04-01", testHandler: testTomorrowMidnight,
							zone: "Poland", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-04-01T00:00:00+02:00", expectedDate: "2024-04-01T00:00:00+02:00",
						},
						{
							name: "2024-04-02", testHandler: testTomorrowMidnight,
							zone: "Poland", dstHandover: "2024-04-01T00:00:00+02:00",
							date: "2024-04-02T00:00:00+02:00", expectedDate: "2024-04-02T00:00:00+02:00",
						},
						{
							name: "2024-10-27", testHandler: testTomorrowMidnight,
							zone: "Poland", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-27T00:00:00+02:00", expectedDate: "2024-10-27T00:00:00+02:00",
						},
						{
							name: "2024-10-28", testHandler: testTomorrowMidnight,
							zone: "Poland", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-28T00:00:00+01:00", expectedDate: "2024-10-28T00:00:00+01:00",
						},
						{
							name: "2024-10-29", testHandler: testTomorrowMidnight,
							zone: "Poland", dstHandover: "2024-10-28T00:00:00+01:00",
							date: "2024-10-29T00:00:00+01:00", expectedDate: "2024-10-29T00:00:00+01:00",
						},
					}},
			},
		},
		{
			name: "Portugal",
			children: []testCase{
				{
					name: "LocalMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testLocalMidnight,
							zone: "Portugal", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
						},
						{
							name: "2024-04-01", testHandler: testLocalMidnight,
							zone: "Portugal", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
						},
						{
							name: "2024-04-02", testHandler: testLocalMidnight,
							zone: "Portugal", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
						},
						{
							name: "2024-10-27", testHandler: testLocalMidnight,
							zone: "Portugal", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
						},
						{
							name: "2024-10-28", testHandler: testLocalMidnight,
							zone: "Portugal", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
						},
						{
							name: "2024-10-29", testHandler: testLocalMidnight,
							zone: "Portugal", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
						},
					}},
				{
					name: "YesterdayMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testYesterdayMidnight,
							zone: "Portugal", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
						},
						{
							name: "2024-04-01", testHandler: testYesterdayMidnight,
							zone: "Portugal", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
						},
						{
							name: "2024-04-02", testHandler: testYesterdayMidnight,
							zone: "Portugal", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
						},
						{
							name: "2024-10-27", testHandler: testYesterdayMidnight,
							zone: "Portugal", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
						},
						{
							name: "2024-10-28", testHandler: testYesterdayMidnight,
							zone: "Portugal", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
						},
						{
							name: "2024-10-29", testHandler: testYesterdayMidnight,
							zone: "Portugal", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
						},
					}},
				{
					name: "TomorrowMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testTomorrowMidnight,
							zone: "Portugal", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
						},
						{
							name: "2024-04-01", testHandler: testTomorrowMidnight,
							zone: "Portugal", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
						},
						{
							name: "2024-04-02", testHandler: testTomorrowMidnight,
							zone: "Portugal", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
						},
						{
							name: "2024-10-27", testHandler: testTomorrowMidnight,
							zone: "Portugal", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
						},
						{
							name: "2024-10-28", testHandler: testTomorrowMidnight,
							zone: "Portugal", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
						},
						{
							name: "2024-10-29", testHandler: testTomorrowMidnight,
							zone: "Portugal", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
						},
					}},
			},
		},
		{
			name: "US",
			children: []testCase{
				{
					name: "Alaska",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "US/Alaska", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-10T00:00:00-09:00", expectedDate: "2024-03-10T00:00:00-09:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "US/Alaska", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-11T00:00:00-08:00", expectedDate: "2024-03-11T00:00:00-08:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "US/Alaska", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-12T00:00:00-08:00", expectedDate: "2024-03-12T00:00:00-08:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "US/Alaska", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-03T00:00:00-08:00", expectedDate: "2024-11-03T00:00:00-08:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "US/Alaska", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-04T00:00:00-09:00", expectedDate: "2024-11-04T00:00:00-09:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "US/Alaska", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-05T00:00:00-09:00", expectedDate: "2024-11-05T00:00:00-09:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "US/Alaska", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-10T00:00:00-09:00", expectedDate: "2024-03-10T00:00:00-09:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "US/Alaska", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-11T00:00:00-08:00", expectedDate: "2024-03-11T00:00:00-08:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "US/Alaska", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-12T00:00:00-08:00", expectedDate: "2024-03-12T00:00:00-08:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "US/Alaska", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-03T00:00:00-08:00", expectedDate: "2024-11-03T00:00:00-08:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "US/Alaska", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-04T00:00:00-09:00", expectedDate: "2024-11-04T00:00:00-09:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "US/Alaska", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-05T00:00:00-09:00", expectedDate: "2024-11-05T00:00:00-09:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "US/Alaska", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-10T00:00:00-09:00", expectedDate: "2024-03-10T00:00:00-09:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "US/Alaska", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-11T00:00:00-08:00", expectedDate: "2024-03-11T00:00:00-08:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "US/Alaska", dstHandover: "2024-03-11T00:00:00-08:00",
									date: "2024-03-12T00:00:00-08:00", expectedDate: "2024-03-12T00:00:00-08:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "US/Alaska", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-03T00:00:00-08:00", expectedDate: "2024-11-03T00:00:00-08:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "US/Alaska", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-04T00:00:00-09:00", expectedDate: "2024-11-04T00:00:00-09:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "US/Alaska", dstHandover: "2024-11-04T00:00:00-09:00",
									date: "2024-11-05T00:00:00-09:00", expectedDate: "2024-11-05T00:00:00-09:00",
								},
							}},
					},
				},
				{
					name: "Aleutian",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "US/Aleutian", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-10T00:00:00-10:00", expectedDate: "2024-03-10T00:00:00-10:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "US/Aleutian", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-11T00:00:00-09:00", expectedDate: "2024-03-11T00:00:00-09:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "US/Aleutian", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-12T00:00:00-09:00", expectedDate: "2024-03-12T00:00:00-09:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "US/Aleutian", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-03T00:00:00-09:00", expectedDate: "2024-11-03T00:00:00-09:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "US/Aleutian", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-04T00:00:00-10:00", expectedDate: "2024-11-04T00:00:00-10:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "US/Aleutian", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-05T00:00:00-10:00", expectedDate: "2024-11-05T00:00:00-10:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "US/Aleutian", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-10T00:00:00-10:00", expectedDate: "2024-03-10T00:00:00-10:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "US/Aleutian", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-11T00:00:00-09:00", expectedDate: "2024-03-11T00:00:00-09:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "US/Aleutian", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-12T00:00:00-09:00", expectedDate: "2024-03-12T00:00:00-09:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "US/Aleutian", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-03T00:00:00-09:00", expectedDate: "2024-11-03T00:00:00-09:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "US/Aleutian", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-04T00:00:00-10:00", expectedDate: "2024-11-04T00:00:00-10:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "US/Aleutian", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-05T00:00:00-10:00", expectedDate: "2024-11-05T00:00:00-10:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "US/Aleutian", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-10T00:00:00-10:00", expectedDate: "2024-03-10T00:00:00-10:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "US/Aleutian", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-11T00:00:00-09:00", expectedDate: "2024-03-11T00:00:00-09:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "US/Aleutian", dstHandover: "2024-03-11T00:00:00-09:00",
									date: "2024-03-12T00:00:00-09:00", expectedDate: "2024-03-12T00:00:00-09:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "US/Aleutian", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-03T00:00:00-09:00", expectedDate: "2024-11-03T00:00:00-09:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "US/Aleutian", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-04T00:00:00-10:00", expectedDate: "2024-11-04T00:00:00-10:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "US/Aleutian", dstHandover: "2024-11-04T00:00:00-10:00",
									date: "2024-11-05T00:00:00-10:00", expectedDate: "2024-11-05T00:00:00-10:00",
								},
							}},
					},
				},
				{
					name: "Central",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "US/Central", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "US/Central", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "US/Central", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "US/Central", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "US/Central", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "US/Central", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "US/Central", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "US/Central", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "US/Central", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "US/Central", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "US/Central", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "US/Central", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "US/Central", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "US/Central", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "US/Central", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "US/Central", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "US/Central", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "US/Central", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
					},
				},
				{
					name: "East-Indiana",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "US/East-Indiana", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "US/East-Indiana", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "US/East-Indiana", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "US/East-Indiana", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "US/East-Indiana", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "US/East-Indiana", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "US/East-Indiana", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "US/East-Indiana", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "US/East-Indiana", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "US/East-Indiana", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "US/East-Indiana", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "US/East-Indiana", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "US/East-Indiana", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "US/East-Indiana", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "US/East-Indiana", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "US/East-Indiana", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "US/East-Indiana", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "US/East-Indiana", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
					},
				},
				{
					name: "Eastern",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "US/Eastern", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "US/Eastern", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "US/Eastern", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "US/Eastern", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "US/Eastern", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "US/Eastern", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "US/Eastern", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "US/Eastern", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "US/Eastern", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "US/Eastern", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "US/Eastern", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "US/Eastern", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "US/Eastern", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "US/Eastern", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "US/Eastern", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "US/Eastern", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "US/Eastern", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "US/Eastern", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
					},
				},
				{
					name: "Indiana-Starke",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "US/Indiana-Starke", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "US/Indiana-Starke", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "US/Indiana-Starke", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "US/Indiana-Starke", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "US/Indiana-Starke", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "US/Indiana-Starke", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "US/Indiana-Starke", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "US/Indiana-Starke", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "US/Indiana-Starke", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "US/Indiana-Starke", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "US/Indiana-Starke", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "US/Indiana-Starke", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "US/Indiana-Starke", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-10T00:00:00-06:00", expectedDate: "2024-03-10T00:00:00-06:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "US/Indiana-Starke", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-11T00:00:00-05:00", expectedDate: "2024-03-11T00:00:00-05:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "US/Indiana-Starke", dstHandover: "2024-03-11T00:00:00-05:00",
									date: "2024-03-12T00:00:00-05:00", expectedDate: "2024-03-12T00:00:00-05:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "US/Indiana-Starke", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-03T00:00:00-05:00", expectedDate: "2024-11-03T00:00:00-05:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "US/Indiana-Starke", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-04T00:00:00-06:00", expectedDate: "2024-11-04T00:00:00-06:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "US/Indiana-Starke", dstHandover: "2024-11-04T00:00:00-06:00",
									date: "2024-11-05T00:00:00-06:00", expectedDate: "2024-11-05T00:00:00-06:00",
								},
							}},
					},
				},
				{
					name: "Michigan",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "US/Michigan", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "US/Michigan", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "US/Michigan", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "US/Michigan", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "US/Michigan", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "US/Michigan", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "US/Michigan", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "US/Michigan", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "US/Michigan", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "US/Michigan", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "US/Michigan", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "US/Michigan", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "US/Michigan", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-10T00:00:00-05:00", expectedDate: "2024-03-10T00:00:00-05:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "US/Michigan", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-11T00:00:00-04:00", expectedDate: "2024-03-11T00:00:00-04:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "US/Michigan", dstHandover: "2024-03-11T00:00:00-04:00",
									date: "2024-03-12T00:00:00-04:00", expectedDate: "2024-03-12T00:00:00-04:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "US/Michigan", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-03T00:00:00-04:00", expectedDate: "2024-11-03T00:00:00-04:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "US/Michigan", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-04T00:00:00-05:00", expectedDate: "2024-11-04T00:00:00-05:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "US/Michigan", dstHandover: "2024-11-04T00:00:00-05:00",
									date: "2024-11-05T00:00:00-05:00", expectedDate: "2024-11-05T00:00:00-05:00",
								},
							}},
					},
				},
				{
					name: "Mountain",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "US/Mountain", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "US/Mountain", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "US/Mountain", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "US/Mountain", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "US/Mountain", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "US/Mountain", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "US/Mountain", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "US/Mountain", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "US/Mountain", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "US/Mountain", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "US/Mountain", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "US/Mountain", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "US/Mountain", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-10T00:00:00-07:00", expectedDate: "2024-03-10T00:00:00-07:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "US/Mountain", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-11T00:00:00-06:00", expectedDate: "2024-03-11T00:00:00-06:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "US/Mountain", dstHandover: "2024-03-11T00:00:00-06:00",
									date: "2024-03-12T00:00:00-06:00", expectedDate: "2024-03-12T00:00:00-06:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "US/Mountain", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-03T00:00:00-06:00", expectedDate: "2024-11-03T00:00:00-06:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "US/Mountain", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-04T00:00:00-07:00", expectedDate: "2024-11-04T00:00:00-07:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "US/Mountain", dstHandover: "2024-11-04T00:00:00-07:00",
									date: "2024-11-05T00:00:00-07:00", expectedDate: "2024-11-05T00:00:00-07:00",
								},
							}},
					},
				},
				{
					name: "Pacific",
					children: []testCase{
						{
							name: "LocalMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testLocalMidnight,
									zone: "US/Pacific", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testLocalMidnight,
									zone: "US/Pacific", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testLocalMidnight,
									zone: "US/Pacific", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testLocalMidnight,
									zone: "US/Pacific", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testLocalMidnight,
									zone: "US/Pacific", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testLocalMidnight,
									zone: "US/Pacific", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
						{
							name: "YesterdayMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testYesterdayMidnight,
									zone: "US/Pacific", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testYesterdayMidnight,
									zone: "US/Pacific", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testYesterdayMidnight,
									zone: "US/Pacific", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testYesterdayMidnight,
									zone: "US/Pacific", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testYesterdayMidnight,
									zone: "US/Pacific", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testYesterdayMidnight,
									zone: "US/Pacific", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
						{
							name: "TomorrowMidnight", children: []testCase{
								{
									name: "2024-03-10", testHandler: testTomorrowMidnight,
									zone: "US/Pacific", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-10T00:00:00-08:00", expectedDate: "2024-03-10T00:00:00-08:00",
								},
								{
									name: "2024-03-11", testHandler: testTomorrowMidnight,
									zone: "US/Pacific", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-11T00:00:00-07:00", expectedDate: "2024-03-11T00:00:00-07:00",
								},
								{
									name: "2024-03-12", testHandler: testTomorrowMidnight,
									zone: "US/Pacific", dstHandover: "2024-03-11T00:00:00-07:00",
									date: "2024-03-12T00:00:00-07:00", expectedDate: "2024-03-12T00:00:00-07:00",
								},
								{
									name: "2024-11-03", testHandler: testTomorrowMidnight,
									zone: "US/Pacific", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-03T00:00:00-07:00", expectedDate: "2024-11-03T00:00:00-07:00",
								},
								{
									name: "2024-11-04", testHandler: testTomorrowMidnight,
									zone: "US/Pacific", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-04T00:00:00-08:00", expectedDate: "2024-11-04T00:00:00-08:00",
								},
								{
									name: "2024-11-05", testHandler: testTomorrowMidnight,
									zone: "US/Pacific", dstHandover: "2024-11-04T00:00:00-08:00",
									date: "2024-11-05T00:00:00-08:00", expectedDate: "2024-11-05T00:00:00-08:00",
								},
							}},
					},
				},
			},
		},
		{
			name: "WET",
			children: []testCase{
				{
					name: "LocalMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testLocalMidnight,
							zone: "WET", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
						},
						{
							name: "2024-04-01", testHandler: testLocalMidnight,
							zone: "WET", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
						},
						{
							name: "2024-04-02", testHandler: testLocalMidnight,
							zone: "WET", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
						},
						{
							name: "2024-10-27", testHandler: testLocalMidnight,
							zone: "WET", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
						},
						{
							name: "2024-10-28", testHandler: testLocalMidnight,
							zone: "WET", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
						},
						{
							name: "2024-10-29", testHandler: testLocalMidnight,
							zone: "WET", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
						},
					}},
				{
					name: "YesterdayMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testYesterdayMidnight,
							zone: "WET", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
						},
						{
							name: "2024-04-01", testHandler: testYesterdayMidnight,
							zone: "WET", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
						},
						{
							name: "2024-04-02", testHandler: testYesterdayMidnight,
							zone: "WET", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
						},
						{
							name: "2024-10-27", testHandler: testYesterdayMidnight,
							zone: "WET", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
						},
						{
							name: "2024-10-28", testHandler: testYesterdayMidnight,
							zone: "WET", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
						},
						{
							name: "2024-10-29", testHandler: testYesterdayMidnight,
							zone: "WET", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
						},
					}},
				{
					name: "TomorrowMidnight", children: []testCase{
						{
							name: "2024-03-31", testHandler: testTomorrowMidnight,
							zone: "WET", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-03-31T00:00:00Z", expectedDate: "2024-03-31T00:00:00Z",
						},
						{
							name: "2024-04-01", testHandler: testTomorrowMidnight,
							zone: "WET", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-01T00:00:00+01:00", expectedDate: "2024-04-01T00:00:00+01:00",
						},
						{
							name: "2024-04-02", testHandler: testTomorrowMidnight,
							zone: "WET", dstHandover: "2024-04-01T00:00:00+01:00",
							date: "2024-04-02T00:00:00+01:00", expectedDate: "2024-04-02T00:00:00+01:00",
						},
						{
							name: "2024-10-27", testHandler: testTomorrowMidnight,
							zone: "WET", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-27T00:00:00+01:00", expectedDate: "2024-10-27T00:00:00+01:00",
						},
						{
							name: "2024-10-28", testHandler: testTomorrowMidnight,
							zone: "WET", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-28T00:00:00Z", expectedDate: "2024-10-28T00:00:00Z",
						},
						{
							name: "2024-10-29", testHandler: testTomorrowMidnight,
							zone: "WET", dstHandover: "2024-10-28T00:00:00Z",
							date: "2024-10-29T00:00:00Z", expectedDate: "2024-10-29T00:00:00Z",
						},
					}},
			},
		},
	}
)

// testTimeIsMidnight tests that got is pointing to Midnight.
// This accounts for some Time Zones where when DST occurs and there is no Midnight when the DST transition occurs.
func testTimeIsMidnight(t *testing.T, timeZone string, localTime, got time.Time) {
	// We would expect midnight to occur at 00:00:00
	if !IsMidnight(got) {
		t.Errorf("%s got %s for %q",
			localTime.Format(time.RFC3339),
			got.Format(time.RFC3339),
			timeZone)
	}
}

func testTimeZonesInner(t *testing.T, tests []testCase) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if len(test.children) > 0 {
				testTimeZonesInner(t, test.children)
			} else {
				test.testHandler(t, test)
			}
		})
	}
}

func testTimeConversion(t *testing.T, tc testCase, f func(time.Time) time.Time, test func(*testing.T, testCase, time.Time, time.Time)) {
	loc, err := time.LoadLocation(tc.zone)
	if err != nil {
		t.Fatalf("unable to load %q: %v", tc.zone, err)
	}

	date := ParseTimeIn(tc.date, loc)
	tm := date.Add(6 * time.Hour)
	got := f(tm)
	testTimeIsMidnight(t, tc.zone, tm, got)
	if test != nil {
		test(t, tc, tm, got)
	}
}
func testLocalMidnight(t *testing.T, tc testCase) {
	testTimeConversion(t, tc, LocalMidnight, nil)
}
func testYesterdayMidnight(t *testing.T, tc testCase) {

	testTimeConversion(t, tc, YesterdayMidnight, func(t *testing.T, _ testCase, tm, got time.Time) {
		// Now check the date is yesterday
		midnight := LocalMidnight(tm)
		if !IsMidnight(midnight) {
			t.Errorf("%q is not midnight", midnight.String())
		}

		if !got.Before(midnight) {
			t.Errorf("%q is not yesterday for %q", got.String(), midnight.String())
		}
	})
}
func testTomorrowMidnight(t *testing.T, tc testCase) {

	testTimeConversion(t, tc, TomorrowMidnight, func(t *testing.T, _ testCase, tm, got time.Time) {
		// Now check the date is tomorrow
		midnight := LocalMidnight(tm)
		if !IsMidnight(midnight) {
			t.Errorf("%q is not midnight", midnight.String())
		}

		ty, tmn, td := midnight.Date()
		tomorrow := LocalMidnight(time.Date(ty, tmn, td+1, 4, 0, 0, 0, midnight.Location()))

		if !IsMidnight(tomorrow) {
			t.Errorf("tomorrow not midnight %q", tomorrow.String())
		}

		if got.Before(midnight) || tomorrow.Before(got) {
			t.Errorf("Not tomorrow, got %q expected %q from %q %q", got.String(), midnight.String(), tm.String(), tomorrow.String())
		}
	})
}

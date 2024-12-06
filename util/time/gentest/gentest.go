// main this is a utility which generates the util/time/timezone_test tests based on the
// timezone database on the local machine.
//
// This test tests several functions against specific dates where their results could be affected by
// the transition to or from local Daylight Savings Time.
//
// The functions tested are:
//
// LocalMidnight to get the start of the current day (may not be 00:00:00, e.g. Cairo)
//
// YesterdayMidnight to get the start of the previous day
//
// TomorrowMidnight to get the start of the following day

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
	"unicode"
)

var (
	testCases = []testCase{
		{
			name: "LocalMidnight",
			test: `testTimeConversion(t, tc, LocalMidnight, nil)`,
		},
		{
			name: "YesterdayMidnight",
			test: `
testTimeConversion(t, tc, YesterdayMidnight, func(t *testing.T, _ testCase, tm, got time.Time) {
	// Now check the date is yesterday
	midnight := LocalMidnight(tm)
	if !IsMidnight(midnight) {
		t.Errorf("%q is not midnight", midnight.String())
	}

	if !got.Before(midnight) {
		t.Errorf("%q is not yesterday for %q", got.String(), midnight.String())
	}
})`,
		},
		{
			name: "TomorrowMidnight",
			test: `
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
})`,
		},
	}

	// Limit this to Linux as this could differ on other platforms.
	// I do know that Darwin (MacOS) does restrict some time zones that Linux provides.
	//
	// more info: https://ssoready.com/blog/engineering/truths-programmers-timezones/
	header = `//go:build linux

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
    // TimeZone test cases generated at ` + time.Now().Format(time.RFC3339) + `
    testCases=[]testCase{`

	footer = `    }
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
}`
)

type testCase struct {
	name string
	test string
}

func main() {
	availZones := getAvailableTimeZones()
	sort.SliceStable(availZones, func(i, j int) bool { return availZones[i] < availZones[j] })

	var tests []*dstTest
	for _, zone := range availZones {
		if loc, err := time.LoadLocation(zone); err == nil {
			e := findDST(zone, loc)
			if e != nil {
				tests = append(tests, e)
			} else {
				fmt.Printf("test %q\n", zone)
			}
		}
	}

	r := &dstTest{depth: 0, childKeys: make(map[string]*dstTest)}
	for _, t := range tests {
		keys := strings.Split(t.name, "/")
		r.add(keys, t)
	}

	out := r.append([]string{header})
	out = append(out, footer)
	for _, test := range testCases {
		out = append(out, fmt.Sprintf("func test%s(t *testing.T, tc testCase) {", test.name), test.test, "}")
	}

	f, err := os.Create("/tmp/timezone_test.go")
	if err == nil {
		defer f.Close()
		for _, s := range out {
			if s != "" {
				_, err = f.Write([]byte(s + "\n"))
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
		}
	}
}

type dstTest struct {
	depth     int
	name      string // name with no path
	childKeys map[string]*dstTest
	children  []*dstTest // Child entries
	zone      string     // full zone name at this point, "" for none
	toDst     time.Time  // time of transition to DST
	fromDst   time.Time  // time of transition from DST
}

func (d *dstTest) add(k []string, e *dstTest) {
	if (d.depth + 1) == len(k) {
		d.children = append(d.children, e)
		e.name = k[d.depth]
		e.depth = d.depth
		return
	}

	key := k[d.depth]
	c, exists := d.childKeys[key]
	if !exists {
		c = &dstTest{
			name:      key,
			depth:     d.depth + 1,
			childKeys: make(map[string]*dstTest),
		}
		d.childKeys[key] = c
		d.children = append(d.children, c)
	}
	c.add(k, e)
}

func (d *dstTest) append(a []string) []string {
	// Work out what we will be doing
	var children []string

	if len(d.children) == 0 {
		for _, testCase := range testCases {
			children = d.genTimeConversionTest(children, testCase.name, d.toDst, d.fromDst)
		}
	} else {
		for _, c := range d.children {
			children = c.append(children)
		}
	}

	// If nothing then ignore this instance
	if len(children) == 0 {
		return a
	}

	if d.name != "" {
		// Non root entries get a full entry
		a = append(a,
			"{",
			fmt.Sprintf("name:%q,", d.name),
			"children:[]testCase{",
		)
	}

	// Now include the work
	a = append(a, children...)

	if d.name != "" {
		a = append(a, "},", "},")
	}
	return a
}

func (d *dstTest) genTimeConversionTest(a []string, name string, dates ...time.Time) []string {
	var dts []time.Time
	for _, dt := range dates {
		if !dt.IsZero() {
			dts = append(dts, dt)
		}
	}
	if len(dts) == 0 {
		return a
	}

	return d.group(a, name, func(a []string) []string {
		for _, date := range dts {
			for i := -1; i < 2; i++ {
				dt := date.AddDate(0, 0, i)

				a = d.genTest(a,
					dt.Format(time.DateOnly),
					"test"+name,
					dt,
					func(a []string, t time.Time, d *dstTest) []string {
						return append(a,
							fmt.Sprintf("date:%q, expectedDate:%q,",
								t.Format(time.RFC3339),
								t.Format(time.RFC3339),
							),
						)
					})
				dt = dt.AddDate(0, 0, 1)
			}
		}
		return a
	})
}

func (d *dstTest) group(a []string, n string, f func([]string) []string) []string {
	a = append(a, "{", fmt.Sprintf("name:%q, children:[]testCase{", n))
	a = f(a)
	return append(a, "}},")
}

func (d *dstTest) genTest(a []string, n, testCaseHandler string, t time.Time, f func(a []string, t time.Time, d *dstTest) []string) []string {
	if t.IsZero() {
		return a
	}

	a = append(a,
		"{",
		fmt.Sprintf("name:%q, testHandler:%s,", n, testCaseHandler),
		fmt.Sprintf("zone:%q, dstHandover:%q,", d.zone, t.Format(time.RFC3339)),
	)
	a = f(a, t, d)
	return append(a, "},")
}

func findDST(name string, loc *time.Location) *dstTest {
	r := &dstTest{name: name, zone: loc.String()}
	tm := time.Date(2024, time.January, 1, 0, 0, 0, 0, loc)
	for tm.Year() == 2024 {
		next := tm.AddDate(0, 0, 1)
		if !next.After(tm) {
			fmt.Printf("Failed %q %v %v", name, tm, next)
			return nil
		}

		if next.IsDST() != tm.IsDST() {
			if next.IsDST() {
				r.toDst = next
			} else {
				r.fromDst = next
			}
		}
		tm = next
	}
	return r
}

func getAvailableTimeZones() []string {
	var timeZones []string
	for _, zd := range []string{
		// Update path according to your OS
		"/usr/share/zoneinfo/",
		"/usr/share/lib/zoneinfo/",
		"/usr/lib/locale/TZ/",
	} {
		timeZones = walkTzDir(zd, timeZones)

		for idx, zone := range timeZones {
			timeZones[idx] = strings.ReplaceAll(zone, zd+"/", "")
		}
	}

	// Skip Lord Howe Island as it fails due to daylight savings being 30 minutes & the general population
	// for the island is around 6, so not of much use spending time to fix this one.
	var r []string
	for _, zone := range timeZones {
		if zone != "Australia/Lord_Howe" && zone != "Australia/LHI" {
			r = append(r, zone)
		}
	}

	return r
}

func walkTzDir(path string, zones []string) []string {
	fileInfos, err := os.ReadDir(path)
	if err != nil {
		return zones
	}

	isAlpha := func(s string) bool {
		for _, r := range s {
			if !unicode.IsLetter(r) {
				return false
			}
		}
		return true
	}

	for _, info := range fileInfos {
		if info.Name() != strings.ToUpper(info.Name()[:1])+info.Name()[1:] {
			continue
		}

		if !isAlpha(info.Name()[:1]) {
			continue
		}

		newPath := path + "/" + info.Name()

		if info.IsDir() {
			zones = walkTzDir(newPath, zones)
		} else {
			zones = append(zones, newPath)
		}
	}

	return zones
}

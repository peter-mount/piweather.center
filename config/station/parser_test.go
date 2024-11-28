package station

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewParser(t *testing.T) {
	type test struct {
		name        string
		script      string
		expectError string
	}
	tests := []struct {
		name  string
		tests []test
	}{
		// Station
		{
			name: "station",
			tests: []test{
				// Station name
				{
					// Expect error with invalid station name
					name:   "name",
					script: `station( "home" )`,
				},
				{
					// Expect error with invalid station name
					name:        "name",
					expectError: "station id must not contain",
					script:      `station( "home." )`,
				},
				{
					// Expect error with invalid station name
					name:        "name",
					expectError: "station id must not contain",
					script:      `station( "home station" )`,
				},
				{
					// Expect error with invalid station name
					name:        "name",
					expectError: "station id must not contain",
					script:      `station( "home_station" )`,
				},
				{
					// Expect no errors as names are trimmed
					name:   "name trimmed",
					script: `station( " home " )`,
				},
				// Location
				{
					name:   "location",
					script: `station("home" location("London" "51.5" "0.5"))`,
				},
				{
					name:   "location",
					script: `station("home" location("London" "-51.5" "-0.5"))`,
				},
				{
					name:   "location",
					script: `station("home" location("London" "N51.5" "E0.5"))`,
				},
				{
					name:   "location",
					script: `station("home" location("London" "N51:30" "E0:30"))`,
				},
				{
					name:   "location",
					script: `station("home" location("London" "51.5N" "0.5E"))`,
				},
				{
					name:   "location",
					script: `station("home" location("London" "51:30N" "0:30E"))`,
				},
				{
					name:   "location",
					script: `station("home" location("London" "S51.5" "W0.5"))`,
				},
				{
					name:   "location",
					script: `station("home" location("London" "S51:30" "W0:30"))`,
				},
				{
					name:   "location",
					script: `station("home" location("London" "51.5S" "0.5W"))`,
				},
				{
					name:   "location",
					script: `station("home" location("London" "51:30S" "0:30W"))`,
				},
			},
		},
		// General syntax
		{
			name: "syntax",
			tests: []test{
				{
					script: `
station( "home"
	dashboard( "home" live update "0 */15 * * * *"
		container(
			col(
				class "col-bordered"
				value("Temp_Max_Min_Feels Like_Max_Min_Dew Point"
					"ecowitt.temp"
					"pseudo1.maxtemp"
					"pseudo1.mintemp"
					"pseudo1.feelslike"
					"pseudo1.maxfeelslike"
					"pseudo1.minfeelslike"
					"pseudo1.dewpoint"
				)
			)
			container(
				row(
					gauge( "Temperature" "ecowitt.temp" "pseudo1.maxtemp" "pseudo1.mintemp" )
				)
			)
		)
	)
	dashboard( "stats" 
		multiValue( "*" )
	)
)`,
				},
			},
		},
		// Dashboard
		{
			name: "dashboard",
			tests: []test{
				// Dashboard name
				{
					// Expect error as names are required
					name:        "name",
					expectError: "dashboard name is required",
					script:      `station( "home" dashboard( "" ) dashboard( "stats" ) )`,
				},
				{
					// Expect no errors as names are trimmed
					name:   "name trimmed",
					script: `station( "home" dashboard( " home " ) )`,
				},
				{
					// Expect no errors as names are unique
					name:   "name unique",
					script: `station( "home" dashboard( "home" ) dashboard( "stats" ) )`,
				},
				{
					// Expect error as names are not unique
					name:        "name not unique",
					expectError: "already exists at",
					script:      `station( "home" dashboard( "home" ) dashboard( "stats" ) dashboard( "home" ) )`,
				},
				{
					// Expect error with invalid dashboard name
					name:        "name invalid",
					expectError: "dashboard name must not contain",
					script:      `station( "home" dashboard( "hello.world") )`,
				},
				{
					// Expect error with invalid dashboard name
					name:        "name invalid",
					expectError: "dashboard name must not contain",
					script:      `station( "home" dashboard( "hello world") )`,
				},
				{
					// Expect error with invalid dashboard name
					name:        "name invalid",
					expectError: "dashboard name must not contain",
					script:      `station( "home" dashboard( "hello_world") )`,
				},
			},
		},
		// value
		{
			name: "value",
			tests: []test{
				{
					// Expect error as names are not unique
					script: `
station( "home"
    dashboard("home"
        live
        update "0 */15 * * * *"
        container(
            col( //"col-bordered"
                value("Temp_Max_Min_Feels Like_Max_Min_Dew Point"
                    "home.ecowitt.temp"
                    "home.pseudo1.maxtemp"
                    "home.pseudo1.mintemp"
                    "home.pseudo1.feelslike"
                    "home.pseudo1.maxfeelslike"
                    "home.pseudo1.minfeelslike"
                    "home.pseudo1.dewpoint"
                )
            )
        )
    )
)
`,
				},
			},
		},
		// multiValue
		{
			name: "multiValue",
			tests: []test{
				// These are the same - match all
				{
					name:   "all",
					script: `station("home" dashboard("home" multiValue( "" )))`,
				},
				{
					name:   "all",
					script: `station("home" dashboard("home" multiValue( "*" )))`,
				},
				{
					name:   "all",
					script: `station("home" dashboard("home" multiValue( "**" )))`,
				},
				// Prefix matches
				{
					name:   "prefix",
					script: `station("home" dashboard("home" multiValue( "*text" )))`,
				},
				{
					name:   "prefix",
					script: `station("home" dashboard("home" multiValue( "*sensor.temp" )))`,
				},
				{
					name:        "prefix",
					expectError: "pattern must not include",
					script:      `station("home" dashboard("home" multiValue( "*metric*invalid" )))`,
				},
				{
					name:        "prefix",
					expectError: "pattern must not include",
					script:      `station("home" dashboard("home" multiValue( "*metric*invalid" )))`,
				},
				// Suffix mattches
				{
					name:   "suffix",
					script: `station("home" dashboard("home" multiValue( "text*" )))`,
				},
				{
					name:   "suffix",
					script: `station("home" dashboard("home" multiValue( "sensor.temp*" )))`,
				},
				{
					name:        "suffix",
					expectError: "pattern must not include",
					script:      `station("home" dashboard("home" multiValue( "metric*invalid*" )))`,
				},
				// Contains
				{
					name:   "contains",
					script: `station("home" dashboard("home" multiValue( "*text*" )))`,
				},
				{
					name:   "contains",
					script: `station("home" dashboard("home" multiValue( "*sensor.temp*" )))`,
				},
				{
					name:        "contains",
					expectError: "pattern must not include",
					script:      `station("home" dashboard("home" multiValue( "***" )))`,
				},
				{
					name:        "contains",
					expectError: "pattern must not include",
					script:      `station("home" dashboard("home" multiValue( "*metric*invalid*" )))`,
				},
				// Equality which should always fail
				{
					name:        "equality",
					expectError: "No wildcard provided",
					script:      `station("home" dashboard("home" multiValue( "metric_valid" )))`,
				},
				{
					name:        "equality",
					expectError: "No wildcard provided",
					script:      `station("home" dashboard("home" multiValue( "metric.invalid" )))`,
				},
				{
					name:        "equality",
					expectError: "pattern must not include",
					script:      `station("home" dashboard("home" multiValue( "metric*invalid" )))`,
				},
				{
					name:        "equality",
					expectError: "pattern must not include",
					script:      `station("home" dashboard("home" multiValue( "metric*invalid" )))`,
				},
				{
					name:        "equality",
					expectError: "pattern must not include",
					script:      `station("home" dashboard("home" multiValue( "metric invalid" )))`,
				},
			},
		},
		// gauge
		{
			name: "gauge",
			tests: []test{
				{
					name:   "gauge",
					script: `station("home" dashboard("home" gauge( "label" "ecowitt.temp" )))`,
				},
				{
					name:        "gauge no metrics",
					expectError: "No metrics provided",
					script:      `station("home" dashboard("home" gauge( "label" )))`,
				},
			},
		},
		// metric
		{
			name: "metric",
			tests: []test{
				{
					name:   "plain",
					script: `station("home" dashboard("dash" value( "" "sensor.test") ))`,
				},
				{
					name:   "plain",
					script: `station("home" dashboard("dash" value( "" "sensor.test.sub") ))`,
				},
				{
					name:   "plain",
					script: `station("home" dashboard("dash" value( "" "sensor.test_sub") ))`,
				},
				{
					name:   "valid unit",
					script: `station("home" dashboard("dash" value( "label" "sensor.test" unit "celsius") ))`,
				},
				{
					name:        "invalid unit",
					expectError: "unsupported unit",
					script:      `station("home" dashboard("dash" value( "label" "sensor.test" unit "invalid") ))`,
				},
			},
		},
	}

	p := NewParser()

	for i, tt := range tests {
		testName := tt.name
		if testName == "" {
			testName = fmt.Sprintf("test_%d", i)
		}

		t.Run(testName, func(t *testing.T) {
			for j, tte := range tt.tests {
				subTestName := tte.name
				if subTestName == "" {
					subTestName = fmt.Sprintf("test_%d", j)
				}

				t.Run(subTestName, func(t *testing.T) {
					_, err := p.ParseString(testName+"/"+subTestName, tte.script)

					if tte.expectError == "" {
						if err != nil {
							t.Fatal("parse returned", err)
						}
					} else {
						if err == nil {
							t.Fatal("expected error")
						}
						if !strings.Contains(err.Error(), tte.expectError) {
							t.Fatalf("Expected %q got %v", tte.expectError, err)
						}
					}
				})

			}
		})

	}
}

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
		multivalue( "*" )
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
		// multivalue
		{
			name: "multivalue",
			tests: []test{
				// These are the same - match all
				{
					name:   "all",
					script: `station("home" dashboard("home" multivalue( "" )))`,
				},
				{
					name:   "all",
					script: `station("home" dashboard("home" multivalue( "*" )))`,
				},
				{
					name:   "all",
					script: `station("home" dashboard("home" multivalue( "**" )))`,
				},
				// Prefix matches
				{
					name:   "prefix",
					script: `station("home" dashboard("home" multivalue( "*text" )))`,
				},
				{
					name:   "prefix",
					script: `station("home" dashboard("home" multivalue( "*sensor.temp" )))`,
				},
				{
					name:        "prefix",
					expectError: "pattern must not include",
					script:      `station("home" dashboard("home" multivalue( "*metric*invalid" )))`,
				},
				{
					name:        "prefix",
					expectError: "pattern must not include",
					script:      `station("home" dashboard("home" multivalue( "*metric*invalid" )))`,
				},
				// Suffix mattches
				{
					name:   "suffix",
					script: `station("home" dashboard("home" multivalue( "text*" )))`,
				},
				{
					name:   "suffix",
					script: `station("home" dashboard("home" multivalue( "sensor.temp*" )))`,
				},
				{
					name:        "suffix",
					expectError: "pattern must not include",
					script:      `station("home" dashboard("home" multivalue( "metric*invalid*" )))`,
				},
				// Contains
				{
					name:   "contains",
					script: `station("home" dashboard("home" multivalue( "*text*" )))`,
				},
				{
					name:   "contains",
					script: `station("home" dashboard("home" multivalue( "*sensor.temp*" )))`,
				},
				{
					name:        "contains",
					expectError: "pattern must not include",
					script:      `station("home" dashboard("home" multivalue( "***" )))`,
				},
				{
					name:        "contains",
					expectError: "pattern must not include",
					script:      `station("home" dashboard("home" multivalue( "*metric*invalid*" )))`,
				},
				// Equality which should always fail
				{
					name:        "equality",
					expectError: "No wildcard provided",
					script:      `station("home" dashboard("home" multivalue( "metric.invalid" )))`,
				},
				{
					name:        "equality",
					expectError: "pattern must not include",
					script:      `station("home" dashboard("home" multivalue( "metric*invalid" )))`,
				},
				{
					name:        "equality",
					expectError: "pattern must not include",
					script:      `station("home" dashboard("home" multivalue( "metric*invalid" )))`,
				},
				{
					name:        "equality",
					expectError: "pattern must not include",
					script:      `station("home" dashboard("home" multivalue( "metric invalid" )))`,
				},
				{
					name:        "equality",
					expectError: "pattern must not include",
					script:      `station("home" dashboard("home" multivalue( "metric_invalid" )))`,
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

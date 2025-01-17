package view

import (
	"github.com/peter-mount/go-kernel/v2/rest"
	"log"
	"net/http"
	"strconv"
	"time"
)

func getDateParameters(r *rest.Rest) (stationId, dash string, year, month, day, status int) {
	status = http.StatusOK
	now := time.Now()
	ny, nm, nd := now.Date()
	stationId = r.Var("stationId")
	dash = r.Var("dash")
	ys, ms, ds := r.Var("year"), r.Var("month"), r.Var("day")

	var err error
	if ys != "" {
		year, err = strconv.Atoi(ys)
	}

	if err == nil && ms != "" {
		month, err = strconv.Atoi(ms)
	}

	if err == nil && ds != "" {
		day, err = strconv.Atoi(ds)
	}

	switch {
	case err != nil:
		status = http.StatusInternalServerError

	case
		// Ensure if we have month we have year or if we have day then we have month & year
		ms != "" && ys == "",
		ds != "" && ms == "",
		// Dates in the future are invalid
		year > ny,
		year == ny && month > int(nm),
		year == ny && month == int(nm) && day > nd,
		// TODO Dates before the station has data
		year < 2003,
		// Validate month range 1...12
		ms != "" && month < 1 || month > 12,
		ds != "" && day < 1:
		status = http.StatusNotFound

	case ms != "" && ds != "":
		// This checks that day is within the month. It works because time.Date() will normalise the date,
		// so if we have April 31 being requested, this will return May 1st and as that differs to the request we fail it.
		dt := time.Date(year, time.Month(month), day, 0, 0, 0, 0, now.Location())
		dy, dm, dd := dt.Date()
		if dy != year || int(dm) != month || dd != day {
			status = http.StatusNotFound
		}
	}

	return
}

func (s *Service) yearIndex(r *rest.Rest) error {
	stationId, dash, year, _, _, status := getDateParameters(r)
	if status != http.StatusOK {
		r.Status(status)
		return nil
	}

	log.Println(stationId, dash, year)
	return nil
}

func (s *Service) monthIndex(r *rest.Rest) error {
	stationId, dash, year, month, _, status := getDateParameters(r)
	if status != http.StatusOK {
		r.Status(status)
		return nil
	}

	log.Println(stationId, dash, year, month)
	return nil
}

func (s *Service) dayDashboard(r *rest.Rest) error {
	stationId, dash, year, month, day, status := getDateParameters(r)
	if status != http.StatusOK {
		r.Status(status)
		return nil
	}

	log.Println(stationId, dash, year, month, day)
	return nil
}

// Package time handles common calculations to complement go's internal time package.
//
// It provides functions to return the start Time of the day specified by go's Time,
// dealing with time zones, and the varying length of days when a time zone switches
// between Daylight Saving and Standard Time.
//
// For example, LocalMidnight() returns the time.Time representing the local Midnight for the
// day for a specific time.Time.
//
// Notes:
//
// Contrary to popular belief, a day is not 24 hours long.
// Most are, but when a time zone switches to daylight savings then that day is only 23 hours long.
// When switching back to standard time then that day is 25 hours long.
//
// Also, midnight is not always 00:00:00.
// For the following 10 time zones when moving to daylight savings, the day starts at 01:00:00
//
// "Africa/Cairo",
// "America/Asuncion",
// "America/Havana",
// "America/Santiago",
// "America/Scoresbysund",
// "Asia/Beirut",
// "Atlantic/Azores",
// "Chile/Continental",
// "Cuba",
// "Egypt"
//
// Testing:
//
// The tests which checks that the various functions work currently only work on Linux or Unix based
// systems as it locates and reads the time zone database from the operating system to get a list of
// all available time zones and runs the tests against each one.
//
// Known issues:
//
// As of 2024 May 14,
// the "Australia/LHI" and "Australia/Lord_How" time zones currently fail for dates where the
// DST/ST transition occurs, and needs more research but seems to be down to the transition only being
// 30 minutes rather than 1 hour.

package time

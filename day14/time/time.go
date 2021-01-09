package main

import (
    "fmt"
    "time"
)

func main() {
    p := fmt.Println

	now := time.Now()
	//We’ll start by getting the current time.


    p(now)

    then := time.Date(
		//You can build a time struct by providing the year, month, day, etc. Times are always associated with a Location, i.e. time zone.
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

	p(then.Year())
	//You can extract the various components of the time value as expected.
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())

	p(then.Weekday())
	//The Monday-Sunday Weekday is also available.

	//These methods compare two times, testing if the first occurs before, after, or at the same time as the second, respectively.
    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

	//The Sub methods returns a Duration representing the interval between two times.
    diff := now.Sub(then)
    p(diff)

    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())

	//We can compute the length of the duration in various units.
    p(then.Add(diff))
    p(then.Add(-diff))
}
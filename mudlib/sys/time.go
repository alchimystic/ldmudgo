package mudlib_sys

type TimeFieldIdx int

const (
	TM_SEC   TimeFieldIdx = 0 /* Seconds (0..59) */
	TM_MIN   TimeFieldIdx = 1 /* Minutes (0..59) */
	TM_HOUR  TimeFieldIdx = 2 /* Hours (0..23) */
	TM_MDAY  TimeFieldIdx = 3 /* Day of the month (1..31) */
	TM_MON   TimeFieldIdx = 4 /* Month of the year (0..11) */
	TM_YEAR  TimeFieldIdx = 5 /* Year (e.g.  2001) */
	TM_WDAY  TimeFieldIdx = 6 /* Day of the week (Sunday = 0) */
	TM_YDAY  TimeFieldIdx = 7 /* Day of the year (0..365) */
	TM_ISDST TimeFieldIdx = 8 /* TRUE: Daylight saving time */
	TM_MAX   int          = int(TM_ISDST) + 1
)

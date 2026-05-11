package mudlib_sys

type TrimFlag int

const (
	TRIM_LEFT  TrimFlag = 0x01
	TRIM_RIGHT TrimFlag = 0x02
	TRIM_BOTH  TrimFlag = (TRIM_LEFT | TRIM_RIGHT)
)

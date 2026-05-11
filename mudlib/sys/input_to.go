package mudlib_sys

type InputToFlag int

const (
	INPUT_NOECHO      InputToFlag = 1
	INPUT_CHARMODE    InputToFlag = 2
	INPUT_PROMPT      InputToFlag = 4
	INPUT_NO_TELNET   InputToFlag = 8
	INPUT_APPEND      InputToFlag = 16
	INPUT_IGNORE_BANG InputToFlag = 128
)

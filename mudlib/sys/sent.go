package mudlib_sys

type SentValue int

const (
	SENT_PLAIN      SentValue = 0
	SENT_SHORT_VERB SentValue = 1
	SENT_NO_SPACE   SentValue = 2
	SENT_NO_VERB    SentValue = 3
	SENT_MARKER     SentValue = 4
)

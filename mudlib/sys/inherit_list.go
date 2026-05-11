package mudlib_sys

type InheritListFlag int

const (
	INHLIST_FLAT        InheritListFlag = 0x00
	INHLIST_TREE        InheritListFlag = 0x01
	INHLIST_TAG_VIRTUAL InheritListFlag = 0x02
)

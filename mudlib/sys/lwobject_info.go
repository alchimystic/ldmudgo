package mudlib_sys

type LWObjectInfoField int

const (
	LI_LWOBJECT_REFS           LWObjectInfoField = -1
	LI_DATA_SIZE               LWObjectInfoField = -2
	LI_DATA_SIZE_TOTAL         LWObjectInfoField = -3
	LI_NO_INHERIT              LWObjectInfoField = -10
	LI_NO_CLONE                LWObjectInfoField = -11
	LI_NO_LIGHTWEIGHT          LWObjectInfoField = -12
	LI_SHARE_VARIABLES         LWObjectInfoField = -13
	LI_PROG_REFS               LWObjectInfoField = -20
	LI_NUM_FUNCTIONS           LWObjectInfoField = -30
	LI_NUM_VARIABLES           LWObjectInfoField = -31
	LI_NUM_STRINGS             LWObjectInfoField = -32
	LI_NUM_INHERITED           LWObjectInfoField = -33
	LI_NUM_INCLUDED            LWObjectInfoField = -34
	LI_SIZE_FUNCTIONS          LWObjectInfoField = -35
	LI_SIZE_VARIABLES          LWObjectInfoField = -36
	LI_SIZE_STRINGS            LWObjectInfoField = -37
	LI_SIZE_STRINGS_DATA       LWObjectInfoField = -38
	LI_SIZE_STRINGS_DATA_TOTAL LWObjectInfoField = -39
	LI_SIZE_INHERITED          LWObjectInfoField = -40
	LI_SIZE_INCLUDED           LWObjectInfoField = -41
	LI_PROG_SIZE               LWObjectInfoField = -42
	LI_PROG_SIZE_TOTAL         LWObjectInfoField = -43
)

package mudlib_sys

type ObjectInfoField int

const (
	OI_ONCE_INTERACTIVE        ObjectInfoField = -1
	OI_RESET_STATE             ObjectInfoField = -2
	OI_WILL_CLEAN_UP           ObjectInfoField = -3
	OI_LAMBDA_REFERENCED       ObjectInfoField = -4
	OI_REPLACED                ObjectInfoField = -5
	OI_NO_INHERIT              ObjectInfoField = -10
	OI_NO_CLONE                ObjectInfoField = -11
	OI_NO_SHADOW               ObjectInfoField = -12
	OI_SHARE_VARIABLES         ObjectInfoField = -13
	OI_NO_LIGHTWEIGHT          ObjectInfoField = -14
	OI_SWAPPED                 ObjectInfoField = -20
	OI_PROG_SWAPPED            ObjectInfoField = -21
	OI_VAR_SWAPPED             ObjectInfoField = -22
	OI_SWAP_NUM                ObjectInfoField = -23
	OI_NEXT_RESET_TIME         ObjectInfoField = -30
	OI_NEXT_CLEANUP_TIME       ObjectInfoField = -31
	OI_LAST_REF_TIME           ObjectInfoField = -32
	OI_OBJECT_NEXT             ObjectInfoField = -40
	OI_OBJECT_PREV             ObjectInfoField = -41
	OI_OBJECT_POS              ObjectInfoField = -42
	OI_SHADOW_NEXT             ObjectInfoField = -50
	OI_SHADOW_PREV             ObjectInfoField = -51
	OI_SHADOW_ALL              ObjectInfoField = -52
	OI_OBJECT_REFS             ObjectInfoField = -60
	OI_TICKS                   ObjectInfoField = -61
	OI_GIGATICKS               ObjectInfoField = -62
	OI_DATA_SIZE               ObjectInfoField = -63
	OI_DATA_SIZE_TOTAL         ObjectInfoField = -64
	OI_PROG_REFS               ObjectInfoField = -70
	OI_NUM_FUNCTIONS           ObjectInfoField = -71
	OI_NUM_VARIABLES           ObjectInfoField = -72
	OI_NUM_STRINGS             ObjectInfoField = -73
	OI_NUM_INHERITED           ObjectInfoField = -74
	OI_NUM_INCLUDED            ObjectInfoField = -75
	OI_SIZE_FUNCTIONS          ObjectInfoField = -76
	OI_SIZE_VARIABLES          ObjectInfoField = -77
	OI_SIZE_STRINGS            ObjectInfoField = -78
	OI_SIZE_STRINGS_DATA       ObjectInfoField = -79
	OI_SIZE_STRINGS_DATA_TOTAL ObjectInfoField = -80
	OI_SIZE_INHERITED          ObjectInfoField = -81
	OI_SIZE_INCLUDED           ObjectInfoField = -82
	OI_PROG_SIZE               ObjectInfoField = -83
	OI_PROG_SIZE_TOTAL         ObjectInfoField = -84
)

package mudlib_sys

type WizListIdx int

const (
	WL_NAME           WizListIdx = 0  /* Wizard name */
	WL_COMMANDS       WizListIdx = 1  /* Number of commands executed */
	WL_COST           WizListIdx = 2  /* Weighted evalcost spent on this wizard */
	WL_GIGACOST       WizListIdx = 3  /* Weighted giga-evalcost spent on this wizard */
	WL_TOTAL_COST     WizListIdx = 4  /* Total evalcost spent on this wizard */
	WL_TOTAL_GIGACOST WizListIdx = 5  /* Total giga-evalcost spent on this wizard */
	WL_HEART_BEATS    WizListIdx = 6  /* Heartbeats spent on this wizard */
	WL_CALL_OUT       WizListIdx = 7  /* unimplemented */
	WL_ARRAY_TOTAL    WizListIdx = 8  /* Arrays accounted for */
	WL_MAPPING_TOTAL  WizListIdx = 9  /* Mappings accounted for */
	WL_STRUCT_TOTAL   WizListIdx = 10 /* Struct elements accounted for */
	WL_EXTRA          WizListIdx = 11 /* Extra Wizinfo, if set */
	WL_SIZE           int        = int(WL_EXTRA) + 1
)

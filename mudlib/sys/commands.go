package mudlib_sys

type AddActionFlag int

const (
	AA_VERB     AddActionFlag = 0
	AA_SHORT    AddActionFlag = 1
	AA_NOSPACE  AddActionFlag = 2
	AA_IMM_ARGS AddActionFlag = 3
)

type QueryActionBitFlag int

const (
	QA_VERB       QueryActionBitFlag = 1
	QA_TYPE       QueryActionBitFlag = 2
	QA_SHORT_VERB QueryActionBitFlag = 4
	QA_OBJECT     QueryActionBitFlag = 8
	QA_FUNCTION   QueryActionBitFlag = 16
)

type CommandActionIdx int

const (
	CMD_VERB    CommandActionIdx = 0
	CMD_TEXT    CommandActionIdx = 1
	CMD_ORIGIN  CommandActionIdx = 2
	CMD_PLAYER  CommandActionIdx = 4
	CMD_FAIL    CommandActionIdx = 5
	CMD_FAILOBJ CommandActionIdx = 6
	CMD_SIZE    CommandActionIdx = 7
)

type MatchCommandIdx int

const (
	CMDM_VERB   MatchCommandIdx = 0
	CMDM_ARG    MatchCommandIdx = 1
	CMDM_OBJECT MatchCommandIdx = 2
	CMDM_FUN    MatchCommandIdx = 3
	CMDM_SIZE   MatchCommandIdx = 4
)

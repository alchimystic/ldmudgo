package mudlib_sys

type InteractiveInfoField int

const (
	II_IP_NAME       InteractiveInfoField = -1
	II_IP_NUMBER     InteractiveInfoField = -2
	II_IP_PORT       InteractiveInfoField = -3
	II_IP_ADDRESS    InteractiveInfoField = -4
	II_MUD_PORT      InteractiveInfoField = -5
	II_MCCP_STATS    InteractiveInfoField = -10
	II_INPUT_PENDING InteractiveInfoField = -20
	II_EDITING       InteractiveInfoField = -21
	II_IDLE          InteractiveInfoField = -22
	II_NOECHO        InteractiveInfoField = -23
	II_CHARMODE      InteractiveInfoField = -24
	II_SNOOP_NEXT    InteractiveInfoField = -30
	II_SNOOP_PREV    InteractiveInfoField = -31
	II_SNOOP_ALL     InteractiveInfoField = -32
)

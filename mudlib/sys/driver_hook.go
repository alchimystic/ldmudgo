package mudlib_sys

type LpcDriverHook int

const (
	H_MOVE_OBJECT0            LpcDriverHook = 0
	H_MOVE_OBJECT1            LpcDriverHook = 1
	H_LOAD_UIDS               LpcDriverHook = 2
	H_CLONE_UIDS              LpcDriverHook = 3
	H_CREATE_SUPER            LpcDriverHook = 4
	H_CREATE_OB               LpcDriverHook = 5
	H_CREATE_CLONE            LpcDriverHook = 6
	H_RESET                   LpcDriverHook = 7
	H_CLEAN_UP                LpcDriverHook = 8
	H_MODIFY_COMMAND          LpcDriverHook = 9
	H_NOTIFY_FAIL             LpcDriverHook = 10
	H_NO_IPC_SLOT             LpcDriverHook = 11
	H_INCLUDE_DIRS            LpcDriverHook = 12
	H_TELNET_NEG              LpcDriverHook = 13
	H_NOECHO                  LpcDriverHook = 14
	H_ERQ_STOP                LpcDriverHook = 15
	H_MODIFY_COMMAND_FNAME    LpcDriverHook = 16
	H_COMMAND                 LpcDriverHook = 17
	H_SEND_NOTIFY_FAIL        LpcDriverHook = 18
	H_AUTO_INCLUDE            LpcDriverHook = 19
	H_DEFAULT_METHOD          LpcDriverHook = 20
	H_DEFAULT_PROMPT          LpcDriverHook = 21
	H_PRINT_PROMPT            LpcDriverHook = 22
	H_REGEXP_PACKAGE          LpcDriverHook = 23
	H_MSG_DISCARDED           LpcDriverHook = 24
	H_FILE_ENCODING           LpcDriverHook = 25
	H_LWOBJECT_UIDS           LpcDriverHook = 26
	H_CREATE_LWOBJECT         LpcDriverHook = 27
	H_CREATE_LWOBJECT_COPY    LpcDriverHook = 28
	H_CREATE_LWOBJECT_RESTORE LpcDriverHook = 29
	H_AUTO_INCLUDE_EXPRESSION LpcDriverHook = 30
	H_AUTO_INCLUDE_BLOCK      LpcDriverHook = 31
	NUM_DRIVER_HOOKS          int           = int(H_AUTO_INCLUDE_BLOCK) + 1
)

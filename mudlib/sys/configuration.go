package mudlib_sys

type ConfigureInteractiveOpt int

const (
	IC_MAX_WRITE_BUFFER_SIZE        ConfigureInteractiveOpt = 0
	IC_SOCKET_BUFFER_SIZE           ConfigureInteractiveOpt = 1
	IC_COMBINE_CHARSET_AS_STRING    ConfigureInteractiveOpt = 2
	IC_COMBINE_CHARSET_AS_ARRAY     ConfigureInteractiveOpt = 3
	IC_CONNECTION_CHARSET_AS_STRING ConfigureInteractiveOpt = 4
	IC_CONNECTION_CHARSET_AS_ARRAY  ConfigureInteractiveOpt = 5
	IC_QUOTE_IAC                    ConfigureInteractiveOpt = 6
	IC_TELNET_ENABLED               ConfigureInteractiveOpt = 7
	IC_MCCP                         ConfigureInteractiveOpt = 8
	IC_PROMPT                       ConfigureInteractiveOpt = 9
	IC_MAX_COMMANDS                 ConfigureInteractiveOpt = 10
	IC_MODIFY_COMMAND               ConfigureInteractiveOpt = 11
	IC_ENCODING                     ConfigureInteractiveOpt = 12
)

type ConfigureObjectOpt int

const (
	OC_COMMANDS_ENABLED ConfigureObjectOpt = 0
	OC_HEART_BEAT       ConfigureObjectOpt = 1
	OC_EUID             ConfigureObjectOpt = 2
)

type ConfigureObjectLwOpt int

const (
	LC_EUID ConfigureObjectLwOpt = 0
)

type ConfigureDriverOpt int

const (
	DC_MEMORY_LIMIT           ConfigureDriverOpt = 0
	DC_ENABLE_HEART_BEATS     ConfigureDriverOpt = 1
	DC_LONG_EXEC_TIME         ConfigureDriverOpt = 2
	DC_DATA_CLEAN_TIME        ConfigureDriverOpt = 3
	DC_TLS_CERTIFICATE        ConfigureDriverOpt = 4
	DC_TLS_DHE_PARAMETER      ConfigureDriverOpt = 5
	DC_TLS_CIPHERLIST         ConfigureDriverOpt = 6
	DC_EXTRA_WIZINFO_SIZE     ConfigureDriverOpt = 7
	DC_DEFAULT_RUNTIME_LIMITS ConfigureDriverOpt = 8
	DC_SWAP_COMPACT_MODE      ConfigureDriverOpt = 9
	DC_SWAP_TIME              ConfigureDriverOpt = 10
	DC_SWAP_VAR_TIME          ConfigureDriverOpt = 11
	DC_CLEANUP_TIME           ConfigureDriverOpt = 12
	DC_RESET_TIME             ConfigureDriverOpt = 13
	DC_DEBUG_FILE             ConfigureDriverOpt = 14
	DC_FILESYSTEM_ENCODING    ConfigureDriverOpt = 15
	DC_SIGACTION_SIGHUP       ConfigureDriverOpt = 20
	DC_SIGACTION_SIGINT       ConfigureDriverOpt = 21
	DC_SIGACTION_SIGUSR1      ConfigureDriverOpt = 22
	DC_SIGACTION_SIGUSR2      ConfigureDriverOpt = 23
)

type DcSigActionOpt int

const (
	DCS_DEFAULT         DcSigActionOpt = 0
	DCS_IGNORE          DcSigActionOpt = 1
	DCS_TERMINATE       DcSigActionOpt = 2
	DCS_SHUTDOWN        DcSigActionOpt = 3
	DCS_INFORM_MASTER   DcSigActionOpt = 4
	DCS_RELOAD_MASTER   DcSigActionOpt = 5
	DCS_THROW_EXCEPTION DcSigActionOpt = 6
)

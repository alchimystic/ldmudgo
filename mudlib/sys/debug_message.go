package mudlib_sys

type DebugMessageBitMask int

const (
	DMSG_DEFAULT DebugMessageBitMask = 0
	DMSG_STDOUT  DebugMessageBitMask = (1 << 0)
	DMSG_STDERR  DebugMessageBitMask = (1 << 1)
	DMSG_LOGFILE DebugMessageBitMask = (1 << 2)
	DMSG_STAMP   DebugMessageBitMask = (1 << 3)
	DMSG_TARGET  DebugMessageBitMask = (DMSG_STDOUT | DMSG_STDERR | DMSG_LOGFILE)
)

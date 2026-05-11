package mudlib_sys

type LPCSignal int

const (
	LPC_SIGHUP  LPCSignal = 1
	LPC_SIGINT  LPCSignal = 2
	LPC_SIGTERM LPCSignal = 15
	LPC_SIGUSR1 LPCSignal = 16
	LPC_SIGUSR2 LPCSignal = 17
)

type Signal int

const (
	SIGHUP  Signal = Signal(LPC_SIGHUP)
	SIGINT  Signal = Signal(LPC_SIGINT)
	SIGTERM Signal = Signal(LPC_SIGTERM)
	SIGUSR1 Signal = Signal(LPC_SIGUSR1)
	SIGUSR2 Signal = Signal(LPC_SIGUSR2)
)

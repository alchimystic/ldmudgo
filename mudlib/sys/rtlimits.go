package mudlib_sys

type RuntimeLimitIdx int

const (
	LIMIT_EVAL         RuntimeLimitIdx = 0
	LIMIT_ARRAY        RuntimeLimitIdx = 1
	LIMIT_MAPPING      RuntimeLimitIdx = 2
	LIMIT_MAPPING_KEYS RuntimeLimitIdx = LIMIT_MAPPING
	LIMIT_MAPPING_SIZE RuntimeLimitIdx = 3
	LIMIT_BYTE         RuntimeLimitIdx = 4
	LIMIT_FILE         RuntimeLimitIdx = 5
	LIMIT_CALLOUTS     RuntimeLimitIdx = 6
	LIMIT_COST         RuntimeLimitIdx = 7
	LIMIT_MEMORY       RuntimeLimitIdx = 8
	LIMIT_MAX          RuntimeLimitIdx = 9
)

type RuntimeLimit int

const (
	LIMIT_UNLIMITED RuntimeLimit = 0  /* No limit */
	LIMIT_KEEP      RuntimeLimit = -1 /* Keep the old limit setting */
	LIMIT_DEFAULT   RuntimeLimit = -2
)

type MemoryLimit int

const (
	MALLOC_SOFT_LIMIT MemoryLimit = 1
	MALLOC_HARD_LIMIT MemoryLimit = 2
)

type LowMemoryCondition int

const (
	NO_MALLOC_LIMIT_EXCEEDED   LowMemoryCondition = 0
	SOFT_MALLOC_LIMIT_EXCEEDED LowMemoryCondition = LowMemoryCondition(MALLOC_SOFT_LIMIT)
	HARD_MALLOC_LIMIT_EXCEEDED LowMemoryCondition = LowMemoryCondition(MALLOC_HARD_LIMIT)
)

type AvailableReserveFlag int

const (
	USER_RESERVE_AVAILABLE   AvailableReserveFlag = 0x1
	MASTER_RESERVE_AVAILABLE AvailableReserveFlag = 0x2
	SYSTEM_RESERVE_AVAILABLE AvailableReserveFlag = 0x4
)

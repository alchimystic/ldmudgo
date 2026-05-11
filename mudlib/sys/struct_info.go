package mudlib_sys

type StructInfoType int

const (
	SINFO_FLAT   StructInfoType = 0
	SINFO_NESTED StructInfoType = 1
)

type StructInfoIdx int

const (
	SI_NAME      StructInfoIdx = 0
	SI_PROG_NAME StructInfoIdx = 1
	SI_PROG_ID   StructInfoIdx = 2
	SI_BASE      StructInfoIdx = 3
	SI_MEMBER    StructInfoIdx = 4
	SI_MAX       int           = int(SI_MEMBER) + 1
)

type StructInfoMember int

const (
	SIM_NAME  StructInfoMember = 0
	SIM_TYPE  StructInfoMember = 1
	SIM_EXTRA StructInfoMember = 2
	SIM_MAX   int              = int(SIM_EXTRA) + 1
)

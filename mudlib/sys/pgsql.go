package mudlib_sys

type PGExtFun int

const (
	PGRES_EMPTY_QUERY    PGExtFun = 0 /* Unimplemented */
	PGRES_COMMAND_OK     PGExtFun = 1
	PGRES_TUPLES_OK      PGExtFun = 2
	PGRES_COPY_OUT       PGExtFun = 3 /* Unimplemented */
	PGRES_COPY_IN        PGExtFun = 4 /* Unimplemented */
	PGRES_BAD_RESPONSE   PGExtFun = 5
	PGRES_NONFATAL_ERROR PGExtFun = 6
	PGRES_FATAL_ERROR    PGExtFun = 7
	PGRES_NOTICE         PGExtFun = 99
)

type PGResultCode int

const (
	PGCONN_SUCCESS PGResultCode = 100
	PGCONN_FAILED  PGResultCode = 101
	PGCONN_ABORTED PGResultCode = 102
)

type PGResultType int

const (
	PG_RESULT_ARRAY PGResultType = 0
	PG_RESULT_MAP   PGResultType = 1
)

package mudlib_sys

type TraceMask int

const (
	TRACE_NOTHING    TraceMask = 0   /* Stop tracing */
	TRACE_CALL       TraceMask = 1   /* Trace all lfun calls */
	TRACE_CALL_OTHER TraceMask = 2   /* Trace inter-object calls */
	TRACE_RETURN     TraceMask = 4   /* Trace function returns */
	TRACE_ARGS       TraceMask = 8   /* Print function arguments and results */
	TRACE_EXEC       TraceMask = 16  /* Trace all executed instructions */
	TRACE_HEART_BEAT TraceMask = 32  /* Trace heartbeat code */
	TRACE_APPLY      TraceMask = 64  /* Trace (internal) applies */
	TRACE_OBJNAME    TraceMask = 128 /* Print the object names */
)

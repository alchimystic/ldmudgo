package mudlib_sys

type StringPrepProfile int

const (
	STRINGPREP_NAMEPREP          StringPrepProfile = 1
	STRINGPREP_SASLPREP          StringPrepProfile = 2
	STRINGPREP_PLAIN             StringPrepProfile = 3
	STRINGPREP_TRACE             StringPrepProfile = 4
	STRINGPREP_KERBEROS5         StringPrepProfile = 5
	STRINGPREP_XMPP_NODEPREP     StringPrepProfile = 6
	STRINGPREP_XMPP_RESOURCEPREP StringPrepProfile = 7
	STRINGPREP_ISCSI             StringPrepProfile = 8
)

type StringPrepFlag int

const (
	STRINGPREP_NO_NFKC_FLAG       StringPrepFlag = (1 << 0)
	STRINGPREP_NO_BIDI_FLAG       StringPrepFlag = (1 << 1)
	STRINGPREP_NO_UNASSIGNED_FLAG StringPrepFlag = (1 << 2)
	STRINGPREP_FLAG_MAX           int            = int(STRINGPREP_NO_UNASSIGNED_FLAG)
)

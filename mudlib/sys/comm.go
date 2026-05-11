package mudlib_sys

type NetConnectCode int

const (
	NC_SUCCESS       NetConnectCode = 0
	NC_EUNKNOWNHOST  NetConnectCode = 1
	NC_ENOSOCKET     NetConnectCode = 2
	NC_ENOBIND       NetConnectCode = 3
	NC_ENOCONNECT    NetConnectCode = 4
	NC_ECONNREFUSED  NetConnectCode = 5
	NC_EMCONN        NetConnectCode = 6
	NC_ENORESSOURCES NetConnectCode = 7
)

package mudlib_sys

type ServiceRequestType int

const (
	ERQ_RLOOKUP   ServiceRequestType = 0
	ERQ_EXECUTE   ServiceRequestType = 1
	ERQ_FORK      ServiceRequestType = 2
	ERQ_AUTH      ServiceRequestType = 3
	ERQ_SPAWN     ServiceRequestType = 4
	ERQ_SEND      ServiceRequestType = 5
	ERQ_KILL      ServiceRequestType = 6
	ERQ_OPEN_UDP  ServiceRequestType = 7
	ERQ_OPEN_TCP  ServiceRequestType = 8
	ERQ_LISTEN    ServiceRequestType = 9
	ERQ_ACCEPT    ServiceRequestType = 10
	ERQ_LOOKUP    ServiceRequestType = 11
	ERQ_RLOOKUPV6 ServiceRequestType = 12
)

const ERQ_CB_STRING = (1 << 31)

type RequestExecuteType int

const (
	ERQ_OK           RequestExecuteType = 0
	ERQ_SIGNALED     RequestExecuteType = 1
	ERQ_E_NOTFOUND   RequestExecuteType = 2
	ERQ_E_UNKNOWN    RequestExecuteType = 3
	ERQ_E_ARGLENGTH  RequestExecuteType = 4
	ERQ_E_ARGFORMAT  RequestExecuteType = 5
	ERQ_E_ARGNUMBER  RequestExecuteType = 6
	ERQ_E_ILLEGAL    RequestExecuteType = 7
	ERQ_E_PATHLEN    RequestExecuteType = 8
	ERQ_E_FORKFAIL   RequestExecuteType = 9
	ERQ_E_TICKET     RequestExecuteType = 11
	ERQ_E_INCOMPLETE RequestExecuteType = 12
	ERQ_E_WOULDBLOCK RequestExecuteType = 13
	ERQ_E_PIPE       RequestExecuteType = 14
	ERQ_STDOUT       RequestExecuteType = 15
	ERQ_STDERR       RequestExecuteType = 16
	ERQ_EXITED       RequestExecuteType = 17
	ERQ_E_NSLOTS     RequestExecuteType = 18
)

type RequestHandle int

const (
	ERQ_HANDLE_RLOOKUP     RequestHandle = (-1)
	ERQ_HANDLE_KEEP_HANDLE RequestHandle = (-2)
	ERQ_HANDLE_RLOOKUPV6   RequestHandle = (-3)
)

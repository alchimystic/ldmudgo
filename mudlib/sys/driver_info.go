package mudlib_sys

type DriverInfoField int

const (
	DI_BOOT_TIME                                   DriverInfoField = -1
	DI_CURRENT_RUNTIME_LIMITS                      DriverInfoField = -10
	DI_EVAL_NUMBER                                 DriverInfoField = -11
	DI_MUD_PORTS                                   DriverInfoField = -20
	DI_UDP_PORT                                    DriverInfoField = -21
	DI_MEMORY_RESERVE_USER                         DriverInfoField = -30
	DI_MEMORY_RESERVE_MASTER                       DriverInfoField = -31
	DI_MEMORY_RESERVE_SYSTEM                       DriverInfoField = -32
	DI_TRACE_CURRENT                               DriverInfoField = -40
	DI_TRACE_CURRENT_DEPTH                         DriverInfoField = -41
	DI_TRACE_CURRENT_AS_STRING                     DriverInfoField = -42
	DI_TRACE_LAST_ERROR                            DriverInfoField = -43
	DI_TRACE_LAST_ERROR_AS_STRING                  DriverInfoField = -44
	DI_TRACE_LAST_UNCAUGHT_ERROR                   DriverInfoField = -45
	DI_TRACE_LAST_UNCAUGHT_ERROR_AS_STRING         DriverInfoField = -46
	DI_NUM_FUNCTION_NAME_CALLS                     DriverInfoField = -100
	DI_NUM_FUNCTION_NAME_CALL_HITS                 DriverInfoField = -101
	DI_NUM_FUNCTION_NAME_CALL_MISSES               DriverInfoField = -102
	DI_NUM_OBJECTS_LAST_PROCESSED                  DriverInfoField = -103
	DI_NUM_HEARTBEAT_TOTAL_CYCLES                  DriverInfoField = -104
	DI_NUM_HEARTBEAT_ACTIVE_CYCLES                 DriverInfoField = -105
	DI_NUM_HEARTBEATS_LAST_PROCESSED               DriverInfoField = -106
	DI_NUM_STRING_TABLE_STRINGS_ADDED              DriverInfoField = -110
	DI_NUM_STRING_TABLE_STRINGS_REMOVED            DriverInfoField = -111
	DI_NUM_STRING_TABLE_LOOKUPS_BY_VALUE           DriverInfoField = -112
	DI_NUM_STRING_TABLE_LOOKUPS_BY_INDEX           DriverInfoField = -113
	DI_NUM_STRING_TABLE_LOOKUP_STEPS_BY_VALUE      DriverInfoField = -114
	DI_NUM_STRING_TABLE_LOOKUP_STEPS_BY_INDEX      DriverInfoField = -115
	DI_NUM_STRING_TABLE_HITS_BY_VALUE              DriverInfoField = -116
	DI_NUM_STRING_TABLE_HITS_BY_INDEX              DriverInfoField = -117
	DI_NUM_STRING_TABLE_COLLISIONS                 DriverInfoField = -118
	DI_NUM_REGEX_LOOKUPS                           DriverInfoField = -120
	DI_NUM_REGEX_LOOKUP_HITS                       DriverInfoField = -121
	DI_NUM_REGEX_LOOKUP_MISSES                     DriverInfoField = -122
	DI_NUM_REGEX_LOOKUP_COLLISIONS                 DriverInfoField = -123
	DI_NUM_MESSAGES_OUT                            DriverInfoField = -200
	DI_NUM_PACKETS_OUT                             DriverInfoField = -201
	DI_NUM_PACKETS_IN                              DriverInfoField = -202
	DI_SIZE_PACKETS_OUT                            DriverInfoField = -203
	DI_SIZE_PACKETS_IN                             DriverInfoField = -204
	DI_LOAD_AVERAGE_COMMANDS                       DriverInfoField = -300
	DI_LOAD_AVERAGE_LINES                          DriverInfoField = -301
	DI_LOAD_AVERAGE_PROCESSED_OBJECTS              DriverInfoField = -302
	DI_LOAD_AVERAGE_PROCESSED_OBJECTS_RELATIVE     DriverInfoField = -303
	DI_LOAD_AVERAGE_PROCESSED_HEARTBEATS_RELATIVE  DriverInfoField = -304
	DI_NUM_ACTIONS                                 DriverInfoField = -400
	DI_NUM_CALLOUTS                                DriverInfoField = -401
	DI_NUM_HEARTBEATS                              DriverInfoField = -402
	DI_NUM_SHADOWS                                 DriverInfoField = -403
	DI_NUM_OBJECTS                                 DriverInfoField = -404
	DI_NUM_OBJECTS_SWAPPED                         DriverInfoField = -405
	DI_NUM_OBJECTS_IN_LIST                         DriverInfoField = -406
	DI_NUM_OBJECTS_IN_TABLE                        DriverInfoField = -407
	DI_NUM_OBJECTS_DESTRUCTED                      DriverInfoField = -408
	DI_NUM_OBJECTS_NEWLY_DESTRUCTED                DriverInfoField = -409
	DI_NUM_OBJECT_TABLE_SLOTS                      DriverInfoField = -410
	DI_NUM_PROGS                                   DriverInfoField = -411
	DI_NUM_PROGS_SWAPPED                           DriverInfoField = -412
	DI_NUM_PROGS_UNSWAPPED                         DriverInfoField = -413
	DI_NUM_ARRAYS                                  DriverInfoField = -414
	DI_NUM_MAPPINGS                                DriverInfoField = -415
	DI_NUM_MAPPINGS_CLEAN                          DriverInfoField = -416
	DI_NUM_MAPPINGS_HASH                           DriverInfoField = -417
	DI_NUM_MAPPINGS_HYBRID                         DriverInfoField = -418
	DI_NUM_STRUCTS                                 DriverInfoField = -419
	DI_NUM_STRUCT_TYPES                            DriverInfoField = -420
	DI_NUM_VIRTUAL_STRINGS                         DriverInfoField = -421
	DI_NUM_STRINGS                                 DriverInfoField = -422
	DI_NUM_STRINGS_TABLED                          DriverInfoField = -423
	DI_NUM_STRINGS_UNTABLED                        DriverInfoField = -424
	DI_NUM_STRING_TABLE_SLOTS                      DriverInfoField = -425
	DI_NUM_STRING_TABLE_SLOTS_USED                 DriverInfoField = -426
	DI_NUM_REGEX                                   DriverInfoField = -427
	DI_NUM_REGEX_TABLE_SLOTS                       DriverInfoField = -428
	DI_NUM_LVALUES                                 DriverInfoField = -429
	DI_NUM_NAMED_OBJECT_TYPES                      DriverInfoField = -430
	DI_NUM_NAMED_OBJECT_TYPES_TABLE_SLOTS          DriverInfoField = -431
	DI_NUM_LWOBJECTS                               DriverInfoField = -432
	DI_NUM_COROUTINES                              DriverInfoField = -433
	DI_NUM_LPC_PYTHON_REFS                         DriverInfoField = -434
	DI_NUM_PYTHON_LPC_REFS                         DriverInfoField = -435
	DI_SIZE_ACTIONS                                DriverInfoField = -450
	DI_SIZE_CALLOUTS                               DriverInfoField = -451
	DI_SIZE_HEARTBEATS                             DriverInfoField = -452
	DI_SIZE_SHADOWS                                DriverInfoField = -453
	DI_SIZE_OBJECTS                                DriverInfoField = -454
	DI_SIZE_OBJECTS_SWAPPED                        DriverInfoField = -455
	DI_SIZE_OBJECT_TABLE                           DriverInfoField = -456
	DI_SIZE_PROGS                                  DriverInfoField = -457
	DI_SIZE_PROGS_SWAPPED                          DriverInfoField = -458
	DI_SIZE_PROGS_UNSWAPPED                        DriverInfoField = -459
	DI_SIZE_ARRAYS                                 DriverInfoField = -460
	DI_SIZE_MAPPINGS                               DriverInfoField = -461
	DI_SIZE_STRUCTS                                DriverInfoField = -462
	DI_SIZE_STRUCT_TYPES                           DriverInfoField = -463
	DI_SIZE_STRINGS                                DriverInfoField = -464
	DI_SIZE_STRINGS_TABLED                         DriverInfoField = -465
	DI_SIZE_STRINGS_UNTABLED                       DriverInfoField = -466
	DI_SIZE_STRING_TABLE                           DriverInfoField = -467
	DI_SIZE_STRING_OVERHEAD                        DriverInfoField = -468
	DI_SIZE_REGEX                                  DriverInfoField = -469
	DI_SIZE_BUFFER_FILE                            DriverInfoField = -470
	DI_SIZE_BUFFER_SWAP                            DriverInfoField = -471
	DI_SIZE_NAMED_OBJECT_TYPES_TABLE               DriverInfoField = -472
	DI_SIZE_LWOBJECTS                              DriverInfoField = -473
	DI_SIZE_COROUTINES                             DriverInfoField = -474
	DI_NUM_SWAP_BLOCKS                             DriverInfoField = -500
	DI_NUM_SWAP_BLOCKS_FREE                        DriverInfoField = -501
	DI_NUM_SWAP_BLOCKS_REUSE_LOOKUPS               DriverInfoField = -502
	DI_NUM_SWAP_BLOCKS_REUSE_LOOKUP_STEPS          DriverInfoField = -503
	DI_NUM_SWAP_BLOCKS_FREE_LOOKUPS                DriverInfoField = -505
	DI_NUM_SWAP_BLOCKS_FREE_LOOKUP_STEPS           DriverInfoField = -506
	DI_SIZE_SWAP_BLOCKS                            DriverInfoField = -507
	DI_SIZE_SWAP_BLOCKS_FREE                       DriverInfoField = -508
	DI_SIZE_SWAP_BLOCKS_REUSED                     DriverInfoField = -509
	DI_SWAP_RECYCLE_PHASE                          DriverInfoField = -510
	DI_MEMORY_ALLOCATOR_NAME                       DriverInfoField = -600
	DI_NUM_SYS_ALLOCATED_BLOCKS                    DriverInfoField = -610
	DI_NUM_LARGE_BLOCKS_ALLOCATED                  DriverInfoField = -611
	DI_NUM_LARGE_BLOCKS_FREE                       DriverInfoField = -612
	DI_NUM_LARGE_BLOCKS_WASTE                      DriverInfoField = -613
	DI_NUM_SMALL_BLOCKS_ALLOCATED                  DriverInfoField = -614
	DI_NUM_SMALL_BLOCKS_FREE                       DriverInfoField = -615
	DI_NUM_SMALL_BLOCKS_WASTE                      DriverInfoField = -616
	DI_NUM_SMALL_BLOCK_CHUNKS                      DriverInfoField = -617
	DI_NUM_UNMANAGED_BLOCKS                        DriverInfoField = -618
	DI_NUM_FREE_BLOCKS_AVL_NODES                   DriverInfoField = -619
	DI_SIZE_SYS_ALLOCATED_BLOCKS                   DriverInfoField = -630
	DI_SIZE_LARGE_BLOCKS_ALLOCATED                 DriverInfoField = -631
	DI_SIZE_LARGE_BLOCKS_FREE                      DriverInfoField = -632
	DI_SIZE_LARGE_BLOCKS_WASTE                     DriverInfoField = -633
	DI_SIZE_LARGE_BLOCK_OVERHEAD                   DriverInfoField = -634
	DI_SIZE_SMALL_BLOCKS_ALLOCATED                 DriverInfoField = -635
	DI_SIZE_SMALL_BLOCKS_FREE                      DriverInfoField = -636
	DI_SIZE_SMALL_BLOCKS_WASTE                     DriverInfoField = -637
	DI_SIZE_SMALL_BLOCK_OVERHEAD                   DriverInfoField = -638
	DI_SIZE_SMALL_BLOCK_CHUNKS                     DriverInfoField = -639
	DI_SIZE_UNMANAGED_BLOCKS                       DriverInfoField = -640
	DI_SIZE_MEMORY_USED                            DriverInfoField = -641
	DI_SIZE_MEMORY_UNUSED                          DriverInfoField = -642
	DI_SIZE_MEMORY_OVERHEAD                        DriverInfoField = -643
	DI_NUM_INCREMENT_SIZE_CALLS                    DriverInfoField = -650
	DI_NUM_INCREMENT_SIZE_CALL_SUCCESSES           DriverInfoField = -651
	DI_SIZE_INCREMENT_SIZE_CALL_DIFFS              DriverInfoField = -652
	DI_NUM_REPLACEMENT_MALLOC_CALLS                DriverInfoField = -653
	DI_SIZE_REPLACEMENT_MALLOC_CALLS               DriverInfoField = -654
	DI_NUM_MEMORY_DEFRAGMENTATION_CALLS_FULL       DriverInfoField = -655
	DI_NUM_MEMORY_DEFRAGMENTATION_CALLS_TARGETED   DriverInfoField = -656
	DI_NUM_MEMORY_DEFRAGMENTATION_CALL_TARGET_HITS DriverInfoField = -657
	DI_NUM_MEMORY_DEFRAGMENTATION_BLOCKS_INSPECTED DriverInfoField = -658
	DI_NUM_MEMORY_DEFRAGMENTATION_BLOCKS_MERGED    DriverInfoField = -659
	DI_NUM_MEMORY_DEFRAGMENTATION_BLOCKS_RESULTING DriverInfoField = -660
	DI_MEMORY_EXTENDED_STATISTICS                  DriverInfoField = -670
	DI_STATUS_TEXT_MEMORY                          DriverInfoField = -700
	DI_STATUS_TEXT_TABLES                          DriverInfoField = -701
	DI_STATUS_TEXT_SWAP                            DriverInfoField = -702
	DI_STATUS_TEXT_MALLOC                          DriverInfoField = -703
	DI_STATUS_TEXT_MALLOC_EXTENDED                 DriverInfoField = -704
	DI_NUM_SIMUL_EFUNS_TABLED                      DriverInfoField = -900
)

type DriverInfoMemExtStatIdx int

const (
	DIM_ES_MAX_ALLOC   DriverInfoMemExtStatIdx = 0
	DIM_ES_CUR_ALLOC   DriverInfoMemExtStatIdx = 1
	DIM_ES_MAX_FREE    DriverInfoMemExtStatIdx = 2
	DIM_ES_CUR_FREE    DriverInfoMemExtStatIdx = 3
	DIM_ES_AVG_XALLOC  DriverInfoMemExtStatIdx = 4
	DIM_ES_AVG_XFREE   DriverInfoMemExtStatIdx = 5
	DIM_ES_FULL_SLABS  DriverInfoMemExtStatIdx = 6
	DIM_ES_FREE_SLABS  DriverInfoMemExtStatIdx = 7
	DIM_ES_TOTAL_SLABS DriverInfoMemExtStatIdx = 8
	DIM_ES_MAX         DriverInfoMemExtStatIdx = 9
)

type DumpDriverInfoValue int

const (
	DDI_OBJECTS            DumpDriverInfoValue = 0
	DDI_OBJECTS_DESTRUCTED DumpDriverInfoValue = 1
	DDI_OPCODES            DumpDriverInfoValue = 2
	DDI_MEMORY             DumpDriverInfoValue = 3
)

type DriverInfoTraceIdx int

const (
	TRACE_TYPE     DriverInfoTraceIdx = 0
	TRACE_NAME     DriverInfoTraceIdx = 1
	TRACE_PROGRAM  DriverInfoTraceIdx = 2
	TRACE_OBJECT   DriverInfoTraceIdx = 3
	TRACE_LOC      DriverInfoTraceIdx = 4
	TRACE_EVALCOST DriverInfoTraceIdx = 5
	TRACE_MAX      DriverInfoTraceIdx = 6
)

type DriverInfoTraceType int

const (
	TRACE_TYPE_SYMBOL DriverInfoTraceType = 0
	TRACE_TYPE_SEFUN  DriverInfoTraceType = 1
	TRACE_TYPE_EFUN   DriverInfoTraceType = 2
	TRACE_TYPE_LAMBDA DriverInfoTraceType = 3
	TRACE_TYPE_LFUN   DriverInfoTraceType = 4
)

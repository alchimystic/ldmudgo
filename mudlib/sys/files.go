package mudlib_sys

type FileSizeCode int

const (
	FSIZE_NOFILE FileSizeCode = -1
	FSIZE_DIR    FileSizeCode = -2
)

type GetDirFlag int

const (
	GETDIR_EMPTY    GetDirFlag = (0)
	GETDIR_NAMES    GetDirFlag = (0x01)
	GETDIR_SIZES    GetDirFlag = (0x02)
	GETDIR_DATES    GetDirFlag = (0x04)
	GETDIR_ACCESS   GetDirFlag = (0x40)
	GETDIR_MODES    GetDirFlag = (0x80)
	GETDIR_PATH     GetDirFlag = (0x10)
	GETDIR_UNSORTED GetDirFlag = (0x20)
	GETDIR_ALL      GetDirFlag = (0xDF)
)

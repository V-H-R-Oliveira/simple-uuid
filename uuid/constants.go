package uuid

const (
	UUID_TIMESTAMP = 122192928000000000
	V1             = ^(1 << 15) & ^(1 << 14) & ^(1 << 13)
	V4             = ^(1 << 15) & ^(1 << 13) & ^(1 << 12)
	SET_3MSB       = ^(1 << 15) & ^(1 << 14) & ^(1 << 13)
	SET_4MSB       = SET_3MSB & ^(1 << 12)
	DCE            = ^(1 << 14)
	MICROSOFT      = ^(1 << 13)
)

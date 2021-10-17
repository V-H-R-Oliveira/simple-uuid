package uuid

const (
	V4 = ^(1 << 15) & ^(1 << 13) & ^(1 << 12)
	SET_3MSB = ^(1 << 15) & ^(1 << 14) & ^(1 << 13)
	SET_4MSB = SET_3MSB & ^(1 << 12)
)

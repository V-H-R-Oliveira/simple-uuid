package uuid

const (
	UUID_TIMESTAMP    = 122192928000000000
	V1                = ^(1 << 15) & ^(1 << 14) & ^(1 << 13)
	V4                = ^(1 << 15) & ^(1 << 13) & ^(1 << 12)
	V3                = ^(1 << 15) & ^(1 << 14)
	V5                = ^(1 << 15) & ^(1 << 13)
	SET_3MSB          = ^(1 << 15) & ^(1 << 14) & ^(1 << 13)
	SET_4MSB          = ^(1 << 15) & ^(1 << 14) & ^(1 << 13) & ^(1 << 12)
	DCE_VARIANT       = ^(1 << 14)
	MICROSOFT_VARIANT = ^(1 << 13)
)

// test constants

const (
	dce            = "dce"
	microsoft      = "microsoft"
	defaultVariant = ""
)

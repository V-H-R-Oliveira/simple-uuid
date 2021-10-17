package uuid

type UUID struct {
	TimeLow            uint32
	TimeMid            uint16
	TimeHighAndVersion uint16
	ClockAndVariant    uint16
	Node               []byte
}

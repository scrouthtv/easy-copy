package device

type Device interface {
	Usage() SpaceUsage
	Name() string
	Equal(other Device) bool
}

type SpaceUsage struct {
	Total uint64
	Free uint64
}

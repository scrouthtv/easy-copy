package device

// Device implements a file holder.
type Device interface {

	// Usage returns the current space usage of this device.
	Usage() SpaceUsage

	// Name returns a user-friendly name for the device.
	Name() string

	// Equal indicates whether two devices are equal.
	// It is used for considering renaming a file instead of
	// copying + deleting it.
	Equal(other Device) bool
}

// SpaceUsage holds the space usage of a device.
// Unit is 1 byte.
type SpaceUsage struct {
	Total uint64
	Free  uint64
}

// UnknownUsage is returned if the space usage of the requested
// device could not be determined.
var UnknownUsage = SpaceUsage{Total: 0, Free: 0}

package device

import "testing"

func TestPartitionReader(t *testing.T) {
	reloadDevices()

	for id, part := range currentPartTable {
		t.Logf("Device #%d (%s) is mounted at %s", id, part.dev, part.mnt)
	}
}

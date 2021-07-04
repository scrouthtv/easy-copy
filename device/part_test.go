package device

import "testing"

func TestPartitionFinder(t *testing.T) {
	fs := []string{
		"/home/lenni/cut.mp4",
		"/home/lenni/out.mp4",
		"/tmp/service-list.new",
		"/usr/share/",
		"/this is my folder/adsf",
	}

	for i, f := range fs {
		dev := GetDevice(f)
		t.Log(fs[i], "on", dev.Name(), "free:", dev.Usage().Free)
	}
}

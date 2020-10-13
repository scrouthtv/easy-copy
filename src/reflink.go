//+build !linux

import "errors";

var notSupportedError error = errors.New("operation not supported");

func reflink(srcPath string, dstPath string, progressStorage *uint64) error {
	return notSupportedError;
}

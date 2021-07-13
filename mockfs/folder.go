package mockfs

type MockFolder struct {
	subfolders []*MockFolder
	files      []*MockFile
}

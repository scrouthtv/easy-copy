package stub

import "easy-copy/flags"

type stubImpl struct {}

func New() flags.Settings {
	return &stubImpl{}
}

func (s *stubImpl) ParseLine() {}

func (s *stubImpl) LoadConfig() error {
	return nil
}

func (s *stubImpl) Sources() []string {
	return nil
}

func (s *stubImpl) Target() string {
	return ""
}

func (s *stubImpl) Mode() flags.Mode {
	return flags.ModeCopy
}

func (s *stubImpl) Verbosity() flags.Verbose {
	return flags.VerbDebug
}

func (s *stubImpl) OnConflict() flags.Conflict {
	return flags.ConflictAsk
}

func (s *stubImpl) OnSymlink() flags.Symlink {
	return flags.SymlinkIgnore
}

func (s *stubImpl) DoLSColors() bool {
	return true
}

func (s *stubImpl) Dryrun() bool {
	return true
}

func (s *stubImpl) Parallel() bool {
	return true
}

func (s *stubImpl) SetOnConflict(c flags.Conflict) {}

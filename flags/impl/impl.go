package impl

import "easy-copy/flags"

type settingsImpl struct {
	sources []string
	target  string

	mode       flags.Mode
	verbosity  flags.Verbose
	onConflict flags.Conflict
	onSymlink  flags.Symlink
	doLScolors bool
	dryrun     bool
}

func New() flags.Settings {
	return &settingsImpl{
		sources:    []string{},
		target:     "",
		mode:       flags.ModeCopy,
		verbosity:  flags.VerbNotice,
		onConflict: flags.ConflictAsk,
		onSymlink:  flags.SymlinkDeref,
		doLScolors: true,
		dryrun:     false,
	}
}

func (s *settingsImpl) Sources() []string {
	return s.sources
}

func (s *settingsImpl) Target() string {
	return s.target
}

func (s *settingsImpl) Mode() flags.Mode {
	return s.mode
}

func (s *settingsImpl) Verbosity() flags.Verbose {
	return s.verbosity
}

func (s *settingsImpl) OnConflict() flags.Conflict {
	return s.onConflict
}

func (s *settingsImpl) OnSymlink() flags.Symlink {
	return s.onSymlink
}

func (s *settingsImpl) DoLSColors() bool {
	return s.doLScolors
}

func (s *settingsImpl) Dryrun() bool {
	return s.dryrun
}

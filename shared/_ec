#compdef ec

backtick="'"

_arguments -s -S \
	'(-f --force -i --interactive -n --no-clobber)'{-f,--force}'[Overwrite existing files without asking]' \
	'(-f --force -i --interactive -n --no-clobber)'{-i,--interactive}'[Ask before overwriting a file]' \
	'(-f --force -i --interactive -n --no-clobber)'{-n,--no-clobber}'[Skip existing files]' \
	'(--no-config)'--no-config"[Don't read any config file]" \
	'(-V --verbose)'{-V,--verbose}'[Verbose mode]' \
	'(-q --quiet)'{-q,--quiet}'[Quiet mode]' \
	'(--color)'--color'=-[Whether to colorize the output]::when:(always never auto)' \
	'(-e --extended-colors)'{-e,--extended-colors}'=-[Use more colors: colorize file names using $LS_COLORS]' \
	'(-h --help)'{-h,--help}'[Print help and exit]' \
	'(-v --version)'{-v,--version}'[Print version information and exit]' \
	'(--copying)'--copying'[Print redistribution information and exit]' \
	'(--warranty)'--warranty'[Print warranty information and exit]' \
	'*:file:_files'

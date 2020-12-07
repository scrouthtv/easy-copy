complete -c ec -s f -l force -d "Overwrite existing files without asking"
complete -c ec -s i -l interactive -d "Ask before overwriting a file"
complete -c ec -s n -l no-clobber -d "Skip existing files"
complete -c ec -l no-config -d "Don't read any config file"
complete -c ec -s V -l verbose -d "Verbose mode"
complete -c ec -s q -l quiet -d "Quiet mode"
complete -c ec -l color -d "Whether to colorize the output" -f -a "always never auto"
complete -c ec -s h -l help -d "Print help and exit"
complete -c ec -s v -l version -d "Print version information and exit"
complete -c ec -l copying -d "Print redistribution information and exit"
complete -c ec -l warranty -d "Print warranty information and exit"

#!/usr/bin/env bash

# $COMP_WORDS is an array of all words that are entered
# $COMP_CWORD is the position of the cursor (0 being ec)

_ec() {
	# after the first letter following the = is inserted, the new word
	# is stripped of the beginning equal sign. wont fix for now
	#if [[ ${COMP_WORDS[$COMP_CWORD]} =~ ^\=.* ]]; then
	#ase "${COMP_WORDS[$COMP_CWORD - 1]}" in
	#--color)
	#	COMPREPLY=($(compgen -W "=always =auto =never" \
	#		\"${COMP_WORDS[$COMP_CWORD]}\"))
	#	echo \"${COMP_WORDS[$COMP_CWORD]}\"
	#	echo $COMPREPLY
	#	;;
	#sac
	if [[ ${COMP_WORDS[$COMP_CWORD]} =~ ^-.* ]]; then
		opts=
		opts+="-f --force "
		opts+="-i --interactive "
		opts+="-n --no-clobber "
		opts+="--no-config "
		opts+="-V --verbose "
		opts+="--color "
		COMPREPLY=($(compgen -W "$opts" \"${COMP_WORDS[$COMP_CWORD]}\"))
	else
		COMPREPLY=($(compgen -f \"${COMP_WORDS[$COMP_CWORD]}\"))
	fi
}

complete -F _ec ec

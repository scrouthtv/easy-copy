# easy-copy
Aims to provide a user-friendly &amp; fast alternative to cp, mv and rf.

![gif of easy-copy moving a folder](./.github/docs/demo01.gif)

## Warning
This project is currently under development.
Use it at your own risk. I take zero responsibility for any damage linked to this program in any way.

## Features
 - **No more lockups**: If *EasyCopy* encounters conflicts, it just moves on with the next file. The fancy GUI is also in a seperate thread so that it does not slow the copy process down.
 - **Know what's happening**: A terminal UI shows all current information at a glance, such as progress, speed and remaining time.
 - **No dependencies**: *EasyCopy* uses zero third-party dependencies. Only builtin packages and packages from `golang.org/x/`.
 - **Uses sane defaults**: Recurses into directories without having to add a flag. Checks by default if enough space is available.
 - **Configurable**: Many common options (e.g. color or verbosity) can be set in a config file. 
 - **Blazingly fast**: While adding extra features, *EasyCopy* stays as fast as proven tools by forking second-priority tasks and using Go.
 - **Modular compilation**: If a very thin executable is wanted, support for different modules can be stripped at compilation.

## Contributing
 - Get help with *EasyCopy* by opening an issue on the project's page.
 - Feel free to fork *EasyCopy* and open PRs with new features.

# Copyright
Copyright &#169; 2021 The *EasyCopy* authors.
This software is licensed under GNU GPL v3.0.
This means that you are free to change and redistribute *EasyCopy* as a whole or any part of it, as long as the source to any derived work is as well publicly disclosed and licensed under GNU GPL v3.0.
Absolutely no warranty is provided for this software.
For more information read the LICENSE file that is distributed with *EasyCopy*.

This software uses third-party source code.
 - Salvatore Sanfilippo's `kilo.c`, accessible on the internet via [https://github.com/snaptoken/kilo-src/blob/sections/kilo.c](GitHub). It is licensed under BSD 2-Clause "Simplified", a copy of which comes included with EasyCopy, named `LICENSE.2`.

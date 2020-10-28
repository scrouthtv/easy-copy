# easy-copy
Aims to provide a user-friendly &amp; fast alternative to cp, mv and rf.

<!-- Put a catching GIF here -->

## Warning
This project is currently under development.
Use it at your own risk. I take zero responsibility for any damage linked to this program in any way.

## Features
 - **No more lockups**: In difference to normal tools, *EasyCopy* forks tasks such as asking the user to overwrite files or searching for source files.
 - **Know what's happening**: A terminal UI shows all current information at one glance, such as progress, speed and remaining time.
 - **No dependencies**: Zero third-party packages are used at anytime. This way you can trust *EasyCopy* even when running as `sudo`.
 - **Uses sane defaults**: Recurses into directories without having to add a flag. Checks by default if enough space is available.
 - **Configurable**: *EasyCopy* supports all features `cp` or `mv` has. *EasyCopy* even supports a config file for setting default options.
 - **Blazingly fast**: While adding extra features, *EasyCopy* stays as fast as proven tools by forking second-priority tasks and using Go.
 - **No need to overcomplicate things**: *EasyCopy* does the very same thing as `cp` or `mv` on Unix. 
 - **Modular compilation**: If a very thin executable is wanted, support for different modules can be stripped at compilation.
 - **Support**: *EasyCopy* is developed for and regularly tested on Linux, Windows and macOS.

### Keep that slow drive running
<!-- Threading -->

### Simple and elegant UI
<!-- some gif or so -->

### Zero third-party dependencies
<!-- how many sol, which go packages are used -->

###  Defaults, differences to coreutils `cp`
There are only some memorable differences:
<!-- ... -->

### Configuration
<!-- show a basic configuration file, time how long it takes to read this -->

### No compromises when copying files
<!-- time the different copy methods -->

### Kept simple
<!-- https://www.gnu.org/prep/standards/html_node/Command_002dLine-Interfaces.html -->
*EasyCopy* complies to the GNU Coding Standards about Command-Line interfaces. That means that using the cli tools is very intuitive. For example, these are all the same:
```
 ~ ec -fV foo/ bar/
 ~ ec -Vf foo/ bar/
 ~ ec --verbose foo/ bar/ -f
 ~ ec -f --verbose -- foo/ bar/
```

### Modularity
<!-- show different configurations and time them -->

### Software you can trust in

## Differences to other tools

## Contributing
 - Get help with *EasyCopy* by opening an issue on the project's page.
 - Feel free to fork *EasyCopy* and open PRs with new features.

# Copyright
Copyright &#169; 2020 The *EasyCopy* authors.
This software is licensed under GNU GPL v3.0.
This means that you are free to change and redistribute *EasyCopy* as a whole or any part of it, as long as the source to any derived work is as well publicly disclosed and licensed under GNU GPL v3.0.
Absolutely no warranty is provided for this software.
For more information read the LICENSE file that is distributed with *EasyCopy*.

This software uses third-party source code.
 - Salvatore Sanfilippo's `kilo.c`, accessible on the internet via [https://github.com/snaptoken/kilo-src/blob/sections/kilo.c](GitHub). It is licensed under BSD 2-Clause "Simplified", a copy of which comes included with EasyCopy, named `LICENSE.2`.
 - Karpelès Lab Inc.'s `reflink` project, accessible on the internet via [https://github.com/KarpelesLab/reflink/](GitHub). It is licensed under the MIT License, a copy of which comes included with EasyCopy, named `LICENSE.3`.

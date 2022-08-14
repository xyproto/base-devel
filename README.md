# base-devel

A small utility for listing directories in `PATH` that is named `base-devel` to help test package manager behavior.

* It lists all directories in `PATH` that are named `.bin` or `bin`.
* If the `-n` flag is passed, directories in `PATH` that are not named `.bin` or `bin` are listed.

## Installation

For Go 1.18 or later:

    go install github.com/xyproto/base-devel@latest

## General information

* License: GPL2
* Version: 0.0.1

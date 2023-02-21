# Incompatibilities with the Previous Version

Here we list the incompatibilities that you may find when moving a
program from Lua 5.3 to Lua 5.4.

You can avoid some incompatibilities by compiling Lua with appropriate
options (see file `luaconf.h`). However, all these compatibility options
will be removed in the future. More often than not, compatibility issues
arise when these compatibility options are removed. So, whenever you
have the chance, you should try to test your code with a version of Lua
compiled with all compatibility options turned off. That will ease
transitions to newer versions of Lua.

Lua versions can always change the C API in ways that do not imply
source-code changes in a program, such as the numeric values for
constants or the implementation of functions as macros. Therefore, you
should never assume that binaries are compatible between different Lua
versions. Always recompile clients of the Lua API when using a new
version.

Similarly, Lua versions can always change the internal representation of
precompiled chunks; precompiled chunks are not compatible between
different Lua versions.

The standard paths in the official distribution may change between
versions.


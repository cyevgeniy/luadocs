# The Auxiliary Library

The *auxiliary library* provides several convenient functions to
interface C with Lua. While the basic API provides the primitive
functions for all interactions between C and Lua, the auxiliary library
provides higher-level functions for some common tasks.

All functions and types from the auxiliary library are defined in header
file `lauxlib.h` and have a prefix `luaL_`.

All functions in the auxiliary library are built on top of the basic
API, and so they provide nothing that cannot be done with that API.
Nevertheless, the use of the auxiliary library ensures more consistency
to your code.

Several functions in the auxiliary library use internally some extra
stack slots. When a function in the auxiliary library uses less than
five slots, it does not check the stack size; it simply assumes that
there are enough slots.

Several functions in the auxiliary library are used to check C function
arguments. Because the error message is formatted for arguments (e.g.,
\"`bad argument #1`\"), you should not use these functions for other
stack values.

Functions called `luaL_check*` always raise an error if the check is not
satisfied.


# The Standard Libraries

The standard Lua libraries provide useful functions that are implemented
in C through the C API. Some of these functions provide essential
services to the language (e.g., [`type`](#pdf-type) and
[`getmetatable`](#pdf-getmetatable)); others provide access to outside
services (e.g., I/O); and others could be implemented in Lua itself, but
that for different reasons deserve an implementation in C (e.g.,
[`table.sort`](#pdf-table.sort)).

All libraries are implemented through the official C API and are
provided as separate C modules. Unless otherwise noted, these library
functions do not adjust its number of arguments to its expected
parameters. For instance, a function documented as `foo(arg)` should not
be called without an argument.

The notation **fail** means a false value representing some kind of
failure. (Currently, **fail** is equal to **nil**, but that may change
in future versions. The recommendation is to always test the success of
these functions with `(not status)`, instead of `(status == nil)`.)

Currently, Lua has the following standard libraries:

-   basic library ([§6.1](#6.1));
-   coroutine library ([§6.2](#6.2));
-   package library ([§6.3](#6.3));
-   string manipulation ([§6.4](#6.4));
-   basic UTF-8 support ([§6.5](#6.5));
-   table manipulation ([§6.6](#6.6));
-   mathematical functions ([§6.7](#6.7)) (sin, log, etc.);
-   input and output ([§6.8](#6.8));
-   operating system facilities ([§6.9](#6.9));
-   debug facilities ([§6.10](#6.10)).

Except for the basic and the package libraries, each library provides
all its functions as fields of a global table or as methods of its
objects.

To have access to these libraries, the C host program should call the
[`luaL_openlibs`](#luaL_openlibs) function, which opens all standard
libraries. Alternatively, the host program can open them individually by
using [`luaL_requiref`](#luaL_requiref) to call
[`luaopen_base`]{#pdf-luaopen_base} (for the basic library),
[`luaopen_package`]{#pdf-luaopen_package} (for the package library),
[`luaopen_coroutine`]{#pdf-luaopen_coroutine} (for the coroutine
library), [`luaopen_string`]{#pdf-luaopen_string} (for the string
library), [`luaopen_utf8`]{#pdf-luaopen_utf8} (for the UTF-8 library),
[`luaopen_table`]{#pdf-luaopen_table} (for the table library),
[`luaopen_math`]{#pdf-luaopen_math} (for the mathematical library),
[`luaopen_io`]{#pdf-luaopen_io} (for the I/O library),
[`luaopen_os`]{#pdf-luaopen_os} (for the operating system library), and
[`luaopen_debug`]{#pdf-luaopen_debug} (for the debug library). These
functions are declared in [`lualib.h`]{#pdf-lualib.h}.


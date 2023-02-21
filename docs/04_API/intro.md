# The Application Program Interface

This section describes the C API for Lua, that is, the set of
C functions available to the host program to communicate with Lua. All
API functions and related types and constants are declared in the header
file [`lua.h`]{#pdf-lua.h}.

Even when we use the term \"function\", any facility in the API may be
provided as a macro instead. Except where stated otherwise, all such
macros use each of their arguments exactly once (except for the first
argument, which is always a Lua state), and so do not generate any
hidden side-effects.

As in most C libraries, the Lua API functions do not check their
arguments for validity or consistency. However, you can change this
behavior by compiling Lua with the macro
[`LUA_USE_APICHECK`]{#pdf-LUA_USE_APICHECK} defined.

The Lua library is fully reentrant: it has no global variables. It keeps
all information it needs in a dynamic structure, called the *Lua state*.

Each Lua state has one or more threads, which correspond to independent,
cooperative lines of execution. The type [`lua_State`](#lua_State)
(despite its name) refers to a thread. (Indirectly, through the thread,
it also refers to the Lua state associated to the thread.)

A pointer to a thread must be passed as the first argument to every
function in the library, except to [`lua_newstate`](#lua_newstate),
which creates a Lua state from scratch and returns a pointer to the
*main thread* in the new state.


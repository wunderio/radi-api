RADI API
---------------

The Radi API is a set of interfaces, builders and base implementations for a
Radi consumers Apps.  Radi consumers are expected to use the interfaces as
implemented by various handlers to build applications that can execute radi
operations through predictable patterns.

Most of the actual implementation comes in Handler packages which provide both 
builders and handlers, which can configure and provide operations. The API 
itself provides access to those operations.

# Concepts

## API

The Core API is a minimal functionality top level interface that should five
any consumer access to a list of operations that can themselves be executed.

The sole goal of an API implementation is to give access to operations.

In this API is also a large number of definitions for other pieces of 
architecture which the API uses.


### Projects

Projects are API implementations that know how to build themselves.  They can
be constructed with whatever tools the need.  They can then be asked to receive
and activate builders as needed.

The goal is to provide a more complex API implementation, in a tool that 
knows abuilt builders and handlers.  It can then be given Builders objects
and given instructions (with configuration) on activating those builders.
This gives us an object which can be configured dynamically based on stored
configurations, which means less manual building needed for an API consumer.

Having a Project gives your API consumer a vendor independent way of deciding 
what radi-handlers
to include, what vendor builders to include, and how to configure them.

The typical usage of a Project is:

1. create the Project
2. add various Builders to it, restricting available implementations (and 
   including whatever go libs are needed for it)
3. activate builders with configuration about what implementations that builder
   should activate.
4. Check it with validation

5. Get operations out of it (it is now a functioning API)

## Builders

Builders are vendor provided tools that know how to create handlers.  They
are optional go-between structs that Projects can use to activate handler 
functionality.  Builders can be used on their own, but are meant to be targets
for Projects.
Builders are pre-package variants of handler constructers, that when given
instructions and some settings, they can themselves wrap and activate a number
of handlers.

While Projects are your API consumer provided code, Builders are the vendor 
provided side of the building process.  Builders exist because often Handlers
have more complicated construction processes, and the Builders can streamline
though by centralizing and reducing the requirements.

Typical usage of a builder:

1. create the builder
2. run the Activate method repeatedly, telling the builder what implementations
   to turn on,
3. validate the builder

4. Get operations out of it (it is now a functioning API)

## Handlers

Handlers are the real center of the vendor provided radi-api implementations.
A handlers is meant to be the bottom level producer of operations.
A vendor radi implementation can use many different handlers, to provide 
different combinations of operations as needed.
They tend to be pretty simple objects, but can have some complex 
implementations that may lead to difficult building should probably be built
through a builder.

Typcial usage of a handler:

1. create the handler
2. validate the handler

3. Get operations out of the handler (it is probably now a functioning API)

## Operations

Operations are the core executive components of the radi-api.  They are the
individual structs, that allow API consumers to perform individual tasks.
They are reasonably descriptive, and are very much discoverable in terms 
of what configurations they allow, and how they execute.

## Property

A Property provides operations with tools to receive input from API consumers
and to provide output in return.  They are exectuable elements which are 
expected to handler their own threading etc.

## Result.

A Result is a self managing threadsafe response to any validation and
execution.
A result knows if the process it describes is completed, whether or not the
process was successful, and tracks errors recorded during execution.

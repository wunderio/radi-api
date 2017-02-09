# API Constructors and Builders

This optional API toolset provides a more automated and preconfigured method
for build an API implementation.

The idea is that handlers provide builder code that allow more automated
construction from settings, and then one of the various constructors is used
to decide which builders to user, and what settings to provide to the builders.

to recap:

1. A project provides various Builders which can be used to create handlers
from settings
2. A constructor is used to decide which builders to load, and then decides
what implementations of the builder to activate.

This separates the code for enabling handlers, from the decision on which 
handlers to enable and which settings to use with them.

About settings, most handlers still pull in their own settings concepts from
configwrappers, but settings from the building stage are still available for
them.

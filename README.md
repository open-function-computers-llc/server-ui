# server-ui
We have a few scripts that help us maintain and host various websites. This go program is a standalone app that helps us see and visualize everything happening on a single node.

## nutshell
This web app shouldn't be on an open port. Either park a firewall in front of this thing, or run the risk that all your controls are open to the world. We'll probably add more refined controls and security in the future. For now, just choose a port that makes sense for your use case and have fun!

## DB
This program is intended to start small, and stay that way. Therefore we are using in memory SQLite databases for all of the propigated data.

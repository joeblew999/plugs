# todo

Add domain fronting off cloudflare, using simple task file calls, and adjustment to the github pages setup.

---

https://github.com/titpetric/task-ui might work good ?

---

Move it all to https://github.com/joeblew999/plugs

- /Users/apple/workspace/go/src/github.com/joeblew999/plugs is there.

---

See if we can get it to a json file driving the plugin system too.

conduitio plugins
- https://github.com/orgs/conduitio-labs/repositories

benthos plugins 

Not sure of best way, but will require a way to pull the code over git, use go.work and then build each of their entry points. Some have many entry points too.  also what version we want.

I suspect that the cli, can have a dev aspect to help manage this json file and then run it.

It will be a golang thing, and when the cli is done can expose it ion the taskfile.

they are all golang proejcts btw so pretty easy. 

---

GUI as well as CLI

https://github.com/guigui-gui/guigui is really easy to use.

We can combine it into the CLI, so that it opens automatically.
We need to carefully design it so that the CLi and GUI use the same code and stay dry.

Because many users are not developers i favout the GUI opening off the binary, and then devs can somehow use the same binary to work at the CLI level. 


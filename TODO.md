# todo

---

https://docs.viam.com/dev/

We are headinfg to this.

go-judge will allow a real time codign experience of these plugins, and so we will need to incldue this as a server.



---


https://github.com/titpetric/task-ui might work good ? 

---

Work out how to add things on our hugo website to show the apps that we run. plugins is just one and will grow quickly.

Also this plugins repo is really a base that ALL our deployed projects could use. this is why we have the task file.  the only thing missing is a fly.io cli and docker and a server example. 

The server could be a nats servers that the plugin system uses, and so we end up with a generic and extensible deployment system. 

We need to think hard about this because these thngs become rerally hard later. You example, we dont release the atsk file at the monent and yet each proejct will need it.  Its the self simialr pattern ? 



---

See if we can get it to a json file driving the plugin system too.

conduitio plugins
- https://github.com/orgs/conduitio-labs/repositories

benthos plugins 

Not sure of best way, but will require a way to pull the code over git, use go.work and then build each of their entry points. Some have many entry points too.  also what version we want.

I suspect that the cli, can have a dev aspect to help manage this json file and then run it.

It will be a golang thing, and when the cli is done can expose it ion the taskfile.

they are all golang proejcts btw so pretty easy. 

All these pluigns have. no CGO, so build matrix is no problem.


---

GUI as well as CLI

https://github.com/guigui-gui/guigui is really easy to use.

https://github.com/go-via/via is my realtime htmx patterns web gui system.

I dont know what my users prefer


We can combine it into the CLI, so that it opens automatically.
We need to carefully design it so that the CLi and GUI use the same code and stay dry.

Because many users are not developers i favout the GUI opening off the binary, and then devs can somehow use the same binary to work at the CLI level. 


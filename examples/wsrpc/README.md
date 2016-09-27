# wsrpc

NOT finished !!

TODO:
- Need embedding to be done properly.
- web server running on goroutine needs to properly pass signals and logs back to main routine.


This is just a simple example an application with a Frontend (running in the embedded browser) and a Backend running in the same process serving http & web-sockets.

try it by using:
````
chmod +x ./t-buildrun.sh
./t-buildrun
````


## How it works
The main point is to show a basic outline of a real world app.
This code is far from best practice, but just shows the concepts.

In the "browser" folder
- Gopherjs is used to bind functionality to the static HTML. 
- Web sockets are setup to send and receive.


In the "webserver" folder
- The embedded browser is setup.
- Web sockets are setup to send and receive data

In the "common" folder is the generic functionality, independent of the business logic.

In the "shared" folder are shared types that both the Client and Server rely on.





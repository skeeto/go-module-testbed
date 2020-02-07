# Go Module Testbed

This script creates and hosts a testbed environment for experimenting
with Go modules. It runs both a static web server to host module paths,
and Git daemon to host the modules themselves. Both are bound to
127.0.0.1 and are not externally accessible. Security is disabled for
modules hosted by the testbed, so no certificates are required.

Requires **Go 1.14** or later.

## Usage

Execute the testbed script to initialize a testbed directory and host
the testbed servers within it. The default testbed directory is the
current working directory, and an optional argument supplies an
alternate testbed directory.

    ./go-module-testbed [dir [options...]]
      -p port    HTTP server port

The testbed directory will be populated with an `activate` shell script
and a example module named `127.0.0.1/example` at v1.0.0.

The web server *must* be available on port 80 which, unfortunately,
requires mucking about with system configuration. The web server is
hosted by default on port 8001, so on Linux a temporary `iptables`
configuration will do the trick:

     iptables -t nat -I OUTPUT -p tcp -d 127.0.0.1 --dport 80 -j REDIRECT --to-ports 8001

With the servers running successfully, start a new shell and source the
`activate` script in the testbed directory. All Go commands will be
isolated to the testbed, and modules can be fetched from `src/` by
module paths starting with `127.0.0.1/`. To check if it's working:

    $ go get 127.0.0.1/example/cmd/demo
    go: downloading 127.0.0.1/example v1.0.0
    go: found 127.0.0.1/example/cmd/demo in 127.0.0.1/example v1.0.0
    $ demo
    Example v1.0.0

Import path responses are stored in `www/`. The example can serve as a
template when creating new modules.

## Transparent and Obvious

The HTTP server could be smarter and handle `?go-get=1` queries itself.
However, one of the goals of this tool is to experiment with different
kinds of responses, and a static file host is the simplest way to go
about it. You can always host a different HTTP server on port 80 to
experiment with other configurations.

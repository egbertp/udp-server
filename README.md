# UDP server

This GoLang mini-project creates a binary file that listens on an UDP port and prints the output of the UDP port to the stdout.

I use this `udp-server` project to verify the correct working of my infrastructure.

## Usage

Download:

```
mkdir $GOPATH/src/github.com/egbertp/udp-server
cd $GOPATH/src/github.com/egbertp/udp-server
git clone https://github.com/egbertp/udp-server.git
```

### Build and Run

Install dependencies using [Glide Package Management for Go](https://glide.sh/)

```
$ glide install
```

Build the binary
```
$ make build
```

Release the app
```
$ make release
```

Run the app
```
$ ./udp-server
```

### Use tha app

Point your browser to `http://localhost:7000/`
```
{
	message: "Hello world"
}
```
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
$ ./udp-server --port 37059
```

```
2018/02/16 14:33:57 UDP sever running on port: 37059
```

Send UDP package on OSx
```
echo -n "Hello" | nc -cu 127.0.0.1 37059
```
You will see:
```
2018/02/16 14:34:54 Received  Hello  from  127.0.0.1:58102
```


Send UDP package on Linux
```
echo "This is my data" > /dev/udp/127.0.0.1/37059
```
You will see:
```
2018/02/16 14:34:54 Received  This is my data  from  127.0.0.1:58102
```


Send UDP packages using my [udp-client](https://github.com/egbertp/udp-client) application
```
./udp-client --address 127.0.0.1 --port 37059

2018/02/16 14:03:40 Sending packet with payload: 0 to 127.0.0.1:37059
2018/02/16 14:03:41 Sending packet with payload: 1 to 127.0.0.1:37059
2018/02/16 14:03:42 Sending packet with payload: 2 to 127.0.0.1:37059
2018/02/16 14:03:43 Sending packet with payload: 3 to 127.0.0.1:37059
2018/02/16 14:03:44 Sending packet with payload: 4 to 127.0.0.1:37059
2018/02/16 14:03:45 Sending packet with payload: 5 to 127.0.0.1:37059
2018/02/16 14:03:46 Sending packet with payload: 6 to 127.0.0.1:37059
```

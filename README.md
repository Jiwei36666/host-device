# host-device
A CNI Plugin to move a specific network device into a containers network namespace

$ go get github.com/containernetworking/cni

$ go get github.com/containernetworking/plugins

$ go get github.com/Jiwei36666/host-device


move or copy Jiwei36666/host-device to $GOPATH/src/github.com/containernetworking/plugins/plugins/main/

$ cp -rf host-device $GOPATH/src/github.com/containernetworking/plugins/plugins/main/


goto $GOPATH/src/github.com/containernetworking/plugins/;./build.sh

$ cd $GOPATH/src/github.com/containernetworking/plugins/;./build.sh

$ copy 10-mynet.conf to /etc/cni/net.d/ (create the folder if it doens't exist).

## Follow https://github.com/containernetworking/cni to run some example.

$ CNI_PATH=$GOPATH/src/github.com/containernetworking/plugins/bin

$ cd $GOPATH/src/github.com/containernetworking/cni/scripts

$ sudo CNI_PATH=$CNI_PATH ./priv-net-run.sh ifconfig

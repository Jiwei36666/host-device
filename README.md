# host-device
A CNI Plugin to move a specific network device into a containers network namespace
```
$ go get github.com/containernetworking/cni

$ go get github.com/containernetworking/plugins

$ go get github.com/Jiwei36666/host-device
```

move or copy Jiwei36666/host-device to $GOPATH/src/github.com/containernetworking/plugins/plugins/main/
```
$ cp -rf host-device $GOPATH/src/github.com/containernetworking/plugins/plugins/main/
```

goto $GOPATH/src/github.com/containernetworking/plugins/;./build.sh
```
$ cd $GOPATH/src/github.com/containernetworking/plugins/;./build.sh

$ copy 10-mynet.conf to /etc/cni/net.d/ (create the folder if it doens't exist).
```

## Follow https://github.com/containernetworking/cni to run some example.
```

$ CNI_PATH=$GOPATH/src/github.com/containernetworking/plugins/bin

$ cd $GOPATH/src/github.com/containernetworking/cni/scripts

$ sudo CNI_PATH=$CNI_PATH ./priv-net-run.sh ifconfig
```

# appendix 
Create device for LS1088 DPAA2 EVB scenario. 
```

# ls-addmux -m=DPDMUX_METHOD_MAC -d=2 dpmac.3 
# restool dpdmux info dpdmux.0
dpdmux version: 6.0
dpdmux id: 0
plugged state: plugged
endpoints:
interface 0:
	connection: dpmac.3
	link state: up
interface 1:
	connection: none
	link state: n/a
interface 2:
	connection: none
	link state: n/a
dpdmux_attr.options value is: 0x2
	DPDMUX_OPT_BRIDGE_EN
# ls-addni --no-link
# restool dprc connect dprc.1 --endpoint1=dpdmux.0.1 --endpoint2=dpni.1
# ls-addni --no-link 
# restool dprc connect dprc.1 --endpoint1=dpdmux.0.2 --endpoint2=dpni.2
# restool dpdmux info dpdmux.0

dpdmux version: 6.0
dpdmux id: 0
plugged state: plugged
endpoints:
interface 0:
	connection: dpmac.3
	link state: up
interface 1:
	connection: dpni.1
	link state: down
interface 2:
	connection: dpni.2
	link state: down
dpdmux_attr.options value is: 0x2
	DPDMUX_OPT_BRIDGE_EN
DPDMUX address table method: DPDMUX_METHOD_MAC
DPDMUX manipulation type: DPDMUX_MANIP_NONE
number of interfaces (excluding the uplink interface): 2
frame storage memory size: 0
```

Iperf logs for different container(different host)

```
root@428489f4e391:/# iperf -c 100.10.1.1
------------------------------------------------------------
Client connecting to 100.10.1.1, TCP port 5001
TCP window size: 85.0 KByte (default)
------------------------------------------------------------
[  3] local 100.10.1.2 port 58960 connected with 100.10.1.1 port 5001
[ ID] Interval       Transfer     Bandwidth
[  3]  0.0-10.0 sec  1.29 GBytes  1.11 Gbits/sec
root@428489f4e391:/# iperf -c 100.10.1.1 -t 60
------------------------------------------------------------
Client connecting to 100.10.1.1, TCP port 5001
TCP window size: 85.0 KByte (default)
------------------------------------------------------------
[  3] local 100.10.1.2 port 58962 connected with 100.10.1.1 port 5001
[ ID] Interval       Transfer     Bandwidth
[  3]  0.0-60.0 sec  7.78 GBytes  1.11 Gbits/sec
root@428489f4e391:/# iperf -c 100.10.1.1 -t 60
------------------------------------------------------------
Client connecting to 100.10.1.1, TCP port 5001
TCP window size: 85.0 KByte (default)
------------------------------------------------------------
[  3] local 100.10.1.2 port 58964 connected with 100.10.1.1 port 5001
[ ID] Interval       Transfer     Bandwidth
[  3]  0.0-60.0 sec  7.26 GBytes  1.04 Gbits/sec
root@428489f4e391:/#
root@428489f4e391:/#
root@428489f4e391:/#
root@428489f4e391:/# iperf -c 100.10.1.1 -t 60
------------------------------------------------------------
Client connecting to 100.10.1.1, TCP port 5001
TCP window size: 85.0 KByte (default)
------------------------------------------------------------
[  3] local 100.10.1.2 port 58966 connected with 100.10.1.1 port 5001
[ ID] Interval       Transfer     Bandwidth
[  3]  0.0-60.0 sec  7.92 GBytes  1.13 Gbits/sec
root@428489f4e391:/# iperf -c 100.10.1.1 -t 60
------------------------------------------------------------
Client connecting to 100.10.1.1, TCP port 5001
TCP window size: 85.0 KByte (default)
------------------------------------------------------------
[  3] local 100.10.1.2 port 58968 connected with 100.10.1.1 port 5001
[ ID] Interval       Transfer     Bandwidth
[  3]  0.0-60.0 sec  7.26 GBytes  1.04 Gbits/sec
root@428489f4e391:/#
```

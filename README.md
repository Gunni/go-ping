# go-ping

* go-ping is a simple program to perform TCP handshakes on-demand.

* go-ping will output easy-to-parse output.

* go-ping will perform a full TCP handshake before the next iteration in its loop.

* go-ping was created to be nicer on TCP stacks and to not piss off DDoS protection systems.

# Example output

    1;82.221.37.18:80;0.000995
    2;82.221.37.18:80;0.000834
    3;82.221.37.18:80;0.001014
    4;82.221.37.18:80;0.000766
    5;82.221.37.18:80;0.000713

# Example ping tcpdump (1 ping)

    IP a.51589 > b.http: Flags [S], seq 312853653, win 14600, options [mss 1460,sackOK,TS val 2552163837 ecr 0,nop,wscale 7], length 0
    IP b.http > a.51589: Flags [S.], seq 3580832727, ack 312853654, win 28960, options [mss 1380,sackOK,TS val 1253625379 ecr 2552163837,nop,wscale 7], length 0
    IP a.51589 > b.http: Flags [.], ack 1, win 115, options [nop,nop,TS val 2552163838 ecr 1253625379], length 0
    IP a.51589 > b.http: Flags [F.], seq 1, ack 1, win 115, options [nop,nop,TS val 2552163839 ecr 1253625379], length 0
    IP b.http > a.51589: Flags [F.], seq 1, ack 2, win 227, options [nop,nop,TS val 1253625380 ecr 2552163839], length 0
    IP a.51589 > b.http: Flags [.], ack 2, win 115, options [nop,nop,TS val 2552163840 ecr 1253625380], length 0`

# TODO

https://git.system.is/kjarni/go-ping/issues

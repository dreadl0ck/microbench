#!/bin/bash

# clean
ip link del tap0

iptables -F

ip tuntap add tap0 mode tap
ip addr add 145.100.106.17/28 dev tap0

ip link set tap0 up

#brctl addif docker0 tap0

sh -c "echo 1 > /proc/sys/net/ipv4/ip_forward"
iptables -t nat -A POSTROUTING -o eno1 -j MASQUERADE
iptables -A FORWARD -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT
iptables -A FORWARD -i tap0 -o eno1 -j ACCEPT

ifconfig tap0

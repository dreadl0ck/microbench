#!/bin/bash

gw_ip=$1
num=$2

if [ -z "$gw_ip" ]; then
    echo "you must pass an ip as parameter #1"
    echo "usage: ./create_tap.sh <ip> <num>"
    exit 1
fi

if [ -z "$num" ]; then
    echo "you must pass a number for the tap as parameter #2"
    echo "usage: ./create_tap.sh <ip> <num>"
    exit 1
fi

echo "creating tap$num"

# clean
ip link del "tap$num"

#iptables -F

ip tuntap add "tap$num" mode tap
ip addr add "$gw_ip"/28 dev "tap$num"

ip link set "tap$num" up

#brctl addif docker0 tap0

# sh -c "echo 1 > /proc/sys/net/ipv4/ip_forward"
iptables -t nat -A POSTROUTING -o eno1 -j MASQUERADE
iptables -A FORWARD -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT
iptables -A FORWARD -i "tap$num" -o eno1 -j ACCEPT

ifconfig "tap$num"


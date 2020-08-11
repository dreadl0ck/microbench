#!/bin/bash

for i in {0..300}; do
  ip link del "tap$i"
done

ip link
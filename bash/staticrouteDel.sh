#!/bin/bash

localnets=(192.168.0.0/24 172.16.8.0/24)
iproute=192.168.0.1

for net in localnets; do
	route del -net $net gw $iproute
done
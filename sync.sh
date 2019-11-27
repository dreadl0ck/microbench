#!/bin/bash

# make gopath writeable for current user: chown -R pmieden /root/go/
scp -r ../firebench pmieden@oslo.studlab.os3.nl:/home/pmieden/go/src/github.com/dreadl0ck

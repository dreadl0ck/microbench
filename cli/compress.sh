#!/bin/bash

rm -f experiment_logs.zip

chmod -R 755 experiment_logs
chown -R pmieden:pmieden experiment_logs

zip experiment_logs.zip -r experiment_logs
cp experiment_logs.zip /home/pmieden/

echo "done"
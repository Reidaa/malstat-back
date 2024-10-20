#!/usr/bin/env sh


cd ../../ && make build && cd - || exit
ansible-playbook deploy.yml -vv 
cd ../../ && make clean
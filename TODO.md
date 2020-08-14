# TODOs

- note that only linux is supported atm
- note that rootfs need to be created per machine, because they contain network interface information and addresses 
- add notes for general setup and experiment workflow:

server:

    su root
    cd path/to/microbench
    cli/run_experiments.sh

workstation:    
    
    fetch_logs.sh
    go run cli/resultparse/main.go

- remove all os3 identifiers

- Go logger: dont log relative time, use proper timestamps
- microbench: use relative paths to get rid of the need for $GOPATH to be set
- don't run microbench as root: mount operations in scripts require root priv:

https://unix.stackexchange.com/questions/96625/how-to-allow-non-superusers-to-mount-any-filesystem
    
    Modify the entry in /etc/fstab corresponding to the filesystem you want to mount, adding the flag user to this entry. 
    Non-privilege users would then be able to mount 
    See man mount for more details.

- os user creation
- networking
- structured JSON logging
- code cleanup and documentation
- create jail if not exist
- use jailer time

- make num cpus and RAM configurable for both engine types, and include current config in log files
- firecracker make cpu template configurable and add tests with C3 type
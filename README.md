# firecracker setup

## check for hardware virtualisation support

    egrep -c '(vmx|svm)' /proc/cpuinfo

> must be > 0

## install KVM and docker

    # apt install qemu-kvm libvirt-bin bridge-utils
    # apt install vim tree curl docker.io make gcc

fetch kernel and rootfs:

    wget https://s3.amazonaws.com/spec.ccfc.min/img/hello/kernel/hello-vmlinux.bin
    wget https://s3.amazonaws.com/spec.ccfc.min/img/hello/fsfiles/hello-rootfs.ext4

build firecracker from source:

    git clone https://github.com/firecracker-microvm/firecracker
    cd firecracker
    tools/devtool build
    toolchain="$(uname -m)-unknown-linux-musl"

run:

    cd firecracker && \
    toolchain="$(uname -m)-unknown-linux-musl" && \
    rm -f /tmp/firecracker.socket && \
    build/cargo_target/${toolchain}/debug/firecracker --api-sock /tmp/firecracker.socket

In your **second shell** prompt:

- get the kernel and rootfs, if you don't have any available:

  ```bash
    arch=`uname -m`
    dest_kernel="hello-vmlinux.bin"
    dest_rootfs="hello-rootfs.ext4"
    image_bucket_url="https://s3.amazonaws.com/spec.ccfc.min/img"

    if [ ${arch} = "x86_64" ]; then
            kernel="${image_bucket_url}/hello/kernel/hello-vmlinux.bin"
            rootfs="${image_bucket_url}/hello/fsfiles/hello-rootfs.ext4"
    elif [ ${arch} = "aarch64" ]; then
            kernel="${image_bucket_url}/aarch64/ubuntu_with_ssh/kernel/vmlinux.bin"
            rootfs="${image_bucket_url}/aarch64/ubuntu_with_ssh/fsfiles/xenial.rootfs.ext4"
    else
            echo "Cannot run firecracker on $arch architecture!"
            exit 1
    fi

    echo "Downloading $kernel..."
    curl -fsSL -o $dest_kernel $kernel

    echo "Downloading $rootfs..."
    curl -fsSL -o $dest_rootfs $rootfs

    echo "Saved kernel file to $dest_kernel and root block device to $dest_rootfs."
  ```

- set the guest kernel:

  ```bash
    arch=`uname -m`
    kernel_path="$(readlink -f hello-vmlinux.bin)"
    file $kernel_path

    if [ ${arch} = "x86_64" ]; then
      curl --unix-socket /tmp/firecracker.socket -i \
          -X PUT 'http://localhost/boot-source'   \
          -H 'Accept: application/json'           \
          -H 'Content-Type: application/json'     \
          -d "{
                \"kernel_image_path\": \"${kernel_path}\",
                \"boot_args\": \"console=ttyS0 reboot=k panic=1 pci=off\"
           }"
    elif [ ${arch} = "aarch64" ]; then
        curl --unix-socket /tmp/firecracker.socket -i \
          -X PUT 'http://localhost/boot-source'   \
          -H 'Accept: application/json'           \
          -H 'Content-Type: application/json'     \
          -d "{
                \"kernel_image_path\": \"${kernel_path}\",
                \"boot_args\": \"keep_bootcon console=ttyS0 reboot=k panic=1 pci=off\"
           }"
    else
        echo "Cannot run firecracker on $arch architecture!"
        exit 1
    fi
  ```

- set the guest rootfs:

  ```bash
    rootfs_path="$(readlink -f hello-rootfs.ext4)"
    file $rootfs_path

    curl --unix-socket /tmp/firecracker.socket -i \
      -X PUT 'http://localhost/drives/rootfs' \
      -H 'Accept: application/json'           \
      -H 'Content-Type: application/json'     \
      -d "{
            \"drive_id\": \"rootfs\",
            \"path_on_host\": \"${rootfs_path}\",
            \"is_root_device\": true,
            \"is_read_only\": false
       }"
  ```

- start the guest machine:

  ```bash
  curl --unix-socket /tmp/firecracker.socket -i \
      -X PUT 'http://localhost/actions'       \
      -H  'Accept: application/json'          \
      -H  'Content-Type: application/json'    \
      -d '{
          "action_type": "InstanceStart"
       }'
  ```

```bash
  curl --unix-socket /tmp/firecracker.socket -i \
      -X PUT 'http://localhost/actions'       \
      -H  'Accept: application/json'          \
      -H  'Content-Type: application/json'    \
      -d '{
          "action_type": "SendCtrlAltDel"
       }'
  ```

Going back to your first shell, you should now see a serial TTY prompting you
to log into the guest machine. If you used our `hello-rootfs.ext4` image,
you can login as `root`, using the password `root`.

When you're done, issuing a `reboot` command inside the guest will actually
shutdown Firecracker gracefully. This is due to the fact that Firecracker
doesn't implement guest power management.

### VM Configuration via API

**Note**: the default microVM will have 1 vCPU and 128 MiB RAM. If you wish to
customize that (say, 2 vCPUs and 1024MiB RAM), you can do so before issuing
the `InstanceStart` call, via this API command:

```bash
curl --unix-socket /tmp/firecracker.socket -i  \
    -X PUT 'http://localhost/machine-config' \
    -H 'Accept: application/json'            \
    -H 'Content-Type: application/json'      \
    -d '{
        "vcpu_count": 2,
        "mem_size_mib": 1024,
        "ht_enabled": false
    }'
```

#### Configuring the microVM without sending API requests

If you'd like to boot up a guest machine without using the API socket, you can do that 
by passing the parameter `--config-file` to the Firecracker process. The command for 
starting Firecracker with this option will look like this:

```bash
./firecracker --api-sock /tmp/firecracker.socket --config-file 
<path_to_the_configuration_file>
```

`path_to_the_configuration_file` should represent the path to a file that contains a 
JSON which stores the entire configuration for all of the microVM's resources. The 
JSON **must** contain the configuration for the guest kernel and rootfs, as these 
are mandatory, but all of the other resources are optional, so it's your choice 
if you want to configure them or not. Because using this configuration method will 
also start the microVM, you need to specify all desired pre-boot configurable resources 
in that JSON. The names of the resources are the ones from the `firecracker.yaml` file 
and the names of their fields are the same that are used in API requests. 
You can find an example of configuration file at `tests/framework/vm_config.json`. 
After the machine is booted, you can still use the socket to send API requests
for post-boot operations.

### Install firectl tool

#### Install latest go on ubuntu

Install via snap:

    apt install snapd
    snap install go --classic

In ~/.bashrc:

    export PATH="$PATH:/snap/bin"
    export GOPATH="/home/pmieden/go"
    export GOBIN="$GOPATH/bin"
    export PATH="$PATH:$GOBIN:$HOME/.cargo/bin"

Reload

    source ~/.bashrc

Firectl:

    go get github.com/firecracker-microvm/firectl

fails with several errors. Go into the directory and install manually using go modules:

    github.com/firecracker-microvm/firectl:$ go install

### Create Disk Image

    qemu-img create -f qcow2 file.qcow2 100M
    mkfs.ext4 file.qcow2
    mount file.qcow2 /mnt
    ...
    umount /mnt

### StartVM with firectl

StartVM:

    firectl \
    --kernel=/tmp/hello-vmlinux.bin \
    --root-drive=/tmp/hello-rootfs.ext4 \
    --kernel-opts="console=ttyS0 noapic reboot=k panic=1 pci=off nomodules rw" \
    --add-drive=file.qcow2:rw

or

    ./firectl \
    --kernel=hello-vmlinux.bin \
    --root-drive=hello-rootfs.ext4 \
    --add-drive=file.qcow2:rw

### Mount disk image

    fdisk -l
    mount /dev/vdb /mnt

### Mount Volume and Network Setup

See: https://medium.com/@s8sg/quick-start-with-firecracker-and-firectl-in-ubuntu-f58aeedae04b

### Deploy Go Executable in AWS Lambda

Documentation: https://docs.aws.amazon.com/lambda/latest/dg/lambda-go-how-to-create-deployment-package.html

Download the Lambda library for Go with go get, and compile your executable.

    ~/my-function$ go get github.com/aws/aws-lambda-go/lambda
    ~/my-function$ GOOS=linux go build main.go

Create a deployment package by packaging the executable in a ZIP file, and use the AWS CLI to create a function. The handler parameter must match the name of the executable containing your handler.


    ~/my-function$ zip function.zip main
    ~/my-function$ aws lambda create-function --function-name my-function --runtime go1.x \
    --zip-file fileb://function.zip --handler main \
    --role arn:aws:iam::123456789012:role/execution_role

Install aws tools:

    brew install awscli
    aws configure

Find role name in AWS Identity and Access Management (IAM) and query API:

    dreadbook:websrv alien$ aws iam get-role --role-name test-role-ct1y0gwd
    {
        "Role": {
            "Path": "/service-role/",
            "RoleName": "test-role-ct1y0gwd",
            "RoleId": "AROA5NS7SA6O5UEFKEZGH",
            "Arn": "arn:aws:iam::922544637853:role/service-role/test-role-ct1y0gwd",
            "CreateDate": "2019-11-14T16:49:14Z",
            "AssumeRolePolicyDocument": {
                "Version": "2012-10-17",
                "Statement": [
                    {
                        "Effect": "Allow",
                        "Principal": {
                            "Service": "lambda.amazonaws.com"
                        },
                        "Action": "sts:AssumeRole"
                    }
                ]
            },
            "MaxSessionDuration": 3600
        }
    }

Deploy:

    dreadbook:websrv alien$ aws lambda create-function --function-name my-function --runtime go1.x   --zip-file fileb://function.zip --handler main   --role arn:aws:iam::922544637853:role/service-role/test-role-ct1y0gwd
    {
        "FunctionName": "my-function",
        "FunctionArn": "arn:aws:lambda:eu-west-1:922544637853:function:my-function",
        "Runtime": "go1.x",
        "Role": "arn:aws:iam::922544637853:role/service-role/test-role-ct1y0gwd",
        "Handler": "main",
        "CodeSize": 3821143,
        "Description": "",
        "Timeout": 3,
        "MemorySize": 128,
        "LastModified": "2019-11-14T20:18:39.549+0000",
        "CodeSha256": "lzi65uyX7IbnW/S0xWrydlmuEKDQ7Dzg8h10i0dwGOg=",
        "Version": "$LATEST",
        "TracingConfig": {
            "Mode": "PassThrough"
        },
        "RevisionId": "d4bc2377-7930-499a-8c46-6740b24ca9f2"
    }

> Create an API Gateway in Lambda Console.

Invoke:

    aws lambda invoke --function-name hello --payload '{ "key": "value" }' response.json


### Custom RootFS for experiments

#### Compile websrv on ubuntu host

    apt install musl-tools
    root@oslo:~/go/src/github.com/dreadl0ck/os3/ls/websrv# CC=musl-gcc go build --ldflags '-linkmode external -extldflags "-static"'

#### Setup FS

Source: https://github.com/firecracker-microvm/firecracker/blob/master/docs/rootfs-and-kernel-setup.md

- create_rootfs.sh (outside)
- init_alpine.sh (run inside container):

Hint: You can also copy the binary from the outside into the container:

    docker cp ~/go/src/github.com/dreadl0ck/os3/ls/websrv/websrv $(docker ps -q --filter ancestor=alpine):/my-rootfs/usr/bin

Umount on server:
    
    sync
    umount /tmp/my-rootfs

You should now have a kernel image (vmlinux) and a rootfs image (rootfs.ext4), that you can boot with Firecracker.


### IP space

Alpine Network setup: https://wiki.alpinelinux.org/wiki/Configure_Networking

https://github.com/firecracker-microvm/firecracker/blob/master/docs/network-setup.md

    Address:   145.100.106.16        10010001.01100100.01101010.0001 0000
    Netmask:   255.255.255.240 = 28  11111111.11111111.11111111.1111 0000
    Wildcard:  0.0.0.15              00000000.00000000.00000000.0000 1111
    =>
    Network:   145.100.106.16/28     10010001.01100100.01101010.0001 0000 (Class B)
    Broadcast: 145.100.106.31        10010001.01100100.01101010.0001 1111
    HostMin:   145.100.106.17        10010001.01100100.01101010.0001 0001
    HostMax:   145.100.106.30        10010001.01100100.01101010.0001 1110
    Hosts/Net: 14
    
#### Run Firecracker Unit Tests

    cd firecracker
    tools/devtool test 2>&1 > firecracker_tests.log 

#### Run Rust Linter

    rustup target add x86_64-unknown-linux-musl
    cargo clippy
    ...
    
### Automate microVM start
    
    dreadl0ck
    cd firebench
    cli/create_rootfs.sh && bin/firebench 145.100.106.18

### Benchmark kernel boot time

Parse */var/log/boot.msg* or */var/log/kern.log* for information regarding kernel boot time.

See: https://unix.stackexchange.com/questions/500732/how-to-find-out-time-taken-by-linux-system-for-cold-boot

Those do not seem to be present in alpine.

Check *cat /proc/uptime* and *dmesg* output.

Alternatively:

- https://wiki.archlinux.org/index.php/Bootchart
- http://people.redhat.com/berrange/systemtap/bootprobe/

What values are we interested in?

- Kernel boot time (via kernel logs)
- Service startup time (via kernel logs)
- Network stack reachability (via ping from outside)
- Time until reachability of a static webservice (via HTTP GET from outside)

### Firecracker Linter Output

- 9x warning: unsafe function's docs miss `# Safety` (all in memory_model)
- 1x deprecated symbol used (in devices)

### Firecracker Test Output

What security aspects are tested?

- seccomp
- sec_audit
- jailing

Search test logs:

    $ grep "security" logs/firecracker_tests.log
    integration_tests/security/test_jail.py ..
    integration_tests/security/test_sec_audit.py .
    integration_tests/security/test_seccomp.py .Hello, world!
    5.99s setup    integration_tests/security/test_seccomp.py::test_seccomp_ls
    1.29s call     integration_tests/security/test_sec_audit.py::test_cargo_audit
    1.26s teardown integration_tests/security/test_seccomp.py::test_seccomp_applies_to_all_threads[ubuntu]
    0.55s setup    integration_tests/security/test_jail.py::test_default_chroot[ubuntu_with_ssh]
    0.54s setup    integration_tests/security/test_jail.py::test_empty_jailer_id[ubuntu_with_ssh]
    0.29s call     integration_tests/security/test_jail.py::test_default_chroot[ubuntu_with_ssh]
    0.19s call     integration_tests/security/test_seccomp.py::test_seccomp_applies_to_all_threads[ubuntu]
    0.09s setup    integration_tests/security/test_seccomp.py::test_seccomp_applies_to_all_threads[ubuntu]
    0.07s teardown integration_tests/security/test_jail.py::test_default_chroot[ubuntu_with_ssh]
    0.07s call     integration_tests/security/test_jail.py::test_empty_jailer_id[ubuntu_with_ssh]
    0.07s teardown integration_tests/security/test_jail.py::test_empty_jailer_id[ubuntu_with_ssh]
    0.01s call     integration_tests/security/test_seccomp.py::test_seccomp_ls

Performance measurements:

    $ grep "performance" logs/firecracker_tests.log
    integration_tests/performance/test_boottime.py FFFF
    integration_tests/performance/test_process_startup_time.py Process startup time is: 23428 us (6601 CPU us)
    integration_tests/performance/test_boottime.py:30: in test_single_microvm_boottime_no_network
    integration_tests/performance/test_boottime.py:98: in _test_microvm_boottime
    integration_tests/performance/test_boottime.py:46: in test_multiple_microvm_boottime_no_network
    integration_tests/performance/test_boottime.py:98: in _test_microvm_boottime
    integration_tests/performance/test_boottime.py:62: in test_multiple_microvm_boottime_with_network
    integration_tests/performance/test_boottime.py:138: in _configure_vm
    integration_tests/performance/test_boottime.py:80: in test_single_microvm_boottime_with_network
    integration_tests/performance/test_boottime.py:98: in _test_microvm_boottime
    3.99s call     integration_tests/performance/test_boottime.py::test_multiple_microvm_boottime_no_network[minimal, 10 instance(s)]
    2.95s call     integration_tests/performance/test_boottime.py::test_multiple_microvm_boottime_with_network[minimal, 10 instance(s)]
    0.81s call     integration_tests/performance/test_boottime.py::test_single_microvm_boottime_with_network[minimal]
    0.64s call     integration_tests/performance/test_boottime.py::test_single_microvm_boottime_no_network[minimal]
    0.57s call     integration_tests/performance/test_process_startup_time.py::test_startup_time[ubuntu]
    0.12s setup    integration_tests/performance/test_boottime.py::test_multiple_microvm_boottime_with_network[minimal, 10 instance(s)]
    0.12s setup    integration_tests/performance/test_boottime.py::test_multiple_microvm_boottime_no_network[minimal, 10 instance(s)]
    0.11s setup    integration_tests/performance/test_boottime.py::test_single_microvm_boottime_with_network[minimal]
    0.11s setup    integration_tests/performance/test_boottime.py::test_single_microvm_boottime_no_network[minimal]
    0.10s setup    integration_tests/performance/test_process_startup_time.py::test_startup_time[ubuntu]
    0.05s teardown integration_tests/performance/test_boottime.py::test_multiple_microvm_boottime_no_network[minimal, 10 instance(s)]
    0.03s teardown integration_tests/performance/test_boottime.py::test_multiple_microvm_boottime_with_network[minimal, 10 instance(s)]
    0.01s teardown integration_tests/performance/test_boottime.py::test_single_microvm_boottime_no_network[minimal]
    0.01s teardown integration_tests/performance/test_process_startup_time.py::test_startup_time[ubuntu]
    0.01s teardown integration_tests/performance/test_boottime.py::test_single_microvm_boottime_with_network[minimal]

- integration_tests/performance/test_process_startup_time.py Process startup time is: 23428 us (6601 CPU us)
- Total: 5 failed, 64 passed, 1 skipped in 946.56s (0:15:46)
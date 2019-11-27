#!/bin/bash

apk update
apk add openrc vim util-linux

# Set up a login terminal on the serial console (ttyS0):
ln -s agetty /etc/init.d/agetty.ttyS0
echo ttyS0 > /etc/securetty
rc-update add agetty.ttyS0 default

# Make sure special file systems are mounted on boot:
rc-update add devfs boot
rc-update add procfs boot
rc-update add sysfs boot

# Networking
rc-update add networking boot

# set default password
passwd -d "root"

# Now add startup item inside container
cat > /etc/init.d/serve <<EOF
#!/sbin/openrc-run
command="direct-fs"
command_background="yes"
pidfile="/run/$RC_SVCNAME/$RC_SVCNAME.pid"
EOF

# Make executable
chmod 755 /etc/init.d/serve

# register to launch at startup
rc-update add serve

# Network config
cat > /etc/network/interfaces <<EOF
auto lo
iface lo inet loopback

auto eth0
iface eth0 inet static
    address 145.100.106.18
    netmask 255.255.255.240
    gateway 145.100.106.17
EOF

# Optionally add a local.d script for setup tasks during boot time
# docs: https://wiki.gentoo.org/wiki//etc/local.d
#rc-update add local default
# cat > /etc/local.d/setup.start <<EOF
#     #!/bin/ash
#     echo "hello"
# EOF
# chmod 755 /etc/local.d/setup.start

echo "firebench" > /etc/hostname
#hostname -F /etc/hostname

echo "nameserver 9.9.9.9" > /etc/resolv.conf

hexdump -C /bin/ash > /etc/hexdump

# Then, copy the newly configured system to the rootfs image
for d in bin etc lib root sbin usr; do tar c "/$d" | tar x -C /my-rootfs; done
for dir in dev proc run sys var; do mkdir /my-rootfs/${dir}; done

mkdir /my-rootfs/run/network
mkdir /my-rootfs/var/run
mkdir /my-rootfs/lib/modules
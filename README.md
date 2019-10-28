# go-dromaius

`go-dromaius` is a tool that wraps commands from [libvirt](https://libvirt.org/)'s `virsh`. It allows a hypervisor administrator to
give users access to their virtual machine(s) without giving full access to _all_ virtual machines.

## Rationale

Usually, a hypervisor administrator can use `polkit(8)` to restrict access to users. See for example
[this answer](https://serverfault.com/questions/797526/restrict-access-to-kvm-virtual-machines-to-specific-users/845973#845973) to a
Serverfault question on restricting access to KVM virtual machines. Unfortunately, all examples assume a `polkit` version >= 0.106,
whereas this version is still not available in [Ubuntu 19.04 (disco)](https://packages.ubuntu.com/disco/libpolkit-backend-1-0) nor
[Debian 10 (Buster)](https://packages.debian.org/buster/libpolkit-backend-1-0). The restrictions available in `polkit` < 0.106 are
not fine-grained enough to restrict access to _one_ virtual machine. See also
[these](https://serverfault.com/questions/949393/limiting-access-to-one-libvirt-domain-via-polkit-older-pksa-format)
[posts](https://unix.stackexchange.com/questions/367731/libvirt-debian-restrict-user-domain-access).

## Supported commands

`go-dromaius` supports the following commands:

* `start`
* `reboot`
* `shutdown`
* `destroy`
* `reset`
* `console`

## Building

* `go get github.com/creack/pty`
* `go get golang.org/x/crypto/ssh/terminal`
* `cd cmd/dromaius && go build`

(`go.mod` will follow)

## Setup

* Create a user, e.g.:
```
# adduser --system --disabled-password --shell /bin/bash --ingroup nogroup --home /var/lib/dromaius dromaius
```
* Prepare the `.ssh/` directory:
```
# mkdir -p /var/lib/dromaius/.ssh
# touch /var/lib/dromaius/.ssh/authorized_keys
# chown dromaius:nogroup /var/lib/dromaius/.ssh/
# chmod 700 /var/lib/dromaius/.ssh/
# chmod 600 /var/lib/dromaius/.ssh/authorized_keys
```
* By default (in Debian), the `libvirt-sock` UNIX socket has mode 777, allowing read/write access to all. If your configuration is different, make sure the `dromaius` user is in the correct group, and the socket has the correct user/group/mode settings
* Per supplied SSH key, add the following line to `/var/lib/dromaius/.ssh/authorized_keys`:
```
command="/var/lib/dromaius/dromaius -cmd interactive -host <full hostname of the VM>",no-agent-forwarding,no-port-forwarding,no-X11-forwarding <ssh-key>
```
* Set the allowed hosts in `/var/lib/dromaius/dromaius.json` (see `dromaius.json.example`)
* Let the user use `ssh dromaius@vm-host` to access the interactive console.

# tlsdate
go implementation of tlsdate from ioerror, see https://github.com/ioerror/tlsdate

I copied the crypto/tls package from golang and modified it in order to add needed functionality.
I added the following lines:

* https://github.com/scusi/tlsdate/blob/master/tls/handshake_client.go#L17
* https://github.com/scusi/tlsdate/blob/master/tls/handshake_client.go#L22-L27
* https://github.com/scusi/tlsdate/blob/master/tls/handshake_client.go#L159-L187

This code is more a proof of concept than production code.
It does work fine but it is much more simplistic than the original tlsdate.
Other than the original tlsdate it does NOT:
- check system time on startup and sets it to COMPILE_TIME if older.
- check extracted time to be sane
- support extraction of HTTP Request Timestamp extraction
- can run as a deamon
- come with an init script

This code can set the clock but unlike the original it uses the system 
commands ```sudo``` and ```date``` to set the date. That also means that
you might need to enter your credentials when ```sudo``` is executed,
or need to configure ```sudo``` accordingly, so no password is required.

Also my commandline flags are different from the original, see USAGE.

## Usage

 ```tlsdate -host="posteo.de" -set=true```

sets the system-clock to the date 'posteo.de' returned in a tls session negotiation.
default port used is 443, but you can specify a port, see next example.

 ```tlsdate -host="mail.whatever.net" -port="993" -skipVerify=true```

the above command does
- just print the time, does not set the clock.
- not check the certificate (skipVerify).
- use port 993 (IMAPS) instead of default (443/HTTPS)

The 'skipVerify' option is needed e.g. if:
- remote cert is a self-signed cert.
- remote cert is signed by a CA which is not listed in tlsdate-ca-roots.conf

## Install

In case you have a working go environment on your system, 
a simple go get command will do the job.

 ```go get github.com/scusi/tlsdate```

If you have not, here is how you can: https://golang.org/doc/install.

## Commits

If you want to commit to this code feel free to send me pull requests.
I prefer lots of small commits that do change one thing rather than 
one huge commit with a dozen of changes hard to follow.

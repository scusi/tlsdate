# tlsdate
go implementation of tlsdate from ioerror, see https://github.com/ioerror/tlsdate

this code is more a proof of concept than production code.

I copied the crypto/tls package from golang and modified it in order to add needed functionality.
Changes i made are basically:

* https://github.com/scusi/tlsdate/blob/master/tls/handshake_client.go#L22-L27
* https://github.com/scusi/tlsdate/blob/master/tls/handshake_client.go#L159-L187

Also my commandline flags are different from the original.

## Usage

 ```tlsdate -host="posteo.de" -set=true```

sets the system-clock to the date 'posteo.de' returned in a tls session negotiation.
default port used is 443, but you can specify a port, see next example.

 ```tlsdate -host="mail.whatever.net" -port="993" -skipVerify=true```

prints systemtime from 'myhost.noip.net', does not set the clock.
Certificate validation is turned off (skipVerify).
Uses IMAPS (993) instead of HTTPS (443) port.

The 'skipVerify' option is needed e.g. if:
- remote cert is a self-signed cert.
- remote cert is signed by a CA which is not listed in tlsdate-ca-roots.conf

## Install

In case you have a working go environment on your system, 
a simple go get command will do the job.

 ```go get github.com/scusi/tlsdate```

If you have not, here is how you can: https://golang.org/doc/install.

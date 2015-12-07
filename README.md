# tlsdate
go implementation of tlsdate from ioerror

original tlsdate from ioerror is at https://github.com/ioerror/tlsdate

this code is more a proof of concept than production code.

i copied the crypto/tls package from golang and modified it in order to add needed functionality.

Also my commandline flags are different from the original.

 tlsdate -host="posteo.de" -port=443 -set=true

 tlsdate -host="myhost.noip.net" -port="443" -skipVerify=true


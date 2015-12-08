// tlsdate re-implements tlsdate from jacob applebaum in go
//
// tlsdate is based on the fact that TLS versions before TLS 1.2 did
// encode the time into the first 32 bits of the hello message random field.
//
// My hacked version of the tls module just reads out those 4 byte
// from the servers helo message random field and
// writes them into the exported Variable 'tls.ServerHeloTime'.
//
// I also introduced 'tls.TLSDATE_VERBOSE', if set to true verbose
// log messages will be printed.
//
// The magic is in tls/handshake_client.go line 155 and following.
//
// USAGE:
//
//  tlsdate -host="posteo.de" -set=true
//
// above command will set the local system time to the given time from "posteo.de:443" tls connection.
//
//  tlsdate -host="mail.whatever.net" -port="993" -skipVerify=true
//
// above command would skip certificate checking, and uses port 993 (IMAPS) instead of the default
//
package main

import (
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	// a hacked version of the golang crypto/tls package
	"github.com/scusi/tlsdate/tls"
	"os/exec"
	"time"
)

var remoteHost string
var remotePort string
var verbose bool
var set bool
var skipVerify bool

func init() {
	flag.StringVar(&remoteHost, "host", "", "host to connect to")
	flag.StringVar(&remotePort, "port", "443", "port to connect to, on remote host")
	flag.BoolVar(&verbose, "verbose", false, "verbose mode")
	flag.BoolVar(&set, "set", false, "sets system time to tls time, when true")
	flag.BoolVar(&skipVerify, "skipVerify", false, "skips certificate verification, when true")
}

func main() {
	flag.Parse()
	tls.TLSDATE_VERBOSE = verbose
	if remoteHost == "" {
		err := fmt.Errorf("No remoteHost given, use -host switch.\n")
		log.Fatal(err)
	}
	if remotePort == "" {
		err := fmt.Errorf("No remotePort given, use -port switch.\n")
		log.Fatal(err)
	}
	hostPort := net.JoinHostPort(remoteHost, remotePort)
	// setup tls config
	conf := new(tls.Config)
	// set tls versions correctly in order to get time info
	conf.MinVersion = tls.VersionSSL30
	conf.MaxVersion = tls.VersionTLS11

	// setup a pool of trusted CA certificates
	// as a basis we use original tlsdate data, tlsdate-ca-roots.conf
	// from https://raw.githubusercontent.com/ioerror/tlsdate/master/ca-roots/tlsdate-ca-roots.conf
	//
	rootPEM, err := ioutil.ReadFile("tlsdate-ca-roots.conf")
	if err != nil {
		log.Fatal(err)
	}
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
		panic("failed to parse root certificate")
	}
	// add cert pool to tls config
	conf.RootCAs = roots

	// if skipVerify is set turn off cert verification
	if skipVerify == true {
		conf.InsecureSkipVerify = true
	}

	// initiate a tls connection to remoteHost
	conn, err := tls.Dial("tcp", hostPort, conf)
	if err != nil {
		log.Fatal(err)
	}
	// after we have a connection tls.ServerHeloTime should be set
	// TODO: sanity checking of provided input
	t := int64ToTime(tls.ServerHeloTime)
	// print the time
	fmt.Printf("TLS extracted timestamp: %s\n", t)
	// if variable 'set' is true, we set system time
	if set == true {
		err := setTime(t)
		if err != nil {
			log.Println(err)
			log.Printf("time was not set, due to an error\n")
		}
		if verbose == true {
			log.Printf("Systemtime set to %s\n", t)
		}
	}
	conn.Close()
}

func int64ToTime(i int64) (t *time.Time) {
	ts := time.Unix(i, 0)
	t = &ts
	return t
}

// setTime sets the system time to given time.
// uses 'sudo date mmddHHMMccyy.ss'
func setTime(t *time.Time) (err error) {
	// date mmddHHMMccyy.ss
	dateA := t.Format("010215042006.05")
	cmd := exec.Command("sudo", "date", dateA)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// +build windows

package main

const (
	rcFname     = "rc.conf"
	directFname = "direct.conf"
	proxyFname  = "proxy.conf"
	rejectFname = "reject.conf"
	// CNIPFname is china ip list
	CNIPFname = "china_ip_list.conf"

	newLine = "\r\n"
)

func getDefaultRcFile() string {
	// put the configuration file in the same directory of meow executable
	// This is not a reliable way to detect binary directory, but it works for double click and run
	// return path.Join(path.Dir(os.Args[0]), rcFname)

	return rcFname
}

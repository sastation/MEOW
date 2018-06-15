// +build !windows

package main

const (
	rcFname     = "rc.conf"
	directFname = "direct.conf"
	proxyFname  = "proxy.conf"
	rejectFname = "reject.conf"
	CNIPFname   = "china_ip_list.conf"

	newLine = "\n"
)

func getDefaultRcFile() string {
	//return path.Join(path.Join(getUserHomeDir(), ".meow", rcFname))

	return rcFname
}

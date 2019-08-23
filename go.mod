module upward_ssh

require (
	github.com/howeyc/gopass v0.0.0-20170109162249-bf9dde6d0d2c
	github.com/pkg/errors v0.8.1
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190605123033-f99c8df09eb5
	golang.org/x/net => github.com/golang/net v0.0.0-20190606173856-1492cefac77f
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190606203320-7fc4e5ec1444
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190706070813-72ffa07ba3db
)

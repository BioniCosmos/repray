package cert

import _ "unsafe"

// from: https://github.com/XTLS/Xray-core/blob/f67167bb3bbfcb0f45393f37005a51024e58864b/transport/internet/tls/unsafe.go
//
//go:linkname errNoCertificates crypto/tls.errNoCertificates
var errNoCertificates error

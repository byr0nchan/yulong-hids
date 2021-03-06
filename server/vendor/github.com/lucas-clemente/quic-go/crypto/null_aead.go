package crypto

import "github.com/lucas-clemente/quic-go/internal/protocol"

// NewNullAEAD creates a NullAEAD
func NewNullAEAD(p protocol.Perspective, v protocol.VersionNumber) AEAD {
	if v == protocol.VersionTLS {
		return &nullAEADFNV64a{}
	}
	return &nullAEADFNV128a{
		perspective: p,
		version:     v,
	}
}

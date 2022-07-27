// Copyright 2016-2017 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sockmap

import (
	"fmt"
	"net"
	"unsafe"

	"github.com/cilium/cilium/common/types"
	"github.com/cilium/cilium/pkg/bpf"
	"github.com/cilium/cilium/pkg/byteorder"
	"github.com/cilium/cilium/pkg/logging"
)

const (
	SockmapKeyIPv4 uint8 = 1
	SockmapKeyIPv6 uint8 = 2
)

type SockmapKey struct {
	SIP    types.IPv6
	DIP    types.IPv6
	Family uint8
	Pad0   uint8
	Pad1   uint16
	SPort  uint32
	DPort  uint32
}

type SockmapValue struct {
	fd uint32
}

func (k SockmapKey) ToIP() (net.IP, net.IP) { //(net.IP, net.IP) {
	switch k.Family {
	case SockmapKeyIPv4:
		return k.SIP[:4], k.DIP[:4]
	case SockmapKeyIPv6:
		return k.SIP[:], k.DIP[:]
	}
	return nil, nil
}

func (key SockmapKey) String() string {
	dip, sip := key.ToIP()

	return fmt.Sprintf("%s:%d->%s:%d",
		sip.String(), byteorder.HostToNetwork(key.SPort),
		dip.String(), byteorder.HostToNetwork(key.DPort))
}

func (v SockmapValue) String() string {
	return fmt.Sprintf("%d", v.fd)
}

// GetValuePtr returns the unsafe pointer to the BPF value.
func (v SockmapValue) GetValuePtr() unsafe.Pointer { return unsafe.Pointer(&v) }

// GetKeyPtr returns the unsafe pointer to the BPF key
func (k SockmapKey) GetKeyPtr() unsafe.Pointer { return unsafe.Pointer(&k) }

// NewValue returns a new empty instance of the structure representing the BPF
// map value
func (k SockmapKey) NewValue() bpf.MapValue { return &SockmapValue{} }

func NewSockmapKey(dip net.IP, sip net.IP, sport uint32, dport uint32) SockmapKey {
	result := SockmapKey{}

	if sip4 := sip.To4(); sip4 != nil {
		result.Family = bpf.EndpointKeyIPv4
		copy(result.SIP[:], sip4)
	} else {
		result.Family = bpf.EndpointKeyIPv6
		copy(result.SIP[:], sip)
	}

	if dip4 := dip.To4(); dip4 != nil {
		result.Family = bpf.EndpointKeyIPv4 // throw error
		copy(result.SIP[:], dip4)
	} else {
		result.Family = bpf.EndpointKeyIPv6
		copy(result.DIP[:], dip)
	}

	result.DPort = dport
	result.SPort = sport
	return result
}

var log = logging.DefaultLogger

const (
	MapName = "sock_ops_map"

	// MaxEntries represents the maximum number of endpoints in the map
	MaxEntries = 65535
)

var (
	// Sockmap represents the BPF map for sockets
	Sockmap = bpf.NewMap(MapName,
		bpf.MapTypeSockHash,
		int(unsafe.Sizeof(SockmapKey{})),
		4,
		MaxEntries,
		0,
		func(key []byte, value []byte) (bpf.MapKey, bpf.MapValue, error) {
			k, v := SockmapKey{}, SockmapValue{}

			if err := bpf.ConvertKeyValue(key, value, &k, &v); err != nil {
				return nil, nil, err
			}

			return k, v, nil
		},
		func(key []byte) (bpf.MapKey, error) {
			k := SockmapKey{}
			if err := bpf.ConvertKey(key, &k); err != nil {
				return nil, err
			}

			return k, nil
		},
	)
)

func init() {
	bpf.OpenAfterMount(Sockmap)
}

// EndpointFrontend is the interface to implement for an object to synchronize
// with the endpoint BPF map
type EndpointFrontend interface {
	// GetBPFKeys must return a slice of SockmapKey which all represent the endpoint
	GetBPFKeys() []SockmapKey
}

// DeleteEntry deletes a single map entry
func DeleteEntry(dip net.IP, sip net.IP, sport uint32, dport uint32) error {
	return Sockmap.Delete(NewSockmapKey(dip, sip, sport, dport))
}

// DeleteElement deletes the endpoint using all keys which represent the
// endpoint. It returns the number of errors encountered during deletion.
func DeleteElement(f EndpointFrontend) []error {
	errors := []error{}
	for _, k := range f.GetBPFKeys() {
		if err := Sockmap.Delete(k); err != nil {
			errors = append(errors, fmt.Errorf("Unable to delete key %v in endpoint BPF map: %s", k, err))
		}
	}

	return errors
}

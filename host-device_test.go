// Copyright 2017 CNI authors
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

package main

import (
	"fmt"

	"github.com/containernetworking/cni/pkg/ns"
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/testutils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vishvananda/netlink"
)

var _ = Describe("base functionality", func() {
	var originalNS ns.NetNS

	BeforeEach(func() {
		var err error
		originalNS, err = ns.NewNS()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		originalNS.Close()
	})

	It("Works with a valid config", func() {
		err := originalNS.Do(func(ns.NetNS) error {
			defer GinkgoRecover()

			err := netlink.LinkAdd(&netlink.Dummy{
				LinkAttrs: netlink.LinkAttrs{
					Name: "dummy0",
				},
			})
			Expect(err).NotTo(HaveOccurred())

			link, err := netlink.LinkByName("dummy0")
			Expect(err).NotTo(HaveOccurred())

			err = netlink.LinkSetUp(link)
			Expect(err).NotTo(HaveOccurred())

			return nil
		})
		Expect(err).NotTo(HaveOccurred())
		ifname := "eth0"
		conf := `{
			"name": "cni-plugin-host-device-test",
			"type": "host-device",
			"device": "dummy0"
		}`
		conf = fmt.Sprintf(conf, ifname, originalNS.Path())
		args := &skel.CmdArgs{
			ContainerID: "dummy",
			Netns:       originalNS.Path(),
			IfName:      ifname,
			StdinData:   []byte(conf),
		}
		_, _, err = testutils.CmdAddWithResult(originalNS.Path(), "dummy0", []byte(conf), func() error { return cmdAdd(args) })
		Expect(err).NotTo(HaveOccurred())

	})

	It("fails an invalid config", func() {
		conf := `{
	"cniVersion": "0.3.0",
	"name": "cni-plugin-sample-test",
	"type": "host-device"
}`

		args := &skel.CmdArgs{
			ContainerID: "dummy",
			Netns:       originalNS.Path(),
			IfName:      "eth0",
			StdinData:   []byte(conf),
		}
		_, _, err := testutils.CmdAddWithResult(originalNS.Path(), "eth0", []byte(conf), func() error { return cmdAdd(args) })
		Expect(err).To(MatchError("anotherAwesomeArg must be specified"))

	})

})

/*
 * Copyright 2020 The Multicluster-Scheduler Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package node

import (
	"flag"
)

const (
	DefaultNodeName        = "admiralty"
	DefaultEnableNodeLease = false
)

type Opts struct {
	NodeName        string
	EnableNodeLease bool
}

func (c *Opts) BindFlags() {
	flag.StringVar(&c.NodeName, "nodename", DefaultNodeName, "")
	flag.BoolVar(&c.EnableNodeLease, "enable-node-lease", DefaultEnableNodeLease, "use node leases (1.13) for node heartbeats")
}

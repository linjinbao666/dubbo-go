/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package kubernetes

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/config_center"
	"dubbo.apache.org/dubbo-go/v3/remoting"
)

func Test_DataChange(t *testing.T) {
	listener := NewRegistryDataListener(&MockDataListener{})
	url, _ := common.NewURL("tri%3A%2F%2F10.244.0.24%3A20001%2Fgrpc.reflection.v1alpha.ServerReflection%3Fanyhost%3Dtrue%26app.version%3D3.0.0%26application%3Ddubbo.io%26bean.name%3DXXX_serverReflectionServer%26cluster%3Dfailover%26environment%3Ddev%26export%3Dtrue%26interface%3Dgrpc.reflection.v1alpha.ServerReflection%26loadbalance%3Drandom%26message_size%3D4%26metadata-type%3Dlocal%26methods%3DServerReflectionInfo%26module%3Dsample%26name%3Ddubbo.io%26organization%3Ddubbo-go%26owner%3Ddubbo-go%26pid%3D13%26registry%3Dkubernetes%26registry.role%3D3%26release%3Ddubbo-golang-3.0.0%26service.filter%3Decho%2Cmetrics%2Ctoken%2Caccesslog%2Ctps%2Cgeneric_service%2Cexecute%2Cpshutdown%26side%3Dprovider%26timestamp%3D1648000702")
	listener.AddInterestedURL(url)
	int := listener.DataChange(remoting.Event{Path: "/dubbo/grpc.reflection.v1alpha.ServerReflection/providers/tri%3A%2F%2F10.244.0.18%3A20001%2Fgrpc.reflection.v1alpha.ServerReflection%3Fanyhost%3Dtrue%26app.version%3D3.0.0%26application%3Ddubbo.io%26bean.name%3DXXX_serverReflectionServer%26cluster%3Dfailover%26environment%3Ddev%26export%3Dtrue%26interface%3Dgrpc.reflection.v1alpha.ServerReflection%26loadbalance%3Drandom%26message_size%3D4%26metadata-type%3Dlocal%26methods%3DServerReflectionInfo%26module%3Dsample%26name%3Ddubbo.io%26organization%3Ddubbo-go%26owner%3Ddubbo-go%26pid%3D15%26registry%3Dkubernetes%26registry.role%3D3%26release%3Ddubbo-golang-3.0.0%26service.filter%3Decho%2Cmetrics%2Ctoken%2Caccesslog%2Ctps%2Cgeneric_service%2Cexecute%2Cpshutdown%26side%3Dprovider%26timestamp%3D1647958533"})

	assert.Equal(t, true, int)
}

type MockDataListener struct{}

func (*MockDataListener) Process(configType *config_center.ConfigChangeEvent) {}

func TestDataChange(t *testing.T) {

	listener := NewRegistryDataListener(&MockDataListener{})
	url, _ := common.NewURL("tri%3A%2F%2F10.244.0.24%3A20001%2Fgrpc.reflection.v1alpha.ServerReflection%3Fanyhost%3Dtrue%26app.version%3D3.0.0%26application%3Ddubbo.io%26bean.name%3DXXX_serverReflectionServer%26cluster%3Dfailover%26environment%3Ddev%26export%3Dtrue%26interface%3Dgrpc.reflection.v1alpha.ServerReflection%26loadbalance%3Drandom%26message_size%3D4%26metadata-type%3Dlocal%26methods%3DServerReflectionInfo%26module%3Dsample%26name%3Ddubbo.io%26organization%3Ddubbo-go%26owner%3Ddubbo-go%26pid%3D13%26registry%3Dkubernetes%26registry.role%3D3%26release%3Ddubbo-golang-3.0.0%26service.filter%3Decho%2Cmetrics%2Ctoken%2Caccesslog%2Ctps%2Cgeneric_service%2Cexecute%2Cpshutdown%26side%3Dprovider%26timestamp%3D1648000702")
	listener.AddInterestedURL(url)
	if !listener.DataChange(remoting.Event{Path: "/dubbo/grpc.reflection.v1alpha.ServerReflection/providers/tri%3A%2F%2F10.244.0.18%3A20001%2Fgrpc.reflection.v1alpha.ServerReflection%3Fanyhost%3Dtrue%26app.version%3D3.0.0%26application%3Ddubbo.io%26bean.name%3DXXX_serverReflectionServer%26cluster%3Dfailover%26environment%3Ddev%26export%3Dtrue%26interface%3Dgrpc.reflection.v1alpha.ServerReflection%26loadbalance%3Drandom%26message_size%3D4%26metadata-type%3Dlocal%26methods%3DServerReflectionInfo%26module%3Dsample%26name%3Ddubbo.io%26organization%3Ddubbo-go%26owner%3Ddubbo-go%26pid%3D15%26registry%3Dkubernetes%26registry.role%3D3%26release%3Ddubbo-golang-3.0.0%26service.filter%3Decho%2Cmetrics%2Ctoken%2Caccesslog%2Ctps%2Cgeneric_service%2Cexecute%2Cpshutdown%26side%3Dprovider%26timestamp%3D1647958533"}) {
		t.Fatal("data change not ok")
	}
}

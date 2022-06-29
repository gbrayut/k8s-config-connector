// Copyright 2022 Google LLC. All Rights Reserved.
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
// GENERATED BY gen_go_data.go
// gen_go_data -package beta -var YAML_hub blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networkconnectivity/beta/hub.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networkconnectivity/beta/hub.yaml
var YAML_hub = []byte("info:\n  title: NetworkConnectivity/Hub\n  description: The NetworkConnectivity Hub resource\n  x-dcl-struct-name: Hub\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Hub\n    parameters:\n    - name: Hub\n      required: true\n      description: A full instance of a Hub\n  apply:\n    description: The function used to apply information about a Hub\n    parameters:\n    - name: Hub\n      required: true\n      description: A full instance of a Hub\n  delete:\n    description: The function used to delete a Hub\n    parameters:\n    - name: Hub\n      required: true\n      description: A full instance of a Hub\n  deleteAll:\n    description: The function used to delete all Hub\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Hub\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Hub:\n      title: Hub\n      x-dcl-id: projects/{{project}}/locations/global/hubs/{{name}}\n      x-dcl-locations:\n      - global\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - project\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The time the hub was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: An optional description of the hub.\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Optional labels in key:value format. For more information about\n            labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'Immutable. The name of the hub. Hub names must be unique.\n            They use the following form: `projects/{project_number}/locations/global/hubs/{hub_id}`'\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        routingVpcs:\n          type: array\n          x-dcl-go-name: RoutingVpcs\n          readOnly: true\n          description: The VPC network associated with this hub's spokes. All of the\n            VPN tunnels, VLAN attachments, and router appliance instances referenced\n            by this hub's spokes must belong to this VPC network. This field is read-only.\n            Network Connectivity Center automatically populates it based on the set\n            of spokes attached to the hub.\n          x-kubernetes-immutable: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: HubRoutingVpcs\n            properties:\n              uri:\n                type: string\n                x-dcl-go-name: Uri\n                description: The URI of the VPC network.\n                x-dcl-references:\n                - resource: Compute/Network\n                  field: selfLink\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: HubStateEnum\n          readOnly: true\n          description: 'Output only. The current lifecycle state of this hub. Possible\n            values: STATE_UNSPECIFIED, CREATING, ACTIVE, DELETING'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - CREATING\n          - ACTIVE\n          - DELETING\n        uniqueId:\n          type: string\n          x-dcl-go-name: UniqueId\n          readOnly: true\n          description: Output only. The Google-generated UUID for the hub. This value\n            is unique across all hub resources. If a hub is deleted and another with\n            the same name is created, the new hub is assigned a different unique_id.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The time the hub was last updated.\n          x-kubernetes-immutable: true\n")

// 4637 bytes
// MD5: 3c0f40874a63ad73778f7e98e33512ba

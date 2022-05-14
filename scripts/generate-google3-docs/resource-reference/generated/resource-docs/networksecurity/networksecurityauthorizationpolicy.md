{# AUTOGENERATED. DO NOT EDIT. #}

{% extends "config-connector/_base.html" %}

{% block page_title %}NetworkSecurityAuthorizationPolicy{% endblock %}
{% block body %}


<table>
<thead>
<tr>
<th><strong>Property</strong></th>
<th><strong>Value</strong></th>
</tr>
</thead>
<tbody>
<tr>
<td>{{gcp_name_short}} Service Name</td>
<td>Network Security</td>
</tr>
<tr>
<td>{{gcp_name_short}} Service Documentation</td>
<td><a href="/traffic-director/docs/">/traffic-director/docs/</a></td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Name</td>
<td>v1beta1/projects.locations.authorizationPolicies</td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Documentation</td>
<td><a href="/traffic-director/docs/reference/network-security/rest/v1beta1/projects.locations.authorizationPolicies">/traffic-director/docs/reference/network-security/rest/v1beta1/projects.locations.authorizationPolicies</a></td>
</tr>
<tr>
<td>{{product_name_short}} Resource Short Names</td>
<td>gcpnetworksecurityauthorizationpolicy<br>gcpnetworksecurityauthorizationpolicies<br>networksecurityauthorizationpolicy</td>
</tr>
<tr>
<td>{{product_name_short}} Service Name</td>
<td>networksecurity.googleapis.com</td>
</tr>
<tr>
<td>{{product_name_short}} Resource Fully Qualified Name</td>
<td>networksecurityauthorizationpolicies.networksecurity.cnrm.cloud.google.com</td>
</tr>

<tr>
    <td>Can Be Referenced by IAMPolicy/IAMPolicyMember</td>
    <td>Yes</td>
</tr>

<tr>
    <td>IAMPolicy/IAMPartialPolicy Supports Conditions</td>
    <td>Yes</td>
</tr>
<tr>
    <td>IAMPolicyMember Supports Conditions</td>
    <td>No</td>
</tr>

<tr>
    <td>Supports IAM Audit Configs</td>
    <td>No</td>
</tr>
<tr>
    <td>IAM External Reference Format</td>
    <td>
        
        <p>{% verbatim %}projects/{{project}}/locations/{{location}}/authorizationPolicies/{{name}}{% endverbatim %}</p>
        
    </td>
</tr>


</tbody>
</table>

## Custom Resource Definition Properties



### Spec
#### Schema
  ```yaml
  action: string
  description: string
  location: string
  projectRef:
    external: string
    name: string
    namespace: string
  resourceID: string
  rules:
  - destinations:
    - hosts:
      - string
      httpHeaderMatch:
        headerName: string
        regexMatch: string
      methods:
      - string
      ports:
      - integer
    sources:
    - ipBlocks:
      - string
      principals:
      - string
  ```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td>
            <p><code>action</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Required. The action to take when a rule match is found. Possible values are "ALLOW" or "DENY". Possible values: ACTION_UNSPECIFIED, ALLOW, DENY{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>description</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Optional. Free-text description of the resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>location</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. The location for the resource{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Immutable. The Project that this resource belongs to.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The project for the resource

Allowed value: The Google Cloud resource name of a `Project` resource (format: `projects/{{name}}`).{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>resourceID</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rules</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Optional. List of rules to match. If not set, the action specified in the ‘action’ field will be applied without any additional rule checks.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rules[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rules[].destinations</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Optional. List of attributes for the traffic destination. If not set, the action specified in the ‘action’ field will be applied without any rule checks for the destination.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rules[].destinations[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rules[].destinations[].hosts</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Required. List of host names to match. Matched against HOST header in http requests. Each host can be an exact match, or a prefix match (example, “mydomain.*”) or a suffix match (example, *.myorg.com”) or a presence(any) match “*”.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rules[].destinations[].hosts[]</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rules[].destinations[].httpHeaderMatch</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Optional. Match against key:value pair in http header. Provides a flexible match based on HTTP headers, for potentially advanced use cases.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rules[].destinations[].httpHeaderMatch.headerName</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Required. The name of the HTTP header to match. For matching against the HTTP request's authority, use a headerMatch with the header name ":authority". For matching a request's method, use the headerName ":method".{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rules[].destinations[].httpHeaderMatch.regexMatch</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Required. The value of the header must match the regular expression specified in regexMatch. For regular expression grammar, please see: en.cppreference.com/w/cpp/regex/ecmascript For matching against a port specified in the HTTP request, use a headerMatch with headerName set to Host and a regular expression that satisfies the RFC2616 Host header's port specifier.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rules[].destinations[].methods</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Optional. A list of HTTP methods to match. Should not be set for gRPC services.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rules[].destinations[].methods[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rules[].destinations[].ports</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">list (integer)</code></p>
            <p>{% verbatim %}Required. List of destination ports to match.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rules[].destinations[].ports[]</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rules[].sources</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Optional. List of attributes for the traffic source. If not set, the action specified in the ‘action’ field will be applied without any rule checks for the source.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rules[].sources[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rules[].sources[].ipBlocks</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Optional. List of CIDR ranges to match based on source IP address. Single IP (e.g., "1.2.3.4") and CIDR (e.g., "1.2.3.0/24") are supported.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rules[].sources[].ipBlocks[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rules[].sources[].principals</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Optional. List of peer identities to match for authorization. Each peer can be an exact match, or a prefix match (example, “namespace/*”) or a suffix match (example, */service-account”) or a presence match “*”.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rules[].sources[].principals[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>


<p>{% verbatim %}* Field is required when parent field is specified{% endverbatim %}</p>


### Status
#### Schema
  ```yaml
  conditions:
  - lastTransitionTime: string
    message: string
    reason: string
    status: string
    type: string
  createTime: string
  observedGeneration: integer
  updateTime: string
  ```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>conditions</code></td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Conditions represent the latest available observation of the resource's current state.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[]</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].lastTransitionTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Last time the condition transitioned from one status to another.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].message</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Human-readable message indicating details about last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].reason</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Unique, one-word, CamelCase reason for the condition's last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].status</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Status is the status of the condition. Can be True, False, Unknown.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].type</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Type is the type of the condition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>createTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. The timestamp when the resource was created.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedGeneration</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>updateTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. The timestamp when the resource was updated.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>

## Sample YAML(s)

### Typical Use Case
  ```yaml
  # Copyright 2021 Google LLC
  #
  # Licensed under the Apache License, Version 2.0 (the "License");
  # you may not use this file except in compliance with the License.
  # You may obtain a copy of the License at
  #
  #     http://www.apache.org/licenses/LICENSE-2.0
  #
  # Unless required by applicable law or agreed to in writing, software
  # distributed under the License is distributed on an "AS IS" BASIS,
  # WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  # See the License for the specific language governing permissions and
  # limitations under the License.
  
  apiVersion: networksecurity.cnrm.cloud.google.com/v1beta1
  kind: NetworkSecurityAuthorizationPolicy
  metadata:
    labels:
      key-one: value-one
    name: networksecurityauthorizationpolicy-sample
  spec:
    projectRef:
       # Replace ${PROJECT_ID?} with your project ID
       external: "projects/${PROJECT_ID?}"
    location: global
    action: ALLOW
    description: Test Authorization Policy
    rules:
    - sources:
      - ipBlocks:
        - "1.2.3.4"
        principals:
        - "*"
      destinations:
      - hosts:
        - "demo-service"
        ports:
        - 8080
        methods:
        - "POST"
    - sources:
      - ipBlocks:
        - "1.2.3.5"
        principals:
        - "*"
      destinations:
      - hosts:
        - "test-service"
        ports:
        - 8081
        methods:
        - "GET"
  ```


{% endblock %}
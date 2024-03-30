#!/bin/bash
# Copyright 2024 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


scriptpath=$(realpath $0)
base=$(dirname $scriptpath)

cr=$1
namespace=$2
expected=$3

tmpfile=$(mktemp expanded-output.XXXXXXXX)
kubectl get plans $cr -n $namespace -o json | jq ".spec.stages"   > $tmpfile
echo ".spec.stages:"
cat $tmpfile
diff $tmpfile $expected || echo " ***************** FAILED ********************"
rm $tmpfile
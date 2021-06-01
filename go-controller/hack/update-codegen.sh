#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

crds=$(ls pkg/crd 2> /dev/null)
if [ -z "${crds}" ]; then
  exit
fi

if  ! ( command -v controller-gen > /dev/null ); then
  echo "controller-gen not found, installing sigs.k8s.io/controller-tools"
  olddir="${PWD}"
  builddir="$(mktemp -d)"
  cd "${builddir}"
  GO111MODULE=on go get -u sigs.k8s.io/controller-tools/cmd/controller-gen
  cd "${olddir}"
  if [[ "${builddir}" == /tmp/* ]]; then #paranoia
      rm -rf "${builddir}"
  fi
fi

if  ! ( command -v deepcopy-gen > /dev/null ); then
  echo "deepcopy-gen not found, installing k8s.io/code-generator"
  olddir="${PWD}"
  builddir="$(mktemp -d)"
  cd "${builddir}"
  GO111MODULE=on go get -u k8s.io/code-generator/cmd/deepcopy-gen
  cd "${olddir}"
  if [[ "${builddir}" == /tmp/* ]]; then #paranoia
      rm -rf "${builddir}"
  fi
fi

if  ! ( command -v client-gen > /dev/null ); then
  echo "client-gen not found, installing k8s.io/code-generator"
  olddir="${PWD}"
  builddir="$(mktemp -d)"
  cd "${builddir}"
  GO111MODULE=on go get -u k8s.io/code-generator/cmd/client-gen
  cd "${olddir}"
  if [[ "${builddir}" == /tmp/* ]]; then #paranoia
      rm -rf "${builddir}"
  fi
fi

if  ! ( command -v informer-gen > /dev/null ); then
  echo "client-gen not found, installing k8s.io/code-generator"
  olddir="${PWD}"
  builddir="$(mktemp -d)"
  cd "${builddir}"
  GO111MODULE=on go get -u k8s.io/code-generator/cmd/informer-gen
  cd "${olddir}"
  if [[ "${builddir}" == /tmp/* ]]; then #paranoia
      rm -rf "${builddir}"
  fi
fi

if  ! ( command -v lister-gen > /dev/null ); then
  echo "client-gen not found, installing k8s.io/code-generator"
  olddir="${PWD}"
  builddir="$(mktemp -d)"
  cd "${builddir}"
  GO111MODULE=on go get -u k8s.io/code-generator/cmd/lister-gen
  cd "${olddir}"
  if [[ "${builddir}" == /tmp/* ]]; then #paranoia
      rm -rf "${builddir}"
  fi
fi

for crd in ${crds}; do
  echo "Generating deepcopy funcs for $crd"
  GO111MODULE=on deepcopy-gen \
    --input-dirs github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd/$crd/v1 \
    -O zz_generated.deepcopy \
    --go-header-file /dev/null \
    --bounding-dirs github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd


  echo "Generating clientset for $crd"
  GO111MODULE=on client-gen \
    --clientset-name "${CLIENTSET_NAME_VERSIONED:-versioned}" \
    --input-base "" \
    --input github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd/$crd/v1 \
    --go-header-file /dev/null \
    --output-package github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd/$crd/v1/apis/clientset \
    "$@"

  echo "Generating listers for $crd"
  GO111MODULE=on lister-gen \
    --input-dirs github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd/$crd/v1 \
    --go-header-file /dev/null \
    --output-package github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd/$crd/v1/apis/listers \
    "$@"

  echo "Generating informers for $crd"
  GO111MODULE=on informer-gen \
    --input-dirs github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd/$crd/v1 \
    --versioned-clientset-package github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd/$crd/v1/apis/clientset/versioned \
    --listers-package  github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd/$crd/v1/apis/listers \
    --go-header-file /dev/null \
    --output-package github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd/$crd/v1/apis/informers \
    "$@"
done

echo "Generating CRDs"
mkdir -p _output/crds
controller-gen crd:crdVersions="v1"  paths=./pkg/crd/... output:crd:dir=_output/crds
echo "Editing egressFirewall CRD"
## We desire that only egressFirewalls with the name "default" are accepted by the apiserver. The only
## way that we can put a pattern for validation on the name of the object which is embedded in
## metav1.ObjectMeta it is required that we add it after the generation of the CRD.
sed -i -e':begin;$!N;s/.*metadata:\n.*type: object/&\n            properties:\n              name:\n                type: string\n                pattern: ^default$/;P;D' \
	_output/crds/k8s.ovn.org_egressfirewalls.yaml
## It is also required that we restrict the number of properties on the 'to' section of the egressfirewall
## so that either 'dnsName' or 'cidrSelector is set in the crd and currently kubebuilder does not support
## adding validation to objects only to the fields
sed -i -e ':begin;$!N;s/                          type: string\n.*type: object/&\n                      minProperties: 1\n                      maxProperties: 1/;P;D' \
	_output/crds/k8s.ovn.org_egressfirewalls.yaml

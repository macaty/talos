#!/bin/bash

set -eou pipefail

export TALOS_IMG="docker.io/autonomy/talos:${TAG}"
export TMP="$(mktemp -d)"
export OSCTL="${PWD}/build/osctl-linux-amd64"
export TALOSCONFIG="${TMP}/talosconfig"
export KUBECONFIG="${TMP}/kubeconfig"


cleanup() {
	${OSCTL} cluster destroy --name integration
	rm -rf ${TMP}
}
trap cleanup EXIT

./hack/test/osctl-cluster-create.sh

exit 0

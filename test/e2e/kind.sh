#!/usr/bin/env bash
set -euo pipefail

source test/e2e/aliases.sh

kind_setup() {
  i=$1

  CLUSTER=cluster$i

  if ! kind get clusters | grep $CLUSTER; then
    kind create cluster --name $CLUSTER --wait 5m
  fi
  kind get kubeconfig --name $CLUSTER --internal >kubeconfig-$CLUSTER
  k $i apply -f test/e2e/must-run-as-non-root.yaml
}

if [[ "${BASH_SOURCE[0]:-}" == "${0}" ]]; then
  kind_setup $1
fi

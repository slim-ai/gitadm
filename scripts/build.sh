#!/usr/bin/env bash

set -euo pipefail

die () {
    echo >&2 "$@"
    exit 1
}

APPNAME=$1
TARGET_OS=$2
TARGET_ARCH=$3
[ -z "${APPNAME:-}" ] && { die "parameter 1, APPNAME is not defined"; }
[ -z "${TARGET_OS:-}" ] && { die "parameter 2, TARGET_OS is not defined"; }
[ -z "${TARGET_ARCH:-}" ] && { die "parameter 3, TARGET_ARCH is not defined"; }



BUILD_TIME="$(date -u '+%Y-%m-%d_%I:%M:%S%p')"
TAG="current"
REVISION="current"

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
SDIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
BDIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

if hash git 2>/dev/null && [ -e ${BDIR}/.git ]; then
  TAG="$(git describe --tags 2>/dev/null || true)"
  [[ -z "$TAG" ]] && TAG="notag"
  REVISION="$(git rev-parse --short HEAD)"
fi

LD_FLAGS="-s -w -X github.com/slim-ai/gitadm/pkg/build.Time=${BUILD_TIME} -X github.com/slim-ai/gitadm/pkg/build.Rev=${REVISION} -X github.com/slim-ai/gitadm/pkg/build.Tag=${TAG}"


mkdir -p ${BDIR}/bin/
(
  env GOOS=${TARGET_OS} GOARCH=${TARGET_ARCH} CGO_ENABLED=0 go mod vendor
  cd ${BDIR}
  case "${TARGET_OS}" in
      darwin*)
        machine=mac
        ;;
      *)
        machine="${TARGET_OS}"
        ;;
  esac
  echo -n "cross-compiling ${BDIR}/bin/${APPNAME}-${machine}..."
  env GOOS=${TARGET_OS} GOARCH=${TARGET_ARCH} CGO_ENABLED=0 go build -mod=mod -trimpath -ldflags="${LD_FLAGS}" ${GO_REBUILD:-"-a"} -tags 'netgo osusergo' -o "${BDIR}/bin/${APPNAME}-${machine}"
  echo  "✅"
)

#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
if [ "X$1" != "X" ]; then
    RUNTIME_ROOT=$1
else
    RUNTIME_ROOT=${CURDIR}
fi

if [ "X$RUNTIME_ROOT" == "X" ]; then
    echo "There is no RUNTIME_ROOT support."
    echo "Usage: ./bootstrap.sh $RUNTIME_ROOT"
    exit -1
fi

PORT=$2

RUNTIME_CONF_ROOT=$RUNTIME_ROOT/conf
RUNTIME_LOG_ROOT=$RUNTIME_ROOT/log

if [ ! -d $RUNTIME_LOG_ROOT/app ]; then
    mkdir -p $RUNTIME_LOG_ROOT/app
fi

if [ ! -d $RUNTIME_LOG_ROOT/rpc ]; then
    mkdir -p $RUNTIME_LOG_ROOT/rpc
fi

if [ ! -f $CURDIR/settings.py ]; then
    echo "there is no settings.py exist."
    exit -1
fi

PRODUCT=$(cd $CURDIR; python -c "import settings; print (settings.PRODUCT)")
SUBSYS=$(cd $CURDIR; python -c "import settings; print (settings.SUBSYS)")
MODULE=$(cd $CURDIR; python -c "import settings; print (settings.MODULE)")

if [ -z "$PRODUCT" ] || [ -z "$SUBSYS" ] || [ -z "$MODULE" ]; then
    echo "Support PRODUCT SUBSYS MODULE PORT in settings.py"
    exit -1
fi

if [ "$IS_HOST_NETWORK" == "1" ]; then
    export RUNTIME_SERVICE_PORT=$PORT0
    export RUNTIME_DEBUG_PORT=$PORT1
fi

SVC_NAME=${PRODUCT}.${SUBSYS}.${MODULE}

BinaryName=${PRODUCT}.${SUBSYS}.${MODULE}

echo "$CURDIR/bin/${BinaryName} $args"

exec $CURDIR/bin/${BinaryName} $args

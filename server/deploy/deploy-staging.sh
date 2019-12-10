#!/bin/sh -x

productionSlot=$(helm get values cyan-api-server | awk '/productionSlot/{print $2}')
if [ $productionSlot = "blue" ]; then
    stagingSlot="green"
else
    stagingSlot="blue"
fi
helm upgrade cyan-api-server server --set $stagingSlot.enabled=true,$stagingSlot.tag=175 --reuse-values
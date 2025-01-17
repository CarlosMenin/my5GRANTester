#!/usr/bin/env bash

set -euo pipefail

CONFIG_DIR="/my5G-RANTester/config"

DEREG_AFTER=${DEREG_AFTER:-3600}

for c in ${CONFIG_DIR}/config.yml; do
    # grep variable names (format: ${VAR}) from template to be rendered
    VARS=$(grep -oP '@[a-zA-Z0-9_]+@' ${c} | sort | uniq | xargs)
    echo "Now setting these variables '${VARS}'"

    # create sed expressions for substituting each occurrence of ${VAR}
    # with the value of the environment variable "VAR"
    EXPRESSIONS=""
    for v in ${VARS}; do
        NEW_VAR=$(echo $v | sed -e "s#@##g")
        if [[ -z ${!NEW_VAR+x} ]]; then
            echo "Error: Environment variable '${NEW_VAR}' is not set." \
                "Config file '$(basename $c)' requires all of $VARS."
            exit 1
        fi
        EXPRESSIONS="${EXPRESSIONS};s|${v}|${!NEW_VAR}|g"
    done
    EXPRESSIONS="${EXPRESSIONS#';'}"

    # render template and inline replace config file
    sed -i "${EXPRESSIONS}" ${c}
done
echo "Done setting the configuration"
echo "Running tester"

TEST=${TEST:-""}
NUM_UE=${NUM_UE:-1}
DELAY=${DELAY:-1}
TIME=${TIME:-1}
INTERVAL=${INTERVAL:-1}
CONSTANT=${CONSTANT:-1}

if [ "$TEST" == "parallel" ]; then
    ./app load-test-$TEST -n $NUM_UE -d $DELAY -t $TIME -a
elif [ "$TEST" == "division" ]; then
    ./app load-test-$TEST -n $NUM_UE -d $DELAY -t $TIME -u $INTERVAL -i $CONSTANT -a
elif [ "$TEST" == "decrement" ]; then
    ./app load-test-$TEST -n $NUM_UE -d $DELAY -t $TIME -u $INTERVAL -i $CONSTANT -a
else
    echo "Your test name is incorrect"
fi

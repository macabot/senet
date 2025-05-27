#!/bin/bash

set -euo pipefail

list() {
    curl -sS "https://senet.metered.live/api/v2/turn/credentials?secretKey=${METERED_SECRET_KEY}" | jq
}

create() {
    local unixTimestamp
    unixTimestamp=$(date +%s)
    curl \
        -sS \
        --request POST \
        --header "Content-Type: application/json" \
        --data "{\"label\": \"${unixTimestamp}\"}" \
        "https://senet.metered.live/api/v1/turn/credential?secretKey=${METERED_SECRET_KEY}" | jq
}

deleteByLabel() {
    local label="$1"
    echo "$label"
    curl \
        -sS \
        --request DELETE \
        "https://senet.metered.live/api/v2/turn/credential/by_label?secretKey=${METERED_SECRET_KEY}&label=${label}"
    echo ""
}

deleteOld() {
    list | jq -c '
        (
            [.data[] | select(has("label") and (.label | test("^[0-9]+$"))) | .label | tonumber] | max
        ) as $maxLabel |

        .data[] | select(
            has("label") and
            (.label | test("^[0-9]+$")) and
            ((.label | tonumber) < $maxLabel)
        )
    ' | while read -r credential; do
        deleteByLabel "$(jq -r '.label' <<<"$credential")"
    done
}

root_dir=$(dirname "$0")
if [[ -f "${root_dir}/.env" ]]; then
    set -a
    # shellcheck disable=SC1091
    source "${root_dir}/.env"
    set +a
else
    echo "WARN: ${root_dir}/.env not found." >&2
fi

action="${1:-}"
case $action in
list)
    list
    ;;
create)
    create
    ;;
deleteOld)
    deleteOld
    ;;
*)
    echo "Invalid action." >&2
    exit 1
    ;;
esac

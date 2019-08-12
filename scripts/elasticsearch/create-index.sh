#!/usr/bin/env bash

URI=$1
INDEX_NAME=$2
VERSION=$3
MAPPING_FILE=$4
CREATE_ALIAS=$5

[[ ${URI} == '' ]] || [[ ${INDEX_NAME} == '' ]] || [[ ${VERSION} == '' ]] || [[ ${MAPPING_FILE} == '' ]] && echo 'URI INDEX_NAME VERSION MAPPING_FILE CREATE_ALIAS(0|1) required' && exit 1

INDEX_WITH_VERSION=${INDEX_NAME}-v${VERSION}

echo "Mapping >> ${INDEX}"

curl -X PUT "${URI}/${INDEX_WITH_VERSION}?pretty" -H 'Content-Type: application/json' -d "@${MAPPING_FILE}"

[[ ${CREATE_ALIAS} == '' ]] && echo 'Create new index done.' && exit 0

curl -X POST "${URI}/_aliases?pretty" -H 'Content-Type: application/json' -d "
{
  \"actions\": [
    {
      \"add\": {
        \"index\": \"${INDEX_WITH_VERSION}\",
        \"alias\": \"${INDEX_NAME}\"
      }
    }
  ]
}
"
echo "Create alias for ${INDEX_WITH_VERSION}"

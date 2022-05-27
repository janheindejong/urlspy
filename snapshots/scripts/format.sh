#!/bin/bash

set -ex

export PREFIX=""
export SOURCE_FILES="app"

${PREFIX}autoflake --in-place --recursive $SOURCE_FILES
${PREFIX}isort $SOURCE_FILES
${PREFIX}black $SOURCE_FILES

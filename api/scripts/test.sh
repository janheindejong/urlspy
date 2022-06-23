#!/bin/bash

set -ex

export PREFIX=""
export SOURCE_FILES="app"

${PREFIX}mypy $SOURCE_FILES
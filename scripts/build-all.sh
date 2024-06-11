#!/usr/bin/env bash
#
# Copyright (c) nexB Inc. and others. All rights reserved.
# ScanCode is a trademark of nexB Inc.
# SPDX-License-Identifier: Apache-2.0
# See http://www.apache.org/licenses/LICENSE-2.0 for the license text.
# See https://github.com/nexB/dependency-inspector for support or download.
# See https://aboutcode.org for more information about nexB OSS projects.
#
    
platforms=(
    "linux/arm64" 
    "linux/amd64" 
    "darwin/arm64" 
    "darwin/amd64" 
    "windows/arm64" 
    "windows/amd64"
    )

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name='deplock-'$GOOS'-'$GOARCH
    
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi    

    env GOOS=$GOOS GOARCH=$GOARCH go build -o ./build/$output_name

    if [ $? -ne 0 ]; then
        echo "An error occurred during the build process!"
        echo "GOOS: $GOOS"
        echo "GOARCH: $GOARCH"
        echo "Aborting build script execution."
        exit 1
    fi

done
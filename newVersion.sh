#!/bin/bash
VERSION=$1
MESSAGE=$2
git add .
git commit -m "$MESSAGE"
git tag -a "v$VERSION" -m "Version v$VERSION"
git push origin main "v$VERSION"
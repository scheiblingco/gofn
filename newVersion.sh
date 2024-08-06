#!/bin/bash
VERSION=$1
MESSAGE=$2

if [ -z "$VERSION" ]; then
  echo "Please provide a version number"
  exit 1
fi

if [ -z "$MESSAGE" ]; then
  echo "Please provide a commit message"
  exit 1
fi

git add .
git commit -m "$MESSAGE"
git tag -a "v$VERSION" -m "Version v$VERSION"
git push origin main "v$VERSION"
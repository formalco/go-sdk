#!/usr/bin/env bash

# Usage: ./bump_version.sh -r <patch|minor|major> - Increments version by chosen release type.
function help() {
  echo "Usage: ./bump_version.sh -r <patch|minor|major> - Increments version by chosen release type."
  echo ""
  echo "Flags:"
  echo "-r: release type (patch, minor, major)"
  echo "-h: help"
}

# get version from version file
VERSION=$(cat VERSION | tr -d '\n')

# r flag is for release
while getopts ":r:h" opt; do
  case $opt in
    r)
      RELEASE=$OPTARG
      ;;
    h)
      help
      exit 0
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      ;;
  esac
done

# if release flag is not set, exit
if [ -z "$RELEASE" ]; then
  echo "No release flag set. Aborting."
  help
  exit 1
fi

# bump version (version, release)
function bump_version(){
  major=$(echo "$1" | cut -d'.' -f1)
  minor=$(echo "$1" | cut -d'.' -f2)
  patch=$(echo "$1" | cut -d'.' -f3)

  case "$2" in
    "major")
      major=$((major+1))
      minor=0
      patch=0
      ;;
    "minor")
      minor=$((minor+1))
      patch=0
      ;;
    "patch")
      patch=$((patch+1))
      ;;
  esac

  echo "$major.$minor.$patch"
}

function confirm() {
	read -r -p "$@ [Y/n]: " confirm

	case "$confirm" in
		[Nn][Oo]|[Nn])
			echo "Aborting."
			exit
			;;
	esac
}

# convert release flag to lowercase
RELEASE=$(echo "$RELEASE" | tr '[:upper:]' '[:lower:]')
# if release flag is not patch, minor, or major, exit
if [ "$RELEASE" != "patch" ] && [ "$RELEASE" != "minor" ] && [ "$RELEASE" != "major" ]; then
  echo "Invalid release flag. Aborting."
  exit 1
fi

# update version
VERSION=$(bump_version "$VERSION" "$RELEASE")
confirm "Updating version  to $VERSION ?"

# update version in version file
echo $VERSION > ./VERSION 

echo $VERSION
export VERSION=$VERSION

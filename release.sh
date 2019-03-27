#!/bin/sh

VERSION=$1

test -z "$VERSION" &&{
	echo "Usage: $0 <version>"
	exit 1
}


make clean all

git tag $VERSION

export GITHUB_TOKEN=`cat .releasetoken`

gothub="go run github.com/itchio/gothub"
gothub_release="$gothub release --user thomasrudin-mt --repo mapserver"
gothub_upload="$gothub upload --user thomasrudin-mt --repo mapserver"

$gothub_release --tag $VERSION --name "Version $VERSION"

FILES="mapserver-linux-arm mapserver-linux-x86 mapserver-linux-x86_64 mapserver-mod.zip mapserver-windows-x86-64.exe mapserver-windows-x86.exe"

for file in $FILES
do
	$gothub_upload --tag $VERSION --name "$file" --file output/$file
done

git push --tags




# CGOを使ってみよう
このリポジトリのROOTがGOPATHという前提

資料には書かなかったiPhone Simulator向けのビルド
```
CGO_ENABLED=1 \
GOOS=darwin \
GOARCH=amd64 \
CC=`xcrun --sdk iphoneos -f clang` \
CXX=`xcrun --sdk iphoneos -f clang` \
CGO_CFLAGS="-isysroot `xcrun --sdk iphonesimulator --show-sdk-path` \
    -arch x86_64 -miphoneos-version-min=7.0" \
CGO_LDFLAGS="-isysroot `xcrun --sdk iphonesimulator --show-sdk-path` \
    -arch x86_64 -miphoneos-version-min=7.0" \
go build -pkgdir=$GOPATH/pkg/gomobile/pkg_darwin_amd64 -buildmode=c-archive \
    -tags=ios -o lib/ios/amd64/libaobayama.a c_aobayama
```
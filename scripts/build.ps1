$ROOTDIR=split-path -parent $PSScriptRoot
$BINDIR="$ROOTDIR\bin"
$env:GOPATH=$ROOTDIR

go build -o "$BINDIR\compile.exe" "compile/cli"
if ($LASTEXITCODE -ne 0) {
  Exit $LASTEXITCODE
}

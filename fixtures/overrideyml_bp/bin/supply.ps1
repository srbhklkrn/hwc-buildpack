Param(
  [Parameter(Mandatory=$True,Position=1)]
    [string]$BuildDir,
  [Parameter(Mandatory=$True,Position=2)]
    [string]$CacheDir,
  [Parameter(Mandatory=$True,Position=3)]
    [string]$DepsDir,
  [Parameter(Mandatory=$True,Position=4)]
    [string]$DepsIdx
)

echo "-----> OverrideYML Buildpack"
(Get-Content "$PSScriptRoot\..\override.yml").replace('BUILDPACK_DIR', "$PSScriptRoot\..") | Set-Content "$DepsDir/$DepsIdx/override.yml"

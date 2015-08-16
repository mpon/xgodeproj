# xgodeproj

## Description

parse xcodeproj/project.pbxproj

## Usage

Execute this command in your root dir for Xcode project or workspace.

It automatically finds the `project.pbxproj` file to parse.

### show all section name

```bash
$ xgodeproj show
```

### show section information

```bash
$ xgodeproj show --section <section name>
```

Implemented section

- [x] PBXBuildFile
- [ ] PBXContainerItemProxy
- [x] PBXFileReference
- [ ] PBXFrameworksBuildPhase
- [ ] PBXGroup
- [x] PBXNativeTarget
- [ ] PBXProject
- [x] PBXResourcesBuildPhase
- [x] PBXSourcesBuildPhase
- [ ] PBXTargetDependency
- [x] PBXVariantGroup
- [ ] XCBuildConfiguration
- [ ] XCConfigurationList

## Install

To install, use `go get`:

```bash
$ go get -d github.com/mpon/xgodeproj
```

## Contribution

1. Fork ([https://github.com/mpon/xgodeproj/fork](https://github.com/mpon/xgodeproj/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[mpon](https://github.com/mpon)

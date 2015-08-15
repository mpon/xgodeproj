# xgodeproj

## Description

parse xcodeproj/project.pbxproj

## Usage

### show all section name

```bash
$ xgodeproj show <your_project.pbxproj>
```

### show section

```bash
$ xgodeproj show <your_project.pbxproj> --section <section name>
```

Implemented section

- [ ] PBXBuildFile
- [ ] PBXContainerItemProxy
- [x] PBXFileReference
- [ ] PBXFrameworksBuildPhase
- [ ] PBXGroup
- [ ] PBXNativeTarget
- [ ] PBXProject
- [ ] PBXResourcesBuildPhase
- [ ] PBXSourcesBuildPhase
- [ ] PBXTargetDependency
- [ ] PBXVariantGroup
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

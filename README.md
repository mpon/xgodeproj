# xgodeproj

## Description

This library parses xcodeproj/project.pbxproj file.

You can see the project hierarchy from cli.

```bash
$ xgodeproj show --section PBXGroup
+ Sample
    AppDelegate.swift
    ViewController.swift
    Images.xcassets
  + Supporting Files
      Info.plist
+ SampleTests
    SampleTests.swift
  + Supporting Files
      Info.plist
+ Products
    Sample.app
    SampleTests.xctest
```

See already implemented list to parse pbxproj.

- [x] PBXBuildFile
- [ ] PBXContainerItemProxy
- [x] PBXFileReference
- [ ] PBXFrameworksBuildPhase
- [x] PBXGroup
- [x] PBXNativeTarget
- [ ] PBXProject
- [x] PBXResourcesBuildPhase
- [x] PBXSourcesBuildPhase
- [ ] PBXTargetDependency
- [x] PBXVariantGroup
- [ ] XCBuildConfiguration
- [ ] XCConfigurationList

## Usage

Execute this command in your root dir for Xcode project or workspace.
It automatically and recursively finds the `project.pbxproj` file to parse.
Also, you can specify project name when you have some projects in workspace.

### show all section name

```bash
$ xgodeproj show
$ xgodeproj show --project Sample
```

### show section information

```bash
$ xgodeproj show --section <section name>
$ xgodeproj show --section <section name> --project Sample
```

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

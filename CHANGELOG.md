# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.0]

### Changed

- Updated minimum Go version to 1.21
- Replaced `github.com/pkg/errors` with standard library error wrapping (`fmt.Errorf` with `%w`)
- Replaced deprecated `io/ioutil` usage with `io` and `os` equivalents
- Migrated CI from Travis CI to GitHub Actions
- Added automated release workflow via GitHub Actions
- Added Codecov integration for coverage reporting

[Unreleased]: https://github.com/adlio/trello/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/adlio/trello/releases/tag/v0.2.0

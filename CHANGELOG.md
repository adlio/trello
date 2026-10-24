# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.14.0] - 2026-10-24

### Changed

- Raised minimum Go version to 1.26 (required by `golang.org/x/time` v0.16.0)
- Bumped `golang.org/x/time` from v0.10.0 to v0.16.0
- Bumped CI actions: `actions/checkout` v5→v7, `actions/setup-go` v6→v7, `codecov/codecov-action` v5→v7

## [1.13.0] - 2026-10-23

### Changed

- Updated minimum Go version to 1.21
- Replaced `github.com/pkg/errors` with standard library error wrapping (`fmt.Errorf` with `%w`)
- Replaced deprecated `io/ioutil` usage with `io` and `os` equivalents
- Migrated CI from Travis CI to GitHub Actions
- Added automated release workflow via GitHub Actions
- Added Codecov integration for coverage reporting

### Fixed

- Client context (cancellation and deadlines) is now honored on all request methods — GET, PUT, POST, multipart POST, and DELETE — not just GET. (#101)

[Unreleased]: https://github.com/adlio/trello/compare/v1.14.0...HEAD
[1.14.0]: https://github.com/adlio/trello/compare/v1.13.0...v1.14.0
[1.13.0]: https://github.com/adlio/trello/releases/tag/v1.13.0

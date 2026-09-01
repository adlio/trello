# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- CI now tests against Go 1.21 through the latest stable release, plus macOS and
  Windows, so the minimum version declared in `go.mod` is actually verified
- CI verifies that `go.mod`/`go.sum` are tidy and that module checksums verify
- Scheduled `govulncheck` workflow (weekly, on `main`, and on dependency changes)
- Dependabot configuration for Go modules and GitHub Actions
- Aggregate "All checks passed" status check for branch protection

### Fixed

- Release notes are no longer interpolated into a shell command. Changelog
  entries containing backticks or `$(...)` were previously evaluated by the
  shell, which mangled the published notes and allowed command execution;
  notes are now passed to `gh` via `--notes-file`
- Markdown link-reference definitions are no longer appended to release notes
- Releases now fail loudly when `CHANGELOG.md` has no section for the tag,
  instead of publishing an empty release body

### Changed

- Release workflow builds and tests the tagged commit before publishing
- Pre-release tags (e.g. `v1.0.0-rc.1`) are now marked as pre-releases
- Workflows declare least-privilege `permissions` and per-job timeouts, and
  redundant in-progress CI runs are cancelled automatically

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

# gh-itkdev

[GitHub CLI](https://docs.github.com/en/github-cli) extension for ITK Development.

## Installation

```shell
gh extension install rimi-itk/gh-itkdev
```

**Note**: This may not (yet) work (cf. <https://github.com/rimi-itk/gh-itkdev/actions/workflows/release.yml>) and you
may have to [install from source](#installing-from-source).

### Installing from source

Assuming [Go is installed](https://go.dev/doc/install), you can [install from
source](https://cli.github.com/manual/gh_extension_install) by running

``` shell
git clone https://github.com/rimi-itk/gh-itkdev /tmp/gh-itkdev
cd /tmp/gh-itkdev
task build
gh extension install .
```

## Usage

### `gh itkdev changelog`

Manage changelog based on [keep a changelog](https://keepachangelog.com/en/1.1.0/):

```shell
gh itkdev changelog --help
```

Create changelog:

```shell
gh itkdev changelog create
```

Add pull request to changelog:

```shell
gh itkdev changelog add-pull-request
```

Prepare a new release:

```shell
gh itkdev changelog add-release «tag»
```

## Configuration

Default option values can be set in `.gh-itkdev.yaml` in the project folder or in `$HOME`.

``` yaml
changelog:
  «sub command»:
    «option»: «value»
```

### Example

Set the default branch for `add-release` to `main`:

``` yaml
changelog:
  add-release:
    base: main
```

Use `gh itkdev config` to dump the current config.

## Development

``` shell
task --list-all
```

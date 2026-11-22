# GoUtils

English | [中文](README_zh.md)

[![GitHub Action](https://github.com/MegaShow/goutils/actions/workflows/main.yaml/badge.svg)](https://github.com/MegaShow/goutils/actions/workflows/main.yaml)
[![CodeCov](https://codecov.io/gh/MegaShow/goutils/graph/badge.svg?token=VI2BCE8X5H)](https://codecov.io/gh/MegaShow/goutils)
[![GitHub Release](https://img.shields.io/github/v/release/megashow/goutils)](https://github.com/megashow/goutils/releases)
[![Go Reference](https://pkg.go.dev/badge/go.icytown.com/utils.svg)](https://pkg.go.dev/go.icytown.com/utils)

> English version is translated by tool.

`go.icytown.com/utils` is MegaShow's Go utility library, built on Go 1.23 and generics. It provides common data structures, utilities, and cache implementations.

## Installation

```sh
go get go.icytown.com/utils@latest
```

The latest version requires Go 1.23 or later. For older Go versions, use:

- v0.1.9: Go 1.21+

## Usage

For detailed usage instructions, see [GoDoc](https://pkg.go.dev/go.icytown.com/utils).

GoUtils provides the following utility packages.

| Package | Functions or Data Structures |
| -- | -- |
| ucond | If, IfFunc, Must(0-4), Not |
| ucrypto | MD5Hex, SHA1Hex, SHA256Hex, SHA512Hex |
| umath | CeilFloat, FloorFloat, RoundFloat |
| uobject | Default, Indirect, IndirectOr, IsNotZero, IsZero, Ptr |
| uruntime | GetFuncFullName, GetFuncName |
| uslice | All, Find, FindLast, Filter, GroupBy, Map, Of, ToMap, Unique, UniqueFunc |
| usync | Singleflight |
| uversion | SemVer |

GoUtils also provides commonly used container implementations.

| Package | Containers |
| -- | -- |
| maps | SyncMap |
| pools | Pool |
| sets | HashSet |
| stacks | ArrayStack |
| tuples | Pair, Triple |

In addition, GoUtils provides cache implementations.

- SimpleCache (experimental): a simple map-based cache implementation.

## License

GoUtils is distributed under the MIT License.

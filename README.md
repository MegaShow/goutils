# GoUtils

English | [中文](README_zh.md)

[![GitHub Action](https://github.com/MegaShow/goutils/actions/workflows/main.yaml/badge.svg)](https://github.com/MegaShow/goutils/actions/workflows/main.yaml)
[![CodeCov](https://codecov.io/gh/MegaShow/goutils/graph/badge.svg?token=VI2BCE8X5H)](https://codecov.io/gh/MegaShow/goutils)
[![GitHub Release](https://img.shields.io/github/v/release/megashow/goutils)](https://github.com/megashow/goutils/releases)
[![Go Reference](https://pkg.go.dev/badge/go.icytown.com/utils.svg)](https://pkg.go.dev/go.icytown.com/utils)

> English version is translated by tool.

`go.icytown.com/utils` is MegaShow's Golang tool library, based on Go 1.21 and type generic. It implements common data structures, tools, cache and etc.

## Installation

```sh
go get go.icytown.com/utils@latest
```

## Usages

For details, please refer to [GoDoc](https://pkg.go.dev/go.icytown.com/utils)。

GoUtils provides the following tools.

| Package | Functions or Data Structures |
| -- | -- |
| ucond | If, IfFunc, Not |
| ucrypto | MD5Hex, SHA1Hex, SHA256Hex, SHA512Hex |
| umath | CeilFloat, FloorFloat, RoundFloat |
| uobject | Default, Indirect, IndirectOr, IsNotZero, IsZero, Ptr |
| uruntime | GetFuncFullName, GetFuncName |
| uslice | Find, FindLast, Filter, GroupBy, Map, Of, ToMap, Unique, UniqueFunc |
| usync | Singleflight |
| uversion | SemVer |

GoUtils alse provides some commonly used containers.

| Package | Containers |
| -- | -- |
| maps | SyncMap |
| pools | Pool |
| sets | HashSet |
| stacks | ArrayStack |
| tuples | Pair, Triple |

In addition, GoUtils alse provides some implementation of cache.

- SimpleCache (Experimental)：a simple cache implementation based on map.

## License

GoUtils is distributed under the MIT License.

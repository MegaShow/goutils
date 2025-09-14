# GoUtils

[English](README.md) | 中文

[![GitHub Action](https://github.com/MegaShow/goutils/actions/workflows/main.yaml/badge.svg)](https://github.com/MegaShow/goutils/actions/workflows/main.yaml)
[![CodeCov](https://codecov.io/gh/MegaShow/goutils/graph/badge.svg?token=VI2BCE8X5H)](https://codecov.io/gh/MegaShow/goutils)
[![GitHub Release](https://img.shields.io/github/v/release/megashow/goutils)](https://github.com/megashow/goutils/releases)
[![Go Reference](https://pkg.go.dev/badge/go.icytown.com/utils.svg)](https://pkg.go.dev/go.icytown.com/utils)

`go.icytown.com/utils` 是 MegaShow 的 Golang 工具库，基于 Go 1.21 和泛型能力，实现了常用的数据结构、工具、缓存封装等。

## 安装

```sh
go get go.icytown.com/utils@latest
```

## 使用

详细的使用说明请查阅 [GoDoc](https://pkg.go.dev/go.icytown.com/utils)。

GoUtils 提供了以下工具库。

| 工具库 | 函数或数据结构 |
| -- | -- |
| ucond | If, IfFunc, Must(0-4), Not |
| ucrypto | MD5Hex, SHA1Hex, SHA256Hex, SHA512Hex |
| umath | CeilFloat, FloorFloat, RoundFloat |
| uobject | Default, Indirect, IndirectOr, IsNotZero, IsZero, Ptr |
| uruntime | GetFuncFullName, GetFuncName |
| uslice | Find, FindLast, Filter, GroupBy, Map, Of, ToMap, Unique, UniqueFunc |
| usync | Singleflight |
| uversion | SemVer |

GoUtils 还提供了一些常用的容器封装。

| 容器库 | 容器 |
| -- | -- |
| maps | SyncMap |
| pools | Pool |
| sets | HashSet |
| stacks | ArrayStack |
| tuples | Pair, Triple |

除此之外，GoUtils 还提供了一些常用的缓存封装。

- SimpleCache (实验性)：基于 map 的简单缓存实现。

## 协议

GoUtils 基于 MIT 协议开源。

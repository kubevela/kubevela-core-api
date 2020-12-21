# KubeVela Core API

API types that work for KubeVela Core CRDs.

## Purpose

This library is the canonical location of the KubeVela Core API definition.

The code is synced from [kubevela/apis](https://github.com/oam-dev/kubevela/tree/master/apis) every release.

You can use this separated package if you want:

* use it as SDK and build your own user interface.
* avoid conflicts of `go.mod` by reducing dependency from KubeVela.

## Usage

Refer to [test/main.go](test/main.go) as example.
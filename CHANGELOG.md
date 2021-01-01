# Changelog
All notable changes to this project will be documented in this file. See [conventional commits](https://www.conventionalcommits.org/) for commit guidelines.

- - -
## 0.8.0 - 2021-01-01


### Refactoring

[c39ef5](https://github.com/pjvds/tunl/commit/c39ef5f292ea86ce06b1a7b3abfb2b29315b86d3) - use default tls config from tls package - [pjvds](https://github.com/pjvds)

[4df146](https://github.com/pjvds/tunl/commit/4df146c64fb73e81b3ba5f2d4cae4b95ac404be7) - use a single listener in context - [pjvds](https://github.com/pjvds)


### Features

[67b2a5](https://github.com/pjvds/tunl/commit/67b2a579148839792d78c55d8067ed8645741561) - remove tls muxer - [pjvds](https://github.com/pjvds)


- - -
## 0.7.3 - 2021-01-01


### Bug Fixes

[aded6a](https://github.com/pjvds/tunl/commit/aded6af937f4bbab683c9ada3b9f94d32c5958e4) - dial tls from client when required - [pjvds](https://github.com/pjvds)


- - -
## 0.7.2 - 2021-01-01


### Bug Fixes

[d56a0a](https://github.com/pjvds/tunl/commit/d56a0a0243b87b25dcaa6f16060fd5ba39504a24) - remove unused listener - [pjvds](https://github.com/pjvds)


- - -
## 0.7.1 - 2021-01-01


### Bug Fixes

[91a853](https://github.com/pjvds/tunl/commit/91a8531f4386f242060af86bd268466898043a90) - fix nil pointer in mux creation path - [pjvds](https://github.com/pjvds)


- - -
## 0.7.0 - 2021-01-01


### Features

[e3f582](https://github.com/pjvds/tunl/commit/e3f582b8ca97d6f4274a1b48cb308b918873ff35) - change default host from http to https - [pjvds](https://github.com/pjvds)


### Bug Fixes

[4f7289](https://github.com/pjvds/tunl/commit/4f72892cd2014f030dee4a8b22c0458b197e5aeb) - fix wrong mux created when using TLS - [pjvds](https://github.com/pjvds)


- - -
## 0.6.0 - 2021-01-01


### Bug Fixes

[6096a8](https://github.com/pjvds/tunl/commit/6096a8d37cb480748d424461d2fd1a616e5e17a3) - improve fix unexpected response error message and include full status - [pjvds](https://github.com/pjvds)


### Features

[ca793c](https://github.com/pjvds/tunl/commit/ca793c84d24c8514ae3e9005d87135499ef1347b) - add TLS hosting - [pjvds](https://github.com/pjvds)


- - -
## 0.5.0 - 2020-12-31


### Features

[54bf04](https://github.com/pjvds/tunl/commit/54bf042cb801b125b0d0d5f61cfb468ffc583c41) - expose version set during build - [pjvds](https://github.com/pjvds)


- - -
## 0.4.3-alpha - 2020-12-29


### Miscellaneous Chores

[439cc3](https://github.com/pjvds/tunl/commit/439cc329d18c3bf5a4e1ba467091d5dae5cbe6f3) - fix docker hub login - [pjvds](https://github.com/pjvds)

[77188e](https://github.com/pjvds/tunl/commit/77188e7f6a858220067f4c805ad27335caa1f136) - publich docker images to docker hub - [pjvds](https://github.com/pjvds)


- - -
## 0.4.1-alpha - 2020-12-29


### Miscellaneous Chores

[7f49c8](https://github.com/pjvds/tunl/commit/7f49c81bd05953a57af1f11bd0120946c9bfb179) - use gitHub Packages for docker images - [pjvds](https://github.com/pjvds)


- - -
## 0.4.0-alpha - 2020-12-29


### Features

[674908](https://github.com/pjvds/tunl/commit/6749087efe5c90a6831bfecc7b168f6b824b25cf) - move host flag to app instead of commands - [pjvds](https://github.com/pjvds)


### Miscellaneous Chores

[22b848](https://github.com/pjvds/tunl/commit/22b848511cb24f4fc3f5f4f9257275914435305b) - use build artifact for docker images - [pjvds](https://github.com/pjvds)

[5eea74](https://github.com/pjvds/tunl/commit/5eea7482b9ca0324ac1eb577d3a4d8f82e9b0bd4) - fix docker image templates - [pjvds](https://github.com/pjvds)

[692bc1](https://github.com/pjvds/tunl/commit/692bc12d2cee44f552196c25b4460ccef33846b0) - distribute tunl as docker image - [pjvds](https://github.com/pjvds)


- - -
## 0.3.0-alpha - 2020-12-29


### Bug Fixes

[c5e4fa](https://github.com/pjvds/tunl/commit/c5e4fa5a5693a6c5ae63fb84012842c27cba6ff5) - files command not compatible with latest version changes - [pjvds](https://github.com/pjvds)


### Features

[cefb6d](https://github.com/pjvds/tunl/commit/cefb6df3d6b04006109512e1a9830db39ae86604) - improve application input validdation and messages - [pjvds](https://github.com/pjvds)


- - -
## 0.2.0-alpha - 2020-12-29


### Miscellaneous Chores

[bc640d](https://github.com/pjvds/tunl/commit/bc640d956507a6b7c08c94c476209136fa162551) - 0.1.0-alpha - [pjvds](https://github.com/pjvds)


### Features

[df73b2](https://github.com/pjvds/tunl/commit/df73b280be0dd059992d229a1b8d18881a688af1) - use go-vhost for multiplexing - [pjvds](https://github.com/pjvds)


- - -
## 0.1.0-alpha - 2020-12-29


### Features

[e95480](https://github.com/pjvds/tunl/commit/e95480af8bebbae2b40979d2fd8d93772e299d7f) - improve error log messages for handshake procedure - [pjvds](https://github.com/pjvds)


### Miscellaneous Chores

[51f26c](https://github.com/pjvds/tunl/commit/51f26ca58e7f52ceb1fafdc9be960e5a345e0098) - 0.1.0-alpha - [pjvds](https://github.com/pjvds)

[457ae5](https://github.com/pjvds/tunl/commit/457ae5374d2ab471f3e0839741d99c13cf07eafa) - add github url and authors definition to cog config - [pjvds](https://github.com/pjvds)

[5ba1de](https://github.com/pjvds/tunl/commit/5ba1de5f95503737e580f523f4b4c0e049a3b7ff) - add cocogitto and goreleaser configurations - [pjvds](https://github.com/pjvds)

[252ab9](https://github.com/pjvds/tunl/commit/252ab9d8cc4e5528bfd87dfdca01c2285c8ecc64) - add release pipeline - [pjvds](https://github.com/pjvds)


### Bug Fixes

[627fb6](https://github.com/pjvds/tunl/commit/627fb6ae3d83e1147812dc8405f9264deeb32ab7) - do not fatal exit when creation of mux connection fails - [pjvds](https://github.com/pjvds)

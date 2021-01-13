# Changelog
All notable changes to this project will be documented in this file. See [conventional commits](https://www.conventionalcommits.org/) for commit guidelines.

- - -
## 0.21.0 - 2021-01-13


### Miscellaneous Chores

[b6c79b](https://github.com/pjvds/tunl/commit/b6c79b2502ec2dc6a3bd97312f04936e147469c7) - disable snapcrafts in release pipeline - [pjvds](https://github.com/pjvds)

[673de9](https://github.com/pjvds/tunl/commit/673de94282d69d6f7b10eacc7a60704be0cf2b29) - improve readme with minor changes - [pjvds](https://github.com/pjvds)

[055c7a](https://github.com/pjvds/tunl/commit/055c7ad658fb18062ea9aa57e56a51febbd2ea1b) - only require network plug for snapcrafts - [pjvds](https://github.com/pjvds)


### Features

[758383](https://github.com/pjvds/tunl/commit/75838382db8fcfe0f004b6bd437c8d6106464bd4) - improve command and flag usage descriptions - [pjvds](https://github.com/pjvds)

[d013ea](https://github.com/pjvds/tunl/commit/d013eacee40168555829c1ccdfd7ddc1c2af86a0) - skip TLS verification for local address by default - [pjvds](https://github.com/pjvds)


### Documentation

[07cca3](https://github.com/pjvds/tunl/commit/07cca3d87d17c3312771e9384c437a5d8d180d7c) - add a link to the latest binary release - [pjvds](https://github.com/pjvds)

[724a09](https://github.com/pjvds/tunl/commit/724a0972f3ed17998356925bcf1519b3b09fa68e) - switch to TLS for install script - [pjvds](https://github.com/pjvds)

[f6be3d](https://github.com/pjvds/tunl/commit/f6be3d5623ae012fd1cf74d93dd2ff96eabb72a1) - add curl and wget installation description - [pjvds](https://github.com/pjvds)

[bcb8f5](https://github.com/pjvds/tunl/commit/bcb8f5d007ea8159f4a531724287518c79cd529d) - shorten example headings - [pjvds](https://github.com/pjvds)

[b120e7](https://github.com/pjvds/tunl/commit/b120e7fbd6ceeecf3d953de6077f7f4ffc2d06cb) - shorting example headings - [pjvds](https://github.com/pjvds)


- - -
## 0.20.0 - 2021-01-12


### Miscellaneous Chores


### Refactoring

[a086b3](https://github.com/pjvds/tunl/commit/a086b3b2426c1f5a02aa23cbc0b381e969ee1787) - move addressing to server package - [pjvds](https://github.com/pjvds)


### Features

[bdfb8d](https://github.com/pjvds/tunl/commit/bdfb8dd103e5e93e251091c75492e7075376998c) - print public address to stdout and target to stderr - [pjvds](https://github.com/pjvds)

[b3b854](https://github.com/pjvds/tunl/commit/b3b854f67bc4b2c809dc79778a2d9bb75db540ef) - improve usage description of commands - [pjvds](https://github.com/pjvds)


### Bug Fixes

[61806c](https://github.com/pjvds/tunl/commit/61806c2fcdff7629b6d8fe3c96629f9dfcecc87d) - incorrect exit result for tunnel accept - [pjvds](https://github.com/pjvds)


- - -
## 0.19.2 - 2021-01-12


### Bug Fixes

[36a202](https://github.com/pjvds/tunl/commit/36a202abb11e18e8fee586f0a3a27ea3667dd6ed) - autocomplete files not included in archives - [pjvds](https://github.com/pjvds)


- - -
## 0.19.1 - 2021-01-11


### Bug Fixes

[5ffd0f](https://github.com/pjvds/tunl/commit/5ffd0fdd5646419dd15e931e83701e1e6bb5a5f0) - autocomplete not included in release packages - [pjvds](https://github.com/pjvds)


- - -
## 0.19.0 - 2021-01-11


### Documentation

[b120e7](https://github.com/pjvds/tunl/commit/b120e7fbd6ceeecf3d953de6077f7f4ffc2d06cb) - shorting example headings - [pjvds](https://github.com/pjvds)


### Bug Fixes

[6e2069](https://github.com/pjvds/tunl/commit/6e2069ae6f3becaea5b088971c4909b891bb49d3) - explicitly close target connection when proxy is done - [pjvds](https://github.com/pjvds)


### Features

[2fa41d](https://github.com/pjvds/tunl/commit/2fa41d0b8775752b1fc5000a5a2b67e7db553b95) - add docker command to create tunnel directly to a docker container - [pjvds](https://github.com/pjvds)

[d3adac](https://github.com/pjvds/tunl/commit/d3adacd0d713ba9c15d51a626386d6a6acef83ab) - add usage text to commands definition - [pjvds](https://github.com/pjvds)

[a4323b](https://github.com/pjvds/tunl/commit/a4323b71be4146b05687b5e68848fff0e1191f4f) - add bash and zsh auto completion - [pjvds](https://github.com/pjvds)

[3dbbf3](https://github.com/pjvds/tunl/commit/3dbbf3996e3619cc06a38a08964ff345bebdbd49) - host tcp tunnel on tcp subdomain - [pjvds](https://github.com/pjvds)



- - -
## 0.18.0 - 2021-01-07


### Features

[dfd4ad](https://github.com/pjvds/tunl/commit/dfd4ad9421bfb90a87d9fb639e86c61d2ca38fa6) - support multiple certificates - [pjvds](https://github.com/pjvds)


- - -
## 0.17.2 - 2021-01-07


### Bug Fixes

[331a30](https://github.com/pjvds/tunl/commit/331a30a4bca481be2d4650a42dcbb15309373cac) - missing newline character in tunnel address printing - [pjvds](https://github.com/pjvds)


- - -
## 0.17.1 - 2021-01-07


[1e8f79](https://github.com/pjvds/tunl/commit/1e8f7905a63bb4ee5e2e68cfc14a02867ac1697b) - address not printed after tunnel creation - [pjvds](https://github.com/pjvds)


### Documentation

[b3e1de](https://github.com/pjvds/tunl/commit/b3e1ded93aa8710854f7d5223420d12defc3932e) - short first example heading - [pjvds](https://github.com/pjvds)

[16993e](https://github.com/pjvds/tunl/commit/16993ed237a4c12a93bbe8e0920faa70ccb16289) - improve examples - [pjvds](https://github.com/pjvds)

[681994](https://github.com/pjvds/tunl/commit/6819941cb4584ce3b604a613965c0a89f37cdde3) - add README - [pjvds](https://github.com/pjvds)


- - -
## 0.17.0 - 2021-01-07


### Features

[a3059c](https://github.com/pjvds/tunl/commit/a3059c407d1d4ade77a6dfc8d331bf00f0d68fa8) - support raw TCP tunnel - [pjvds](https://github.com/pjvds)

[d12de6](https://github.com/pjvds/tunl/commit/d12de6f98ff2fdf0b5231e2a4def08edea1b89de) - support and release binaries for arm and arm64 architecture - [pjvds](https://github.com/pjvds)

[4ec500](https://github.com/pjvds/tunl/commit/4ec5007c3b9f045eba775f55e09b8f9efc8f186e) - all command input errors are written to stderr instread of - [pjvds](https://github.com/pjvds)

[324a99](https://github.com/pjvds/tunl/commit/324a9912c5565a1f29314bca3b2218a9c9771636) - reconnect tunnel when connection get interupted - [pjvds](https://github.com/pjvds)

### Bug Fixes

[e185e6](https://github.com/pjvds/tunl/commit/e185e6bd423929c76226c6c742f75c1dd70d339f) - snap package not requesting correct plugs - [pjvds](https://github.com/pjvds)


- - -
## 0.16.2 - 2021-01-06

### Miscellaneous Chores

[795544](https://github.com/pjvds/tunl/commit/795544985a9770a59013a67355b35693076bad0d) - lower package confinement to devmode - [pjvds](https://github.com/pjvds)

### Refactoring

[b30eb4](https://github.com/pjvds/tunl/commit/b30eb4966ffeb35e578d9546a4c6ada802ff2f64) - move all tunnel and networking logic to own package and - [pjvds](https://github.com/pjvds)

- - -
## 0.16.1 - 2021-01-05


### Refactoring

[4da164](https://github.com/pjvds/tunl/commit/4da16456b3213fed0f65bff4009d451702f64db7) - split commands and package logic - [pjvds](https://github.com/pjvds)


### Miscellaneous Chores

[cd64cc](https://github.com/pjvds/tunl/commit/cd64cc7ad6f6b7e52f4e7aa1dee449446c9e768b) - shorten usage description - [pjvds](https://github.com/pjvds)

[1c6e06](https://github.com/pjvds/tunl/commit/1c6e06928f09ee9460c536475b1a96bc044cb5ae) - use zip archives for windows builds - [pjvds](https://github.com/pjvds)

[836ffe](https://github.com/pjvds/tunl/commit/836ffe1fcafbd3bdb397ca27ebc90348c9355703) - add apache 2.0 license file - [pjvds](https://github.com/pjvds)

[f4965b](https://github.com/pjvds/tunl/commit/f4965b9bdccaa41d120ca53787222624882bd3b9) - remove checksum for unused jwt-go version - [pjvds](https://github.com/pjvds)

[c917ec](https://github.com/pjvds/tunl/commit/c917ec3fbf72d0b6236775a24972308c2896cac5) - change release pipeline to publish releases to snap store - [pjvds](https://github.com/pjvds)

[d33e3f](https://github.com/pjvds/tunl/commit/d33e3fef554a5b05cf645c702d01aca160d1788a) - 0.10.0 - [pjvds](https://github.com/pjvds)

[2ec95d](https://github.com/pjvds/tunl/commit/2ec95def75f77dd9510f6dd68df7944e6e71a437) - fix .goreleaser format error - [pjvds](https://github.com/pjvds)

[ff7a9b](https://github.com/pjvds/tunl/commit/ff7a9b5588b9a49248691fc666182042144af0a8) - add favicon assets - [pjvds](https://github.com/pjvds)

[1906a9](https://github.com/pjvds/tunl/commit/1906a94a1b8c1afc003816a88da9048ff2db03be) - publish snapcraft package with correct permissions - [pjvds](https://github.com/pjvds)


### Features

[d6328c](https://github.com/pjvds/tunl/commit/d6328ca2cb8ea4cb81090dddda7b69545e49d8f3) - prefix target url with schema if not exists - [pjvds](https://github.com/pjvds)

[34126a](https://github.com/pjvds/tunl/commit/34126aa4494207364577c7f1739967ffad67f993) - print tunnel explicitly with tunnel address and - [pjvds](https://github.com/pjvds)

[2942f8](https://github.com/pjvds/tunl/commit/2942f894d3b3a2efa277144dbd8d72b36a09c0bb) - make dir argument optional and validate - [pjvds](https://github.com/pjvds)

[cb5955](https://github.com/pjvds/tunl/commit/cb5955143b2937f4276b6c9ec196fd4f28a7e30c) - switch to distroless image - [pjvds](https://github.com/pjvds)

[95a8d4](https://github.com/pjvds/tunl/commit/95a8d4d110be9a845562030c701b47ca36ce00e6) - tunnel can be reclaimed after disconnect - [pjvds](https://github.com/pjvds)

[1664e0](https://github.com/pjvds/tunl/commit/1664e0e0a284b4c78be0ea25974ce398c63dc370) - serve favicon and meta data if not exists in file server location - [pjvds](https://github.com/pjvds)


### Bug Fixes

[e185e6](https://github.com/pjvds/tunl/commit/e185e6bd423929c76226c6c742f75c1dd70d339f) - snap package not requesting correct plugs - [pjvds](https://github.com/pjvds)


- - -
## 0.10.0 - 2021-01-04


### Bug Fixes

[e185e6](https://github.com/pjvds/tunl/commit/e185e6bd423929c76226c6c742f75c1dd70d339f) - snap package not requesting correct plugs - [pjvds](https://github.com/pjvds)


### Miscellaneous Chores

[2ec95d](https://github.com/pjvds/tunl/commit/2ec95def75f77dd9510f6dd68df7944e6e71a437) - fix .goreleaser format error - [pjvds](https://github.com/pjvds)

[ff7a9b](https://github.com/pjvds/tunl/commit/ff7a9b5588b9a49248691fc666182042144af0a8) - add favicon assets - [pjvds](https://github.com/pjvds)

[1906a9](https://github.com/pjvds/tunl/commit/1906a94a1b8c1afc003816a88da9048ff2db03be) - publish snapcraft package with correct permissions - [pjvds](https://github.com/pjvds)


### Features

[cb5955](https://github.com/pjvds/tunl/commit/cb5955143b2937f4276b6c9ec196fd4f28a7e30c) - switch to distroless image - [pjvds](https://github.com/pjvds)

[95a8d4](https://github.com/pjvds/tunl/commit/95a8d4d110be9a845562030c701b47ca36ce00e6) - tunnel can be reclaimed after disconnect - [pjvds](https://github.com/pjvds)

[1664e0](https://github.com/pjvds/tunl/commit/1664e0e0a284b4c78be0ea25974ce398c63dc370) - serve favicon and meta data if not exists in file server location - [pjvds](https://github.com/pjvds)


- - -
## 0.9.1 - 2021-01-04


### Miscellaneous Chores

[4934b3](https://github.com/pjvds/tunl/commit/4934b339e37dc38262f10874f8cdb12237c75276) - add tunl logo - [pjvds](https://github.com/pjvds)

[6cc6ca](https://github.com/pjvds/tunl/commit/6cc6cacff272ac46417429fbb37676be9a5b53f3) - disable snap release - [pjvds](https://github.com/pjvds)


- - -
## 0.9.0 - 2021-01-04


### Miscellaneous Chores

[827af9](https://github.com/pjvds/tunl/commit/827af9d1cd62aa77b5666de7871f645f8066602d) - only run release pipeline on new tags - [pjvds](https://github.com/pjvds)

[f0c375](https://github.com/pjvds/tunl/commit/f0c375546416a82021a88f5685690b158888c0b0) - add snapcraft release - [pjvds](https://github.com/pjvds)


### Bug Fixes

[cb81f2](https://github.com/pjvds/tunl/commit/cb81f29cef1fd1a6003a7e7e1aba407a4be978c1) - host value error messages not including underlying error - [pjvds](https://github.com/pjvds)


### Features

[e3a3e2](https://github.com/pjvds/tunl/commit/e3a3e2c5b22e1473c8315b2ef57fb2c10432f667) - add tunnel address template support - [pjvds](https://github.com/pjvds)


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

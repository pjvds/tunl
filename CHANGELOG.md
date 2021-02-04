# Changelog
All notable changes to this project will be documented in this file. See [conventional commits](https://www.conventionalcommits.org/) for commit guidelines.

- - -
## 0.33.0 - 2021-02-04


### Revert

[dcdf89](https://github.com/pjvds/tunl/commit/dcdf893b274f9f3c2079d409a5d5b0e43a9af8c4) - remove raw tcp tunnel over tls - [pjvds](https://github.com/pjvds)


- - -
## 0.32.1 - 2021-02-04


### Bug Fixes

[81281b](https://github.com/pjvds/tunl/commit/81281b06d3aa323ac52a8bdb09d9147aacf749ff) - fix tls server connection handshake - [pjvds](https://github.com/pjvds)


- - -
## 0.32.0 - 2021-02-04


### Features

[460be1](https://github.com/pjvds/tunl/commit/460be1c0729ce1a4d1644e2bb2edac29b8c3cdd5) - log vhost connection type discovery failure details - [pjvds](https://github.com/pjvds)


- - -
## 0.31.0 - 2021-02-04


### Features

[6ebbe2](https://github.com/pjvds/tunl/commit/6ebbe29ffde398ad2d78afcab1c166baec41868d) - adds ability to setup TLS tunnel - [pjvds](https://github.com/pjvds)

[9b446d](https://github.com/pjvds/tunl/commit/9b446d61fb3e41ec31ff9b025a6e6e935de48635) - adds ability to copy public address to clipboard - [pjvds](https://github.com/pjvds)

[43147c](https://github.com/pjvds/tunl/commit/43147c0fb67bcc3cecc01662f4186ffa2ceab00c) - add TLS flag to open public vhost that can be routed by name - [pjvds](https://github.com/pjvds)


- - -
## 0.30.0 - 2021-02-03


### Bug Fixes

[aa9355](https://github.com/pjvds/tunl/commit/aa9355f896b776c0576a1cf95125014b6aabdb51) - fix new version flow build errors - [pjvds](https://github.com/pjvds)


### Features

[af96a3](https://github.com/pjvds/tunl/commit/af96a3373edbd0620117191b17d7af6b600b3ec4) - join host and port into a single argument - [pjvds](https://github.com/pjvds)

[93388b](https://github.com/pjvds/tunl/commit/93388b29194def5a05dbb1b14168b05800b5816c) - print if new version is available - [pjvds](https://github.com/pjvds)


### Miscellaneous Chores

[ab6907](https://github.com/pjvds/tunl/commit/ab6907d3f7c7cb6d3d4e7de90e5c595cd7c66d8f) - bump go version in go.mod to go 1.16 - [pjvds](https://github.com/pjvds)

[e6a721](https://github.com/pjvds/tunl/commit/e6a72169cec31af5240c3b78d706dd2cb9fcf763) - move to go 1.16 release candidate 1 - [pjvds](https://github.com/pjvds)


- - -
## 0.29.1 - 2021-01-25


### Bug Fixes

[303a82](https://github.com/pjvds/tunl/commit/303a8280f13966bf679c1e6927cde71ee89a827c) - fix potential wrong control port choosen - [pjvds](https://github.com/pjvds)


- - -
## 0.29.0 - 2021-01-25


### Features

[f28a3d](https://github.com/pjvds/tunl/commit/f28a3dacb1088cb2f4cbd120e94729b1fbd4d85a) - increase details in error messages for connect state - [pjvds](https://github.com/pjvds)

[fa76b6](https://github.com/pjvds/tunl/commit/fa76b6c8ba3ae6a44d9a13254b5fd4842e92c018) - remove default port lookup for most used uri schemes - [pjvds](https://github.com/pjvds)

[7542e0](https://github.com/pjvds/tunl/commit/7542e0f06e938dd12346b4baf34fcb67f3353855) - adds honeycomb metrics logging - [pjvds](https://github.com/pjvds)


### Documentation

[e488d0](https://github.com/pjvds/tunl/commit/e488d070e71a5119d28322686d0ca2a9dbd55a5d) - add installation path description to examples - [pjvds](https://github.com/pjvds)


- - -
## 0.28.1 - 2021-01-20


### Bug Fixes

[66ef38](https://github.com/pjvds/tunl/commit/66ef3829324834b81ae8563d4e767da2b679e84b) - fix client version only negotiated in auto-complete - [pjvds](https://github.com/pjvds)


- - -
## 0.28.0 - 2021-01-20


### Miscellaneous Chores

[8d0ac1](https://github.com/pjvds/tunl/commit/8d0ac10c30c3f6ed5e4582032cbd3b4bed165000) - run goreleaser as snapshot build - [pjvds](https://github.com/pjvds)

[5fed77](https://github.com/pjvds/tunl/commit/5fed77d484baefa0417f57a7d5a652f7baaa4db5) - fix build pipeline - [pjvds](https://github.com/pjvds)

[e74ba7](https://github.com/pjvds/tunl/commit/e74ba7d45e526420e9a19a6f708b3d57514629d8) - install Go 1.16 via setup-go action - [pjvds](https://github.com/pjvds)


### Features

[4bee51](https://github.com/pjvds/tunl/commit/4bee51e62db54489b2f951508d7234c95deae78d) - add docker client API version negotiation - [pjvds](https://github.com/pjvds)


- - -
## 0.27.0 - 2021-01-20


### Bug Fixes

[d2d27e](https://github.com/pjvds/tunl/commit/d2d27e70ba3c844ca79af51afb87527fc114e2a0) - fix missing template for single password auth - [pjvds](https://github.com/pjvds)


### Features

[348ded](https://github.com/pjvds/tunl/commit/348ded04e0eb9e9973d5f6810b61b90c60dd527e) - add webdav command to publish a local dir via webdav - [pjvds](https://github.com/pjvds)

[c28fd9](https://github.com/pjvds/tunl/commit/c28fd9249548420286b21678104d409cc74eac28) - add ability to add single password protection - [pjvds](https://github.com/pjvds)

[53f551](https://github.com/pjvds/tunl/commit/53f5513e50b294c434cef91049df508865263ac4) - add http basic auth support - [pjvds](https://github.com/pjvds)

[262259](https://github.com/pjvds/tunl/commit/262259f9c08a428c3565fc2ae3c1ec23589de642) - validate directory before starting tunnel - [pjvds](https://github.com/pjvds)

[10e6cb](https://github.com/pjvds/tunl/commit/10e6cbd7fcbe16e7c7f4805163fe37f25ab4709b) - replace files command with dir - [pjvds](https://github.com/pjvds)

[7c1c6a](https://github.com/pjvds/tunl/commit/7c1c6a834f1fcf7e92aca5b7087c8cdfb385c99d) - add ability to print QR code with the public address - [pjvds](https://github.com/pjvds)


- - -
## 0.26.0 - 2021-01-19


### Features

[be188b](https://github.com/pjvds/tunl/commit/be188b348412da23dcd354c6f0380ae2e1f6175e) - add ability to print QR code with the public address - [pjvds](https://github.com/pjvds)


### Miscellaneous Chores

[fe9ff8](https://github.com/pjvds/tunl/commit/fe9ff85524bb2aab4d06ad475770eb1f913ce0a8) - remove snapcraft install and login from release workflows - [pjvds](https://github.com/pjvds)

[1433ad](https://github.com/pjvds/tunl/commit/1433ad9e93143b27acc757296bdb8cf2c1cba6e2) - switch from release to build command for build workflow - [pjvds](https://github.com/pjvds)


- - -
## 0.26.0 - 2021-01-19


### Features

[be188b](https://github.com/pjvds/tunl/commit/be188b348412da23dcd354c6f0380ae2e1f6175e) - add ability to print QR code with the public address - [pjvds](https://github.com/pjvds)


- - -
## 0.25.0 - 2021-01-16


### Features

[343e67](https://github.com/pjvds/tunl/commit/343e67575f40765e0838bc60bb5c2f04d616c297) - simplify files serving by removing fallback handling - [pjvds](https://github.com/pjvds)


- - -
## 0.24.0 - 2021-01-15


### Bug Fixes

[c76e1d](https://github.com/pjvds/tunl/commit/c76e1de30f0bf32ac6c6be366eb2a7e938cf3f14) - error template not served correctly - [pjvds](https://github.com/pjvds)


### Miscellaneous Chores

[d42b72](https://github.com/pjvds/tunl/commit/d42b72c5249d61a3860e3c2ace6206d6bd1a5106) - bind homebrew tap token - [pjvds](https://github.com/pjvds)

[ad7bf4](https://github.com/pjvds/tunl/commit/ad7bf408a15296cc88a53b9b3152eaf80e7cb746) - add homebrew tap token - [pjvds](https://github.com/pjvds)

[d076a5](https://github.com/pjvds/tunl/commit/d076a5b31da71d001a90f4e8786ce688fa233df8) - release version to brew tap - [pjvds](https://github.com/pjvds)


- - -
## 0.23.0-2 - 2021-01-15


### Miscellaneous Chores

[d076a5](https://github.com/pjvds/tunl/commit/d076a5b31da71d001a90f4e8786ce688fa233df8) - release version to brew tap - [pjvds](https://github.com/pjvds)


- - -
## 0.23.0 - 2021-01-15


### Miscellaneous Chores

[7be460](https://github.com/pjvds/tunl/commit/7be4600bc3b622ad0d7bba93a6fd61c737c10a8d) - build releases with go 1.16beta1 - [pjvds](https://github.com/pjvds)

[68cbca](https://github.com/pjvds/tunl/commit/68cbca283c2314eb11a84102a73987afdad7cd8b) - add build pipeline with go1.16beta1 - [pjvds](https://github.com/pjvds)


### Features

[c9f538](https://github.com/pjvds/tunl/commit/c9f53880c64e2b6eb0c36214d635abf57adbe2ba) - optimize file embeds - [pjvds](https://github.com/pjvds)


- - -
## 0.22.1 - 2021-01-15


### Bug Fixes

[3bf61a](https://github.com/pjvds/tunl/commit/3bf61a9abaa10de4c20bf9215849e48449a1c4a8) - potential runtime panic when launching tunl - [pjvds](https://github.com/pjvds)

[61806c](https://github.com/pjvds/tunl/commit/61806c2fcdff7629b6d8fe3c96629f9dfcecc87d) - incorrect exit result for tunnel accept - [pjvds](https://github.com/pjvds)

[36a202](https://github.com/pjvds/tunl/commit/36a202abb11e18e8fee586f0a3a27ea3667dd6ed) - autocomplete files not included in archives - [pjvds](https://github.com/pjvds)

[5ffd0f](https://github.com/pjvds/tunl/commit/5ffd0fdd5646419dd15e931e83701e1e6bb5a5f0) - autocomplete not included in release packages - [pjvds](https://github.com/pjvds)

[6e2069](https://github.com/pjvds/tunl/commit/6e2069ae6f3becaea5b088971c4909b891bb49d3) - explicitly close target connection when proxy is done - [pjvds](https://github.com/pjvds)

[b59c11](https://github.com/pjvds/tunl/commit/b59c11dff20fad55645740dab0bf2061027b460f) - missing newline character in tunnel address printing - [pjvds](https://github.com/pjvds)

[1e8f79](https://github.com/pjvds/tunl/commit/1e8f7905a63bb4ee5e2e68cfc14a02867ac1697b) - address not printed after tunnel creation - [pjvds](https://github.com/pjvds)

[e185e6](https://github.com/pjvds/tunl/commit/e185e6bd423929c76226c6c742f75c1dd70d339f) - snap package not requesting correct plugs - [pjvds](https://github.com/pjvds)


### Features

[c1a397](https://github.com/pjvds/tunl/commit/c1a397f6ab0a53326860605490938bd2f96d15db) - switch from go-bin to pkger for embedding file and directories - [pjvds](https://github.com/pjvds)

[991d38](https://github.com/pjvds/tunl/commit/991d38998b2caaa71bd9d1ca9b18f32cc34289a6) - add new logo to assets - [pjvds](https://github.com/pjvds)

[58cada](https://github.com/pjvds/tunl/commit/58cada314a71300477e95bb64b134cfb12d73f4c) - serve explicit error page for downstream errors - [pjvds](https://github.com/pjvds)

[758383](https://github.com/pjvds/tunl/commit/75838382db8fcfe0f004b6bd437c8d6106464bd4) - improve command and flag usage descriptions - [pjvds](https://github.com/pjvds)

[d013ea](https://github.com/pjvds/tunl/commit/d013eacee40168555829c1ccdfd7ddc1c2af86a0) - skip TLS verification for local address by default - [pjvds](https://github.com/pjvds)

[bdfb8d](https://github.com/pjvds/tunl/commit/bdfb8dd103e5e93e251091c75492e7075376998c) - print public address to stdout and target to stderr - [pjvds](https://github.com/pjvds)

[b3b854](https://github.com/pjvds/tunl/commit/b3b854f67bc4b2c809dc79778a2d9bb75db540ef) - improve usage description of commands - [pjvds](https://github.com/pjvds)

[2fa41d](https://github.com/pjvds/tunl/commit/2fa41d0b8775752b1fc5000a5a2b67e7db553b95) - add docker command to create tunnel directly to a docker container - [pjvds](https://github.com/pjvds)

[d3adac](https://github.com/pjvds/tunl/commit/d3adacd0d713ba9c15d51a626386d6a6acef83ab) - add usage text to commands definition - [pjvds](https://github.com/pjvds)

[a4323b](https://github.com/pjvds/tunl/commit/a4323b71be4146b05687b5e68848fff0e1191f4f) - add bash and zsh auto completion - [pjvds](https://github.com/pjvds)

[3dbbf3](https://github.com/pjvds/tunl/commit/3dbbf3996e3619cc06a38a08964ff345bebdbd49) - host tcp tunnel on tcp subdomain - [pjvds](https://github.com/pjvds)

[dfd4ad](https://github.com/pjvds/tunl/commit/dfd4ad9421bfb90a87d9fb639e86c61d2ca38fa6) - support multiple certificates - [pjvds](https://github.com/pjvds)

[a3059c](https://github.com/pjvds/tunl/commit/a3059c407d1d4ade77a6dfc8d331bf00f0d68fa8) - support raw TCP tunnel - [pjvds](https://github.com/pjvds)

[d12de6](https://github.com/pjvds/tunl/commit/d12de6f98ff2fdf0b5231e2a4def08edea1b89de) - support and release binaries for arm and arm64 architecture - [pjvds](https://github.com/pjvds)

[4ec500](https://github.com/pjvds/tunl/commit/4ec5007c3b9f045eba775f55e09b8f9efc8f186e) - all command input errors are written to stderr instread of - [pjvds](https://github.com/pjvds)

[324a99](https://github.com/pjvds/tunl/commit/324a9912c5565a1f29314bca3b2218a9c9771636) - reconnect tunnel when connection get interupted - [pjvds](https://github.com/pjvds)

[d6328c](https://github.com/pjvds/tunl/commit/d6328ca2cb8ea4cb81090dddda7b69545e49d8f3) - prefix target url with schema if not exists - [pjvds](https://github.com/pjvds)

[34126a](https://github.com/pjvds/tunl/commit/34126aa4494207364577c7f1739967ffad67f993) - print tunnel explicitly with tunnel address and - [pjvds](https://github.com/pjvds)

[2942f8](https://github.com/pjvds/tunl/commit/2942f894d3b3a2efa277144dbd8d72b36a09c0bb) - make dir argument optional and validate - [pjvds](https://github.com/pjvds)

[cb5955](https://github.com/pjvds/tunl/commit/cb5955143b2937f4276b6c9ec196fd4f28a7e30c) - switch to distroless image - [pjvds](https://github.com/pjvds)

[95a8d4](https://github.com/pjvds/tunl/commit/95a8d4d110be9a845562030c701b47ca36ce00e6) - tunnel can be reclaimed after disconnect - [pjvds](https://github.com/pjvds)

[1664e0](https://github.com/pjvds/tunl/commit/1664e0e0a284b4c78be0ea25974ce398c63dc370) - serve favicon and meta data if not exists in file server location - [pjvds](https://github.com/pjvds)


### Miscellaneous Chores

[17c8d9](https://github.com/pjvds/tunl/commit/17c8d9b987826ca63e546b49bc935ba5f32eb387) - 0.22.0 - [pjvds](https://github.com/pjvds)

[335644](https://github.com/pjvds/tunl/commit/335644fcd90567b7f9afb11b4e02027376dca152) - run pkger via go generate instead of pipeline step - [pjvds](https://github.com/pjvds)

[1ac04b](https://github.com/pjvds/tunl/commit/1ac04b6c20f4c37e04089ca1433936b63f032557) - 0.22.0 - [pjvds](https://github.com/pjvds)

[7accb8](https://github.com/pjvds/tunl/commit/7accb8481be51e57dfef35c542691fcd0d34633c) - run pkger during release - [pjvds](https://github.com/pjvds)

[bc76b1](https://github.com/pjvds/tunl/commit/bc76b1cce1d937ed0fe1952eb7c6cd3349e242fc) - 0.21.0 - [pjvds](https://github.com/pjvds)

[b6c79b](https://github.com/pjvds/tunl/commit/b6c79b2502ec2dc6a3bd97312f04936e147469c7) - disable snapcrafts in release pipeline - [pjvds](https://github.com/pjvds)

[673de9](https://github.com/pjvds/tunl/commit/673de94282d69d6f7b10eacc7a60704be0cf2b29) - improve readme with minor changes - [pjvds](https://github.com/pjvds)

[055c7a](https://github.com/pjvds/tunl/commit/055c7ad658fb18062ea9aa57e56a51febbd2ea1b) - only require network plug for snapcrafts - [pjvds](https://github.com/pjvds)

[63cd73](https://github.com/pjvds/tunl/commit/63cd735c97f0c4fa9370f9851b5c5ed24defbc2c) - 0.20.0 - [pjvds](https://github.com/pjvds)

[3bc9d5](https://github.com/pjvds/tunl/commit/3bc9d51f68f9b60c58d8919fd67547ece192649a) - 0.19.2 - [pjvds](https://github.com/pjvds)

[4fa389](https://github.com/pjvds/tunl/commit/4fa389d4d67494bc272f1265fb86618b8c7de430) - 0.19.1 - [pjvds](https://github.com/pjvds)

[ac9e5a](https://github.com/pjvds/tunl/commit/ac9e5a358359caf6d29f0e4006bc44e0e2ead466) - 0.19.0 - [pjvds](https://github.com/pjvds)

[b320cd](https://github.com/pjvds/tunl/commit/b320cd0a464e28aa892891c26cfbc36f029444b0) - 0.18.0 - [pjvds](https://github.com/pjvds)

[7af2fc](https://github.com/pjvds/tunl/commit/7af2fc3ffde3644d10c2d5c54a37f7c238082241) - 0.17.2 - [pjvds](https://github.com/pjvds)

[3c66bc](https://github.com/pjvds/tunl/commit/3c66bc21faefb91340f81cb1afd3ce00413605fe) - 0.17.1 - [pjvds](https://github.com/pjvds)

[848367](https://github.com/pjvds/tunl/commit/84836745d7a7b33393027b9acc431c81d3adb489) - 0.17.0 - [pjvds](https://github.com/pjvds)

[c00e65](https://github.com/pjvds/tunl/commit/c00e656cc26092722427570ec82707d16db6c282) - make background of logo transparent - [pjvds](https://github.com/pjvds)

[949673](https://github.com/pjvds/tunl/commit/9496739fb6f5a915b98995f98d9a4f272a388e50) - 0.16.2 - [pjvds](https://github.com/pjvds)

[795544](https://github.com/pjvds/tunl/commit/795544985a9770a59013a67355b35693076bad0d) - lower package confinement to devmode - [pjvds](https://github.com/pjvds)

[0e5ac0](https://github.com/pjvds/tunl/commit/0e5ac057be5b4c2849a53a2506b369f082801de3) - 0.16.1 - [pjvds](https://github.com/pjvds)

[cd64cc](https://github.com/pjvds/tunl/commit/cd64cc7ad6f6b7e52f4e7aa1dee449446c9e768b) - shorten usage description - [pjvds](https://github.com/pjvds)

[1c6e06](https://github.com/pjvds/tunl/commit/1c6e06928f09ee9460c536475b1a96bc044cb5ae) - use zip archives for windows builds - [pjvds](https://github.com/pjvds)

[836ffe](https://github.com/pjvds/tunl/commit/836ffe1fcafbd3bdb397ca27ebc90348c9355703) - add apache 2.0 license file - [pjvds](https://github.com/pjvds)

[f4965b](https://github.com/pjvds/tunl/commit/f4965b9bdccaa41d120ca53787222624882bd3b9) - remove checksum for unused jwt-go version - [pjvds](https://github.com/pjvds)

[c917ec](https://github.com/pjvds/tunl/commit/c917ec3fbf72d0b6236775a24972308c2896cac5) - change release pipeline to publish releases to snap store - [pjvds](https://github.com/pjvds)

[d33e3f](https://github.com/pjvds/tunl/commit/d33e3fef554a5b05cf645c702d01aca160d1788a) - 0.10.0 - [pjvds](https://github.com/pjvds)

[2ec95d](https://github.com/pjvds/tunl/commit/2ec95def75f77dd9510f6dd68df7944e6e71a437) - fix .goreleaser format error - [pjvds](https://github.com/pjvds)

[ff7a9b](https://github.com/pjvds/tunl/commit/ff7a9b5588b9a49248691fc666182042144af0a8) - add favicon assets - [pjvds](https://github.com/pjvds)

[1906a9](https://github.com/pjvds/tunl/commit/1906a94a1b8c1afc003816a88da9048ff2db03be) - publish snapcraft package with correct permissions - [pjvds](https://github.com/pjvds)


### Refactoring

[4f6732](https://github.com/pjvds/tunl/commit/4f6732db0ef2d3ec3648f43092a9a22ca973f272) - move template to package and assets - [pjvds](https://github.com/pjvds)

[a086b3](https://github.com/pjvds/tunl/commit/a086b3b2426c1f5a02aa23cbc0b381e969ee1787) - move addressing to server package - [pjvds](https://github.com/pjvds)

[b30eb4](https://github.com/pjvds/tunl/commit/b30eb4966ffeb35e578d9546a4c6ada802ff2f64) - move all tunnel and networking logic to own package and - [pjvds](https://github.com/pjvds)

[4da164](https://github.com/pjvds/tunl/commit/4da16456b3213fed0f65bff4009d451702f64db7) - split commands and package logic - [pjvds](https://github.com/pjvds)


### Documentation

[18554f](https://github.com/pjvds/tunl/commit/18554ff67190a7db854f40be0102629d554840f0) - add docker container example - [pjvds](https://github.com/pjvds)

[07cca3](https://github.com/pjvds/tunl/commit/07cca3d87d17c3312771e9384c437a5d8d180d7c) - add a link to the latest binary release - [pjvds](https://github.com/pjvds)

[724a09](https://github.com/pjvds/tunl/commit/724a0972f3ed17998356925bcf1519b3b09fa68e) - switch to TLS for install script - [pjvds](https://github.com/pjvds)

[f6be3d](https://github.com/pjvds/tunl/commit/f6be3d5623ae012fd1cf74d93dd2ff96eabb72a1) - add curl and wget installation description - [pjvds](https://github.com/pjvds)

[bcb8f5](https://github.com/pjvds/tunl/commit/bcb8f5d007ea8159f4a531724287518c79cd529d) - shorten example headings - [pjvds](https://github.com/pjvds)

[b120e7](https://github.com/pjvds/tunl/commit/b120e7fbd6ceeecf3d953de6077f7f4ffc2d06cb) - shorting example headings - [pjvds](https://github.com/pjvds)

[f7c945](https://github.com/pjvds/tunl/commit/f7c9456eb5fe35b4206e17fb7c5c0b67baf231e5) - add docker to installation - [pjvds](https://github.com/pjvds)

[b3e1de](https://github.com/pjvds/tunl/commit/b3e1ded93aa8710854f7d5223420d12defc3932e) - short first example heading - [pjvds](https://github.com/pjvds)

[16993e](https://github.com/pjvds/tunl/commit/16993ed237a4c12a93bbe8e0920faa70ccb16289) - improve examples - [pjvds](https://github.com/pjvds)

[681994](https://github.com/pjvds/tunl/commit/6819941cb4584ce3b604a613965c0a89f37cdde3) - add README - [pjvds](https://github.com/pjvds)


- - -
## 0.22.0 - 2021-01-15


### Miscellaneous Chores

[335644](https://github.com/pjvds/tunl/commit/335644fcd90567b7f9afb11b4e02027376dca152) - run pkger via go generate instead of pipeline step - [pjvds](https://github.com/pjvds)

[1ac04b](https://github.com/pjvds/tunl/commit/1ac04b6c20f4c37e04089ca1433936b63f032557) - 0.22.0 - [pjvds](https://github.com/pjvds)

[7accb8](https://github.com/pjvds/tunl/commit/7accb8481be51e57dfef35c542691fcd0d34633c) - run pkger during release - [pjvds](https://github.com/pjvds)

[bc76b1](https://github.com/pjvds/tunl/commit/bc76b1cce1d937ed0fe1952eb7c6cd3349e242fc) - 0.21.0 - [pjvds](https://github.com/pjvds)

[b6c79b](https://github.com/pjvds/tunl/commit/b6c79b2502ec2dc6a3bd97312f04936e147469c7) - disable snapcrafts in release pipeline - [pjvds](https://github.com/pjvds)

[673de9](https://github.com/pjvds/tunl/commit/673de94282d69d6f7b10eacc7a60704be0cf2b29) - improve readme with minor changes - [pjvds](https://github.com/pjvds)

[055c7a](https://github.com/pjvds/tunl/commit/055c7ad658fb18062ea9aa57e56a51febbd2ea1b) - only require network plug for snapcrafts - [pjvds](https://github.com/pjvds)

[63cd73](https://github.com/pjvds/tunl/commit/63cd735c97f0c4fa9370f9851b5c5ed24defbc2c) - 0.20.0 - [pjvds](https://github.com/pjvds)

[3bc9d5](https://github.com/pjvds/tunl/commit/3bc9d51f68f9b60c58d8919fd67547ece192649a) - 0.19.2 - [pjvds](https://github.com/pjvds)

[4fa389](https://github.com/pjvds/tunl/commit/4fa389d4d67494bc272f1265fb86618b8c7de430) - 0.19.1 - [pjvds](https://github.com/pjvds)

[ac9e5a](https://github.com/pjvds/tunl/commit/ac9e5a358359caf6d29f0e4006bc44e0e2ead466) - 0.19.0 - [pjvds](https://github.com/pjvds)

[b320cd](https://github.com/pjvds/tunl/commit/b320cd0a464e28aa892891c26cfbc36f029444b0) - 0.18.0 - [pjvds](https://github.com/pjvds)

[7af2fc](https://github.com/pjvds/tunl/commit/7af2fc3ffde3644d10c2d5c54a37f7c238082241) - 0.17.2 - [pjvds](https://github.com/pjvds)

[3c66bc](https://github.com/pjvds/tunl/commit/3c66bc21faefb91340f81cb1afd3ce00413605fe) - 0.17.1 - [pjvds](https://github.com/pjvds)

[848367](https://github.com/pjvds/tunl/commit/84836745d7a7b33393027b9acc431c81d3adb489) - 0.17.0 - [pjvds](https://github.com/pjvds)

[c00e65](https://github.com/pjvds/tunl/commit/c00e656cc26092722427570ec82707d16db6c282) - make background of logo transparent - [pjvds](https://github.com/pjvds)

[949673](https://github.com/pjvds/tunl/commit/9496739fb6f5a915b98995f98d9a4f272a388e50) - 0.16.2 - [pjvds](https://github.com/pjvds)

[795544](https://github.com/pjvds/tunl/commit/795544985a9770a59013a67355b35693076bad0d) - lower package confinement to devmode - [pjvds](https://github.com/pjvds)

[0e5ac0](https://github.com/pjvds/tunl/commit/0e5ac057be5b4c2849a53a2506b369f082801de3) - 0.16.1 - [pjvds](https://github.com/pjvds)

[cd64cc](https://github.com/pjvds/tunl/commit/cd64cc7ad6f6b7e52f4e7aa1dee449446c9e768b) - shorten usage description - [pjvds](https://github.com/pjvds)

[1c6e06](https://github.com/pjvds/tunl/commit/1c6e06928f09ee9460c536475b1a96bc044cb5ae) - use zip archives for windows builds - [pjvds](https://github.com/pjvds)

[836ffe](https://github.com/pjvds/tunl/commit/836ffe1fcafbd3bdb397ca27ebc90348c9355703) - add apache 2.0 license file - [pjvds](https://github.com/pjvds)

[f4965b](https://github.com/pjvds/tunl/commit/f4965b9bdccaa41d120ca53787222624882bd3b9) - remove checksum for unused jwt-go version - [pjvds](https://github.com/pjvds)

[c917ec](https://github.com/pjvds/tunl/commit/c917ec3fbf72d0b6236775a24972308c2896cac5) - change release pipeline to publish releases to snap store - [pjvds](https://github.com/pjvds)

[d33e3f](https://github.com/pjvds/tunl/commit/d33e3fef554a5b05cf645c702d01aca160d1788a) - 0.10.0 - [pjvds](https://github.com/pjvds)

[2ec95d](https://github.com/pjvds/tunl/commit/2ec95def75f77dd9510f6dd68df7944e6e71a437) - fix .goreleaser format error - [pjvds](https://github.com/pjvds)

[ff7a9b](https://github.com/pjvds/tunl/commit/ff7a9b5588b9a49248691fc666182042144af0a8) - add favicon assets - [pjvds](https://github.com/pjvds)

[1906a9](https://github.com/pjvds/tunl/commit/1906a94a1b8c1afc003816a88da9048ff2db03be) - publish snapcraft package with correct permissions - [pjvds](https://github.com/pjvds)


### Documentation

[18554f](https://github.com/pjvds/tunl/commit/18554ff67190a7db854f40be0102629d554840f0) - add docker container example - [pjvds](https://github.com/pjvds)

[07cca3](https://github.com/pjvds/tunl/commit/07cca3d87d17c3312771e9384c437a5d8d180d7c) - add a link to the latest binary release - [pjvds](https://github.com/pjvds)

[724a09](https://github.com/pjvds/tunl/commit/724a0972f3ed17998356925bcf1519b3b09fa68e) - switch to TLS for install script - [pjvds](https://github.com/pjvds)

[f6be3d](https://github.com/pjvds/tunl/commit/f6be3d5623ae012fd1cf74d93dd2ff96eabb72a1) - add curl and wget installation description - [pjvds](https://github.com/pjvds)

[bcb8f5](https://github.com/pjvds/tunl/commit/bcb8f5d007ea8159f4a531724287518c79cd529d) - shorten example headings - [pjvds](https://github.com/pjvds)

[b120e7](https://github.com/pjvds/tunl/commit/b120e7fbd6ceeecf3d953de6077f7f4ffc2d06cb) - shorting example headings - [pjvds](https://github.com/pjvds)

[f7c945](https://github.com/pjvds/tunl/commit/f7c9456eb5fe35b4206e17fb7c5c0b67baf231e5) - add docker to installation - [pjvds](https://github.com/pjvds)

[b3e1de](https://github.com/pjvds/tunl/commit/b3e1ded93aa8710854f7d5223420d12defc3932e) - short first example heading - [pjvds](https://github.com/pjvds)

[16993e](https://github.com/pjvds/tunl/commit/16993ed237a4c12a93bbe8e0920faa70ccb16289) - improve examples - [pjvds](https://github.com/pjvds)

[681994](https://github.com/pjvds/tunl/commit/6819941cb4584ce3b604a613965c0a89f37cdde3) - add README - [pjvds](https://github.com/pjvds)


### Bug Fixes

[61806c](https://github.com/pjvds/tunl/commit/61806c2fcdff7629b6d8fe3c96629f9dfcecc87d) - incorrect exit result for tunnel accept - [pjvds](https://github.com/pjvds)

[36a202](https://github.com/pjvds/tunl/commit/36a202abb11e18e8fee586f0a3a27ea3667dd6ed) - autocomplete files not included in archives - [pjvds](https://github.com/pjvds)

[5ffd0f](https://github.com/pjvds/tunl/commit/5ffd0fdd5646419dd15e931e83701e1e6bb5a5f0) - autocomplete not included in release packages - [pjvds](https://github.com/pjvds)

[6e2069](https://github.com/pjvds/tunl/commit/6e2069ae6f3becaea5b088971c4909b891bb49d3) - explicitly close target connection when proxy is done - [pjvds](https://github.com/pjvds)

[b59c11](https://github.com/pjvds/tunl/commit/b59c11dff20fad55645740dab0bf2061027b460f) - missing newline character in tunnel address printing - [pjvds](https://github.com/pjvds)

[1e8f79](https://github.com/pjvds/tunl/commit/1e8f7905a63bb4ee5e2e68cfc14a02867ac1697b) - address not printed after tunnel creation - [pjvds](https://github.com/pjvds)

[e185e6](https://github.com/pjvds/tunl/commit/e185e6bd423929c76226c6c742f75c1dd70d339f) - snap package not requesting correct plugs - [pjvds](https://github.com/pjvds)


### Features

[c1a397](https://github.com/pjvds/tunl/commit/c1a397f6ab0a53326860605490938bd2f96d15db) - switch from go-bin to pkger for embedding file and directories - [pjvds](https://github.com/pjvds)

[991d38](https://github.com/pjvds/tunl/commit/991d38998b2caaa71bd9d1ca9b18f32cc34289a6) - add new logo to assets - [pjvds](https://github.com/pjvds)

[58cada](https://github.com/pjvds/tunl/commit/58cada314a71300477e95bb64b134cfb12d73f4c) - serve explicit error page for downstream errors - [pjvds](https://github.com/pjvds)

[758383](https://github.com/pjvds/tunl/commit/75838382db8fcfe0f004b6bd437c8d6106464bd4) - improve command and flag usage descriptions - [pjvds](https://github.com/pjvds)

[d013ea](https://github.com/pjvds/tunl/commit/d013eacee40168555829c1ccdfd7ddc1c2af86a0) - skip TLS verification for local address by default - [pjvds](https://github.com/pjvds)

[bdfb8d](https://github.com/pjvds/tunl/commit/bdfb8dd103e5e93e251091c75492e7075376998c) - print public address to stdout and target to stderr - [pjvds](https://github.com/pjvds)

[b3b854](https://github.com/pjvds/tunl/commit/b3b854f67bc4b2c809dc79778a2d9bb75db540ef) - improve usage description of commands - [pjvds](https://github.com/pjvds)

[2fa41d](https://github.com/pjvds/tunl/commit/2fa41d0b8775752b1fc5000a5a2b67e7db553b95) - add docker command to create tunnel directly to a docker container - [pjvds](https://github.com/pjvds)

[d3adac](https://github.com/pjvds/tunl/commit/d3adacd0d713ba9c15d51a626386d6a6acef83ab) - add usage text to commands definition - [pjvds](https://github.com/pjvds)

[a4323b](https://github.com/pjvds/tunl/commit/a4323b71be4146b05687b5e68848fff0e1191f4f) - add bash and zsh auto completion - [pjvds](https://github.com/pjvds)

[3dbbf3](https://github.com/pjvds/tunl/commit/3dbbf3996e3619cc06a38a08964ff345bebdbd49) - host tcp tunnel on tcp subdomain - [pjvds](https://github.com/pjvds)

[dfd4ad](https://github.com/pjvds/tunl/commit/dfd4ad9421bfb90a87d9fb639e86c61d2ca38fa6) - support multiple certificates - [pjvds](https://github.com/pjvds)

[a3059c](https://github.com/pjvds/tunl/commit/a3059c407d1d4ade77a6dfc8d331bf00f0d68fa8) - support raw TCP tunnel - [pjvds](https://github.com/pjvds)

[d12de6](https://github.com/pjvds/tunl/commit/d12de6f98ff2fdf0b5231e2a4def08edea1b89de) - support and release binaries for arm and arm64 architecture - [pjvds](https://github.com/pjvds)

[4ec500](https://github.com/pjvds/tunl/commit/4ec5007c3b9f045eba775f55e09b8f9efc8f186e) - all command input errors are written to stderr instread of - [pjvds](https://github.com/pjvds)

[324a99](https://github.com/pjvds/tunl/commit/324a9912c5565a1f29314bca3b2218a9c9771636) - reconnect tunnel when connection get interupted - [pjvds](https://github.com/pjvds)

[d6328c](https://github.com/pjvds/tunl/commit/d6328ca2cb8ea4cb81090dddda7b69545e49d8f3) - prefix target url with schema if not exists - [pjvds](https://github.com/pjvds)

[34126a](https://github.com/pjvds/tunl/commit/34126aa4494207364577c7f1739967ffad67f993) - print tunnel explicitly with tunnel address and - [pjvds](https://github.com/pjvds)

[2942f8](https://github.com/pjvds/tunl/commit/2942f894d3b3a2efa277144dbd8d72b36a09c0bb) - make dir argument optional and validate - [pjvds](https://github.com/pjvds)

[cb5955](https://github.com/pjvds/tunl/commit/cb5955143b2937f4276b6c9ec196fd4f28a7e30c) - switch to distroless image - [pjvds](https://github.com/pjvds)

[95a8d4](https://github.com/pjvds/tunl/commit/95a8d4d110be9a845562030c701b47ca36ce00e6) - tunnel can be reclaimed after disconnect - [pjvds](https://github.com/pjvds)

[1664e0](https://github.com/pjvds/tunl/commit/1664e0e0a284b4c78be0ea25974ce398c63dc370) - serve favicon and meta data if not exists in file server location - [pjvds](https://github.com/pjvds)


### Refactoring

[4f6732](https://github.com/pjvds/tunl/commit/4f6732db0ef2d3ec3648f43092a9a22ca973f272) - move template to package and assets - [pjvds](https://github.com/pjvds)

[a086b3](https://github.com/pjvds/tunl/commit/a086b3b2426c1f5a02aa23cbc0b381e969ee1787) - move addressing to server package - [pjvds](https://github.com/pjvds)

[b30eb4](https://github.com/pjvds/tunl/commit/b30eb4966ffeb35e578d9546a4c6ada802ff2f64) - move all tunnel and networking logic to own package and - [pjvds](https://github.com/pjvds)

[4da164](https://github.com/pjvds/tunl/commit/4da16456b3213fed0f65bff4009d451702f64db7) - split commands and package logic - [pjvds](https://github.com/pjvds)


- - -
## 0.22.0 - 2021-01-15


### Miscellaneous Chores

[7accb8](https://github.com/pjvds/tunl/commit/7accb8481be51e57dfef35c542691fcd0d34633c) - run pkger during release - [pjvds](https://github.com/pjvds)


### Refactoring

[4f6732](https://github.com/pjvds/tunl/commit/4f6732db0ef2d3ec3648f43092a9a22ca973f272) - move template to package and assets - [pjvds](https://github.com/pjvds)


### Features

[c1a397](https://github.com/pjvds/tunl/commit/c1a397f6ab0a53326860605490938bd2f96d15db) - switch from go-bin to pkger for embedding file and directories - [pjvds](https://github.com/pjvds)

[991d38](https://github.com/pjvds/tunl/commit/991d38998b2caaa71bd9d1ca9b18f32cc34289a6) - add new logo to assets - [pjvds](https://github.com/pjvds)

[58cada](https://github.com/pjvds/tunl/commit/58cada314a71300477e95bb64b134cfb12d73f4c) - serve explicit error page for downstream errors - [pjvds](https://github.com/pjvds)


### Documentation

[18554f](https://github.com/pjvds/tunl/commit/18554ff67190a7db854f40be0102629d554840f0) - add docker container example - [pjvds](https://github.com/pjvds)


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

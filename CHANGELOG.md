# Changelog

## 1.3.2 (2026-02-25)

Full Changelog: [v1.3.1...v1.3.2](https://github.com/oregister/openregister-go/compare/v1.3.1...v1.3.2)

### Bug Fixes

* allow canceling a request while it is waiting to retry ([6d10a38](https://github.com/oregister/openregister-go/commit/6d10a38e76bc69d73f3c6272d9e53370f6f7596a))


### Chores

* **internal:** move custom custom `json` tags to `api` ([cb1f3d6](https://github.com/oregister/openregister-go/commit/cb1f3d6bab23c933cc2b923a5a87df386c93163a))
* **internal:** remove mock server code ([0b2de1b](https://github.com/oregister/openregister-go/commit/0b2de1ba67282af282a2f859ffdc6d7c1ed05c7f))
* update mock server docs ([2301fe5](https://github.com/oregister/openregister-go/commit/2301fe59553331d0b5fd698674077098b2ffa726))

## 1.3.1 (2026-02-11)

Full Changelog: [v1.3.0...v1.3.1](https://github.com/oregister/openregister-go/compare/v1.3.0...v1.3.1)

### Bug Fixes

* **encoder:** correctly serialize NullStruct ([c422983](https://github.com/oregister/openregister-go/commit/c42298381a09a9e261d35af5eaa02a2e80ac2d42))


### Chores

* **internal:** codegen related update ([39c2d79](https://github.com/oregister/openregister-go/commit/39c2d7968a4dad44ec6ca498af5ab78a547733bc))

## 1.3.0 (2026-01-24)

Full Changelog: [v1.2.1...v1.3.0](https://github.com/oregister/openregister-go/compare/v1.2.1...v1.3.0)

### Features

* **client:** add a convenient param.SetJSON helper ([7b3cf6e](https://github.com/oregister/openregister-go/commit/7b3cf6e1bf5988631f4c5606738ee32078c90b35))
* **encoder:** support bracket encoding form-data object members ([2827f3d](https://github.com/oregister/openregister-go/commit/2827f3dd75e2d90aae53bbfdf47fca36378fd13c))


### Bug Fixes

* **docs:** add missing pointer prefix to api.md return types ([e2f9878](https://github.com/oregister/openregister-go/commit/e2f98786d4a20416b0a8b9399aa143af6bedfb94))
* **mcp:** correct code tool API endpoint ([834ffb2](https://github.com/oregister/openregister-go/commit/834ffb2cbcd4da0df4c10a3fe7d0906847d6ebef))
* rename param to avoid collision ([5312e75](https://github.com/oregister/openregister-go/commit/5312e754fa883dec88b2806bac2c88556beecdfc))
* skip usage tests that don't work with Prism ([4506e4d](https://github.com/oregister/openregister-go/commit/4506e4d84145e6174cad8d13afb54368e3c4e94a))


### Chores

* add float64 to valid types for RegisterFieldValidator ([d425d64](https://github.com/oregister/openregister-go/commit/d425d64972a7e0e48e9089badc9cee2d53f0b05d))
* elide duplicate aliases ([7d46ca9](https://github.com/oregister/openregister-go/commit/7d46ca99c2088db7bfa94c2afb9414819b5ea013))
* **internal:** codegen related update ([6a7b041](https://github.com/oregister/openregister-go/commit/6a7b041e5a4509efaa6eaca52a675a0801e70716))
* **internal:** codegen related update ([2ec520c](https://github.com/oregister/openregister-go/commit/2ec520c64827dbf077154853546b9a9cae554aca))
* **internal:** update `actions/checkout` version ([31b43a4](https://github.com/oregister/openregister-go/commit/31b43a484fcddbab1d59613a7e9ffbcd1160173e))


### Documentation

* prominently feature MCP server setup in root SDK readmes ([4967a3c](https://github.com/oregister/openregister-go/commit/4967a3ce3d618a31656191ac6e3c56b7b228c5d2))

## 1.2.1 (2025-11-12)

Full Changelog: [v1.2.0...v1.2.1](https://github.com/oregister/openregister-go/compare/v1.2.0...v1.2.1)

### Chores

* bump gjson version ([4491e3d](https://github.com/oregister/openregister-go/commit/4491e3d35bd622ac030cfd2c4faae308438867e0))
* **internal:** grammar fix (it's -&gt; its) ([9aaf9e4](https://github.com/oregister/openregister-go/commit/9aaf9e4ee87ff56bc825e7797155a8f5888e61c6))

## 1.2.0 (2025-10-16)

Full Changelog: [v1.1.0...v1.2.0](https://github.com/oregister/openregister-go/compare/v1.1.0...v1.2.0)

### Features

* **api:** manual updates ([2faaa13](https://github.com/oregister/openregister-go/commit/2faaa1312f854af5e415e81b604c2e63cf199366))
* **api:** manual updates ([1785d56](https://github.com/oregister/openregister-go/commit/1785d5679e3d076050ba5664991344ad560450e2))

## 1.1.0 (2025-10-02)

Full Changelog: [v1.0.3...v1.1.0](https://github.com/oregister/openregister-go/compare/v1.0.3...v1.1.0)

### Features

* **api:** add contact information ([a0b8232](https://github.com/oregister/openregister-go/commit/a0b823294ee189dc88654435397258a50932f8a8))


### Bug Fixes

* **api:** merged reports ([66a9602](https://github.com/oregister/openregister-go/commit/66a96029ff7592d5ebc4ec2deb186e119cff9e8e))
* bugfix for setting JSON keys with special characters ([1f9ab17](https://github.com/oregister/openregister-go/commit/1f9ab17fa9f03ae14716caf93274d469075f3daa))
* use slices.Concat instead of sometimes modifying r.Options ([f9ea371](https://github.com/oregister/openregister-go/commit/f9ea37121a637cdedb47fe31f2a1e48f1d527a8b))


### Chores

* bump minimum go version to 1.22 ([0f82abd](https://github.com/oregister/openregister-go/commit/0f82abd64a984572148bcf137d39d706fbc1bf10))
* do not install brew dependencies in ./scripts/bootstrap by default ([92dfa5f](https://github.com/oregister/openregister-go/commit/92dfa5f47c457731a5e106b8fd61a92f3bfb87e7))
* update more docs for 1.22 ([a6dd530](https://github.com/oregister/openregister-go/commit/a6dd5304ccdede380458d61ecf9ea29ac7631994))

## 1.0.3 (2025-09-09)

Full Changelog: [v1.0.2...v1.0.3](https://github.com/oregister/openregister-go/compare/v1.0.2...v1.0.3)

### Bug Fixes

* **internal:** unmarshal correctly when there are multiple discriminators ([632477b](https://github.com/oregister/openregister-go/commit/632477b57a2b946f7f719f60d54893eac65c0df9))


### Chores

* **internal:** codegen related update ([07ec8a3](https://github.com/oregister/openregister-go/commit/07ec8a3dac284babafaf1e5405f53bf72c88bead))

## 1.0.2 (2025-08-29)

Full Changelog: [v1.0.1...v1.0.2](https://github.com/oregister/openregister-go/compare/v1.0.1...v1.0.2)

### Features

* **api:** update via SDK Studio ([da2fa61](https://github.com/oregister/openregister-go/commit/da2fa618392b02b18152295c3bb964ec9163fec8))


### Bug Fixes

* close body before retrying ([042445b](https://github.com/oregister/openregister-go/commit/042445b9d2e1a1ee8492302c17bfa78fc46a0de3))

## 1.0.1 (2025-08-27)

Full Changelog: [v1.0.0...v1.0.1](https://github.com/oregister/openregister-go/compare/v1.0.0...v1.0.1)

### Features

* **api:** update via SDK Studio ([84aa280](https://github.com/oregister/openregister-go/commit/84aa28087dd5c3366864605260cb3c0ca12db3b3))

## 1.0.0 (2025-08-27)

Full Changelog: [v0.1.0-alpha.10...v1.0.0](https://github.com/oregister/openregister-go/compare/v0.1.0-alpha.10...v1.0.0)

### Features

* **api:** update via SDK Studio ([7531dd2](https://github.com/oregister/openregister-go/commit/7531dd294c273898eb3c2e4293fc4a94fb4b4d0d))
* **api:** update via SDK Studio ([169a80a](https://github.com/oregister/openregister-go/commit/169a80a5e438929327757400d0eadd7ccc4d919e))
* **api:** update via SDK Studio ([d0fdaee](https://github.com/oregister/openregister-go/commit/d0fdaeefa386f940890e2f15630cafb1244c9ff5))
* **api:** update via SDK Studio ([0c21d9c](https://github.com/oregister/openregister-go/commit/0c21d9c0f20cee64172d6c55e42d09b300d3fb89))
* **api:** update via SDK Studio ([4f77ac5](https://github.com/oregister/openregister-go/commit/4f77ac56686e9b1532f5a21b6f51bf37f29c7f7c))
* **api:** update via SDK Studio ([d4bbd79](https://github.com/oregister/openregister-go/commit/d4bbd79f81ef58b37b75388b70178168ef3fce4d))
* **api:** update via SDK Studio ([cd6d020](https://github.com/oregister/openregister-go/commit/cd6d0207564d039b1a995a757fc07cdb0f19d039))


### Chores

* **internal:** codegen related update ([b7d09b8](https://github.com/oregister/openregister-go/commit/b7d09b8e8401d6d118e5f7ad0d620d209ab4ab39))
* **internal:** update comment in script ([e807b53](https://github.com/oregister/openregister-go/commit/e807b536a6a183af5b835394db1c0166fadd9b69))
* update @stainless-api/prism-cli to v5.15.0 ([5133cc9](https://github.com/oregister/openregister-go/commit/5133cc9b7f27dd9dcd0c06056cc87b31058a9be8))

## 0.1.0-alpha.10 (2025-08-07)

Full Changelog: [v0.1.0-alpha.9...v0.1.0-alpha.10](https://github.com/oregister/openregister-go/compare/v0.1.0-alpha.9...v0.1.0-alpha.10)

### Features

* **api:** update via SDK Studio ([1bf804d](https://github.com/oregister/openregister-go/commit/1bf804d81eea01ff114fc7f24ef2fdde430ad36e))
* **api:** update via SDK Studio ([d06290a](https://github.com/oregister/openregister-go/commit/d06290a50933b266899b228c8cfb22c47bd76ae4))
* **api:** update via SDK Studio ([919ae8f](https://github.com/oregister/openregister-go/commit/919ae8f1a17f0c096b1a9c7c6e7db534c1cabec7))
* **api:** update via SDK Studio ([fe435c3](https://github.com/oregister/openregister-go/commit/fe435c34fc68ed302d619c20d7aa56628e5b9755))
* **api:** update via SDK Studio ([9dd3702](https://github.com/oregister/openregister-go/commit/9dd3702dfd9e5874ed048f9c9479081261dcc6f4))
* **api:** update via SDK Studio ([9e22185](https://github.com/oregister/openregister-go/commit/9e22185c1e6c282410c4f8e35c2843ed024fb551))
* **api:** update via SDK Studio ([351b67d](https://github.com/oregister/openregister-go/commit/351b67d8c7c4c76e83776014173d35ec4ad725ec))
* **api:** update via SDK Studio ([ffadfd8](https://github.com/oregister/openregister-go/commit/ffadfd882679e591a7686bcbcfafc5a4ad8dc145))
* **api:** update via SDK Studio ([735aeca](https://github.com/oregister/openregister-go/commit/735aecaa44a33de089f4bda1bc2519943840e026))
* **api:** update via SDK Studio ([5c81f59](https://github.com/oregister/openregister-go/commit/5c81f59cb92dc4d48a5cea142b083c6cfe3f34d0))
* **api:** update via SDK Studio ([ea890da](https://github.com/oregister/openregister-go/commit/ea890da1160562eb24449455d153c498441f0dea))
* **api:** update via SDK Studio ([fbfb60d](https://github.com/oregister/openregister-go/commit/fbfb60d05b7e9fddefe0a2ae6f727435d698ac78))
* **api:** update via SDK Studio ([75c41d8](https://github.com/oregister/openregister-go/commit/75c41d8c40561a0f6b9738a27a2cbce0902e5ee2))
* **api:** update via SDK Studio ([56d2486](https://github.com/oregister/openregister-go/commit/56d2486fda9bcbb28c30f1b42f06c5bc27aeae3c))
* **api:** update via SDK Studio ([a679486](https://github.com/oregister/openregister-go/commit/a679486eaaf3f95c6dc55366ee16d655095e09a6))
* **api:** update via SDK Studio ([749caa2](https://github.com/oregister/openregister-go/commit/749caa2541c5c49cb677b34a7d0de4a7ecc046c9))
* **client:** support optional json html escaping ([2a291f8](https://github.com/oregister/openregister-go/commit/2a291f86a59191a127f3deb5c60ddb5da287825d))

## 0.1.0-alpha.9 (2025-07-28)

Full Changelog: [v0.1.0-alpha.8...v0.1.0-alpha.9](https://github.com/oregister/openregister-go/compare/v0.1.0-alpha.8...v0.1.0-alpha.9)

### Features

* **api:** update via SDK Studio ([cbad9ec](https://github.com/oregister/openregister-go/commit/cbad9ecbded674a065bcc990ab22d597974a2a15))
* **api:** update via SDK Studio ([f8a174e](https://github.com/oregister/openregister-go/commit/f8a174e9f0c5af2f43d4294a13f54af15d611a10))

## 0.1.0-alpha.8 (2025-07-23)

Full Changelog: [v0.1.0-alpha.7...v0.1.0-alpha.8](https://github.com/oregister/openregister-go/compare/v0.1.0-alpha.7...v0.1.0-alpha.8)

### Features

* **api:** update via SDK Studio ([889a62e](https://github.com/oregister/openregister-go/commit/889a62e6d07717d532d3a5653c17aabaf3726d89))


### Bug Fixes

* **client:** process custom base url ahead of time ([3d3805a](https://github.com/oregister/openregister-go/commit/3d3805aef579b5bb05b4e23cd78d58b96f3b2889))
* don't try to deserialize as json when ResponseBodyInto is []byte ([5dcb8f1](https://github.com/oregister/openregister-go/commit/5dcb8f133eebbc50ab0522d083b2a9b75c277ab2))


### Chores

* **ci:** only run for pushes and fork pull requests ([7fb99ec](https://github.com/oregister/openregister-go/commit/7fb99eca2046bf09587e43c0d984932eab6075f8))
* **internal:** fix lint script for tests ([b0b6545](https://github.com/oregister/openregister-go/commit/b0b65450eaebab5f8275822d18242bffecf8f830))
* lint tests ([b7b4ad8](https://github.com/oregister/openregister-go/commit/b7b4ad8cf0ead498fb9a8a35d8c5f952a3b33c62))
* lint tests in subpackages ([9dfb96f](https://github.com/oregister/openregister-go/commit/9dfb96feebef60217b9b5ed09506f5c20dee6f4e))

## 0.1.0-alpha.7 (2025-06-24)

Full Changelog: [v0.1.0-alpha.6...v0.1.0-alpha.7](https://github.com/oregister/openregister-go/compare/v0.1.0-alpha.6...v0.1.0-alpha.7)

### Features

* **api:** update via SDK Studio ([999a871](https://github.com/oregister/openregister-go/commit/999a871e99585a9fcdec22e7cd9a5c4ded8a782e))

## 0.1.0-alpha.6 (2025-06-24)

Full Changelog: [v0.1.0-alpha.5...v0.1.0-alpha.6](https://github.com/oregister/openregister-go/compare/v0.1.0-alpha.5...v0.1.0-alpha.6)

### Features

* **api:** update via SDK Studio ([8129223](https://github.com/oregister/openregister-go/commit/8129223c54dccd56f42b0bbd54a863ac84433170))
* **api:** update via SDK Studio ([d744c4f](https://github.com/oregister/openregister-go/commit/d744c4f1ef945b4f89a28c6ed88461d8c072231e))
* **api:** update via SDK Studio ([2476f21](https://github.com/oregister/openregister-go/commit/2476f215d0604e9cdedd34e37fa79063ea06932a))
* **client:** add escape hatch for null slice & maps ([0bb7e0d](https://github.com/oregister/openregister-go/commit/0bb7e0d86ea696755e8491dcb4475eeded0c1338))


### Chores

* fix documentation of null map ([6e8a154](https://github.com/oregister/openregister-go/commit/6e8a1542dc39faff86f8232cb564dbe0a35f0d0b))

## 0.1.0-alpha.5 (2025-06-20)

Full Changelog: [v0.1.0-alpha.4...v0.1.0-alpha.5](https://github.com/oregister/openregister-go/compare/v0.1.0-alpha.4...v0.1.0-alpha.5)

### Features

* **api:** update via SDK Studio ([cd0bca8](https://github.com/oregister/openregister-go/commit/cd0bca843416780ca00e83c52a22dfe34f82e0e2))

## 0.1.0-alpha.4 (2025-06-20)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/oregister/openregister-go/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Features

* **api:** update via SDK Studio ([e0a9975](https://github.com/oregister/openregister-go/commit/e0a997505737deb803596410672ab2605cb0e5c7))
* **api:** update via SDK Studio ([eee534e](https://github.com/oregister/openregister-go/commit/eee534ef2a51d1f0503a45501efa4f0b52ce9832))
* **client:** add debug log helper ([761ea6f](https://github.com/oregister/openregister-go/commit/761ea6ffb0515c33ce241261805678b8071886b7))


### Chores

* **ci:** enable for pull requests ([4486c97](https://github.com/oregister/openregister-go/commit/4486c979fc49ffd32bc2b8b953ca41124a9decbf))

## 0.1.0-alpha.3 (2025-06-09)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/oregister/openregister-go/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* **api:** update via SDK Studio ([07c0edd](https://github.com/oregister/openregister-go/commit/07c0eddc3f1682087d1b1feceb915fe1bf7ea8fd))
* **api:** update via SDK Studio ([f5058fb](https://github.com/oregister/openregister-go/commit/f5058fb162816286b237cde28ab50481e3b4025d))
* **api:** update via SDK Studio ([9eb7899](https://github.com/oregister/openregister-go/commit/9eb7899f4e2299de9216614d5e5131cc03f63443))

## 0.1.0-alpha.2 (2025-06-09)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/oregister/openregister-go/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Features

* **api:** update via SDK Studio ([cabc290](https://github.com/oregister/openregister-go/commit/cabc290f4295199eae475bf7a89276fe5abddfd2))

## 0.1.0-alpha.1 (2025-06-07)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/oregister/openregister-go/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** update via SDK Studio ([f644e2f](https://github.com/oregister/openregister-go/commit/f644e2f9378bdb49e150443d9741fe466a9c7c93))
* **api:** update via SDK Studio ([9360416](https://github.com/oregister/openregister-go/commit/936041639f6452df736b1bdbd9ac3c884448afa2))


### Chores

* configure new SDK language ([f7177d3](https://github.com/oregister/openregister-go/commit/f7177d36b03639582a375b8edf9bb1da14c738ce))
* update SDK settings ([1fe6228](https://github.com/oregister/openregister-go/commit/1fe62285e8043a65dd86b5c4612a137bf938f047))

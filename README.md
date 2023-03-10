[![Slack](https://img.shields.io/badge/Slack-Join%20Direktiv-4a154b?style=flat&logo=slack)](https://join.slack.com/t/direktiv-io/shared_invite/zt-zf7gmfaa-rYxxBiB9RpuRGMuIasNO~g)

This repository contains the user interface for [direktiv](https://github.com/direktiv/direktiv).

<p align="center" style="width: 50%; height:50%">
  <img src="assets/images/ui.png" alt="direktiv ui">
</p>

# Setup development environment

- this repo requires a specific node version, to easily install the correct one, please install [nvm](https://github.com/nvm-sh/nvm)
- in the root directory, run `nvm use` to automatically switch to the required version
  - which will be read from `.nvmrc` file
  - please note that you always must run `nvm use` in every terminal session
- run `yarn` to install all dependencies
- create a `.env` file, copy the content from `.env.example` and change the variables to your needs
- run `yarn start` to start the dev server
- when you are using VSCode, make sure to install the recommended extensions for the best dev experience
  - VSCode should suggest you the extensions when you open this folder in VSCode, but you can also check them in the `.vscode/extensions.json` file

# Scripts you might want to run

- `yarn run storybook` opens the storybook and documentation
- `yarn run test` runs the tests in watch mode
- `yarn run check` runs all the checks that will run in CI:
  - `yarm run check:lint` runs the linter
  - `yarm run check:types` checks all types
  - `yarm run check:test` runs the tests in ci mode
  - `yarm run check:prettier` checks if all prettier formatting was applied (`yarn run prettier`)

# Code of Conduct

We have adopted the [Contributor Covenant](https://github.com/direktiv/.github/blob/master/CODE_OF_CONDUCT.md) code of conduct.

# Contributing

Any feedback and contributions are welcome. Read our [contributing guidelines](https://github.com/direktiv/.github/blob/master/CONTRIBUTING.md) for details.

# License

Distributed under the Apache 2.0 License. See `LICENSE` for more information.

# See Also

- The [direktiv.io](https://direktiv.io/) website.
- The direktiv [documentation](https://docs.direktiv.io/).
- The direktiv [blog](https://blog.direktiv.io/).
- The [Godoc](https://godoc.org/github.com/direktiv/direktiv) library documentation.

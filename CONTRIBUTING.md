## Contributing

### License

This project is licensed under the MIT License.

## Guidelines apply from main OpenFaaS repo

See guide for [FaaS](https://github.com/alexellis/faas/blob/master/CONTRIBUTING.md) here.

## Hacking on the faas-cli

## Installation / pre-requirements

* Docker

Install Docker because it is used to build Docker images if you create new functions.

* FaaS - deployed and live

This CLI can build and deploy templated functions, so it's best if you have FaaS started up on your laptop. Head over to http://docs.get-faas.com/ and get up and running with a sample stack in 60 seconds.

* Golang

> Here's how to install Go in 60 seconds.

* Grab Go 1.7.x from https://golang.org/dl/

Then after installing run this command or place it in your `$HOME/.bash_profile`

```bash
export GOPATH=$HOME/go
```

* Now clone / build `faas-cli`:

```
$ mkdir -p $GOPATH/src/github.com/alexellis/
$ cd $GOPATH/src/github.com/alexellis/
$ git clone https://github.com/alexellis/faas-cli
$ cd faas-cli
$ go get -d -v
$ go build
```

### How to update the `brew` formula

The `brew` formula for the faas-cli is part of the official [homebrew-core](https://github.com/Homebrew/homebrew-core/blob/master/Formula/faas-cli.rb) repo on Github. It needs to be updated for each subsequent release.

#### Simple version bumps

If the only change required is a version bump, ie no new tests, or changes to existing tested functionality or build steps, the `brew bump-formula-pr` command can be used to do everything (i.e. forking, committing, pushing) required to bump the version.

For example (supplying both the new version tag and its associated Git sha-256).

```
brew bump-formula-pr --strict faas-cli --tag=<version> --revision=<sha-256>
```

#### Changes requiring new/update tests/build steps

If a new release alters behaviour tested in the Brew Formula, adds new testable behaviors or alters the build steps then you will need to manually raise a PR with an updated Formula, the guidelines for updating brew describe the process in more detail:

https://github.com/Homebrew/homebrew-core/blob/master/CONTRIBUTING.md

After `brew edit` run the build and test the results:

```
$ brew uninstall --force faas-cli ; \
  brew install --build-from-source faas-cli ; \
  brew test faas-cli ; \
  brew audit --strict faas-cli
```

## Update the utility-script

Please raise a PR for the get.sh file held in this repository. It's used when people install via `curl` and `cli.openfaas.com`. The updated file then has to be redeployed to the hosting server.

## Developer DCO (re-iteration from referenced CONTRIBUTING guide)

### Sign your work

The sign-off is a simple line at the end of the explanation for a patch. Your
signature certifies that you wrote the patch or otherwise have the right to pass
it on as an open-source patch. The rules are pretty simple: if you can certify
the below (from [developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
1 Letterman Drive
Suite D4700
San Francisco, CA, 94129

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.

Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

Then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe.smith@email.com>

Use your real name (sorry, no pseudonyms or anonymous contributions.)

If you set your `user.name` and `user.email` git configs, you can sign your
commit automatically with `git commit -s`.

* Please sign your commits with `git commit -s` so that commits are traceable.

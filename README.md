Terraform CodeClimate Provider
==================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- [![Build Status](https://travis-ci.org/babbel/terraform-provider-codeclimate.svg?branch=master)](https://travis-ci.org/babbel/terraform-provider-codeclimate)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.11.13+
- [Go](https://golang.org/doc/install) 1.11.x+ (to build the provider plugin)

Building The Provider
---------------------
Clone repository to: `$GOPATH/src/github.com/babbel/terraform-provider-codeclimate`

```sh
$ mkdir -p $GOPATH/src/github.com/babbel; cd $GOPATH/src/github.com/babbel
$ git clone git@github.com:babbel/terraform-provider-codeclimate
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/babbel/terraform-provider-codeclimate
$ make build
```

Using the provider
----------------------

Currently the provider supports just Repository retrieval, based on the repository name.
It is used then as a data source.

```hcl
provider "codeclimate" {
  api_key = "${var.api_key}"
}

data "codeclimate_repository" "test" {
  repository_slug = "babbel/test"
}
```

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11.x+ is *required*). This provider works using Go Modules.

To compile the provider, run `go build -o terraform-provider-codeclimate`. This will build the provider and put the provider binary in the current directory.

```sh
$ make bin
...
$ terraform-provider-codeclimate
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ go test ./...
```

Github Releases
---------------------------
In order to push a release to Github the feature branch has to be merged into master and then a tag needs to be created with the version name of the provider e.g. **v0.0.1** and pushed.

```sh
git checkout master
git pull origin master
git tag v<semver>
git push origin master --tags
```

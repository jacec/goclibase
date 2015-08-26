# goclibase

#Go CLI Base Repo
[![wercker status](https://app.wercker.com/status/a11ad24642f15ab726057506900c2a76/m "wercker status")](https://app.wercker.com/project/bykey/a11ad24642f15ab726057506900c2a76)

Getting started "base" repo for a golang CLI application, using the great work by HashiCorp

Now with wercker support

If you're looking for a getting started repo for building a CLI based app in
golang, with testing, makefile and wercker CI then feel free to fork this.

##Building Locally (without Wercker)
```
git clone https://github.com/jacec/goclibase.git
cd goclibase
make
```

##Building Locally (with Wercker)
This assumes you have the wercker CLI installed and a reachable Docker instance
```
git clone https://github.com/jacec/goclibase.git
cd goclibase
wercker dev
```

##Building & Text Locally (with Wercker)
This assumes you have the wercker CLI installed and a reachable Docker instance
```
git clone https://github.com/jacec/goclibase.git
cd goclibase
wercker b
```

##Additional Information
Addition information for this repo is available at the following:

The CLI library used in most of HashiCorp's applications
[github.com/mitchellh/cli](github.com/mitchellh/cli)

Example project on how to build and link a wercker account to a golang project
[http://devcenter.wercker.com/quickstarts/building/golang.html](http://devcenter.wercker.com/quickstarts/building/golang.html)

More on wercker
[http://devcenter.wercker.com/index.html](http://devcenter.wercker.com/index.html)

Managing go dependencies with this library
[https://github.com/tools/godep](https://github.com/tools/godep)

Docker Machine (formally Boot2Docker) useful for local wercker CLI development
[https://docs.docker.com/installation/mac/](https://docs.docker.com/installation/mac/)

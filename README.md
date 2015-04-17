**jkl** is a static site generator written in [Go](http://www.golang.org),
based on [Jekyll](https://github.com/mojombo/jekyll)

[![Build Status](https://drone.io/github.com/sanbornm/jkl/status.png)](https://drone.io/drone/jkl/latest)

**This is a fork of [jkl](https://github.com/drone/jkl)** as the original is no longer maintained.

Notable changes in this fork:

* No support for AWS S3 syncing (do one thing and one thing well) use s3cmd or some other tool instead.
* Support for pretty urls
* Short descriptions with the &lt;!--more-&gt; tag
* Fix: display of dates
* Added urlencode template filter


Notable similarities between jkl and Jekyll:

* Directory structure
* Use of YAML front matter in Pages and Posts
* Availability of `site`, `content`, `page` and `posts` variables in templates
* Copies all static files into destination directory

Notable differences between jkl and Jekyll:

* Uses [Go templates](http://www.golang.org/pkg/text/template)
* Only supports YAML front matter in markup files
* No plugin support

Additional features:

* Deploy to S3

Sites built with jkl:

* ~~Drone.io Blog: http://blog.drone.io~~ (now moved to Hugo)
* ~~Drone.io Documentation: http://docs.drone.io~~ (now moved to Hugo)

--------------------------------------------------------------------------------

### Installation

In order to compile with `go build` you will first need to download
the following dependencies:

```
go get github.com/russross/blackfriday
go get launchpad.net/goyaml
go get github.com/howeyc/fsnotify
```
Once you have compiled `jkl` you can install with the following command:

```sh
sudo install -t /usr/local/bin jkl
```

If you are running x64 linux you can download and install the pre-compiled
binary:

### Usage

```
Usage: jkl [OPTION]... [SOURCE]

      --auto           re-generates the site when files are modified
      --base-url       serve website from a given base URL
      --source         changes the dir where Jekyll will look to transform files
      --destination    changes the dir where Jekyll will write files to
      --server         starts a server that will host your _site directory
      --server-port    changes the port that the Jekyll server will run on
  -v, --verbose        runs Jekyll with verbose output
  -h, --help           display this help and exit

Examples:
  jkl                  generates site from current working dir
  jkl --server         generates site and serves at localhost:4000
  jkl /path/to/site    generates site from source dir /path/to/site

```

### Auto Generation

If you are running the website in server mode, with the `--server` flag, you can
also instruct `jkl` to auto-recompile you website by adding the `--auto` flag.

NOTE: this feature is only available on Linux and OSX

### Deployment

Use rsync or s3cmd to sync files to remote server.

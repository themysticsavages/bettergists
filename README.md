# bettergists
An introduction to Golang for me and a featured page for [Sulphurous Gists](https://sulphurous.cf/gists) for you.

## Usage
Make sure you have Golang installed.

```bash
$ git clone https://github.com/themysticsavages/bettergists
$ cd bettergists
$ # set the go111module variable to auto
$ go env -w GO111MODULE=auto
$ go build index.go
$ ./index
YYYY/MM/DD HH:MM:SS Website is up
```

## How does it work?
Bettergists actually uses external stylesheets and scripts, which means you can kinda have the same experience on any other website. 

When one goes to the root page, the latest gist information is taken from Sulphurous' API and slapped onto a section called `'The latest'`. When you search for a user, gist ID, or the gist title itself, the same information in `'The latest'` section is iterated through, matched with the search, and put right below the search box. This all done on the main page, which makes the app is more simplified.

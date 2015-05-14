Forward
=========

Forward is a simple utility written in Go to provide a sharable URL to your local development site. This can be handy for sharing some changes made on your machines local development environment without having to deploy to staging or for testing webhooks.

## Requirements

To do this, you are going to need an ssh server that is routable on the Internet. If you are like me, I happen to have a lightly used Digital Ocean server that does just the trick. Currently, forward only uses public key authentication, so you are going to need to make sure you have your public RSA key in your server's authorized_keys file. If your not sure if you have your public key on your server, generally a command like so will get it added:

`cat ~/.ssh/id_rsa.pub | ssh user@yourserver.com "cat >> ~/.ssh/authorized_keys"`

Also if you have a look at nginx-site, it is a simple nginx config for setting up a reverse proxy.

## Examples

assuming you have a web server you are writing on port 8080 and you are forwarding request from your reverse proxy (aka nginx) to port 5000 on your remote server, you would run forward like so:

`forward --user=mr_awesome --server=myserver.com:22 --local=localhost:8080 --remote=localhost:5000`

An viola, request from myserver.com should be sent your local development site!

## Installation

`go get github.com/acmacalister/forward`

## TODO

- [ ] Pre-built binaries for all the OSes.

## License

forward is licensed under Apache v2.

## Questions, Problems, Fixes?

Feel free to open a Github issue. I'm usually pretty prompt in responding.

## Contact

### Austin Cherry
* https://github.com/acmacalister
* http://twitter.com/acmacalister
* http://austincherry.me
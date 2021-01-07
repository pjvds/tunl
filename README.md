# Tunl

With Tunl you can expose services on your localhost to the public via a fast and secure tunnel.

## Want to share a local http process running at localhost:3000?

```
$ tunl http localhost:3000
https://red-fox.tunl.es/ -> http://localhost:3000
```

## Want to share a directory or some files with a client or a friend? 

```
$ tunl files /the/directory
https://red-fox.tunl.es/ -> /the/directory
```

## Want to make a private MySQL instance available for a remote machine?

```
$ tunl tcp localhost:3306
tunl.es:89311 -> localhost:3306
```

# Installation

Download a prebuild release and put it somewhere in your `$PATH`.

## Arch Linux

```
yay tunl
```

## Ubuntu

```
sudo snap install tunl --edge
```

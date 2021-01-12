# Tunl

With Tunl you can expose services on your localhost to the public via a fast and secure tunnel.

## expose local http process publicly

```
$ tunl http localhost:3000
https://red-fox.tunl.es/ -> http://localhost:3000
```

## share a local directory publicly

```
$ tunl files /the/directory
https://red-fox.tunl.es/ -> /the/directory
```

## expose a local MySQL instance publicly

```
$ tunl tcp localhost:3306
tunl.es:89311 -> localhost:3306
```

# Installation

Download a prebuild release and put it somewhere in your `$PATH`.

## Basic (curl)

```
curl -fsSL http://get.tunl.es | sh
```

## Basic (wget)

```
wget -qO- http://get.tunl.es | sh
```

## Arch Linux

```
yay tunl
```

## Ubuntu

```
sudo snap install tunl --edge
```

## Docker

```
docker run --rm --network=host -v $PWD:$PWD tunl/tunl
```

**warning**: this maps the current working directory and means files outside the current working directory cannot be shared

# PRPL server example

A basic Polymer Starter Kit with a few slight optimizations to test the [prpl-server](https://github.com/CaptainCodeman/prpl-server-go)

## Requirements

* Polymer-cli
* AppEngine SDK for Go
* PRPL-Server-Go

    go get -i github.com/captaincodeman/prpl-server-go

## Usage

First build the client side files

    polymer build

### Symlink build folder to server

From server folder, run:

    ln -s ../build static

### Generate app.yaml

Generate prpl-server config:

    prpl-config --root static \
                --config polymer.json \
                --static-version 20170815 \
                --project-id captain-codeman \
                --project-service prpl \
                --project-version example \
                --template app.yaml.tmpl > app.yaml

### Run

From server folder, run:

    goapp serve

Visit http://localhost:8080

NOTE: No http/2 server push happens but the http Link headers should be added to the responses.

### Deploy

From server folder, run:

    goapp deploy .

NOTE: you won't be able to deploy unless you change the project id to your own. Alternatively, you can remove these and use the `gcloud` command (see server rep for details)

### Standalone

The prpl-server can also be run standalone:

    prpl-server --root static --config static/polymer.json

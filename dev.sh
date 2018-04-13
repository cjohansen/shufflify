#!/bin/bash

export clientId=$(jq -cr '.spotifyClientId' config/dev.json)
export clientSecret=$(jq -cr '.spotifyClientSecret' config/dev.json)
$GOPATH/bin/modd

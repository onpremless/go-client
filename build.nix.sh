#!/usr/bin/env nix-shell
#! nix-shell -i bash --pure
#! nix-shell -p go openapi-generator-cli
#! nix-shell -I nixpkgs=https://github.com/NixOS/nixpkgs/archive/19cbff58383a4ae384dea4d1d0c823d72b49d614.tar.gz

exec ./build.sh

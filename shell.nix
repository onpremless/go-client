{ pkgs ? import <nixpkgs> {} }:
 pkgs.mkShell {
   packages = with pkgs; [
     go
     openapi-generator-cli
     wget
   ];
}


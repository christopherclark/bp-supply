# Buildpack Supply

bp-supply is a CF plugin to create supply buildpacks.

Below is an example usage

```
go build -o bp-supply main.go
cf install-plugin -f bp-supply
cf bp-supply --path=fixtures/simple --version=1.2.3 Simple simple_supply_buildpack.zip
```

# Stuff Products
## About
Stuff is a **fake store** is a microservices-based demo application that illustrates the use of [CF Container Networking](https://github.com/cloudfoundry-incubator/netman-release).

The app uses [Amalgam8](http://amalagam8.io) for [sidecar](https://www.amalgam8.io/docs/sidecar) and Service Discovery / Routing [control plane](https://www.amalgam8.io/docs/control-plane).

This project is used by the [store front app](https://github.com/markstgodard/stuff), please see this project for details on prerequisites and instructions deploying the store.

## Configuration
The scripts in this example uses this [config file](./scripts/cf.cfg) to configure CF domains, app names, etc.
If you wish to use the scripts to deploy the demo apps, please change the values to match your target environment.
The defaults assume [bosh-lite](https://github.com/cloudfoundry/bosh-lite) and that you already are targeting a org and space.

### Deploy app

You may either run this script:
```sh
./scripts/deploy.sh
```

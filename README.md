# Service Fabrik Bosh Release for Cloud Foundry

A bosh release to deploy Service Fabrik which provisions service instances as Docker containers and Bosh deployments.

## 3rd Party
This BOSH release is based on/forked from and therefore includes sources from the following 3rd party BOSH release: https://github.com/cloudfoundry-community/docker-boshrelease, [Apache License Version 2.0](https://github.com/cloudfoundry-community/docker-boshrelease/blob/master/LICENSE).

## Local Development Setup

Please follow the instructions for the [broker](https://github.com/sap/service-fabrik-broker)
on how to set up a Bosh Lite with Cloud Foundry.

### Cloning the Repository

As a service developer you may have your own service bosh release. If not or you want to first deploy the Service Fabrik without modifications, you have to upload the blueprint service bosh release as a prerequisite. The blueprint service  is an/our examplary service specifically created for demonstration and testing purposes:

Assuming your working directory is ~/git:
```shell
cd ~/git
git clone https://github.com/sap/service-fabrik-blueprint-boshrelease
cd blueprint-boshrelease
bosh upload release $(ls -1rv releases/blueprint/blueprint-*.yml | head -1)
```

And now let's clone the actual Service Fabrik bosh release:
```shell
cd ~/git
git clone https://github.com/sap/service-fabrik-boshrelease
cd service-fabrik-boshrelease
```


### Deploy from Sources

If you need to modify the sources (beyond the example manifest):
```shell
./update
bosh deployment examples/bosh-lite-manifest.yml
bosh create release --force && bosh upload release && bosh -n deploy # repeat these commands on every change
```

### Deploy from Release

If you do not need to modify the sources (beyond the example manifest):
```shell
bosh deployment examples/bosh-lite-manifest.yml
bosh upload release $(ls -1rv releases/service-fabrik/service-fabrik-*.yml | head -1) && bosh -n deploy
```

### Register the Broker

You have to do this only once or whenever you modify the catalog. Then of course, use `update-service-broker` instead of `create-service-broker`.

* Registration
```shell
cf create-service-broker service-fabrik-broker broker secret https://service-fabrik-broker.bosh-lite.com/cf # route registered for broker
cf service-brokers # should show the above registered service broker
curl -skH "X-Broker-Api-Version: 2.9" https://broker:secret@service-fabrik-broker.bosh-lite.com/cf/v2/catalog | jq -r ".services[].name" | xargs -L 1 -I {} cf enable-service-access {}
cf service-access # should show all services as enabled, cf marketplace should show the same
```
* In order for the broker to provision bosh director based services, the releases used in the service manifest templates must be manually uploaded to the targetted bosh director (if you have followed this guide you have already uploaded the blueprint release)

### Run a Service Lifecycle

You will need a Cloud Foundry application, let's call it `my-app` (see below). If you have no specific one, you can use our [blueprint-app](https://github.com/sap/service-fabrik-blueprint-app).

```shell
cf create-service blueprint v1.0-container my-service
cf bind-service my-app my-service
# take a look at the generated binding with cf env my-app
cf restart my-app # do this a.) to make binding information available in environment of the app and b.) to activate the security group created with the service
# verify the application sees the service; if you have deployed the above app, run curl -skH "Accept: application/json" "https://my-app.bosh-lite.com/test"
cf unbind-service my-app my-service
cf delete-service -f my-service
```

## License

This project is licensed under the Apache Software License, v. 2 except as noted otherwise in the [LICENSE](LICENSE) file.

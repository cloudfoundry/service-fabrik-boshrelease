[![CLA assistant](https://cla-assistant.io/readme/badge/SAP/service-fabrik-boshrelease)](https://cla-assistant.io/SAP/service-fabrik-boshrelease)

# Service Fabrik Bosh Release for Cloud Foundry

A bosh release to deploy the [Service Fabrik](https://github.com/SAP/service-fabrik-broker) which provisions service instances as Docker containers and Bosh deployments.

## License

This project is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file.

## 3rd Party
This BOSH release is based on/forked from and therefore includes sources from the following 3rd party BOSH release: https://github.com/cloudfoundry-community/docker-boshrelease, [Apache License Version 2.0](https://github.com/cloudfoundry-community/docker-boshrelease/blob/master/LICENSE).

## Local Development Setup

Please follow the instructions for the [broker](https://github.com/SAP/service-fabrik-broker)
on how to set up a Bosh Lite with Cloud Foundry. Then clone this repo.

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
* In order for the broker to provision bosh director based services, the releases used in the service manifest templates must be manually uploaded to the targetted bosh director

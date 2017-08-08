# Service Fabrik Bosh Release for Cloud Foundry

A bosh release to deploy Service Fabrik which provisions service instances as Docker containers and Bosh deployments.

## License

This project is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file.

## 3rd Party
This BOSH release is based on/forked from and therefore includes sources from the following 3rd party BOSH release: https://github.com/cloudfoundry-community/docker-boshrelease, [Apache License Version 2.0](https://github.com/cloudfoundry-community/docker-boshrelease/blob/master/LICENSE).

## Local Development Setup

Please follow the instructions for the [broker](https://github.com/SAP/service-fabrik-broker)
on how to set up a Bosh Lite with Cloud Foundry.

### Cloning the Repository

As a service developer you may have your own service bosh release. If not or you want to first deploy the Service Fabrik without modifications, you have to upload the blueprint service bosh release as a prerequisite. The blueprint service  is an/our examplary service specifically created for demonstration and testing purposes:

Assuming your working directory is ~/git:
```shell
cd ~/git
git clone https://github.com/sap/service-fabrik-blueprint-boshrelease
cd service-fabrik-blueprint-boshrelease
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
bosh create release --force && bosh upload release # repeat these commands on every change
cd templates
./make_manifest warden
bosh -n deploy
```

### Deploy from Release

If you do not need to modify the sources (beyond the example manifest):
```shell
bosh upload release $(ls -1rv releases/service-fabrik/service-fabrik-*.yml | head -1)
cd templates
./make_manifest warden
bosh -n deploy
```

### Register the Broker

You have to do this only once or whenever you modify the catalog. Then of course, use `update-service-broker` instead of `create-service-broker`.

* Registration
```shell
cf create-service-broker service-fabrik-broker broker secret https://10.244.4.2:9293/cf # route registered for broker
cf service-brokers # should show the above registered service broker
curl -skH "X-Broker-Api-Version: 2.9" https://broker:secret@10.244.4.2:9293/cf/v2/catalog | jq -r ".services[].name" | xargs -L 1 -I {} cf enable-service-access {}
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

### Integrating a new service
#### Integrating a new docker service
There are couple steps one will have to follow to integrate a new docker service. Following are the steps to be followed.

1. Create the catalog.yml file for your service which has details of the service catalog. We have provided an example of a Redis service catalog file redis.yml.erb under services folder.

1. Now, if you want to add this service, edit templates/service_definitions.yml and add the following line.
    ```
      - <%= JSON.pretty_generate(YAML.load(ERB.new(File.read("../services/redis.yml.erb")).result)) %>
    ```
1. generate the manifest again and redeploy service-fabrik.

    ```
    cd templates
    ./make_manifest warden
    bosh deploy
    ```
1. Update the broker registration and enable service access.

    ```
    cf update-service-broker service-fabrik-broker broker secret https://10.244.4.2:9293/cf # route registered for broker
    curl -skH "X-Broker-Api-Version: 2.9" https://broker:secret@10.244.4.2:9293/cf/v2/catalog | jq -r ".services[].name" | xargs -L 1 -I {} cf enable-service-access {}
    ```
1. Now cf market place should show the newly added service, and one should be able to create as well.
    ```
        $ cf m
    Getting services from marketplace in org dev / space broker as admin...
    OK

    service     plans                                                                                 description
    blueprint   v1.0-container, v1.0-container-large, v1.0-dedicated-xsmall*, v1.0-dedicated-large*   Blueprint service for internal development, testing, and documentation purposes of the Service Fabrik
    redis       v3.0-container                                                                        Redis in-memory data structure store

    * These service plans have an associated cost. Creating a service instance will incur this cost.

    TIP:  Use 'cf marketplace -s SERVICE' to view descriptions of individual plans of a given service.

    $ cf cs redis v3.0-container r1c
    Creating service instance r1c in org dev / space broker as admin...
    OK

    $ cf csk r1c r1ck
    Creating service key r1ck for service instance r1c as admin...
    OK

    $ cf service-key r1c r1ck
    Getting key r1ck for service instance r1c as admin...

    {
     "hostname": "10.244.4.3",
     "password": "tAzsqs4qqx3VSs32",
     "port": "32843",
     "ports": {
      "6379/tcp": "32843"
     }
    }
    ```


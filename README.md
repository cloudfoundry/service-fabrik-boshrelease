# Service Fabrik Bosh Release for Cloud Foundry

A bosh release to deploy Service Fabrik which provisions service instances as Docker containers and Bosh deployments.

## 3rd Party
This BOSH release is based on/forked from and therefore includes sources from the following 3rd party BOSH release: https://github.com/cloudfoundry-community/docker-boshrelease, [Apache License Version 2.0](https://github.com/cloudfoundry-community/docker-boshrelease/blob/master/LICENSE).

# Table of Contents
1. [Local Development Setup](#local-development-setup)
2. [Cloning the Repository](#cloning-the-repository)
3. [Deploying ApiServer](#deploying-apiServer)
4. [Deploying Service-Fabrik](#deploying-service-fabrik)
5. [Register the Broker](#register-the-broker)
6. [Run a Service Lifecycle](#run-a-service-lifecycle)
7. [How to obtain Support](#how-to-obtain-support)

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

### Deploying ApiServer
Apiserver is a prerequisite for running Service-Fabrik-Broker. It can be deployed on boshlite using provided manifest [file](https://github.com/cloudfoundry-incubator/service-fabrik-boshrelease/blob/master/templates/apiserver.yml)
Prerequisites:
apiserver has dependency on cfcr-etcd release, which have to be first uploaded on bosh before deploying apiserver.
```shell
cd ~/git
git clone https://github.com/cloudfoundry-incubator/cfcr-etcd-release
cd cfcr-etcd-release
bosh -e upload-release
cd ~/git/service-fabrik-boshrelease
bosh –e bosh -d apiserver deploy --vars-store=templates/vars-store.yml templates/apiserver.yml
```
Useful prerequisites: When working with apiserver kubectl comes in handy. To configure kubectl locally, please follow steps mentioned in this [doc](https://github.com/cloudfoundry-incubator/service-fabrik-boshrelease/wiki/Configuring-kubectl-for-local-development)
### Deploying Service-Fabrik

#### Creating and uploading release:

  ##### Deploying from Sources

  If you need to modify the sources (beyond the example manifest):
  ```shell
  ./scripts/update
  bosh create release --force && bosh upload release # repeat these commands on every change
  ```

  #### Deploying from Release

  If you do not need to modify the sources (beyond the example manifest):
  ```shell
  bosh upload release $(ls -1rv releases/service-fabrik/service-fabrik-*.yml | head -1)
  ```

#### Uploading the cloud-config:
Then, we need to upload the cloud-config required for service-fabrik on bosh.

For bosh-lite, you can upload cloud-config in the following manner:
```shell
cd templates
bosh –e bosh upload-cloud-config cloud-config-boshlite.yml
```

For AWS, we need to update the vars-files for the cloud-config. 
The vars file to be edited is `cloud-config-aws-vars.yml`. It can be found in the `templates` directory.
Once the vars file is filled with proper details, the cloud-config can be uploaded:
```shell
cd templates
bosh –e bosh upload-cloud-config --vars-store=cloud-config-aws-vars.yml cloud-config-aws.yml
```

#### Deployment:
Once the release and cloud-config has been uploaded to bosh, we can deploy service-fabrik.

First we need to generate the services catalog file. Run the render script present in the templates directory. In case deployment actions need to be added, actions should be placed inside `services/actions` directory.
```shell
./render
```

Service-fabrik uses various ops-files which provide the configs required during deployment.
The ops-file specific to AWS is `ops-file-aws.yml`. Similary for boshlite, it is `ops-file-boshlite.yml`.
Normally, this shouldn't require any modification.

The common ops-file used for any IaaS is `ops-file.yml`.
This file may require modifications depending on the user's landscape configurations.
An example of such an use-case:
Suppose you need to change the cf's username. Let's assume the new username should be `user`.
The entry present in the ops-file is:
```
- type: replace
  path: /instance_groups/name=broker/jobs/name=service-fabrik-broker/properties/cf/username
  value: admin
```
The `value` parameter here can be modified as: `value: user`.

Service-fabrik also uses vars-store file for its deployment.
This file contains mostly essential certificates and some configurations which are likely to be specific to an user's environment.
The required certificates can be edited/replaced and other configurations like, `cf_url` and `broker_host` can be edited.

Once all the changes are made, service-fabrik can be deployed:

for AWS:
```shell
bosh –e bosh deploy -o services.yml -o actions.yml -o ops-file.yml -o ops-file-aws.yml --vars-store=vars-store.yml deployment.yml
```

for boshlite:
```shell
bosh –e bosh deploy -o services.yml -o actions.yml -o ops-file.yml -o ops-file-boshlite.yml --vars-store=vars-store.yml deployment.yml
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

1. Create a Catalog file for your service which has details of the service catalog. We have provided an example of a Redis service catalog file redis.yml.erb under services folder.

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

#### Integrating a new BOSH based service
Integration of BOSH based service also is similar to Docker based services, except in this case, in order to deploy BOSH based services, bosh release and bosh manifest information is also needed for service fabrik to deploy it on BOSH.

A Catalog file for your service which has details of the service catalog similar to [blueprint.yml.erb](services/blueprint.yml.erb). In this file, we need to provide the bosh release link as shown [here](https://github.com/SAP/service-fabrik-boshrelease/blob/master/services/blueprint.yml.erb#L117)
The manifest is added as ejs file similar to [blueprint-manifest.yml.ejs](services/blueprint-manifest.yml.ejs) file and added in the catalog file as base64 file, as shown [here](https://github.com/SAP/service-fabrik-boshrelease/blob/master/services/blueprint.yml.erb#L115).


service manifest ejs file provides the skeleton of a deployment manifest. If we look at https://github.com/SAP/service-fabrik-boshrelease/blob/master/services/blueprint-manifest.yml.ejs, it contains update section, job section, and properties section. Other things like network and all are taken care of by service fabrik. Network section can be accessed using the following code, as shown in blueprint example.

    
    <%
	const net = spec.networks[0];
	%>
    

Similarly, resource pool can be accessed in the following way:
  
      
      <%= `${p('resource_pool')}_${net.az}` %>
      
      
and disk pool in the following way:

      
	  <%= p('disk_pool') %>
      
  

Secure random numbers can be generated and used for credentials of service, as shown in blueprint, in the following way

    
    properties.blueprint = {
	    admin: {
	      username: SecureRandom.hex(16),
	      password: SecureRandom.hex(16)
	    }
    

If any parameters are passed during service creation, service fabrik directly passes it and it can be access here in the following way.

    
    <% if (spec.parameters.myproperty) { %>
          properties:
              myproperty: <%= spec.parameters.myproperty  %>
        <% } %>
    

Once, these two files are created, same steps as of docker service integration can be followed to deploy BOSH based service as well.

## How to obtain support

If you need any support, have any question or have found a bug, please report it in the [GitHub bug tracking system](https://github.com/sap/service-fabrik-boshrelease/issues). We shall get back to you.

## License

This project is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file.

## Reliable, Scalable Redis on Openshift

The following document describes the deployment of a reliable, multi-node Redis on Openshift.  It deploys a master with replicated slaves, as well as replicated redis sentinels which are use for health checking and failover.

### Prerequisites

This example assumes that you have a Openshift cluster installed and running, and that you have installed the ```oc``` command line tool somewhere in your path.


### Run 

    #make sure you have base image available
    oc create -f https://raw.githubusercontent.com/mjudeikis/redis-openshift/master/openshift/is-base.yaml -n openshift
    #create all components
    oc create -f https://raw.githubusercontent.com/mjudeikis/redis-openshift/master/list.yaml
    #start build and watch 
    oc start-build redis-build

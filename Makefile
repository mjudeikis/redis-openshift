build:
	cd image && podman build . -t quay.io/${USER}/redis-openshift

build-example:
	cd example && go build

build-example-image:
	cd example && podman build . -t quay.io/${USER}/redis-openshift-example

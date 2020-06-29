# solar-system

When the sun is created as a custom resource, the controller creates a deployment for the energy source and a service of NodePort, so it can be accessed outside of the Kubernetes cluster.


in case the deleting namespace stuck at "Terminating" state, please run:

kubectl get namespace "stucked-namespace" -o json \
            | tr -d "\n" | sed "s/\"finalizers\": \[[^]]\+\]/\"finalizers\": []/" \
                        | kubectl replace --raw /api/v1/namespaces/stucked-namespace/finalize -f -

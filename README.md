# Kube watch
Simple tool that watch on events from your Kubernetes cluster and push them as webhooks

# Install 
`go get github.com/olsynt/kube-watch`

# Run
Get you test hebhook url from https://webhook.site
Run `kube-watch run --url {url}` to start watching on events across all namespaces from you current-context in `~/.kube/config`

## More functionallity
`kube watch run --help`
```
   --kube-config value        (default: "/home/olsynt/.kube/config")
   --url value                Url where to sent the hook
   --slack-channel-url value  Sent event to slack channel url
   --watch-type value         Type of event to watch on (Warning, Normal) (default: "ALL")
   --watch-kind value         Kind of K8S resource to watch on (Pod, Service). Alias not supported (default: "ALL")
```

# Run in docker container
`docker run -v ~/.kube/config:/config olsynt/kube-watcher run --url {url}  --kube-config /config`


# Install kube-watch in your cluster
Run `kube-watch install` will apply new deployment in the default namespace of cluster in current-context.
Optional to pass any flag same as `kube-watch run` and they will be applied
The applied deployment looks like:
```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    app: kube-watch
  name: kube-watch
spec:
  replicas: 1
  selector:
    matchLabels:
      app: kube-watch
  template:
    metadata:
      labels:
        app: kube-watch
    spec:
      containers:
      - args:
        - run
        - --in-cluster
        # And all flags passed to `kube-watch run` command
        image: olsynt/kubewatch:command-start
        imagePullPolicy: Always
        name: kube-watch
```
# Todo:
* [ ] Tests!
* [ ] Supoort install on different namespaces
* [ ] Add update/uninstall commands to update the image and remove the deployment
* [ ] Support more integrations
* [ ] Support more complexity
* [ ] Support secret to be added in request header

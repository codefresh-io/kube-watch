# Kube watch
Simple tool that watch on events from your cluster and push them as webhooks

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
## testkube set context

Set context data for Testkube Pro

```
testkube set context <value> [flags]
```

### Options

```
      --agent-prefix string   usually don't need to be changed [required for custom cloud mode] (default "agent")
      --agent-token string    Testkube Cloud agent key [required for centralized mode]
      --agent-uri string      Testkube Cloud agent URI [required for centralized mode]
  -k, --api-key string        API Key for Testkube Cloud
      --api-prefix string     usually don't need to be changed [required for custom cloud mode] (default "api")
      --env-id string         Testkube Cloud environment id [required for centralized mode]
      --feature-logs-v2       Logs v2 feature flag
  -h, --help                  help for context
      --kubeconfig            reset context mode for CLI to default kubeconfig based
      --logs-prefix string    usually don't need to be changed [required for custom cloud mode] (default "logs")
      --logs-uri string       Testkube Cloud logs URI [required for centralized mode]
      --master-insecure       should client connect in insecure mode (will use http instead of https)
  -n, --namespace string      Testkube namespace to use for CLI commands
      --org-id string         Testkube Cloud organization id [required for centralized mode]
      --root-domain string    usually don't need to be changed [required for custom cloud mode] (default "testkube.io")
      --ui-prefix string      usually don't need to be changed [required for custom cloud mode] (default "app")
```

### Options inherited from parent commands

```
  -a, --api-uri string   api uri, default value read from config if set (default "http://localhost:8088")
  -c, --client string    client used for connecting to Testkube API one of proxy|direct (default "proxy")
      --insecure         insecure connection for direct client
      --oauth-enabled    enable oauth
      --verbose          show additional debug messages
```

### SEE ALSO

* [testkube set](testkube_set.md)	 - Set resources


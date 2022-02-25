## testkube testsuites get

Get test suite by name

### Synopsis

Getting test suite from given namespace - if no namespace given "testkube" namespace is used

```
testkube testsuites get <testSuiteName> [flags]
```

### Options

```
  -h, --help   help for get
```

### Options inherited from parent commands

```
      --analytics-enabled    should analytics be enabled (default true)
  -c, --client string        Client used for connecting to testkube API one of proxy|direct (default "proxy")
      --go-template string   in case of choosing output==go pass golang template (default "{{ . | printf \"%+v\"  }}")
  -s, --namespace string     kubernetes namespace (default "testkube")
  -o, --output string        output type one of raw|json|go  (default "raw")
  -v, --verbose              should I show additional debug messages
```

### SEE ALSO

* [testkube testsuites](testkube_testsuites.md)	 - Test suites management commands


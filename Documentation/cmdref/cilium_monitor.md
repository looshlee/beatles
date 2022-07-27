<!-- This file was autogenerated via cilium cmdref, do not edit manually-->

## cilium monitor

Monitoring

### Synopsis


The monitor displays notifications and events emitted by the BPF
programs attached to endpoints and devices. This includes:
  * Dropped packet notifications
  * Captured packet traces
  * Debugging information

```
cilium monitor
```

### Options

```
      --from []uint16         Filter by source endpoint id
      --hex                   Do not dissect, print payload in HEX
      --related-to []uint16   Filter by either source or destination endpoint id
      --to []uint16           Filter by destination endpoint id
  -t, --type []string         Filter by event types [agent capture debug drop l7 trace]
  -v, --verbose               Enable verbose output
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.cilium.yaml)
  -D, --debug           Enable debug messages
  -H, --host string     URI to server-side API
```

### SEE ALSO
* [cilium](cilium.html)	 - CLI

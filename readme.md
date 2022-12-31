# Permacast

## Deployment
### Deploy on fly.io
Deploy the application & allocate n v4 IP address.
```bash
flyctl launch
flyctl ips allocate-v4
```

### Configure DNS
Configure your DNS A/AAAA records with the app's IPs.
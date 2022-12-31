# Permacast

## Deployment
### 1. Deploy on fly.io
Deploy the application & allocate an v4 IP address.
```bash
flyctl launch
```

### 2. SSL & DNS (optional)
If not using Fly's domain `{your-app}.fly.dev`, create an SSL certificate for your own domain.

You will also need to allocate a v4 IP addresss:
```bash
flyctl ips allocate-v4
```

Finally, configure your DNS A/AAAA records with the app's IPs.
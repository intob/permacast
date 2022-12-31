# Permacast

## Deployment
### 1. Deploy on fly.io
Deploy the application & allocate an v4 IP address.
```bash
flyctl launch
flyctl ips allocate-v4
```

### 2. SSL
Create an SSL certificate for your domain.

### 3. DNS
Configure your DNS A/AAAA records with the app's IPs.
# Permacast-svc
This makes consuming content on Arweave easier.

This is an HTTP server that responds to requests for files on Arweave, addressed by the transaction id.

The URL pattern is as simple as `{host}/{txId}`.

I have put this service behind Cloudflare's CDN to reduce latency & solve rate-limits. I have also deployed the app in multiple availability zones. This service is available at `https://svc.permacast.io`.

## For example
This is an image on Arweave:
```
https://svc.permacast.io/kpqV8wgI7BgIe2asY-s2eoAXWURwcsQMfZ7dignX3ME
```
![Image on Arweave](https://svc.permacast.io/kpqV8wgI7BgIe2asY-s2eoAXWURwcsQMfZ7dignX3ME)

## Utility
This service solves some issues with the Arweave endpoint to get transaction data.

The current endpoint does not respect the Content-Type tag. For example, go to:
```
https://arweave.net/tx/kpqV8wgI7BgIe2asY-s2eoAXWURwcsQMfZ7dignX3ME/data
```

## Deployment
### 1. Clone
```bash
git clone https://github.com/intob/permacast-svc
```

### 2. Deploy
Deploy the application. I'm using Fly, but you can deploy it anywhere.
```bash
flyctl launch
```

### 3. SSL & DNS (optional)
If not using Fly's domain `{your-app}.fly.dev`, create an SSL certificate for your own domain.

You will also need to allocate a v4 IP addresss:
```bash
flyctl ips allocate-v4
```

Finally, configure your DNS A/AAAA records with the app's IPs.
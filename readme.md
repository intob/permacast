# Permacast-svc
This makes content on Arweave easier to consume.

This is an HTTP server that responds to requests for files on Arweave addressed by their File-Hash.
The URL pattern is as simple as `{Host}/{File-Hash}`

## For example
This is an image on Arweave:
```
https://svc.permacast.io/e89000c615a420acbfd6b1f58558e4be5625f1bd792892821384756a5cc44ef3
```
![Image on Arweave](https://svc.permacast.io/e89000c615a420acbfd6b1f58558e4be5625f1bd792892821384756a5cc44ef3)

## Utility
The host can be swapped out for any deployment or implementation of this service, making it essentially just a gateway to content on Arweave.

## Perf
I have put this service behind Cloudflare's CDN to reduce latency & solve rate-limits. I have also deployed the app in multiple availability zones.

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
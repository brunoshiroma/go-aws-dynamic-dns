# go-aws-dynamic-dns
Simple client to update a route53 A recordName to public ip.  
It´s work like the dyndns, no-ip, etc

# Using
 1. Configure the [aws credentials](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#creating-the-credentials-file)
 2. Run like below example:
```
go-aws-dynamic-dns -z ZONEID -r my-dynamic-dns.example.com -i https://tiny-credit-7e1b.benchtool.workers.dev
```

## Uses
 * [golang](https://golang.org/doc/install)
 * [aws-cli](https://aws.amazon.com/sdk-for-go/)

### Utils
 * [Cloudflare Workers](https://workers.cloudflare.com/) implement the service (https://tiny-credit-7e1b.benchtool.workers.dev) to get public ip, with very basic script
 ```javascript
 addEventListener('fetch', event => {
  event.respondWith(handleRequest(event.request))
})

/**
 * Respond to the request
 * @param {Request} request
 */
async function handleRequest(request) {
  return new Response(request.headers.get('CF-Connecting-IP'))
}
 ```
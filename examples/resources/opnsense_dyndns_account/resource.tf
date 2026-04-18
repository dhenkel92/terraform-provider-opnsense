// Full Cloudflare account
resource "opnsense_dyndns_account" "cloudflare" {
  description = "Cloudflare dyndns for example.com"
  service     = "cloudflare"
  server      = "api.cloudflare.com"
  username    = "user@example.com"
  password    = "secret-api-token"
  resource_id = "zone-id-123"
  hostnames   = ["home.example.com"]
  zone        = "example.com"
  checkip     = "web_dyndns"
  interface   = "wan"
  force_ssl   = true
  ttl         = 300
}

// Minimal DynDNS2 account
resource "opnsense_dyndns_account" "dyndns2" {
  service  = "dyndns2"
  server   = "members.dyndns.org"
  username = "myuser"
  password = "mypass"
  hostnames = [
    "host1.example.org",
    "host2.example.org",
  ]
  checkip = "web_dyndns"
}

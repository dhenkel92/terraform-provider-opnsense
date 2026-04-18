// Look up a dyndns account by its UUID
data "opnsense_dyndns_account" "example" {
  id = "00000000-0000-0000-0000-000000000000"
}

output "current_ip" {
  value = data.opnsense_dyndns_account.example.current_ip
}

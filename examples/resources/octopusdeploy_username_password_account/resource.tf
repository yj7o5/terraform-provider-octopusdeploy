resource "octopusdeploy_account" "username_password_account" {
  account_type = "UsernamePassword"
  name         = "Username-Password Account (OK to Delete)"
  password     = "###########" # get from secure environment/store
  username     = "[username]"
}
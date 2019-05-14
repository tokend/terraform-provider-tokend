# TokenD Terraform Provider

# What is Terraform? I want my initscripts!

It's OK, but you still need [`terraform`](https://www.terraform.io/downloads.html) available in your `$GOPATH/bin`.
Then install provider with `make` and you are basically ready to init.

The only thing you should know about Terraform at this point is that to make magic happen it needs to store *state*
somewhere. And if you want to provision different environments, you have got to have different state stores. See [workspaces](https://www.terraform.io/docs/state/workspaces.html) for one of the solutions.

# Notes

* Tested with Terraform v0.11


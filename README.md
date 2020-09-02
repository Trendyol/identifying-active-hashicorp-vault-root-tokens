### Based on the official doc: [Vault Tokens](https://www.vaultproject.io/docs/concepts/tokens) 
## Token Accessors
When tokens are created, a token accessor is also created and returned. This accessor is a value that acts as a reference to a token and can only be used to perform limited actions:

* Look up a token's properties (not including the actual token ID)
* Look up a token's capabilities on a path
* Renew the token
* Revoke the token

Finally, the only way to "list tokens" is via the auth/token/accessors command, which actually gives a list of token accessors. While this is still a dangerous endpoint (since listing all of the accessors means that they can then be used to revoke all tokens), it also provides a way to audit and revoke the currently-active set of token
So we are decided to visualize Vault Token's details as a table in order to do that we need to first list all of the token accessors and then by iterating over them we got the details of the token and send them to stdout.

This little program helps us about that.It gives us an output with details Token's Display Name, Creation Time , Expiration Time, Attached Policies and Accessor as a table format.

```bash
$ VAULT_ADDR="" VAULT_TOKEN="" go run main.go
```

## Reference
https://www.greenreedtech.com/identifying-active-hashicorp-vault-root-tokens/

# vault secret client
 messing around with vault
 
## export env vars
```bash
export VAULT_ADDR="https://myvault"
export VAULT_TOKEN=""
export VAULT_NAMESPACE=""
```

## cheater scripts

```bash
. init.sh
cd vault
init
get
build
```

## run

```bash
cd vault
./vault mysecretsengine/data/mysecret

```

## curl
```bash
curl \
    --header "X-Vault-Token: $VAULT_TOKEN" \
    --header "X-Vault-Namespace: ${VAULT_NAMESPACE}" \
    --request GET \
    $VAULT_ADDR/<mysecretengine>/data/<mysecret>
```
# delete-glacier-archives
delete glacier archives from output.json


List vaults
```bash
aws glacier list-vaults --account-id ACCOUNT_ID --region eu-central-1
```


Intiate inventory retrieval
```
aws glacier initiate-job --account-id ACCOUNT_ID --region eu-central-1 --job-parameters '{"Type": "inventory-retrieval"}' --vault-name VAULT_NAME
```


Look at status of inventory retrieval job
```
aws glacier describe-job --account-id ACCOUNT_ID --region eu-central-1 --vault-name Backups --job-id JOB_ID
```


When complete, output the inventory to `~/output.json`
```
aws glacier get-job-output --account-id ACCOUNT_ID --region eu-central-1 --vault-name VAULT_NAME --job-id JOB_ID ~/output.json
```


Update AccountID & VaultName in `main.go`
```
go mod tidy && go run .
```

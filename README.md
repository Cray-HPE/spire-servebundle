# Go API Server for spire-bundle

# Example Request

curl https://api-gateway.vshasta.io/spire-bundle

# Example Response

{ "Domain":"vshasta.io","Server":"spire-vshastaio.vshasta.io","CertBundle":"-----BEGIN CERTIFICATE-----\n[TRUNCATED CERTIFICATE]\n-----END CERTIFICATE-----\n"}

OpenAPI for spire-bundle which serves spire server information and CA bundle. This container is deployed in the spire chart.

Re-running the generator:

```
mkdir -p ./schema
cp ./api/openapi.yaml ./schema
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
  -i /local/schema/openapi.yaml \
  --package-name tokens \
  -g go-server \
  -o /local
rm -rf ./schema
```

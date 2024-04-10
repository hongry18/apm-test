# elastic apm with

## elastic run

### start and setup with tls

```
docker-compose \
-f docker-compose.yml \
-f extensions/fleet/fleet-compose.yml \
-f extensions/fleet/agent-apmserver-compose.yml \
up
```

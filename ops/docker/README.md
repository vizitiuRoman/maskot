# Run application

```bash
docker-compose -f ./ops/docker/docker-compose.yaml -f docker-compose.override.yaml up -d
```

# Customize your docker-compose environment

Make [docker-compose.override.yaml](https://docs.docker.com/compose/extends/) with override configuration. For example:

```yaml
---
version: "3.8"

services:
  app:
    build:
      target: application-dev
```

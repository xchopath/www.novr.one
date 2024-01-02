# Docker-compose YAML Explanation

### Sample 1

```
version: '3'

services:

  <service name>:
    build:
      context: .
      dockerfile: ./path-to/Dockerfile
    image: <image name>:<version>
    container_name: <container name>
    restart: always
    ports:
      - <host port>:<container port>
    volumes:
      - ./file:/path-to/file:ro
      - ./path-to:/etc/path-to
    environment:
      - ENV_VAR=${TEST}
    networks:
      - sample-network
    entrypoint: ["python3", "/app/startup.py"]

networks:
  sample-network:
    name: <network name>
    driver: bridge
```

- Placeholder name `<service name>`.
- Service is built from a Dockerfile located at `./path-to/Dockerfile`.
- Image result name `<image name>:<version>`.
- Container name `<container name>`.
- `restart` option is set to `always`, the container will be restarted automatically if it crashes or restarted.
- Expose a `<container port>` with `<host port>`.
- Volume `./file:/path-to/file:ro` read-only file and mount some path.
- The service configured with an environment variable `ENV_VAR`, which is set to `TEST`.
- The service is connected to `sample-network`.
- Entrypoint `["python3", "/app/startup.py"]`, means the command that should be run when the container is started.
- `sample-network` is defined with a name `<network name>` and using the `bridge` driver.

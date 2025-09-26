# Echo Agent

A minimal FastAPI-based agent that echoes incoming requests. The agent uses the Proompteng runtime SDK to load an `AGENT` instance and expose it over `/invoke`.

## Local Usage

```sh
pip install -r apps/echo-agent/requirements.txt
uvicorn apps.echo-agent.app:app --reload
```

Environment variables:

- `MODEL_PROVIDER`: Provider label surfaced in responses.
- `MODEL_NAME`: Model name surfaced in responses.
- `MEMORY_URI`: Optional URI for the backing memory system.
- `AGENT_MODULE`: Override the module used by the loader (default: `app`).

## Docker Image

Build and tag the example agent:

```sh
docker build -t ghcr.io/proompteng/echo-agent:0.1.0 apps/echo-agent
```

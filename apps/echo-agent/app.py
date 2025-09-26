import os
from typing import Any, Dict

from fastapi import FastAPI, HTTPException
from loguru import logger
from pydantic import BaseModel

from sdk.runtime import AgentRuntime, load_agent

app = FastAPI(title="Proompteng Echo Agent", version="0.1.0")


class InvokeRequest(BaseModel):
    input: Dict[str, Any]
    context: Dict[str, Any] | None = None


class InvokeResponse(BaseModel):
    output: Dict[str, Any]


class EchoAgent:
    """A trivial agent that echoes the payload back."""

    def handle(self, input_payload: Dict[str, Any], context: Dict[str, Any] | None = None) -> Dict[str, Any]:
        context = context or {}
        logger.info("EchoAgent invoked with input={} context={}", input_payload, context)
        return {
            "echo": input_payload,
            "model": {
                "provider": os.getenv("MODEL_PROVIDER", "unknown"),
                "name": os.getenv("MODEL_NAME", "unknown"),
            },
            "memoryUri": os.getenv("MEMORY_URI", ""),
            "context": context,
        }


AGENT: AgentRuntime = EchoAgent()
_cached_agent: AgentRuntime | None = None


def get_agent() -> AgentRuntime:
    global _cached_agent
    if _cached_agent is None:
        module_path = os.getenv("AGENT_MODULE", "app")
        _cached_agent = load_agent(module_path)
        logger.info("Loaded agent module {}", module_path)
    return _cached_agent


@app.post("/invoke", response_model=InvokeResponse)
def invoke(request: InvokeRequest) -> InvokeResponse:
    agent = get_agent()
    try:
        output = agent.handle(request.input, request.context or {})
    except Exception as exc:  # pragma: no cover - defensive logging
        logger.exception("Agent invocation failed")
        raise HTTPException(status_code=500, detail=str(exc)) from exc
    return InvokeResponse(output=output)


@app.get("/healthz")
def healthz() -> Dict[str, str]:
    return {"status": "ok"}


@app.get("/readyz")
def readyz() -> Dict[str, str]:
    get_agent()
    return {"status": "ready"}

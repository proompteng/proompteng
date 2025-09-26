"""Runtime helpers bundled with the example agent."""

from importlib import import_module
from typing import Any, Dict, Protocol


class AgentRuntime(Protocol):
    def handle(self, input_payload: Dict[str, Any], context: Dict[str, Any]) -> Dict[str, Any]:
        """Process the inbound request and return a response payload."""


def load_agent(module_path: str) -> AgentRuntime:
    module = import_module(module_path)
    if not hasattr(module, "AGENT"):
        raise AttributeError(f"Module '{module_path}' must define an AGENT instance")

    agent = getattr(module, "AGENT")
    if not callable(getattr(agent, "handle", None)):
        raise TypeError("AGENT must implement a handle(input_payload, context) method")

    return agent

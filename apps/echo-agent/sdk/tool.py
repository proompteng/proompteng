"""Lightweight tool registry for the example agent."""

from dataclasses import dataclass, field
from typing import Any, Callable, Dict, Optional

import requests


@dataclass
class HTTPTool:
    name: str
    url: str
    method: str = "POST"
    timeout: int = 10
    auth_header: Optional[str] = None

    def call(self, payload: Dict[str, Any]) -> requests.Response:
        headers = {"Content-Type": "application/json"}
        if self.auth_header:
            headers["Authorization"] = self.auth_header
        response = requests.request(
            method=self.method,
            url=self.url,
            json=payload,
            headers=headers,
            timeout=self.timeout,
        )
        response.raise_for_status()
        return response


@dataclass
class ToolRegistry:
    http_tools: Dict[str, HTTPTool] = field(default_factory=dict)
    callables: Dict[str, Callable[..., Any]] = field(default_factory=dict)

    def register_http(self, tool: HTTPTool) -> None:
        self.http_tools[tool.name] = tool

    def register_callable(self, name: str, func: Callable[..., Any]) -> None:
        self.callables[name] = func

    def invoke(self, name: str, *args: Any, **kwargs: Any) -> Any:
        if name in self.callables:
            return self.callables[name](*args, **kwargs)
        if name in self.http_tools:
            payload = kwargs.get("payload", {})
            response = self.http_tools[name].call(payload)
            return response.json()
        raise KeyError(f"Tool '{name}' is not registered")

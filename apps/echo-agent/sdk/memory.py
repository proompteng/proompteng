"""In-memory adapter for the example agent."""

from typing import Any, Dict, Optional, Protocol


class MemoryAdapter(Protocol):
    def put(self, key: str, value: Dict[str, Any]) -> None:
        ...

    def get(self, key: str) -> Optional[Dict[str, Any]]:
        ...

    def query(self, expression: str) -> Any:
        ...


class InMemoryAdapter:
    def __init__(self) -> None:
        self._store: Dict[str, Dict[str, Any]] = {}

    def put(self, key: str, value: Dict[str, Any]) -> None:
        self._store[key] = value

    def get(self, key: str) -> Optional[Dict[str, Any]]:
        return self._store.get(key)

    def query(self, expression: str) -> Any:
        return {"expression": expression, "results": list(self._store.values())}

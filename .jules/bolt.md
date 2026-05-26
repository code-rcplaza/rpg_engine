## 2024-05-26 - Prevent Redundant DB Calls in Builder Methods
**Learning:** The usecase layer has a pattern of delegating name building to sub-functions (like `buildComposite` called from `buildComponent`). This can accidentally introduce redundant N+1 queries if the caller already fetched the required data but doesn't pass it down.
**Action:** Always pass pre-fetched data to delegated sub-functions instead of re-querying the repository in the child function. Ensure builder method signatures accept these pre-loaded components.

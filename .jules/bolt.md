## 2024-05-19 - Redundant Query Check in Composition
**Learning:** `buildComponent` was making redundant queries because it performed a check for existence using the exact same query that was run again inside `buildComposite`. It loaded the data to check, but didn't pass it forward.
**Action:** When delegating to functions that need data, check if the parent function already loaded the required data for a conditional check. If so, update the signature of the delegated function to accept the already-loaded data instead of re-querying it.

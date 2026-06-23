## 2024-11-20 - N+1 query prevention in builder delegations
**Learning:** In the usecase layer, when a parent method (e.g., `buildComponent`) pre-fetches data from the repository and then delegates to a sub-function (e.g., `buildComposite`), the sub-function should not re-query the same data. This is a common pattern in the codebase to prevent redundant N+1 database calls.
**Action:** Always pass pre-fetched components to builder methods instead of re-querying the repository inside the delegated sub-function.

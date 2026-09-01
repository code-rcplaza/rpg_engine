## 2024-05-18 - Passing pre-fetched data instead of re-querying
**Learning:** In the usecase layer, there's a pattern of fetching data (e.g., composite parts) to check existence, and then delegating to a sub-function that re-queries the exact same data. This leads to redundant N+1 database calls.
**Action:** Always pass pre-fetched data to delegated sub-functions instead of re-querying the repository.

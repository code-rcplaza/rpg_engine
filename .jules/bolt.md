## 2024-05-24 - Redundant Database Queries in Builder Methods
**Learning:** Passing only the identifier to builder functions can lead to redundant N+1 database queries when the parent function has already fetched the required data. This is common in component-building patterns where a parent checks for existence and the child re-fetches the same data to build the component.
**Action:** Always pass pre-fetched data directly to delegated sub-functions instead of re-querying the repository by identifier to avoid redundant database calls.

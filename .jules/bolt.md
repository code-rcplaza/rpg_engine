## 2024-08-25 - Avoid redundant repository queries by passing pre-fetched data
**Learning:** In the usecase layer, checking if a domain component exists (like fetching composite parts) and then delegating to another method (like `buildComposite`) that fetches the exact same data leads to redundant N+1 database queries.
**Action:** Always pass pre-fetched data to delegated sub-functions instead of re-querying the repository.

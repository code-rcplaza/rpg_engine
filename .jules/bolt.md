## 2024-06-18 - Avoid N+1 queries by passing pre-fetched data
**Learning:** The usecase layer can unintentionally perform redundant N+1 database calls (e.g., repeatedly querying for composite name parts) if delegated builder functions don't accept pre-fetched data.
**Action:** Always pass pre-fetched data to delegated sub-functions instead of re-querying the repository when the data is already available in the caller function's scope.

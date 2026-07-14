## 2026-07-14 - Eliminate redundant database calls by passing pre-fetched data
**Learning:** Passing already fetched data (like composite name parts) down to delegated sub-functions in the usecase layer prevents redundant N+1 database calls that occur when the sub-function re-fetches the same data.
**Action:** Always check if a sub-function requires data that has already been retrieved in the calling function, and pass it as an argument instead of querying the repository again.

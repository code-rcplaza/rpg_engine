## 2024-05-18 - Prevent N+1 queries in builder methods
**Learning:** Checking for data existence in one function and fetching it again in a delegated builder method causes an unnecessary N+1 database call pattern. In `buildComponent`, the composite "first" parts were fetched to check if a composite name was needed, but then `buildComposite` queried for those exact same parts again.
**Action:** Always pass pre-fetched data directly into delegated builder methods instead of re-querying the repository.

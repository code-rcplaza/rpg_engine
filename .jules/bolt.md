## 2026-07-28 - Avoid Redundant Queries in Delegated Functions
**Learning:** Checking for data existence then re-fetching that same data in a delegated sub-function creates redundant queries (N+1 database calls).
**Action:** Always pass pre-fetched data directly to delegated sub-functions.

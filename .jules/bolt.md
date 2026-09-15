## 2024-05-24 - Redundant Database Queries (N+1 Avoidance)
**Learning:** Found a performance bottleneck where data already fetched in a parent function (e.g. `buildComponent`) was being re-queried in a delegated sub-function (`buildComposite`). This leads to unnecessary N+1 database queries.
**Action:** Always pass pre-fetched data to delegated sub-functions instead of re-querying the repository. Look for opportunities to pass data structures down the call stack to eliminate redundant repository calls.

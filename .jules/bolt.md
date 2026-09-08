## 2024-05-24 - Pre-fetching optimization in UseCase Layer
**Learning:** Checking for data existence (e.g., retrieving `parts` to check `len(parts) > 0`) and then discarding that data only to have a sub-function query it again causes an N+1/redundant database call bottleneck.
**Action:** Always pass pre-fetched data directly to delegated sub-functions as an argument instead of re-querying the repository.

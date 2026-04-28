## 2024-04-28 - Redundant DB queries in builder methods
**Learning:** Delegation to builder methods (like `buildComposite`) in this architecture often results in re-querying data already fetched by the caller (like `buildComponent`).
**Action:** Always pass pre-fetched data to delegated sub-functions instead of re-querying the repository.

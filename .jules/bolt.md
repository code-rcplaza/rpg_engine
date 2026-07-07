## 2025-03-09 - Avoid redundant queries in builder methods
**Learning:** The usecase layer was making redundant database calls when checking for existence before delegating to sub-functions (e.g., `buildComposite` re-querying for `first` composite parts already fetched in `buildComponent`).
**Action:** Always pass pre-fetched data down to delegated sub-functions instead of re-querying the repository to avoid redundant database calls.

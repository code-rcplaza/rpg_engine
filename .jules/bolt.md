## 2025-02-28 - Prevent redundant DB calls in usecase builder methods
**Learning:** Usecase builder methods (like `buildComponent` checking existence before delegating to `buildComposite`) sometimes fetch data to check existence, and then the delegated method fetches the exact same data again, resulting in redundant DB calls.
**Action:** Always pass pre-fetched data to delegated sub-functions instead of re-querying the repository.

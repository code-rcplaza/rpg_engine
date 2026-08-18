## 2024-05-15 - Avoid N+1 DB Queries in Usecase Builders
**Learning:** Calling the repository for the exact same data (`FindCompositeParts(raceID, "first")`) in parent (`buildComponent`) and child (`buildComposite`) functions creates redundant DB queries.
**Action:** Always pass pre-fetched data as arguments to delegated builder functions instead of re-querying the repository.

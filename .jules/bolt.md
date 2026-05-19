## 2024-05-19 - Prevent redundant DB queries when delegating to sub-functions in the usecase layer
**Learning:** In the usecase layer, delegating to sub-functions can inadvertently cause redundant N+1 DB calls if the parent function has already queried the same data. In `name_generator.go`, `buildComponent` queried `FindCompositeParts(raceID, "first")` to check for existence, and then `buildComposite` queried it again.
**Action:** Always pass pre-fetched data to delegated sub-functions instead of re-querying the repository.

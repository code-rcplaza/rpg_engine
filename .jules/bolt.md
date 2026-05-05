## 2025-02-20 - N+1 Query in Use Case Layer

**Learning:** Redundant repository calls can occur when a parent function checks for data existence and then delegates to a child function that re-queries the same data. In `internal/usecase/name_generator.go`, `buildComponent` was querying `FindCompositeParts` to check for parts, and then `buildComposite` queried it again.

**Action:** When a parent function fetches data from a repository and delegates work to a child function, pass the already-fetched data as an argument to the child instead of having the child re-query the repository. This avoids N+1 style redundant database calls.

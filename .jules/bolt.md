## 2024-06-09 - Avoid N+1 queries by passing pre-fetched data
**Learning:** The usecase layer can easily introduce N+1 query problems by making redundant calls to the repository for data it already possesses. For example, `buildComposite` was re-querying for `firstParts` even though `buildComponent` had just fetched them.
**Action:** When a method delegates logic to a sub-function and needs data that was already fetched in the caller method, pass the pre-fetched data as an argument to the sub-function rather than re-querying the database.

## 2024-05-19 - N+1 database call optimization in NameGenerator
**Learning:** Pre-fetched data should be passed to delegated sub-functions to avoid redundant database calls. In `buildComponent`, `parts` were being fetched to check if composite parts exist, and then `buildComposite` was fetching the exact same `first` parts again.
**Action:** When a method fetches data to conditionally call another method, pass that fetched data as an argument to the second method if it also needs it, preventing duplicate repository/database queries.

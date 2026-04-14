## 2026-04-14 - [Redundant DB Query in NameGenerator]
**Learning:** Checking for the existence of composite parts inside  by fetching them, and then subsequently re-fetching the exact same 'first' parts inside  results in an unnecessary database query per generated compound name.
**Action:** When a method fetches data to verify its existence before delegating to a helper method that also needs that data, pass the already-fetched data as an argument to the helper instead of re-fetching.
## 2026-04-14 - [Redundant DB Query in NameGenerator]
**Learning:** Checking for the existence of composite parts inside `buildComponent` by fetching them, and then subsequently re-fetching the exact same 'first' parts inside `buildComposite` results in an unnecessary database query per generated compound name.
**Action:** When a method fetches data to verify its existence before delegating to a helper method that also needs that data, pass the already-fetched data as an argument to the helper instead of re-fetching.

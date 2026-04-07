# Bolt's Journal

## 2024-04-07 - N+1 query elimination in halfling composite last names
**Learning:** Found a redundant database query in `buildComposite` within `NameGenerator` in `internal/usecase/name_generator.go`. `buildComponent` calls `FindCompositeParts` for `"first"` to check if composite parts exist. Then, if they do, it calls `buildComposite`, which unnecessarily fetches `"first"` AGAIN before fetching `"second"`.
**Action:** Pass the `firstParts` retrieved in `buildComponent` directly to `buildComposite` instead of re-fetching it from the DB.

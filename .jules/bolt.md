## 2026-07-21 - Redundant Database Calls in Recursive/Delegated Builders

**Learning:** In the usecase layer, when a method (`buildComponent`) fetches data (`FindCompositeParts("first")`) and delegates to a sub-builder (`buildComposite`), the sub-builder might fetch the exact same data again if it's not passed as an argument. This leads to hidden N+1 queries.

**Action:** Always inspect delegated builder methods to see if they query the repository for data that was already fetched by the caller. Refactor signatures to pass pre-fetched data down the chain instead of re-querying.

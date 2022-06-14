This is a high-level mockup of the Acme Inc. backend project. It is a bit contived,
but it should give one a fair idea of my current skill level.

Points to note:
1. I tried to write it so that the database could be swappable. Here I'd use BigQuery.
    The reason for that is offers all the benefits of a NOSQL type, without any disadvantages.
    However, I'm not sure if it's a good idea to use BigQuery for this project. It's also could-native.
2. Gizmos and Widgets tables would need to be joined on the Descirption field or the SKU number for de-duping.
3. The resulting table db.AcmeTableProducts would be indexed by the ID field (not the unique one, but the one provided by the original tables).
4. Two REST endpoints are created, one for single retrieval and one for batch. 
   1. Batch query is constructed on the fly, and in case there's a max limit to what the DB can take in, pagination is an option.
   2. Functional options are used (I'm a fan). However, they're not protected against bogus entries for the sake of time.
5. I was thinking to implement a basic queue and a caching layer, but I learned that caching for SELECT queries is a fool's errand; modern dbs are optimized enough.
6. Provided the main table is populated, and the cloud auth is configured, the app should be return valid-ish responses.
A demo of an "ideal" architecture for a GO API backend using no ORM.

## Get started
1. This uses sqlite3, so make sure it's installed on your system.
1. Do `sqlite3 demo.db` to start your db.
1. Inside sqlite prompt do: `.read db/schema.sql` to initialize your tables.
1. To populate the db with a few records do: `.read db/seed.sql`.

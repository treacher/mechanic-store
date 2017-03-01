CREATE DATABASE "mechanic-store";
REVOKE connect ON DATABASE "mechanic-store" FROM PUBLIC;
CREATE USER "mechanic-store" WITH PASSWORD '<Fetch Password>';
GRANT ALL PRIVILEGES ON DATABASE "mechanic-store" to "mechanic-store";

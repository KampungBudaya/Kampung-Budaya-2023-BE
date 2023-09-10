#!/bin/sh
set -e

source .env

# Run your SQL migration script using the MariaDB client
echo "Running SQL migrations..."
mysql -h $DB_HOST -u $DB_USERNAME -p$DB_PASSWORD $DB_DATABASE < migration/000001_create_contests_table.up.sql
mysql -h $DB_HOST -u $DB_USERNAME -p$DB_PASSWORD $DB_DATABASE < migration/000002_create_participants_table.up.sql
mysql -h $DB_HOST -u $DB_USERNAME -p$DB_PASSWORD $DB_DATABASE < migration/000003_create_users_table.up.sql
mysql -h $DB_HOST -u $DB_USERNAME -p$DB_PASSWORD $DB_DATABASE < migration/000004_create_roles_table.up.sql
mysql -h $DB_HOST -u $DB_USERNAME -p$DB_PASSWORD $DB_DATABASE < migration/000005_create_user_has_roles_table.up.sql
echo "SQL migrations complete."

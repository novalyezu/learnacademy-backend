SELECT 'CREATE DATABASE db_learn_academy' 
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'db_learn_academy')\gexec

GRANT ALL PRIVILEGES ON DATABASE db_learn_academy TO postgres;

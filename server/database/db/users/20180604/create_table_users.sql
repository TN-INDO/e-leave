CREATE TABLE IF NOT EXISTS users 
(
  "employee_number" int NOT NULL PRIMARY KEY,
  "name" text NOT NULL,
  "gender" text NOT NULL,
  "position" text NOT NULL,
  "start_working_date" text NOT NULL,
  "mobile_phone" text NOT NULL,
  "email" varchar(30) NOT NULL,
  "password" varchar(100) NOT NULL,
  "role" text NOT NULL,
  "supervisor_id" int,
  "created_at" timestamp with time zone NOT NULL default CURRENT_TIMESTAMP,
  "updated_at" timestamp with time zone    
)
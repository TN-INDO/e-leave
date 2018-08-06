CREATE TABLE IF NOT EXISTS leave_request
(
  "id" int PRIMARY KEY NOT NULL,
  "employee_number" int NOT NULL REFERENCES users(employee_number) ON DELETE CASCADE,
  "type_leave_id" int NOT NULL,
  "reason" text NOT NULL,
  "date_from" text NOT NULL,
  "date_to" text NOT NULL,  
  "half_dates" TEXT [],
  "back_on" text NOT NULL,
  "total" float NOT NULL,
  "contact_address" text NOT NULL,
  "contact_number" text NOT NULL,
  "status" text NOT NULL,
  "action_by" text,
  "reject_reason" text,  
  "created_at" date NOT NULL default CURRENT_TIMESTAMP,
  "updated_at" timestamp with time zone
)
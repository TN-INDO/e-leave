CREATE TABLE IF NOT EXISTS user_type_leave
(
  "id" SERIAL PRIMARY KEY,
  "employee_number" int NOT NULL REFERENCES users(employee_number) ON DELETE CASCADE,  
  "type_leave_id" int NOT NULL REFERENCES type_leave(id) ON DELETE CASCADE,
  "leave_remaining" float NOT NULL
);

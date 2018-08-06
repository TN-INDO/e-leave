ALTER TABLE leave_request 
   ADD CONSTRAINT leave_request_type_leave_id_fkey
   FOREIGN KEY (type_leave_id) 
   REFERENCES type_leave(id);
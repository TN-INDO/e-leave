CREATE SEQUENCE leave_request_id_seq        
    NO MINVALUE
    NO MAXVALUE
    START WITH 12345
    INCREMENT BY 1
    CACHE 1;    
        
ALTER sequence leave_request_id_seq OWNED BY leave_request.id;
ALTER TABLE leave_request ALTER COLUMN id SET DEFAULT NEXTVAL('leave_request_id_seq'::regclass);
